package observer

import "fmt"

/* ============== 理论 ============== */
// 观察者模式对于前端来讲是比较熟悉的了，可以和发布订阅模式对比学习
//
// 观察者模式的执行角色是2个：观察者和被观察者
// 对于观察者与被观察者一般都会提供接口，实现多态性和解耦，方便以后拓展
//
// 1. 观察者需要提供一个接收通知的方法
// 2. 被观察者会维护一个观察者列表，当被观察者状态改变时，遍历通知所有的观察者
//

// 1.设计观察者和被观察者接口
type Observer interface {
	Receive(event string)
}
type Notifier interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify(event string)
}

// 具体观察者
type InvestorObserver struct {
	Name string
}

func (invester *InvestorObserver) Receive(event string) {
	fmt.Printf("%s 收到事件通知 %s\n", invester.Name, event)
}

// 股票被观察者
type ShareNotifier struct {
	Price  float64
	oblist []Observer //收集观察者表
}

func (share *ShareNotifier) Register(observer Observer) {
	share.oblist = append(share.oblist, observer)
}

func (share *ShareNotifier) Remove(observer Observer) {
	if len(share.oblist) == 0 {
		return
	}
	for i, ob := range share.oblist {
		if ob == observer {
			share.oblist = append(share.oblist[:i], share.oblist[i+1:]...)
		}
	}
}

// 通知所有观察者
func (share *ShareNotifier) Notify(event string) {
	for _, ob := range share.oblist {
		ob.Receive(event)
	}
}

func NewInvestorObserver(name string) *InvestorObserver {
	return &InvestorObserver{Name: name}
}

func NewShareNotifier(price float64) *ShareNotifier {
	return &ShareNotifier{Price: price}
}
