package factorymethod

// sub operation
type SubOperation struct {
	a int
	b int
}

func (s *SubOperation) SetValue(a, b int) {
	s.a = a
	s.b = b
}

func (s *SubOperation) GetResult() int {
	return s.a - s.b
}

func NewSubOperation() OperationInterface {
	return &SubOperation{}
}
