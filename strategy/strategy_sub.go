package strategy

type OperationSub struct {
}

func (o *OperationSub) Operation(i, j int) int {
	return i - j
}

func NewSubOperation() Strategy {
	return &OperationSub{}
}
