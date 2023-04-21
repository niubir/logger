package logger

import (
	"io"
	"os"
)

type Logger struct {
	opts options

	stdoutWriter io.Writer
	fileWriter   *FileLogger
}

func NewLogger(opts ...Option) (*Logger, error) {
	var optObj options
	for _, opt := range opts {
		opt.f(&optObj)
	}
	optObj.compair()

	l := Logger{
		opts: optObj,
	}

	if err := l.setWrites(); err != nil {
		return nil, err
	}
	return &l, nil
}

func (l *Logger) setWrites() error {
	fileLogger, err := NewFileLoggerByOpts(l.opts)
	if err != nil {
		return err
	}
	l.fileWriter = fileLogger
	if l.opts.stdout {
		l.stdoutWriter = os.Stdout
	}
	return nil
}

func (l *Logger) Debug(msg string, i ...interface{}) { l.log(DebugLevel, msg, i...) }

func (l *Logger) Info(msg string, i ...interface{}) { l.log(InfoLevel, msg, i...) }

func (l *Logger) Warn(msg string, i ...interface{}) { l.log(WarnLevel, msg, i...) }

func (l *Logger) Error(msg string, i ...interface{}) { l.log(ErrorLevel, msg, i...) }

func (l *Logger) log(level level, msg string, i ...interface{}) {
	if level < l.opts.level {
		return
	}
	if l.stdoutWriter != nil {
		l.stdoutWriter.Write([]byte(l.opts.msg(true, level, msg, i...)))
	}
	if l.fileWriter != nil {
		l.fileWriter.Write([]byte(l.opts.msg(false, level, msg, i...)))
	}
}
