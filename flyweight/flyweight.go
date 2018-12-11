package flyweight

import (
	"fmt"
)

// Flyweight Pattern Definition:
// Use sharing to support large numbers of fine-grained objects efficiently.

type Element struct {
	key   string
	color string
}

var m = make(map[string]*Element)

func (e *Element)GetColor() string {
	return e.color
}

func NewElement(key string, color string) *Element {
	if v, ok := m[key]; ok {
		fmt.Printf("element [key: %v] has exist\n", key)
		v.color = color
		return v
	}

	fmt.Printf("create new element [key: %v, color: %v]\n", key, color)
	e := &Element{
		key:   key,
		color: color,
	}
	m[key] = e
	return e
}
