package interpreter

import (
	"testing"
	"fmt"
)

func TestInterpreter(t *testing.T) {
	john := NewBaseExpression("john")
	tall := NewBaseExpression("tall")
	nobody := NewBaseExpression("nobody")
	nameless := NewBaseExpression("nameless")

	fmt.Println("john is tall? ", NewAndExpression(john, tall).Interpret("john is tall?"))
	fmt.Println("john or his brother is here? ", NewOrExpression(nobody, nameless).Interpret("john or his brother"))
}
