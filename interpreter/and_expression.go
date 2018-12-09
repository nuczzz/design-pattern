package interpreter

type AndExpression struct {
	expr1 Expression
	expr2 Expression
}

func (ae *AndExpression) Interpret(data string) bool {
	return ae.expr1.Interpret(data) && ae.expr2.Interpret(data)
}

func NewAndExpression(expr1, expr2 Expression) Expression {
	return &AndExpression{
		expr1: expr1,
		expr2: expr2,
	}
}
