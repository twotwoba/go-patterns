package flyweight

import (
	"fmt"
	"testing"
)

func TestShapeFactory_GetCircle(t *testing.T) {
	factory := NewShapeFactory()
	redCircle := factory.GetCircle("red")
	redCircle.Draw(10, 20, 5)
	redCircle.Draw(50, 60, 8)
	factory.GetCircle("blue")
	if _, ok := factory.circleMap["red"]; !ok {
		t.Error("map为空， 期待为1")
	}
	circle := redCircle.(*Circle)
	fmt.Println(circle.color)
	if circle.color != "red" {
		t.Error("expected color is red")
	}
}
