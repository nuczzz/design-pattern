package decorator

import "fmt"

// Decorator Pattern Definition:
// Attach additional responsibilities to an object dynamically keeping the same interface.
// Decorators provide a flexible alternative to subclassing for extending functionality.

type Human interface {
	Wear()
	Walk()
}

type HumanInstance struct {
}

func (h *HumanInstance) Wear() {
	fmt.Println("human Wear...")
}

func (h *HumanInstance) Walk() {
	fmt.Println("human Walk...")
}

func NewHumanInstance() Human {
	return &HumanInstance{}
}

type BaseDecorator struct {
	human Human
}

func (b *BaseDecorator) Wear() {
	b.human.Wear()
	fmt.Println("BaseDecorator Wear...")
}

func (b *BaseDecorator) Walk() {
	b.human.Wear()
	fmt.Println("BaseDecorator Walk...")
}

func NewBaseDecorator(human Human) Human {
	return &BaseDecorator{
		human: human,
	}
}

type InnerDecorator struct {
	baseDecorator Human
}

func (s *InnerDecorator) Wear() {
	s.baseDecorator.Wear()
	fmt.Println("InnerDecorator Wear...")
}

func (s *InnerDecorator) Walk() {
	s.baseDecorator.Wear()
	fmt.Println("InnerDecorator Walk...")
}

func NewInnerDecorator(human Human) Human {
	return &InnerDecorator{
		baseDecorator: human,
	}
}

type OuterDecorator struct {
	innerDecorator Human
}

func (o *OuterDecorator) Wear() {
	o.innerDecorator.Wear()
	fmt.Println("OuterDecorator Wear...")
}

func (o *OuterDecorator) Walk() {
	o.innerDecorator.Wear()
	fmt.Println("OuterDecorator Walk...")
}

func NewOuterDecorator(human Human) Human {
	return &OuterDecorator{
		innerDecorator: human,
	}
}
