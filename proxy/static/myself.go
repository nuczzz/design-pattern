package static

import "fmt"

type Myself struct {
	Name string
}

func (m *Myself) ReadMsg(msg string) {
	fmt.Println(m.Name + " read " + msg)
}

func (m *Myself) SendMsg() string {
	return "I'm " + m.Name
}
