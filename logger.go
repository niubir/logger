package logger

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

type Logger struct {
	option option

	writers []io.Writer
}

func NewLogger(options ...Option) (*Logger, error) {
	var option option
	for _, o := range options {
		o.f(&option)
	}
	option.init()

	l := Logger{
		option: option,
	}

	if err := l.setWrites(); err != nil {
		return nil, err
	}
	return &l, nil
}

func (l *Logger) Debug(args ...interface{}) {
	l.LevelPrint(DebugLevel, args...)
}
func (l *Logger) Debugln(args ...interface{}) {
	l.LevelPrintln(DebugLevel, args...)
}
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.LevelPrintf(DebugLevel, format, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.LevelPrint(InfoLevel, args...)
}
func (l *Logger) Infoln(args ...interface{}) {
	l.LevelPrintln(InfoLevel, args...)
}
func (l *Logger) Infof(format string, args ...interface{}) {
	l.LevelPrintf(InfoLevel, format, args...)
}

func (l *Logger) Warning(args ...interface{}) {
	l.LevelPrint(WarningLevel, args...)
}
func (l *Logger) Warningln(args ...interface{}) {
	l.LevelPrintln(WarningLevel, args...)
}
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.LevelPrintf(WarningLevel, format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.LevelPrint(ErrorLevel, args...)
}
func (l *Logger) Errorln(args ...interface{}) {
	l.LevelPrintln(ErrorLevel, args...)
}
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.LevelPrintf(ErrorLevel, format, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.LevelPrint(FatalLevel, args...)
	os.Exit(1)
}
func (l *Logger) Fatalln(args ...interface{}) {
	l.LevelPrintln(FatalLevel, args...)
	os.Exit(1)
}
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.LevelPrintf(FatalLevel, format, args...)
	os.Exit(1)
}

func (l *Logger) LevelPrint(level Level, args ...interface{}) {
	if level < l.option.logLevel {
		return
	}
	l.Write([]byte(l.getMessage(level, fmt.Sprint(args...))))
}

func (l *Logger) LevelPrintln(level Level, args ...interface{}) {
	if level < l.option.logLevel {
		return
	}
	l.Write([]byte(l.getMessage(level, fmt.Sprintln(args...))))
}

func (l *Logger) LevelPrintf(level Level, format string, args ...interface{}) {
	if level < l.option.logLevel {
		return
	}
	l.Write([]byte(l.getMessage(level, fmt.Sprintf(format, args...))))
}

func (l *Logger) Write(p []byte) (int, error) {
	var errs []string
	for _, writer := range l.writers {
		_, err := writer.Write(p)
		if err != nil {
			errs = append(errs, err.Error())
		}
	}
	if len(errs) > 0 {
		return 0, errors.New(strings.Join(errs, ","))
	}
	return len(p), nil
}

func (l *Logger) setWrites() error {
	if l.option.withStdout {
		l.writers = append(l.writers, os.Stdout)
	}
	if l.option.withFileout != nil {
		w, err := NewFileLogger(*l.option.withFileout...)
		if err != nil {
			return err
		}
		l.writers = append(l.writers, w)
	}
	return nil
}

func (l *Logger) getMessage(level Level, info string) string {
	var msg string
	if level < l.option.logLevel {
		return msg
	}

	// time
	if l.option.withTimeFormat != "" {
		msg += time.Now().Format(l.option.withTimeFormat) + " "
	}

	// level
	msg += "[" + level.String() + "] "

	// stack
	if l.option.withStack {
		_, file, lineNo, ok := runtime.Caller(3)
		stack := "[" + fmt.Sprintf("%s:%d", file, lineNo) + "]"
		if !ok {
			stack = "[unknow stack]"
		}
		msg += stack + " "
	}

	// info
	msg += info

	return msg
}
