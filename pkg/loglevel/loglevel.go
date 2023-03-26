package loglevel

import "fmt"

type LogLevel int

const (
	None  LogLevel = 1
	Error LogLevel = 2
	Warn  LogLevel = 3
	Info  LogLevel = 4
	Debug LogLevel = 5
	Trace LogLevel = 6
)

func New(logLevel int) (LogLevel, error) {
	switch logLevel {
	case 1:
		return None, nil
	case 2:
		return Error, nil
	case 3:
		return Warn, nil
	case 4:
		return Info, nil
	case 5:
		return Debug, nil
	case 6:
		return Trace, nil
	default:
		return 0, fmt.Errorf("unknown logging level: %d", logLevel)
	}
}
