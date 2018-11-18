package mediator

import (
	"fmt"
)

// Mediator Pattern Definition:
// Define an object that encapsulates how a set of objects interact.
// Mediator promotes loose coupling by keeping objects from referring to each other explicitly,
// and it lets you vary their interaction independently.

type User struct {
	Name string
}

func NewUser(name string) User {
	return User{
		Name: name,
	}
}

type Mediator interface {
	SendMessage(from, to User, msg string)
}

type ChatRoom struct {
}

func (c *ChatRoom) SendMessage(from, to User, msg string) {
	fmt.Printf("[%v]->[%v]: %v\n", from.Name, to.Name, msg)
}

func NewChatRoom() Mediator {
	return &ChatRoom{}
}
