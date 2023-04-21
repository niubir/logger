package logger

import (
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once
var zapLog *zap.Logger

func Init(
	path, fileName, serviceName string, maxSize, maxBackups, maxAge, level int, withStdout bool,
) error {
	once.Do(func() {
		hook := lumberjack.Logger{
			Filename:   path + fileName, // 日志文件路径
			MaxSize:    maxSize,         // megabytes
			MaxBackups: maxBackups,      // 最多保留3个备份
			MaxAge:     maxAge,          // days
			Compress:   true,            // 是否压缩 disabled by default
		}

		var zapLevel zapcore.Level
		switch level {
		case 1:
			zapLevel = zapcore.DebugLevel
		case 2:
			zapLevel = zapcore.InfoLevel
		case 3:
			zapLevel = zapcore.WarnLevel
		case 4:
			zapLevel = zapcore.ErrorLevel
		case 5:
			zapLevel = zapcore.FatalLevel
		default:
			zapLevel = zapcore.InfoLevel
		}

		encoderConfig := zapcore.EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "level",
			TimeKey:       "time",
			NameKey:       "logger",
			CallerKey:     "linenum",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
			},
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeName:     zapcore.FullNameEncoder,
		}

		var writeSyncer zapcore.WriteSyncer
		if withStdout {
			writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
		} else {
			writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))
		}

		core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writeSyncer, zapLevel)

		caller := zap.AddCaller()
		development := zap.Development()
		filed := zap.Fields(zap.String("serviceName", serviceName))

		zapLog = zap.New(core, caller, zap.AddCallerSkip(1), development, filed)
	})
	return nil
}
