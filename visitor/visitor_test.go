package visitor

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	accountBook := NewAccountBook()
	accountBook.AddBill(NewConsume("item1", 10))
	accountBook.AddBill(NewConsume("item2", 12))
	accountBook.AddBill(NewIncome("item3", 12))
	accountBook.AddBill(NewIncome("item4", 12))

	boss := NewBoss()
	accountBook.Show(boss)
	fmt.Println(boss.GetTotalConsume())
	fmt.Println(boss.TotalIncome)

	accountant := NewAccountant()
	accountBook.Show(accountant)
}
