package observer

type Center interface {
	SetState(string)
	Watch(observer Observer)
	Remove(observer Observer)
}

type Subject struct {
	state string
	list []Observer
}

func (s *Subject) SetState(state string) {
	s.state = state
	s.notifyAllObserver(state)
}

func (s *Subject) Watch(observer Observer) {
	s.list = append(s.list, observer)
}

func (s *Subject) Remove(observer Observer) {
	for i := range s.list {
		if s.list[i] == observer {
			s.list = append(s.list[:i], s.list[i+1:]...)
			return
		}
	}
}

func (s *Subject)notifyAllObserver(msg string) {
	for i := range s.list {
		s.list[i].Notify(msg)
	}
}

func NewSubject() Center {
	return &Subject{
		list: make([]Observer, 0),
	}
}
