package abstractfactory

import (
	"fmt"
)

/* ============== 理论 ============== */
// 这里的例子更贴近工厂方法模式，它更关注--对象是由谁创建
// 由不同的具体工厂决定
//
// 简单工厂关注的是创建什么，由参数决定
// 抽象工厂关注的是创建一整套，确保同一个工厂出来的对象互相兼容
//
/*
设计思想
 1. 抽象产品接口
 2. 抽象工厂接口
 3. 具体的工厂和产品struct
 4. 使用具体的工厂来创建产品，并返回接口类型值
*/

type Product interface {
	Describe()
}
type ConcreteProduct struct {
	Name string
}

func (conproduct *ConcreteProduct) Describe() {
	fmt.Println(conproduct.Name)
}

type Factory interface {
	CreateProduct() Product
}
type ConCreteFactory1 struct{}

func (confactory *ConCreteFactory1) CreateProduct1() Product {
	return &ConcreteProduct{Name: "KG"}
}

type ConCreteFactory2 struct{}

func (confactory *ConCreteFactory2) CreateProduct2() Product {
	return &ConcreteProduct{Name: "KG2"}
}

/* ============== 抽象工厂例子 ============== */
// 1. 抽象产品接口 (有两种产品：按钮和复选框)
type Button interface{ Render() }
type Checkbox interface{ Paint() }

// 2. 抽象工厂接口 (能创建一整套产品)
type GUIFactory interface {
	// ⚠️注意，多个创建方法
	CreateButton() Button
	CreateCheckbox() Checkbox
}

// --- Light 主题产品族 ---
type LightButton struct{}

func (b *LightButton) Render() { fmt.Println("Rendering Light Button") }

type LightCheckbox struct{}

func (c *LightCheckbox) Paint() { fmt.Println("Painting Light Checkbox") }

// 3. Light 主题的具体工厂
type LightThemeFactory struct{}

func (f *LightThemeFactory) CreateButton() Button     { return &LightButton{} }
func (f *LightThemeFactory) CreateCheckbox() Checkbox { return &LightCheckbox{} }

// --- Dark 主题产品族 ---
type DarkButton struct{}

func (b *DarkButton) Render() { fmt.Println("Rendering Dark Button") }

type DarkCheckbox struct{}

func (c *DarkCheckbox) Paint() { fmt.Println("Painting Dark Checkbox") }

// 4. Dark 主题的具体工厂
type DarkThemeFactory struct{}

func (f *DarkThemeFactory) CreateButton() Button     { return &DarkButton{} }
func (f *DarkThemeFactory) CreateCheckbox() Checkbox { return &DarkCheckbox{} }
