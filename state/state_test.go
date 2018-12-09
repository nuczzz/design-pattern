package state

import (
	"testing"
	"fmt"
)

func TestState(t *testing.T) {
	content := NewContent()

	ss := NewStartState()
	content.SetState(ss)
	fmt.Println(content.GetState().Info())

	es := NewEndStart()
	content.SetState(es)
	fmt.Println(content.GetState().Info())
}
