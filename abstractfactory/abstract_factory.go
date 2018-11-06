package abstractfactory

// Abstract Factory pattern
// Provide an interface for creating families of related
// or dependent objects without specifying their concrete classes.

type HamburgerInterface interface {
	Eat() string
}

type ColaInterface interface {
	Drink() string
}

type AbstractFactory interface {
	NewHamburger() HamburgerInterface
	NewColaInterface() ColaInterface
}
