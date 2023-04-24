package logger

import "time"

type FileOption struct {
	f func(*fileOption)
}

type fileOption struct {
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
}

func WithFilePath(path string) FileOption {
	return FileOption{func(op *fileOption) {
		op.path = path
	}}
}

func WithFilePrefix(prefix string) FileOption {
	return FileOption{func(op *fileOption) {
		op.prefix = prefix
	}}
}

func WithFileDuration(duration time.Duration) FileOption {
	return FileOption{func(op *fileOption) {
		op.duration = duration
	}}
}

func WithFileMaxByte(maxByte int64) FileOption {
	return FileOption{func(op *fileOption) {
		op.maxByte = maxByte
	}}
}

func (op *fileOption) init() {
	if op.path == "" {
		op.path = "./"
	}
	if op.prefix != "" {
		op.prefix += "_"
	}
	if op.duration < 10*time.Second {
		op.duration = 365 * 24 * time.Hour
	}
	if op.maxByte < 1024 {
		op.maxByte = 1024
	}
}
