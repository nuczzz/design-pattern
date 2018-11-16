package static

import "testing"

func TestProxy(t *testing.T) {
	girl := NewGirl("juice")
	friend := NewFriendProxy(girl)
	myself := NewMyself("li ming", friend)
	myself.SendMsg("[hello]")
}
