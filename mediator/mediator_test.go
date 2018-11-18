package mediator

import "testing"

func TestChatRoomSendMessage(t *testing.T) {
	room := NewChatRoom()
	boy := NewUser("boy")
	girl := NewUser("girl")
	room.SendMessage(boy, girl, "I love you!")
	room.SendMessage(girl, boy, "But I don't love you...")
	room.SendMessage(boy, girl, "...")
}
