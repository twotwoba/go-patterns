package generator

/* ============== 理论 ============== */
// 生成器模式，略使用了go的channel个并发特性，相当于实现其他语言的yield功能
//
// 1. 生产者和消费者被完全解耦
// 2. 不会立即在内存中创建一个包含一百万个整数的数组。内存中在任何时候通常只有一个整数在“传送带”上。
// 这对于处理大数据集或无限流非常高
// 3. 生产和消费可以并发进行，提高了程序的效率
/*
	设计思想：
		函数返回一个只读的 <-chan
		在函数内部开一个goruntine并发生成值放入chan中
*/

func Count(start, end int) <-chan int {
	ch := make(chan int) // 无缓冲chan

	go func(ch chan int) {
		for i := start; i <= end; i++ {
			ch <- i // 发送到通道里，产生阻塞，直到被消费
		}
		close(ch)
	}(ch)

	return ch
}
