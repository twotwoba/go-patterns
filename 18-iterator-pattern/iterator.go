package iterator

import "fmt"

/* ============== 理论 ============== */
// 迭代器模式是一种行为设计模式，它允许你顺序访问一个聚合对象中的元素，
// 而无需暴露其底层表示。
//
// 迭代器模式的主要优点是：
// 1. 提供了一种统一的遍历方式，使得遍历操作与聚合对象的内部表示分离。
// 2. 支持多种遍历方式，如顺序遍历、逆序遍历等。
// 3. 提高了代码的可读性和可维护性。
//
// 迭代器模式的主要缺点是：
// 1. 增加了系统的复杂度，需要引入新的类和接口。
// 2. 可能会影响系统的性能，因为需要创建新的对象和接口。
//

/*
	设计思想：
		1. Iterator结构体
			实现Next()  HasNext()方法
		2. Container容器
			容器实现添加 移除Visitor 和
*/
//创建迭代器
type Iterator struct {
	index int
	Container
}

func (i *Iterator) Next() Visitor {
	fmt.Println(i.index)
	visitor := i.list[i.index]
	i.index += 1
	return visitor
}

func (i *Iterator) HasNext() bool {
	if i.index >= len(i.list) {
		return false
	}
	return true
}

// 创建容器
type Container struct {
	list []Visitor
}

func (c *Container) Add(visitor Visitor) {
	c.list = append(c.list, visitor)
}

func (c *Container) Remove(index int) {
	if index < 0 || index > len(c.list) {
		return
	}
	c.list = append(c.list[:index], c.list[index+1:]...)
}

// 创建Visitor接口
type Visitor interface {
	Visit()
}

// 创建具体的visitor对象
type Teacher struct{}

type Analysis struct{}

func (t *Teacher) Visit() {
	fmt.Println("this is teacher visitor")
}

func (a *Analysis) Visit() {
	fmt.Println("this is analysis visitor")
}

// 工厂方法创建迭代器
func NewIterator() *Iterator {
	return &Iterator{
		index:     0,
		Container: Container{},
	}
}
