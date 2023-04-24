package logger

import (
	"time"
)

const (
	DebugLevel Level = iota + 1
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

var (
	levelNames = map[Level]string{
		DebugLevel:   "DEBUG",
		InfoLevel:    "INFO",
		WarningLevel: "WARNING",
		ErrorLevel:   "ERROR",
		FatalLevel:   "FATAL",
	}
)

type Level int

type Option struct {
	f func(*option)
}

type option struct {
	// default no time
	withTimeFormat string
	// default Info
	logLevel Level
	// default no withStack
	withStack bool

	// default no os.Stdout
	withStdout bool

	// defalut no FileLogger
	withFileout *[]FileOption
}

func WithTime(format ...string) Option {
	timeFormat := time.RFC3339Nano
	if len(format) > 0 && format[0] != "" {
		timeFormat = format[0]
	}
	return Option{func(op *option) {
		op.withTimeFormat = timeFormat
	}}
}

func SetLevel(level Level) Option {
	return Option{func(op *option) {
		op.logLevel = level
	}}
}

func WithStack() Option {
	return Option{func(op *option) {
		op.withStack = true
	}}
}

func WithStdout() Option {
	return Option{func(op *option) {
		op.withStdout = true
	}}
}

func WithFileout(options ...FileOption) Option {
	return Option{func(op *option) {
		op.withFileout = &options
	}}
}

func (l Level) String() string {
	levelName, ok := levelNames[l]
	if !ok {
		return "UNKNOW"
	}
	return levelName
}

func (opts *option) init() {
	if _, ok := levelNames[opts.logLevel]; !ok {
		opts.logLevel = InfoLevel
	}
}
