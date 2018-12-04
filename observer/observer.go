package observer

import "fmt"

// Observer Pattern Definition:
// Define a one-to-many dependency between objects so that when one object changes state,
// all its dependents are notified and updated automatically.

type Observer interface {
	Notify(string)
}

type Instance struct {
	name string
}

func (i *Instance) Notify(msg string) {
	fmt.Printf("[%v]<-[%v]\n", i.name, msg)
}

func NewObserver(name string) Observer {
	return &Instance{
		name: name,
	}
}
