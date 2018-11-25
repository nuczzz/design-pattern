package logger

const (
	_ = iota
	LevelDebug
	levelWarning
	LevelInfo
	LevelError
)

type Logger struct {
	level      int
	nextLogger *Logger
}

func (l *Logger) SetNextLogger(logger *Logger) {
	l.nextLogger = logger
}

func (l *Logger) PrintMessage(level int, msg string) {
	if level <= l.level {

	}
}

func (l *Logger) Write(msg string) {

}
