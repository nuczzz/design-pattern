package computer

type Constituent interface {
	Accept(Visitor)
	GetInfo() string
}
