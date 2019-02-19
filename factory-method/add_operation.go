package factory_method

// add operation
type AddOperation struct {
	a int
	b int
}

func (a *AddOperation) SetValue(v1, v2 int) {
	a.a = v1
	a.b = v2
}

func (a *AddOperation) GetResult() int {
	return a.a + a.b
}

func NewAddOperation() OperationInterface {
	return &AddOperation{}
}
