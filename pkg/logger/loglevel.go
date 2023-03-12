package logger

type LogLevel int

const (
	LogLevelNone  LogLevel = 1
	LogLevelError LogLevel = 2
	LogLevelWarn  LogLevel = 3
	LogLevelInfo  LogLevel = 4
	LogLevelDebug LogLevel = 5
	LogLevelTrace LogLevel = 6
)
