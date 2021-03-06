package command

// Command Pattern Definition:
// Encapsulate a request as an object,thereby letting you parameterize clients with different requests,
// queue or log requests,and support undoable operations.

type Command interface {
	Execute()
}
