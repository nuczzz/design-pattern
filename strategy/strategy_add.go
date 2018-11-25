package strategy

type OperationAdd struct {
}

func (o *OperationAdd)Operation(i, j int) int {
	return i+j
}

func NewAddOperation() Strategy {
	return &OperationAdd{}
}
