package abstractfactory

import "fmt"

type kfcHamburger struct {
	HamburgerInfo string
}

func (kh *kfcHamburger) Eat() string {
	return fmt.Sprintf("Eat: %v", kh.HamburgerInfo)
}

type kfcCola struct {
	ColaInfo string
}

func (kc *kfcCola) Drink() string {
	return fmt.Sprintf("Drink: %v", kc.ColaInfo)
}

type kfc struct {
}

func (k *kfc) NewHamburger() HamburgerInterface {
	return &kfcHamburger{
		HamburgerInfo: "kfc hamburger",
	}
}

func (k *kfc) NewColaInterface() ColaInterface {
	return &kfcCola{
		ColaInfo: "kfc cola",
	}
}

func NewKfcFactory() AbstractFactory {
	return &kfc{}
}
