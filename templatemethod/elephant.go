package templatemethod

import "fmt"

type Elephant struct {
}

func (e *Elephant)OpenDoor() {
	fmt.Println("elephant open the door")
}

func (e *Elephant)Enter() {
	fmt.Println("elephant enter room")
}

func (e *Elephant)CloseDoor() {
	fmt.Println("elephant close the door, oh no...the door crash.")
}

func NewElephant() Animal {
	return &Elephant{}
}
