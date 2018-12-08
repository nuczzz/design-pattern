package computer

import "fmt"

type Visitor interface {
	VisitKeyboard(keyboard *Keyboard)
	VisitorMonitor(monitor *Monitor)
	VisitMouse(mouse *Mouse)
}

type ConstituentVisitor struct {
	name string
}

func (c *ConstituentVisitor) VisitKeyboard(keyboard *Keyboard) {
	fmt.Printf("[%v] visit %v\n", c.name, keyboard.GetInfo())
}

func (c *ConstituentVisitor) VisitorMonitor(monitor *Monitor) {
	fmt.Printf("[%v] visit %v\n", c.name, monitor.GetInfo())
}

func (c *ConstituentVisitor) VisitMouse(mouse *Mouse) {
	fmt.Printf("[%v] visit %v\n", c.name, mouse.GetInfo())
}

func NewVisitor(name string) Visitor {
	return &ConstituentVisitor{name: name}
}
