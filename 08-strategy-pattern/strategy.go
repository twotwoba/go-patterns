package strategy

/* ============== 理论 ============== */
// 策略模式的侧重点是将对象和具体行为解耦，让上下文可以灵活的选择不同的策略
//
// 1. 解耦就离不开接口，需要定义一个策略的接口
// 2. 多种不同的策略需要实现该接口
// 3. 上下文对象内部有组合策略接口，并且有一个具体执行的方法

type Operator interface {
	Apply(int, int) int
}

// 策略 1，实现策略接口
type Addition struct{}

func (add *Addition) Apply(left, right int) int {
	return left + right
}

// 策略2，也实现策略接口
type Multiplication struct{}

func (mu *Multiplication) Apply(left, right int) int {
	return left * right
}

// 策略n ....

// 上下文需要包装上接口
type Operation struct {
	operator Operator
}

func NewOperation(operator Operator) Operation {
	return Operation{operator}
}

// 上下文执行由内部的‘影子替身’ operator 确定具体策略
func (op *Operation) Operate(left, right int) int {
	return op.operator.Apply(left, right)
}
