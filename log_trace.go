package logger

import (
	pkg_errors "github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Debug(msg string, fields ...zap.Field) {
	zapLog.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zapLog.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zapLog.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	// 如果有error field 则增加error堆栈
	for _, field := range fields {
		if field.Type == zapcore.ErrorType {
			err, ok := field.Interface.(error)
			if ok {
				fields = append(fields, zap.Any("stack", pkg_errors.WithStack(err)))
			}
		}
	}
	zapLog.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zapLog.Fatal(msg, fields...)
}
