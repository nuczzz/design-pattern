package bill

import "fmt"

type Accountant struct {
}

func (a *Accountant) ViewConsume(c *ConsumeBill) {
	if c.Item == "item1" {
		fmt.Println("accountant check consume of item1")
		return
	}
	fmt.Println("accountant check pass ", c.Item)
}

func (a *Accountant) ViewIncome(i *IncomeBill) {
	fmt.Println("accountant calculate income tax of ", i.Item)
}

func NewAccountant() *Accountant{
	return &Accountant{}
}
