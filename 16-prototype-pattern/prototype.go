package prototype

/*
通过将已经存在的实例赋值给新的变量来完成clone, 可定制clone对象
*/
type Example struct {
	Description string
}

// 实现Clone
func (e *Example) Clone() *Example {
	res := *e
	return &res
}

func New(des string) *Example {
	return &Example{
		Description: des,
	}
}

// 上面的比较鸡肋。
// 深拷贝 对于指针类型 可能更有用点
// type Details struct {
//     Category string
// }

// type ComplexObject struct {
//     Name    string
//     Friends []string      // 切片是引用类型
//     Details *Details      // 指针是引用类型
// }

// func (co *ComplexObject) Clone() *ComplexObject {
//     // 1. 开始浅拷贝
//     res := *co

//     // 2. 对切片进行深拷贝
//     res.Friends = make([]string, len(co.Friends))
//     copy(res.Friends, co.Friends)

//     // 3. 对指针指向的对象进行深拷贝
//     if co.Details != nil {
//         detailsCopy := *co.Details
//         res.Details = &detailsCopy
//     }

//     return &res
// }
