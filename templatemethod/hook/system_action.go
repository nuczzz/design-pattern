package hook

import "fmt"

type SystemAction struct {
}

func (s *SystemAction) Do() {
	fmt.Println("system prepare start...")
	fmt.Println("system prepare ...")
	fmt.Println("system prepare done")
}

func NewSystemAction() SystemHook {
	return &SystemAction{}
}
