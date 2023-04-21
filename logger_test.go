package logger

import (
	"os"
	"testing"
	"time"
)

const (
	test_log_path = "./logs"
)

func initTestLogPath() {
	_, err := os.Stat(test_log_path)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		if err := os.Mkdir(test_log_path, 0644); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}
}

func TestLogger(t *testing.T) {
	initTestLogPath()
	logger, err := NewLogger(
		WithPath("./logs"),
		WithPrefix("test"),
		WithStdout(true),
		WithTimeFormat(time.RFC3339Nano),
		WithLevel(DebugLevel),
		WithStack(true),
	)
	if err != nil {
		t.Fatal(err)
	}

	logger.Debug("test debug")
	logger.Info("test info")
	logger.Warn("test warn")
	logger.Error("test error")
}

func TestLoggerWithDuration(t *testing.T) {
	initTestLogPath()
	logger, err := NewLogger(
		WithPath("./logs"),
		WithPrefix("test"),
		WithStdout(true),
		WithTimeFormat(time.RFC3339Nano),
		WithLevel(DebugLevel),
		WithStack(true),
		WithDuration(10*time.Second),
	)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		logger.Debug("test debug duration %d", i)
		logger.Info("test info duration %d", i)
		logger.Warn("test warn duration %d", i)
		logger.Error("test error duration %d", i)
		time.Sleep(7 * time.Second)
	}
}
