package logger

const (
	LevelDebug = iota
	LevelInfo
	LevelError
)

type BaseLogger interface {
	Print(level int, msg string)
	Printf(level int, format string, msg ...interface{})
	SetNextLogger(next BaseLogger)
}
