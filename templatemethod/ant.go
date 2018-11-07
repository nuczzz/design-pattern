package templatemethod

import "fmt"

type Ant struct {
}

func (e *Ant)OpenDoor() {
	fmt.Println("ant open the door")
}

func (e *Ant)Enter() {
	fmt.Println("ant enter room")
}

func (e *Ant)CloseDoor() {
	fmt.Println("ant close the door.")
}

func NewAnt() Animal {
	return &Ant{}
}
