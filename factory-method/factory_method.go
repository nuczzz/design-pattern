package factory_method

// Factory Method Pattern Definition:
// Defines an interface for creating objects,
// but let subclasses to decide which class to instantiate.
// Refers the newly created object through a common interface.
type OperationInterface interface {
	SetValue(int, int)
	GetResult() int
}
