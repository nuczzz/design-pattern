package decorator

import (
	"testing"
	"fmt"
)

func TestDecorator(t *testing.T) {
	outer := NewOuterDecorator(NewInnerDecorator(NewBaseDecorator(NewHumanInstance())))
	outer.Wear()
	fmt.Println()
	outer.Walk()
}
