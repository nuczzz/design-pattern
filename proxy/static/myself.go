package static

type Myself struct {
	friend Proxy
	Name   string
}

func (m *Myself) SendMsg(msg string) {
	m.friend.SendMsg2Girl(msg)
}

func NewMyself(name string, friend Proxy) *Myself {
	return &Myself{
		friend: friend,
		Name:   name,
	}
}
