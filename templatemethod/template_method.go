package templatemethod

// Template Method Pattern Definition:
// Define the skeleton of an algorithm in an operation,
// deferring some steps to subclasses.
// Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.
type Animal interface {
	openDoor()
	enter()
	closeDoor()
}

type Template struct {
	Animal
}

func (t *Template) Run() {
	t.openDoor()
	t.enter()
	t.closeDoor()
}
