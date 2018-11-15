package static

import "fmt"

type Girl struct {
	Name string
}

func (g *Girl) ReadMsg(msg string) {
	fmt.Println(g.Name + " read " + msg)
}

func (g *Girl) SendMsg() string {
	return "I'm " + g.Name
}
