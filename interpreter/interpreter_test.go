package interpreter

import (
	"testing"
	"fmt"
)

func TestInterpreter(t *testing.T) {
	john := NewBaseExpression("john")
	tall := NewBaseExpression("tall")
	robert := NewBaseExpression("robert")
	cute := NewBaseExpression("cute")

	fmt.Println("john is tall? ", NewAndExpression(john, tall).Interpret("john is tall?"))
	fmt.Println("robert is ugly? ", NewOrExpression(robert, cute).Interpret("ugly"))
}
