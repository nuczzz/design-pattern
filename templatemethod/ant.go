package templatemethod

import "fmt"

type Ant struct {
}

func (e *Ant)openDoor() {
	fmt.Println("ant open the door")
}

func (e *Ant)enter() {
	fmt.Println("ant enter room")
}

func (e *Ant)closeDoor() {
	fmt.Println("ant close the door.")
}

func NewAnt() Animal {
	return &Ant{}
}
