package builder

// Builder Pattern Definition:
// Separate the construction of a complex object from its representation
// so that the same construction process can create different representations.
type Human struct {
	Head string
	Body string
	Hand string
	Foot string
}

type HumanBuilder interface {
	BuildHead()
	BuildBody()
	BuildHand()
	BuildFoot()
	NewHuman() Human
	PrintHuman()
}
