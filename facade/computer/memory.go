package computer

import "fmt"

type memory struct {
	sn string
}

func (m *memory) start() {
	fmt.Printf("[%v] memory start...\n", m.sn)
}

func (m *memory) stop() {
	fmt.Printf("[%v] memory stop...\n", m.sn)
}

func newMemory(sn string) *memory {
	return &memory{sn: sn}
}
