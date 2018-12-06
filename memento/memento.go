package memento

// Memento Pattern Definition:
// Without violating encapsulation,
// capture and externalize an object's internal state
// so that the object can be restored to this state later.

type Memento interface {
	PopLastState() int
	LastState() int
	AddState(int)
}

const Size = 3

type Stack struct {
	cur   int
	stack []int
}

func (s *Stack)PopLastState() int {
	if s.cur < 0 || s.cur >= Size {
		return -1
	}
	ret := s.stack[s.cur]
	s.cur--
	return ret
}

func (s *Stack) LastState() int {
	if s.cur < 0 || s.cur >= Size {
		return -1
	}
	return s.stack[s.cur]
}

func (s *Stack) AddState(data int) {
	if s.cur < Size-1 {
		s.cur++
		s.stack[s.cur] = data
		return
	}
	s.stack = append(s.stack[1:], data)
}

func NewMemento() Memento {
	return &Stack{
		cur:   -1,
		stack: make([]int, Size),
	}
}
