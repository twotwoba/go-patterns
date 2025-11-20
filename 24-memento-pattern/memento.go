package memento

/*
	跟原型模式有些相识
	设计思想：
		1. 发起人角色 Originator： 记录当前的内部状态， 负责创建和恢复备忘录数据
		2. 备忘录角色 Memento :    存储发起人对象的内部状态
		3. 管理者角色 Caretaker:   负责保存备忘录对象
	*代码上实现逻辑：
		通过备忘录类来专门存储对象状态， 当需要保存状态时， 管理者从发起人拿到状态，保存到备忘录中(即将一个对象的状态保存到另一个对象中)
	使用场景：需要保存和恢复数据的相关状态， 提供回滚操作
*/
/*创建发起人struct*/
//在实际开发过程中，根据需要变为对应的数据结构, 如果多状态的，就保存一个栈
type Originator struct {
	state string
}

func (o *Originator) GetState() string {
	return o.state
}

func (o *Originator) SetState(state string) {
	o.state = state
}

/*创建Memento struct*/
//属性可以为state,或Originator对象
type Memento struct {
	Originator
}

func (m *Memento) GetState() string {
	return m.Originator.state
}

func (m *Memento) SetState(originator Originator) {
	m.Originator = originator
}

/*创建管理者Caretaker*/
//实现创建备忘录和恢复发起者状态方法
type Caretaker struct{}

func (c *Caretaker) CreateMemento(originator Originator) Memento {
	return Memento{originator}
}

func (c *Caretaker) RecoverOriginator(memento Memento) Originator {
	return memento.Originator
}

/*
   封装性：您的实现中，`Memento` 结构体通过匿名嵌入 `Originator`，
   实际上将其内部状态暴露给了所有能接触到 `Memento` 的代码。在 Go 中，
   如果想实现更严格的封装（让 `Caretaker` 无法触及 Memento 的内部），
   通常会使用**包级私有 (package private)** 的特性。
   我们可以将 `Memento` 的字段设为小写（包私有），
   并确保 `Originator` 和 `Memento` 在同一个包中，而 `Caretaker` 在另一个包中。
   这样 `Caretaker` 就无法直接访问 `Memento` 的内部状态了。

   实现方式：上述代码通过直接复制整个 `Originator` 对象的值来创建备忘录。
   这对于简单的结构体是有效的。如果 `Originator` 包含了切片、map或指针等引用类型，
   那么就需要实现**深拷贝**逻辑（类似我们在原型模式中讨论的），否则恢复的只是一个浅拷贝，
   可能会导致问题。
*/
