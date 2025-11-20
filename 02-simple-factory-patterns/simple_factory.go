package factory

import "errors"

/* ============== 理论  ============== */
// 简单工厂模式，不属于23中模式之一，比较简单，
// 就是把对象的创建逻辑集中到一个方法中，根据传入参数决定创建什么
//
// 三要素
// 1. Procuct, 接口，抽象产品行为
// 2. Concrete Product, 多个具体产品，实现了 Product 的接口
// 3. Factory 函数，返回接口，根据 kind 创造出不同的 Concrete Product
//

type Kind int

const (
	Cash Kind = 1 << iota
	Credit
)

// 产品接口
type Payment interface {
	Pay(money float32) error
}

// 产品1,实现产品接口
type CashPay struct {
	Balance float32
}

func (cash *CashPay) Pay(money float32) error {
	if cash.Balance < 0 || cash.Balance < money {
		return errors.New("balance not enough")
	}
	cash.Balance -= money
	return nil
}

// 产品2,实现产品接口
type CreditPay struct {
	Balance float32
}

func (credit *CreditPay) Pay(money float32) error {
	if credit.Balance < 0 || credit.Balance < money {
		return errors.New("balance not enough")
	}
	credit.Balance -= money
	return nil
}

// 工厂函数
func GeneratePayment(k Kind, balance float32) (Payment, error) {
	switch k {
	case Cash:
		cash := new(CashPay)
		cash.Balance = balance
		return cash, nil
	case Credit:
		return &CreditPay{balance}, nil
	default:
		return nil, errors.New("Payment do not support this ")
	}
}
