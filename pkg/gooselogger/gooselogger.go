package gooselogger

import (
	"strings"

	"go.uber.org/zap"
)

type GooseLogger struct {
	lg *zap.SugaredLogger
}

func New(lg *zap.Logger) *GooseLogger {
	return &GooseLogger{lg: lg.Sugar()}
}

func (l *GooseLogger) Fatal(v ...interface{}) {
	l.lg.Info(v...)
}

func (l *GooseLogger) Fatalf(format string, v ...interface{}) {
	l.lg.Fatalf(l.clean(format), v...)
}

func (l *GooseLogger) Print(v ...interface{}) {
	l.lg.Info(v...)
}

func (l *GooseLogger) Println(v ...interface{}) {
	l.lg.Infoln(v...)
}

func (l *GooseLogger) Printf(format string, v ...interface{}) {
	l.lg.Infof(l.clean(format), v...)
}

func (l *GooseLogger) clean(str string) string {
	str = strings.TrimPrefix(str, "goose: ")
	return strings.TrimSuffix(str, "\n")
}
