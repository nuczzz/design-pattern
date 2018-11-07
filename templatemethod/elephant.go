package templatemethod

import "fmt"

type Elephant struct {
}

func (e *Elephant)openDoor() {
	fmt.Println("elephant open the door")
}

func (e *Elephant)enter() {
	fmt.Println("elephant enter room")
}

func (e *Elephant)closeDoor() {
	fmt.Println("elephant close the door, oh no...the door crash.")
}

func NewElephant() Animal {
	return &Elephant{}
}
