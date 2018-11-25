package strategy

import "testing"

func TestStrategy(t *testing.T) {
	opAdd := NewConcentrate(NewAddOperation())
	resultAdd := opAdd.Execute(1, 2)
	t.Log(resultAdd)

	opSub := NewConcentrate(NewSubOperation())
	resultSub := opSub.Execute(1, 2)
	t.Log(resultSub)
}
