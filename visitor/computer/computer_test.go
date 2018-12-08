package computer

import "testing"

func TestVisitor(t *testing.T) {
	computer := NewComputer()
	visitor := NewVisitor("visitor-ufo")
	computer.Accept(visitor)
}