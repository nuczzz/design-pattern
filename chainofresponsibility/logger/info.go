package logger

import "fmt"

type InfoLogger struct {
	level int
	next  BaseLogger
}

func (l *InfoLogger) Print(level int, msg string) {
	if level < l.level {
		return
	} else if level == l.level {
		l.Printf(level, "%s", msg)
	} else {
		l.next.Print(level, msg)
	}
}

func (l *InfoLogger) Printf(level int, format string, msg ...interface{}) {
	if level < l.level {
		return
	} else if level == l.level {
		fmt.Printf("[INFO]: "+format+"\n", msg...)
	} else {
		l.next.Printf(level, format, msg...)
	}
}

func (l *InfoLogger) SetNextLogger(next BaseLogger) {
	l.next = next
}

func NewInfoLogger() BaseLogger {
	return &InfoLogger{
		level: LevelInfo,
	}
}
