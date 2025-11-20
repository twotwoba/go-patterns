package pubsub

import (
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"
)

// TestNewPublisher 验证发布者是否被正确创建
func TestNewPublisher(t *testing.T) {
	p := NewPublisher(10, 100*time.Millisecond)
	if p == nil {
		t.Fatal("NewPublisher returned nil")
	}
	if p.buffer != 10 {
		t.Errorf("expected buffer size 10, got %d", p.buffer)
	}
	if p.timeout != 100*time.Millisecond {
		t.Errorf("expected timeout 100ms, got %v", p.timeout)
	}
	if p.subscribers == nil {
		t.Error("subscribers map was not initialized")
	}
}

// TestSubscribeAndPublish 测试基本的订阅和发布功能
func TestSubscribeAndPublish(t *testing.T) {
	p := NewPublisher(1, 10*time.Millisecond)
	defer p.Close()

	// 订阅所有主题
	sub := p.Subscribe()

	p.Publish("hello")

	msg, ok := <-sub
	if !ok {
		t.Fatal("channel was closed unexpectedly")
	}
	if msg.(string) != "hello" {
		t.Errorf("expected message 'hello', got '%s'", msg.(string))
	}
	if msg != nil {
		fmt.Printf("msg: %v\n", msg)
	}
}

// TestSubscribeTopic 测试基于主题的订阅
func TestSubscribeTopic(t *testing.T) {
	// 可切换 buf 测试其他情况
	p := NewPublisher(1, 10*time.Millisecond)
	defer p.Close()

	// 只订阅包含 "error" 的字符串消息
	stringTopic := func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "error")
		}
		return false
	}
	sub := p.SubscribeTopic(stringTopic)

	// 发布一个不匹配的消息
	p.Publish("this is a normal message")
	// 发布一个匹配的消息
	p.Publish("this is an error message")
	// 发布另一个错误消息
	p.Publish("this is another error message")

	// 检查是否收到了匹配的消息, 阻塞式提取消息
	select {
	case msg := <-sub:
		if !strings.Contains(msg.(string), "error") {
			t.Errorf("received a message that should have been filtered: %s", msg.(string))
		}
		fmt.Printf("msga: %v\n", msg)
	case <-time.After(50 * time.Millisecond):
		t.Fatal("timed out waiting for the correct message")
	}

	// 检查不匹配的消息是否被过滤，非阻塞式提取，有default
	select {
	case msg := <-sub:
		t.Errorf("should not have received a second message, but got: %s", msg.(string))
		fmt.Printf("msgb: %v\n", msg)
	default:
		// 这是预期的行为
	}
}

// TestExit 测试订阅者退出
func TestExit(t *testing.T) {
	p := NewPublisher(1, 10*time.Millisecond)
	defer p.Close()

	sub := p.Subscribe()
	sub2 := p.Subscribe()

	p.Exit(sub)

	// 检查退出的 channel 是否已关闭
	_, ok := <-sub
	if ok {
		t.Fatal("expected channel to be closed after Exit()")
	}

	// 检查其他订阅者是否仍然正常
	p.Publish("hello")
	msg, ok := <-sub2
	if !ok || msg.(string) != "hello" {
		t.Error("second subscriber was affected by the first one's exit")
	}
}

// TestClose 测试关闭发布者
func TestClose(t *testing.T) {
	p := NewPublisher(1, 10*time.Millisecond)

	sub1 := p.Subscribe()
	sub2 := p.Subscribe()

	p.Close()

	if _, ok := <-sub1; ok {
		t.Error("sub1 should be closed")
	}
	if _, ok := <-sub2; ok {
		t.Error("sub2 should be closed")
	}

	if len(p.subscribers) != 0 {
		t.Error("subscribers map should be empty after Close()")
	}
}

// TestPublishTimeout 测试发布超时
func TestPublishTimeout(t *testing.T) {
	// 使用一个非常小的超时时间，且订阅者 channel 无缓冲
	p := NewPublisher(0, 10*time.Millisecond)
	defer p.Close()

	_ = p.Subscribe() // 创建一个订阅者，但我们不去读取它的 channel

	// 这个发布操作应该会超时，但不会永久阻塞
	start := time.Now()
	p.Publish("message that will time out")
	duration := time.Since(start)

	if duration > 100*time.Millisecond {
		t.Errorf("Publish took too long (%v), timeout is not working", duration)
	}
}

// TestConcurrency 测试并发安全
func TestConcurrency(t *testing.T) {
	p := NewPublisher(10, 50*time.Millisecond)
	defer p.Close()

	var wg sync.WaitGroup
	numGoroutines := 100

	// 启动多个 goroutine 并发地订阅、发布和退出
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 每个 goroutine 订阅
			sub := p.Subscribe()

			// 发布一些消息
			p.Publish(id)

			// 尝试读取消息
			<-sub

			// 最后退出
			p.Exit(sub)
		}(i)
	}

	wg.Wait()
}
