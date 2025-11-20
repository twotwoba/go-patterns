package template

import "fmt"

/* ============== 理论 ============== */
// 模板方法模式，核心：在一个函数中定义一个算法的骨架，而将一些步骤延迟到子类中去实现
// 所以就是 --> 子类匿名组合父类实现接口继承

// 1. 定义一个接口，只包含算法中可变的步骤
type Actioner interface {
	BeforeAction()
	GetName() string
}

//  2. 定义模板方法，它是一个独立的函数
//     它接收 Actioner 接口作为参数，定义了算法骨架
func PerformExit(a Actioner) {
	a.BeforeAction()                   // 步骤1 (可变)
	fmt.Println(a.GetName() + " exit") // 步骤2 (固定逻辑的一部分)
}

// --- 具体实现 ---

// 3. 实现一个通用的基础结构体 (可选，用于复用)
type Person struct {
	name string
}

func (p *Person) GetName() string {
	return p.name
}

// 4. Boy 结构体，嵌入 Person 并实现 Actioner 接口
type Boy struct {
	Person
}

func (b *Boy) BeforeAction() {
	fmt.Println("Boy named", b.name, "is preparing to exit...")
}

// 5. Girl 结构体，嵌入 Person 并实现 Actioner 接口
type Girl struct {
	Person
}

func (g *Girl) BeforeAction() {
	fmt.Println("Girl named", g.name, "is waving goodbye...")
}

/* ============== 下面是原仓库的例子 ============== */
/*
	模版方法的核心是父结构体包含接口的引用，同时子类匿名组合父类实现接口继承
	设计思想:
		1. 定义一个接口Shape
		2. 实现父struct, 并接口继承Shape, 同时fu类中包含子类引用，用来调用子类的方法
		3. 实现子struct, 匿名组合父struct， 这样子类也实现接口继承
*/
// //定义接口
// type Shape interface {
// 	SetName(name string)
// 	BeforeAction()
// 	Exit()
// }

// // 定义父类Person
// type Person struct {
// 	name     string
// 	Concrete Shape //具体子类的引用， 因为子类继承了接口
// }

// func (p *Person) SetName(name string) {
// 	p.name = name
// }

// func (p *Person) BeforeAction() {
// 	//将具体的action延迟到子类中执行
// 	p.Concrete.BeforeAction()
// }

// func (p *Person) Exit() {
// 	p.BeforeAction()
// 	fmt.Println(p.name + "exit")
// }

// // 定义具体子类，且实现具体的action
// type Boy struct {
// 	Person //匿名组合实现继承
// }

// // 重写BeforeAction
// func (b *Boy) BeforeAction() {
// 	fmt.Println(b.name)
// }

// type Girl struct {
// 	Person //匿名组合实现继承
// }

// func (g *Girl) BeforeAction() {
// 	fmt.Println(g.name)
// }
