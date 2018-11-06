package factorymethod

import "testing"

func TestFactoryMethod(t *testing.T) {
	add := NewAddOperation()
	sub := NewSubOperation()

	add.SetValue(1, 2)
	sub.SetValue(2, 1)

	t.Log(add.GetResult())
	t.Log(sub.GetResult())
}
