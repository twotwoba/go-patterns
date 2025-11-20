package proxy

import "fmt"

/* ============== 理论 ============== */
// 代理模式就是对源对象进行了一层托管，想要访问源对象都需要经过代理
//
// 1.与装饰着模式一样，想要使用该模式都一定先有一个源对象，这里是被代理结构体
// 2.同样的想要实现代理也得先长得一样，所以需要一个2者都实现的接口
// 3.代理对象包含了真实对象，通过同样的接口拦截不同的行为

// 1. 被代理对象
type Object struct {
	action string
}

// 2. 被代理和代理都得实现的接口
type IObject interface {
	ObjDo(action string)
}

func (obj *Object) ObjDo(action string) {
	fmt.Printf("I can %s \n", action)
}

// 3. 代理对象 属性为 真实对象，通过同样的接口，拦截各种行为
type ProxyObject struct {
	object *Object
}

// 拦截作用
func (p *ProxyObject) ObjDo(action string) {
	// 懒实例化
	if p.object == nil {
		p.object = new(Object)
	}
	//
	if action == "run" {
		p.object.ObjDo(action)
	}
}
