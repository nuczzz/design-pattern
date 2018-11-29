package iterator

import (
	"testing"
	"fmt"
)

func TestIterator(t *testing.T) {
	list := []int{1, 2, 3, 4}
	iterator := NewListIterator(list)
	for iterator.HasNext() {
		fmt.Println(iterator.Value())
		iterator.Next()
	}
}
