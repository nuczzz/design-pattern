package strategy

import "testing"

func TestStrategy(t *testing.T) {
	opAdd := NewConCentrate(NewAddOperation())
	resultAdd := opAdd.Execute(1, 2)
	t.Log(resultAdd)

	opSub := NewConCentrate(NewSubOperation())
	resultSub := opSub.Execute(1, 2)
	t.Log(resultSub)
}
