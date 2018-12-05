package computer

import "fmt"

type cpu struct {
	sn string
}

func (c *cpu) start() {
	fmt.Printf("[%v] cpu start...\n", c.sn)
}

func (c *cpu) stop() {
	fmt.Printf("[%v] cpu stop...\n", c.sn)
}

func newCpu(sn string) *cpu {
	return &cpu{sn: sn}
}
