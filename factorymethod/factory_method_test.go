package factorymethod

import (
	"fmt"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	op1 := ChooseOperation('+')
	op1.SetValue(1, 2)
	t.Log(op1.GetResult())

	op2 := ChooseOperation('-')
	op2.SetValue(3, 2)
	t.Log(op2.GetResult())
}

func ChooseOperation(t byte) OperationInterface {
	switch t {
	case '+':
		return NewAddOperation()
	case '-':
		return NewSubOperation()
	default:
		panic(fmt.Sprintf("not support '%v'", t))
	}
}
