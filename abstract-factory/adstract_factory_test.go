package abstract_factory

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	factory := ChooseFactory('m')
	hamburger := factory.NewHamburger()
	cola := factory.NewColaInterface()

	t.Log(hamburger.Eat())
	t.Log(cola.Drink())
}

func ChooseFactory(t byte) AbstractFactory {
	switch t {
	case 'k':
		return NewKfcFactory()
	case 'm':
		return NewMacDonaldFactory()
	default:
		panic(fmt.Sprintf("not support type '%v'", t))
	}
}
