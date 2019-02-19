package abstract_factory

import "fmt"

type macDonaldHamburger struct {
	HamburgerInfo string
}

func (kh *macDonaldHamburger) Eat() string {
	return fmt.Sprintf("Eat: %v", kh.HamburgerInfo)
}

type macDonaldCola struct {
	ColaInfo string
}

func (kc *macDonaldCola) Drink() string {
	return fmt.Sprintf("Drink: %v", kc.ColaInfo)
}

type macDonald struct {
}

func (k *macDonald) NewHamburger() HamburgerInterface {
	return &macDonaldHamburger{
		HamburgerInfo: "MacDonald hamburger",
	}
}

func (k *macDonald) NewColaInterface() ColaInterface {
	return &macDonaldCola{
		ColaInfo: "MacDonald cola",
	}
}

func NewMacDonaldFactory() AbstractFactory {
	return &macDonald{}
}
