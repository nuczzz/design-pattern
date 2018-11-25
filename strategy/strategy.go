package strategy

// Strategy Pattern Definition:
// Define a family of algorithms,encapsulate each one,
// and make them interchangeable.

type Strategy interface {
	Operation(i int, j int) int
}

type Concentrate struct {
	strategy Strategy
}

func NewConcentrate(s Strategy) *Concentrate {
	return &Concentrate{strategy: s}
}

func (c *Concentrate) Execute(i, j int) int {
	return c.strategy.Operation(i, j)
}
