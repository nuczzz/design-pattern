package iterator

// Iterator Pattern Definition:
// Provide a way to access the elements of an aggregate object sequentially
// without exposing its underlying representation.

type Iterator interface {
	HasNext() bool
	Next()
	Value() interface{}
}

type Array struct {
	list  []int
	index int //index of list
}

func (a *Array) HasNext() bool {
	return a.index < len(a.list)
}

func (a *Array) Next() {
	a.index++
}

func (a *Array) Value() interface{} {
	return a.list[a.index]
}

func NewListIterator(list []int) Iterator {
	return &Array{
		list: list,
		index:  0,
	}
}
