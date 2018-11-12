package builder

import (
	"testing"
	"fmt"
)

type Student Human

func (s *Student) BuildHead() {
	s.Head = "student's head"
}

func (s *Student) BuildBody() {
	s.Body = "student's body"
}

func (s *Student) BuildHand() {
	s.Hand = "student's hand"
}

func (s *Student) BuildFoot() {
	s.Foot = "student's foot"
}

func (s *Student) NewHuman() Human {
	return Human{}
}

func (s *Student) PrintHuman() {
	fmt.Println(s.Head)
	fmt.Println(s.Body)
	fmt.Println(s.Hand)
	fmt.Println(s.Foot)
}

func TestBuilder(t *testing.T) {
	st := &Student{}
	director := NewHumanDirector()
	director.CreateHuman(st)
	st.PrintHuman()
}
