package decorator

import "fmt"

/* ============== 理论 ============== */
// 装饰模式使用对象组合的方式动态改变或增加对象行为，在原对象的基础上增加功能
//
// 1.因此需要用到装饰器模式的时候一定是已经有了一个源对象了-这里是被装饰结构体
// 2.想要装饰源结构体的对象就必须提供接口来让源结构体和装饰结构体看上去一样！
// 3.装饰结构体中, 采用匿名组合的方式将接口Component作为结构体的属性
// 4.装饰方法中改动

// 源结构体
type Fruit struct {
	Count       int
	Description string
}

func (f *Fruit) Describe() string {
	return f.Description
}
func (f *Fruit) GetCount() int {
	return f.Count
}

// 如果想要给 Fruit 类型添加装饰器就必须提供一个 Fruit 实现了的接口！

// 装饰器模式的核心目的是：**在不改变对象自身结构的情况下，动态地给一个对象添加一些额外的职责（功能）**。
// 为了实现这个目的，必须保证两件事：
// a. 装饰器（Decorator）和被装饰的原始对象（Component）**对外看起来是同一种类型**。
// 因此必须需要一个接口，来让这两者同时实现这个接口达到看起来一致的效果
// b. 装饰器内部必须**持有一个被装饰对象的引用**。在结构体中使用接口来承载
type Component interface {
	Describe() string
	GetCount() int
}
type AppleDecorator struct {
	Component // 接口
	Type      string
	Num       int
}

func (apple *AppleDecorator) Describe() string {
	return fmt.Sprintf("%s, %s", apple.Component.Describe(), apple.Type)
}
func (apple *AppleDecorator) GetCount() int {
	return apple.Component.GetCount() + apple.Num
}

func CreateAppleDecorator(c Component, t string, n int) Component {
	return &AppleDecorator{c, t, n}
}
