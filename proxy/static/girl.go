package static

import "fmt"

type Girl struct {
	Name string
}

func (g *Girl) ReadMsg(msg string) {
	fmt.Println(msg + g.Name)
}

func NewGirl(name string) *Girl {
	return &Girl{
		Name: name,
	}
}
