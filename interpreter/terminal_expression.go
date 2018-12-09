package interpreter

import "strings"

type BaseExpression struct {
	data string
}

func (te *BaseExpression) Interpret(data string) bool {
	return strings.Contains(data, te.data)
}

func NewBaseExpression(data string) Expression {
	return &BaseExpression{data: data}
}
