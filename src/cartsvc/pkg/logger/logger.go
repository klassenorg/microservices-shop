package logger

import (
	"go.uber.org/zap"
)

var l *zap.SugaredLogger

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() *Logger {
	return &Logger{l}
}

func Init() {
	logger, _ := zap.NewProduction()
	l = logger.Sugar()
}
