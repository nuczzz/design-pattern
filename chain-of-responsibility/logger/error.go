package logger

import "fmt"

type ErrorLogger struct {
	level int
	next  BaseLogger
}

func (l *ErrorLogger) Print(level int, msg string) {
	if level < l.level {
		return
	} else if level == l.level {
		l.Printf(level, "%s", msg)
	} else {
		l.next.Print(level, msg)
	}
}

func (l *ErrorLogger) Printf(level int, format string, msg ...interface{}) {
	if level < l.level {
		return
	} else if level == l.level {
		fmt.Printf("[ERROR]: "+format+"\n", msg...)
	} else {
		l.next.Printf(level, format, msg...)
	}
}

func (l *ErrorLogger) SetNextLogger(next BaseLogger) {
	l.next = next
}

func NewErrorLogger() BaseLogger {
	return &ErrorLogger{
		level: LevelError,
	}
}
