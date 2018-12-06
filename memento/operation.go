package memento

import "fmt"

type Operation struct {
	m Memento
}

func (o *Operation) WriteWord(w int) {
	fmt.Printf("write word %v on paper\n", w)
	o.m.AddState(w)
}

func (o *Operation) Withdraw() {
	fmt.Println("[ctrl+z] withdraw word")
	o.m.PopLastState()
}

func (o *Operation) ShowLastWord() {
	last := o.m.LastState()
	if last == -1 {
		fmt.Println("no word in memento")
		return
	}

	fmt.Printf("last word is: %v\n", last)
}

func NewOperation() *Operation {
	return &Operation{
		m: NewMemento(),
	}
}
