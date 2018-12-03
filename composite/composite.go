package composite

import "fmt"

// Composite Pattern Definition:
// Compose objects into tree structures to represent part-whole hierarchies.
// Composite lets clients treat individual objects and compositions of objects uniformly.

type Composite interface {
	Add(composite Composite)
	Remove(composite Composite)
	GetEList() []Composite
	Print() string
}

type Employee struct {
	name   string
	dept   string
	salary float64
	eList  []Composite
}

func (e *Employee) Add(composite Composite) {
	e.eList = append(e.eList, composite)
}

func (e *Employee) Remove(composite Composite) {
	for i := range e.eList {
		if e.eList[i] == composite {
			e.eList = append(e.eList[:i], e.eList[i+1:]...)
			return
		}
	}
}

func (e *Employee) GetEList() []Composite {
	return e.eList
}

func (e *Employee) Print() string {
	return fmt.Sprintf("Employee: %v, dept: %v, salary: %v\n", e.name, e.dept, e.salary)
}

func NewEmployee(name, dept string, salary float64) Composite {
	return &Employee{
		name:   name,
		dept:   dept,
		salary: salary,
	}
}
