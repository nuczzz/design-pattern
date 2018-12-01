package composite

// Composite Pattern Definition:
// Compose objects into tree structures to represent part-whole hierarchies.
// Composite lets clients treat individual objects and compositions of objects uniformly.

type Composite interface {
	Name() string
	Description() string
	Price() float64
	Print()
	Add(Composite)
	Remove(int)
	Child(int) Composite
}

