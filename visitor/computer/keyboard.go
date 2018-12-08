package computer

import "fmt"

type Keyboard struct {
	sn string
}

func (k *Keyboard) Accept(visitor Visitor) {
	visitor.VisitKeyboard(k)
}

func (k *Keyboard) GetInfo() string {
	return fmt.Sprintf("keyboard[%v]", k.sn)
}

func NewKeyboard(sn string) Constituent {
	return &Keyboard{sn: sn}
}
