package interpreter

type OrExpression struct {
	expr1 Expression
	expr2 Expression
}

func (oe *OrExpression) Interpret(data string) bool {
	return oe.expr1.Interpret(data) || oe.expr2.Interpret(data)
}

func NewOrExpression(expr1, expr2 Expression) Expression {
	return &OrExpression{
		expr1: expr1,
		expr2: expr2,
	}
}
