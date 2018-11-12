package hook

type ClientAction struct {
}

func (c *ClientAction)BeforeDo(){
}

func (c *ClientAction)AfterDo(){
}

func NewClientAction(before, after func()) ClientHook {
	c := &ClientAction{}
	c.BeforeDo = before
	c.AfterDo = after
	return c
}
