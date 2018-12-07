package visitor

// Visitor Pattern Definition:
// Represent an operation to be performed on the elements of an object structure.
// Visitor lets you define a new operation without changing the classes of the elements on which it operates.

type Bill interface {
	Accept(AccountBookView)
}

type AccountBookView interface {
	ViewConsume(*ConsumeBill)
	ViewIncome(*IncomeBill)
}

type ConsumeBill struct {
	Amount float64
	Item   string
}

func (c *ConsumeBill) Accept(a AccountBookView) {
	a.ViewConsume(c)
}

func NewConsume(item string, amount float64) *ConsumeBill {
	return &ConsumeBill{
		Item:   item,
		Amount: amount,
	}
}

type IncomeBill struct {
	Amount float64
	Item   string
}

func (i *IncomeBill) Accept(a AccountBookView) {
	a.ViewIncome(i)
}

func NewIncome(item string, amount float64) *IncomeBill {
	return &IncomeBill{
		Item:   item,
		Amount: amount,
	}
}
