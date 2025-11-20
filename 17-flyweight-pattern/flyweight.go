package flyweight

import "fmt"

/* ============== 理论 ============== */
// 享元模式是一种结构型设计模式，它通过共享对象来减少内存使用和提高性能。
// 享元模式的核心思想是将对象的内部状态和外部状态分离，内部状态可以被共享，外部状态则需要根据具体情况进行处理。

// 享元接口，所有方法接收外部状态作为参数
type Shape interface {
	Draw(x, y, radius int)
}

// 具体享元，只包含内部共享状态
type Circle struct {
	color string
}

func (c *Circle) Draw(x, y, radius int) {
	fmt.Printf("Drawing a %s circle at (%d, %d) with radius %d\n", c.color, x, y, radius)
}

// 享元工厂
type ShapeFactory struct {
	circleMap map[string]*Circle
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{
		circleMap: make(map[string]*Circle),
	}
}
func (s *ShapeFactory) GetCircle(color string) Shape {
	// 先在池中查找
	if circle, ok := s.circleMap[color]; ok {
		return circle
	}
	// 找不到就创建新的并加入到池中
	circle := &Circle{color: color}
	s.circleMap[color] = circle
	return circle
}

/* ============== 下方是原仓库的例子，但是有些不足 ============== */
// 使用color来共享对象，但是 Radius 也存在于享元接口中，
// 如果对已存在color的对象修改 Radius，会影响到原来的对象，这与享元违背
// 所以需要把 Radius 从享元接口中移除，改为外部状态，这样可以避免修改共享对象影响到其他对象

/*
	享元模式核心是创建一个map属性的结构体
	设计思想:
		1. 创建Shape接口
		2. 创建实现接口Shape的实体struct Circle
		3. 创建ShapeFactory, 属性为Circle的map
*/

// // 享元接口
// type Shape interface {
// 	SetRadius(radius int)
// 	SetColor(color string)
// }

// // 具体共享对象circle,实现shape接口
// type Circle struct {
// 	color  string
// 	radius int
// }

// func (c *Circle) SetRadius(radius int) {
// 	c.radius = radius
// }

// func (c *Circle) SetColor(color string) {
// 	c.color = color
// }

// // 创建享元工厂
// type ShapeFactory struct {
// 	circleMap map[string]Shape
// }

// // GetCircle 对象不存在则创建
// func (sh *ShapeFactory) GetCircle(color string) Shape {
// 	if sh.circleMap == nil {
// 		sh.circleMap = make(map[string]Shape)
// 	}
// 	// 先在池中查找
// 	if shape, ok := sh.circleMap[color]; ok {
// 		return shape
// 	}
// 	// 没找到就创建新的，并放入池中
// 	circle := new(Circle)
// 	circle.SetColor(color)
// 	sh.circleMap[color] = circle
// 	return circle
// }
