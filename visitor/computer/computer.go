package computer

type Computer struct {
	keyboard Constituent
	monitor  Constituent
	mouse    Constituent
}

func NewComputer() *Computer {
	return &Computer{
		keyboard: NewKeyboard("keyboard-1"),
		monitor:  NewMonitor("monitor-2"),
		mouse:    NewMouse("mouse-3"),
	}
}

func (c *Computer) Accept(visitor Visitor) {
	c.keyboard.Accept(visitor)
	c.monitor.Accept(visitor)
	c.mouse.Accept(visitor)
}
