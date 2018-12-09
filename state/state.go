package state

// State Pattern Definition:
// Allow an object to alter its behavior when its internal state changes.
// The object will appear to change its class.

type State interface {
	Do(content *Content)
	Info() string
}

type StartState struct {}

func (ss *StartState) Do(content *Content) {
	content.SetState(ss)
}

func (ss *StartState) Info() string {
	return "start state"
}

func NewStartState() State {
	return &StartState{}
}

type EndState struct {}

func (es *EndState) Do(content *Content) {
	content.SetState(es)
}

func (es *EndState) Info() string {
	return "end state"
}

func NewEndStart() State {
	return &EndState{}
}
