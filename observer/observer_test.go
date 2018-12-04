package observer

import (
	"testing"
	"fmt"
)

func TestObserver(t *testing.T) {
	ob1 := NewObserver("ob1")
	ob2 := NewObserver("ob2")
	ob3 := NewObserver("ob3")
	ob4 := NewObserver("ob4")
	ob5 := NewObserver("ob5")

	subject := NewSubject()
	subject.Watch(ob1)
	subject.Watch(ob2)
	subject.Watch(ob3)
	subject.Watch(ob4)
	subject.Watch(ob5)

	subject.SetState("hello")
	fmt.Println()

	subject.Remove(ob3)
	subject.SetState("world")
}
