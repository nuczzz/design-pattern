package logger

type Logger struct {
	loggerChain BaseLogger
}

func (l *Logger) Debug(msg string) {
	l.loggerChain.Print(LevelDebug, msg)
}

func (l *Logger) Debugf(format string, msg ...interface{}) {
	l.loggerChain.Printf(LevelDebug, format, msg...)
}

func (l *Logger) Info(msg string) {
	l.loggerChain.Print(LevelInfo, msg)
}

func (l *Logger) Infof(format string, msg ...interface{}) {
	l.loggerChain.Printf(LevelInfo, format, msg...)
}

func (l *Logger) Error(msg string) {
	l.loggerChain.Print(LevelError, msg)
}

func (l *Logger) Errorf(format string, msg ...interface{}) {
	l.loggerChain.Printf(LevelError, format, msg...)
}

func newLoggerChain() BaseLogger {
	debugLogger := NewDebugLogger()
	infoLogger := NewInfoLogger()
	errorLogger := NewErrorLogger()

	debugLogger.SetNextLogger(infoLogger)
	infoLogger.SetNextLogger(errorLogger)

	return debugLogger
}

func NewLogger() *Logger {
	return &Logger{
		loggerChain: newLoggerChain(),
	}
}
