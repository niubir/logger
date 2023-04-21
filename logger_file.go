package logger

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	duration_layout = "2006_01_02_15_04_05"
)

type FileLogger struct {
	opts options

	start time.Time
	end   time.Time

	mu       sync.Mutex
	file     *os.File
	fileByte int64
}

func NewFileLogger(opts ...Option) (*FileLogger, error) {
	var optObj options
	for _, opt := range opts {
		opt.f(&optObj)
	}
	optObj.compair()

	return NewFileLoggerByOpts(optObj)
}

func NewFileLoggerByOpts(opts options) (*FileLogger, error) {
	fl := &FileLogger{
		opts: opts,

		mu: sync.Mutex{},
	}
	fl.updateDuration(time.Now())
	return fl, nil
}

func (fw *FileLogger) Write(b []byte) (n int, err error) {
	fw.mu.Lock()
	defer fw.mu.Unlock()

	maxByte := fw.opts.maxByte
	writeByte := int64(len(b))

	if writeByte > maxByte {
		return 0, fmt.Errorf("write byte %d exceeds max byte %d", writeByte, maxByte)
	}

	if fw.file == nil || fw.fileByte+writeByte > maxByte || !fw.inDuration(time.Now()) {
		if err := fw.setFile(); err != nil {
			return 0, err
		}
	}

	return fw.file.Write(b)
}

func (fw *FileLogger) setFile() error {
	fw.updateDuration(time.Now())
	for suffix := 0; ; suffix++ {
		filename := fmt.Sprintf(
			"%s/%s%s",
			fw.opts.path,
			fw.opts.prefix,
			fw.start.Format(duration_layout),
			// suffix,
		)
		if suffix != 0 {
			filename += fmt.Sprintf("__%d", suffix)
		}
		filename += ".log"
		exists, err := fw.fileExists(filename)
		if err != nil {
			return err
		}
		if exists {
			continue
		}
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		fw.file = file
		fw.fileByte = 0
		return nil
	}
}

func (fw *FileLogger) inDuration(t time.Time) bool {
	if t.Before(fw.start) || t.After(fw.end) {
		return false
	}
	return true
}

func (fw *FileLogger) updateDuration(t time.Time) {
	fw.start = t
	fw.end = fw.start.Add(fw.opts.duration)
}

func (fw *FileLogger) fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
