package semaphore

import (
	"errors"
	"time"
)

/* ============== 理论：GO特色模式 ============== */
// 信号量模式，是一种**并发模式 (Concurrency Pattern)**。
// 它主要关注的是**解决多线程或多协程环境下的同步和资源管理问题**。它关心的是程序的**运行时行为**，如何安全、高效地协调多个并发执行流。

// 信号量是一个非常基础且重要的**并发原语 (Concurrency Primitive)**。
// 您可以把它想象成一个用来控制访问特定资源“许可证”或“令牌”的计数器

/*
设计思想：

	1.定义接口包含Acquire和release行为
	2.定义结构体， 包含chan和过期时间属性
	3. 在Acquire中实现channel读入
	4. 在release中channel 读出， 阻塞超时返回错误
*/
var (
	ErrNoTickets      = errors.New("semaphore: could not acquire semaphore")
	ErrIllegalRelease = errors.New("semaphore: can't release semaphore without acquiring it first")
)

type Interface interface {
	Acquire() error
	Release() error
}

// 定义结构体， 信号量使用chan struct{}
type Semaphore struct {
	sem     chan struct{}
	timeout time.Duration
}

func (s *Semaphore) Acquire() error {
	select {
	case s.sem <- struct{}{}: // 尝试发送一个值到 channel
		return nil
	case <-time.After(s.timeout): // 如果超时
		return ErrNoTickets
	}
}

func (s *Semaphore) Release() error {
	select {
	case <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return ErrIllegalRelease
	}
}

func New(tickets int, timeout time.Duration) Interface {
	return &Semaphore{
		sem:     make(chan struct{}, tickets), // 核心！缓冲区的大小，就是信号量的容量
		timeout: timeout,
	}
}
