package computer

import "fmt"

type Monitor struct {
	sn string
}

func (m *Monitor) Accept(visitor Visitor) {
	visitor.VisitorMonitor(m)
}

func (m *Monitor) GetInfo() string {
	return fmt.Sprintf("monitor[%v]", m.sn)
}

func NewMonitor(sn string) Constituent {
	return &Monitor{sn: sn}
}
