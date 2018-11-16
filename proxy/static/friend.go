package static

type Friend struct {
	girl *Girl
}

func (f *Friend)SendMsg2Girl(msg string) {
	msg += " -> friend_proxy -> "
	f.girl.ReadMsg(msg)
}



func NewFriendProxy(girl *Girl) Proxy {
	return &Friend{
		girl: girl,
	}
}
