package bridge

import "fmt"

// Bridge Pattern Definition:
// Decouple an abstraction from its implementation so that the two can vary independently.

type Instance interface {
	Operation()
}

type InstanceA struct {}

func (a *InstanceA)Operation() {
	fmt.Println("instance A operation")
}

func NewInstanceA() Instance {
	return &InstanceA{}
}

type InstanceB struct {}

func (b *InstanceB) Operation() {
	fmt.Println("instance B operation")
}

func NewInstanceB() Instance {
	return &InstanceB{}
}
