package composite

import (
	"testing"
	"fmt"
)

func printAll(composite Composite) {
	for _, e := range composite.GetEList() {
		fmt.Print(e.Print())
		printAll(e)
	}
}

func TestComposite(t *testing.T) {
	ceo := NewEmployee("zhao", "ceo", 1000)
	cto := NewEmployee("qian", "cto", 1000)
	dev1 := NewEmployee("dev1", "dev", 100)
	dev2 := NewEmployee("dev2", "dev", 100)
	stu1 := NewEmployee("stu1", "stu", 50)
	stu2 := NewEmployee("stu2", "stu", 50)

	ceo.Add(cto)
	cto.Add(dev1)
	cto.Add(dev2)
	dev1.Add(stu1)
	dev2.Add(stu2)

	fmt.Print(ceo.Print())
	printAll(ceo)
}
