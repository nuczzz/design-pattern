package composite

import "fmt"

type MenuItem struct {
	name  string
	desc  string
	price float64
}

func (m *MenuItem) Name() string {
	return m.name
}

func (m *MenuItem) Description() string {
	return m.desc
}

func (m *MenuItem) Price() float64 {
	return m.price
}

func (m *MenuItem) Print() {
	fmt.Printf("  %s, ￥%.2f\n", m.name, m.price)
	fmt.Printf("    -- %s\n", m.desc)
}

func (m *MenuItem) Add(c Composite) {
	return
}

func (m *MenuItem) Remove(pos int) {
	return
}

func (m *MenuItem) Child(pos int) Composite {
	return nil
}

func NewMenu(name, desc string, price float64) Composite{
	return &MenuItem{
		name: name,
		desc: desc,
		price: price,
	}
}

