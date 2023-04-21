package logger

import (
	"fmt"
	"runtime"
	"time"
)

const (
	DebugLevel level = iota + 1
	InfoLevel
	WarnLevel
	ErrorLevel
)

type level int

type Option struct {
	f func(*options)
}

type options struct {
	// default ./
	path string
	// default no prefix
	prefix string
	// duration log file
	// greater or equal than 10s
	// default 365 day
	duration time.Duration
	// log file size greater than maxByte, suffix+1
	// greater or equal than 1024
	// default 1024
	maxByte int64
	// os.Stdout
	// default false
	stdout bool

	// time format
	// default none
	timeFormat string
	// default InfoLevel
	level level
	// stack
	// default false
	stack bool

	compaired bool
}

func (l level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	default:
		return "UNKNOW"
	}
}

func (opts *options) compair() {
	if opts.compaired {
		return
	}
	if opts.path == "" {
		opts.path = "./"
	}
	if opts.prefix != "" {
		opts.prefix += "_"
	}
	if opts.level < DebugLevel || opts.level > ErrorLevel {
		opts.level = InfoLevel
	}
	if opts.duration < 10*time.Second {
		opts.duration = 365 * 24 * time.Hour
	}
	if opts.maxByte < 1024 {
		opts.maxByte = 1024
	}
	opts.compaired = true
}

func (opts *options) msg(withColor bool, level level, s string, i ...interface{}) string {
	var msg string
	// time
	if opts.timeFormat != "" {
		msg += time.Now().Format(opts.timeFormat) + "\t"
	}

	// level
	if withColor {
		var (
			levelBackground = 40
			levelShowType   = 1
			levelColor      = 37
		)
		switch level {
		case DebugLevel:
			levelColor = 36
		case InfoLevel:
			levelColor = 32
		case WarnLevel:
			levelColor = 33
		case ErrorLevel:
			levelColor = 31
		}
		msg += fmt.Sprintf("[%c[%d;%d;%dm%s%s%c[0m]\t", 0x1B, levelShowType, levelBackground, levelColor, level.String(), "", 0x1B)
	} else {
		msg += "[" + level.String() + "]\t"
	}

	// stack
	if opts.stack {
		_, file, lineNo, ok := runtime.Caller(3)
		stack := "[" + fmt.Sprintf("%s:%d", file, lineNo) + "]"
		if !ok {
			stack = "[unknow stack]"
		}
		msg += stack + "\t"
	}
	msg += fmt.Sprintf(s, i...) + "\n"

	return msg
}

func WithPath(path string) Option {
	return Option{func(op *options) {
		op.path = path
	}}
}

func WithPrefix(prefix string) Option {
	return Option{func(op *options) {
		op.prefix = prefix
	}}
}

func WithLevel(level level) Option {
	return Option{func(op *options) {
		op.level = level
	}}
}

func WithDuration(duration time.Duration) Option {
	return Option{func(op *options) {
		op.duration = duration
	}}
}

func WithMaxByte(maxByte int64) Option {
	return Option{func(op *options) {
		op.maxByte = maxByte
	}}
}

func WithStdout(stdout bool) Option {
	return Option{func(op *options) {
		op.stdout = stdout
	}}
}

func WithTimeFormat(timeFormat string) Option {
	return Option{func(op *options) {
		op.timeFormat = timeFormat
	}}
}

func WithStack(stack bool) Option {
	return Option{func(op *options) {
		op.stack = stack
	}}
}
