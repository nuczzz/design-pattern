package logger

import "fmt"

type DebugLogger struct {
	level int
	next  BaseLogger
}

func (l *DebugLogger) Print(level int, msg string) {
	if level < l.level {
		return
	} else if level == l.level {
		l.Printf(level, "%s", msg)
	} else {
		l.next.Print(level, msg)
	}
}

func (l *DebugLogger) Printf(level int, format string, msg ...interface{}) {
	if level < l.level {
		return
	} else if level == l.level {
		fmt.Printf("[DEBUG]: "+format+"\n", msg...)
	} else {
		l.next.Printf(level, format, msg...)
	}
}

func (l *DebugLogger) SetNextLogger(next BaseLogger) {
	l.next = next
}

func NewDebugLogger() BaseLogger {
	return &DebugLogger{
		level: LevelDebug,
	}
}
