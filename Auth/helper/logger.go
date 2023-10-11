package helper

import (
	"go.uber.org/zap"
)

var logger = zap.Must(zap.NewProduction())

func Debugf(format string, args ...interface{}) {
	logger.Sugar().Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	logger.Sugar().Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Sugar().Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Sugar().Errorf(format, args...)

}

// Fatalf logs a message at level Fatal on the standard logger.
func Fatalf(format string, args ...interface{}) {
	logger.Sugar().Fatalf(format, args...)

}
