package hook

import "fmt"

// lowercase structure and lowercase function of systemHook
type systemHook interface {
	do()
}

type ClientHook interface {
	BeforeDo()
	AfterDo()
}

type Hook struct {
	systemAction systemHook
	ClientAction ClientHook
}

func (h *Hook) Run() {
	h.ClientAction.BeforeDo()
	h.systemAction.do()
	h.ClientAction.AfterDo()
}

type systemAction struct {
}

func (s *systemAction) do() {
	fmt.Println("system prepare start...")
	fmt.Println("system prepare ...")
	fmt.Println("system prepare done")
}

func newSystemAction() systemHook {
	return &systemAction{}
}

func NewHook() *Hook {
	return &Hook{
		systemAction: newSystemAction(),
	}
}
