package state

type Content struct {
	state State
}

func (c *Content) SetState(state State) {
	c.state = state
}

func (c *Content) GetState() State {
	return c.state
}

func NewContent() *Content {
	return &Content{}
}
