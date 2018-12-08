package bill

type Boss struct {
	TotalIncome float64
	TotalConsume float64
}

func (b *Boss) ViewConsume(c *ConsumeBill) {
	b.TotalConsume += c.Amount
}

func (b *Boss) ViewIncome(i *IncomeBill) {
	b.TotalIncome += i.Amount
}

func (b *Boss) GetTotalIncome() float64 {
	return b.TotalIncome
}

func (b *Boss) GetTotalConsume() float64 {
	return b.TotalConsume
}

func NewBoss() *Boss {
	return &Boss{}
}
