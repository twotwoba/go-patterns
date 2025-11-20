package chain

import "fmt"

/* ============== 理论 ============== */
// 责任链模式，是一种行为设计模式，它允许你将请求沿着链传递，直到某个处理者处理它为止。
// 可以想象员工报销
//

/*
	核心是结构体包含下一个结构体的引用
	状态模式和职责链模式区别：
	    状态模式：通常是**一个**上下文对象，它的行为随着**内部状态**的改变而改变。状态的转换通常是确定的，并且由上下文或状态自身来控制。
		责任链模式：通常有**多个**独立的处理者对象，它们被链接起来处理外部请求。请求在链上的传递是动态的，发送者不知道最终由谁处理。

	设计思想：
		1. 一个Interface接口，用来封装方法集合
		2. 具体struct, 匿名组合接口(对象链中next对象引用)
*/
// 处理者接口
type Interface interface {
	SetNext(next Interface)  // 设置下一个处理者 参数不确定，所以这里使用接口
	HandleEvent(event Event) // 处理方法
}

// 定义ObjectA struct
type ObjectA struct {
	Interface // 持有下一个处理者的引用
	Level     int
	Name      string
}

func (ob *ObjectA) SetNext(next Interface) {
	ob.Interface = next
}

func (ob *ObjectA) HandleEvent(event Event) {
	if ob.Level == event.Level {
		fmt.Printf("%s 处理这个事件 %s\n", ob.Name, event.Name)
	} else {
		if ob.Interface != nil {
			ob.Interface.HandleEvent(event)
		} else {
			fmt.Println("无法处理")
		}
	}
}

// 定义ObjectB struct
type ObjectB struct {
	Interface
	Level int
	Name  string
}

func (ob *ObjectB) SetNext(next Interface) {
	ob.Interface = next
}

func (ob *ObjectB) HandleEvent(event Event) {
	if ob.Level == event.Level {
		fmt.Printf("%s 处理这个事件 %s\n", ob.Name, event.Name)
	} else {
		if ob.Interface != nil {
			ob.Interface.HandleEvent(event)
		} else {
			fmt.Println("无法处理")
		}
	}
}

// 在链上传递的结构体
type Event struct {
	Level int
	Name  string
}
