package pool

import "fmt"

/* ============== 理论：GO特色模式 ============== */
// 对象池模式
/*
根据需求将预测的对象保存到channel中， 用于对象的生成成本大于维持成本
*设计思想

	1.对象结构体
	2.类型为结构体指针的channel
	3.New方法, 创建新的对象放到channel中
*/
type Object struct {
	Name string
}

type Pool chan *Object

func NewPool(count int) *Pool {
	pool := make(Pool, count)
	// 构造完成需要关闭channel, 否则会引起deadline
	// PS：如果你需要的是在整个程序生命周期都存在的话，则不需要下面这行
	defer close(pool)
	// 循环创建对象，放入pool中
	for i := 0; i < count; i++ {
		pool <- new(Object)
	}
	return &pool
}

func (obj *Object) Do() {
	fmt.Println(&obj)
}
