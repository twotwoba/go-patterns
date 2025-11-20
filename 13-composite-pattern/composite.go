package composite

import "fmt"

/* ============== 理论 ============== */
// 核心：**将对象组合成《树形结构》以表示“部分-整体”的层次结构。
// 组合模式使得用户对单个对象（叶子节点）和组合对象（树枝节点）的使用具有一致性。
//
// 组合模式最大的优点**就是，客户端代码可以非常简单。客户端只需要与 `Component` 接口交互，
// 而无需用 `if/else` 来区分它当前处理的是一个叶子还是一个组合对象。
/*
	设计思想：
		struct不依赖interface
		1. 包含角色：
			1). 共同的接口MenuComponent，为root和leaf结构体共有的方法
			2). root结构体(包含leaf列表)和leaf结构体
			3). 将结构体中共同部分的数据抽离，使用匿名组合的方式实现2中的两类结构体
*/
// Menu示例如何使用组合设计模式：Menu和MenuItem

// 抽离共同属性部分
type MenuDesc struct {
	name        string
	description string
}

func (desc *MenuDesc) Name() string {
	return desc.name
}
func (desc *MenuDesc) Description() string {
	return desc.description
}

// MenuItem组合，继承了MenuDesc的方法
type MenuItem struct {
	MenuDesc
	price float32
}

func NewMenuItem(name, description string, price float32) *MenuItem {
	return &MenuItem{
		MenuDesc: MenuDesc{
			name:        name,
			description: description,
		},
		price: price,
	}
}

// 实现MenuItem Price方法和Print()
func (item *MenuItem) Price() float32 {
	return item.price
}
func (item *MenuItem) Print() {
	fmt.Printf("	%s, %0.2f\n", item.name, item.price)
	fmt.Printf("	-- %s\n", item.description)
}

/*
MenuGroup, 这里引入接口MenuComponent, 因为child类型是不确定的
此外，由于接口为Menu和MenuItem的共同接口，所以包含Price和Print方法
*/
type MenuComponent interface {
	Price() float32
	Print()
}
type MenuGroup struct {
	child []MenuComponent
}

// MenuGroup需要实现Add、Remove、Find方法
func (group *MenuGroup) Add(component MenuComponent) {
	group.child = append(group.child, component)
}
func (group *MenuGroup) Remove(id int) {
	group.child = append(group.child[:id], group.child[id+1:]...)
}
func (group *MenuGroup) Find(id int) MenuComponent {
	return group.child[id]
}

// 接下来实现Menu struct, Menu包含共同部分MenuDesc以及列表, 将列表部分分离出来
// Menu结构体
type Menu struct {
	MenuDesc
	MenuGroup
}

// 简单工厂
func NewMenu(name, description string) *Menu {
	return &Menu{
		MenuDesc: MenuDesc{
			name:        name,
			description: description,
		},
	}
}

// 实现Price和Print方法
func (m *Menu) Price() (price float32) {
	for _, v := range m.child {
		price += v.Price()
	}
	return price
}

func (m *Menu) Print() {
	fmt.Printf("%s, %s, ¥%.2f\n", m.name, m.description, m.Price())
	fmt.Println("------------------------")
	for _, v := range m.child {
		v.Print()
	}
	fmt.Println("结束")
}

// MenuItem` 和 `Menu` 都嵌入了 `MenuDesc`，这使得它们可以直接复用 `Name()` 和 `Description()` 方法，避免了代码重复。
// `Menu` 嵌入了 `MenuGroup`，这使得 `Menu` 直接获得了 `Add()`, `Remove()`, `Find()` 这些管理子节点的能力，而无需编写 `m.group.Add()`
