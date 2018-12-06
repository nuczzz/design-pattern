package memento

import (
	"testing"
	"fmt"
)

func TestMemento(t *testing.T) {
	op := NewOperation()
	op.ShowLastWord()
	fmt.Println()

	op.WriteWord(1)
	op.ShowLastWord()
	fmt.Println()

	op.WriteWord(1)
	op.WriteWord(2)
	op.WriteWord(3)
	op.WriteWord(4)
	op.ShowLastWord()
	op.Withdraw()
	op.ShowLastWord()
	op.Withdraw()
	op.ShowLastWord()
	op.Withdraw()
	op.ShowLastWord()
	op.Withdraw()
	op.ShowLastWord()
}
