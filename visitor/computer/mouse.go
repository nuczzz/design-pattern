package computer

import "fmt"

type Mouse struct {
	sn string
}

func (m *Mouse) Accept(visitor Visitor) {
	visitor.VisitMouse(m)
}

func (m *Mouse) GetInfo() string {
	return fmt.Sprintf("mouse[%v]", m.sn)
}

func NewMouse(sn string) Constituent {
	return &Mouse{sn: sn}
}
