package hook

import (
	"fmt"
	"testing"
)

type ClientAction struct {
}

func (c *ClientAction) BeforeDo() {
	fmt.Println("client do something before system prepare...")
}

func (c *ClientAction) AfterDo() {
	fmt.Println("client do something after system prepare...")
}

func TestHookRun(t *testing.T) {
	hook := NewHook()
	hook.ClientAction = &ClientAction{}
	hook.Run()
}
