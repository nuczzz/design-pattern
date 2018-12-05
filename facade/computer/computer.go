package computer

type Computer struct {
	cpu    *cpu
	disk   *disk
	memory *memory
}

func (c *Computer) Start() {
	c.cpu.start()
	c.disk.start()
	c.memory.start()
}

func (c *Computer) Stop() {
	c.memory.stop()
	c.disk.stop()
	c.cpu.stop()
}

func NewComputer(cpuSn, diskSn, memorySn string) *Computer {
	return &Computer{
		cpu:    newCpu(cpuSn),
		disk:   newDisk(diskSn),
		memory: newMemory(memorySn),
	}
}
