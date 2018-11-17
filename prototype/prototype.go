package prototype

import "fmt"

// Prototype Pattern Definition:
// Specify the kinds of objects to create using a prototypical instance,
// and create new objects by copying this prototype.

type Prototype interface {
	SetName(string)
	SetHeight(int)
	SetWeight(int)
	Print()
	Clone() Prototype
}

type Boy struct {
	Name string
	Height int
	Weight int
}

func (b *Boy)SetName(name string) {
	b.Name = name
}

func (b *Boy)SetHeight(height int) {
	b.Height = height
}

func (b *Boy)SetWeight(weight int) {
	b.Weight = weight
}

func (b *Boy)Print(){
	fmt.Printf("name: %v\n", b.Name)
	fmt.Printf("height: %v\n", b.Height)
	fmt.Printf("weight: %v\n", b.Weight)
	fmt.Println()
}

func (b *Boy)Clone() Prototype {
	p := *b
	return &p
}
