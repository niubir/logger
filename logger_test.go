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
		SetLevel(DebugLevel),
		WithTime(time.Stamp),
		WithStack(),
		WithStdout(),
		WithFileout(
			WithFilePath(test_log_path),
			WithFilePrefix("test"),
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	logger.Debugln("test debug")
	logger.Infoln("test info")
	logger.Warningln("test warning")
	logger.Errorln("test error")
}

func TestLoggerWithDuration(t *testing.T) {
	initTestLogPath()
	logger, err := NewLogger(
		SetLevel(DebugLevel),
		WithTime(time.Stamp),
		WithStack(),
		WithStdout(),
		WithFileout(
			WithFilePath(test_log_path),
			WithFilePrefix("test"),
			WithFileDuration(10*time.Second),
		),
	)
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		logger.Debugf("test debug duration %d\n", i)
		logger.Infof("test info duration %d\n", i)
		logger.Warningf("test warning duration %d\n", i)
		logger.Errorf("test error duration %d\n", i)
		time.Sleep(7 * time.Second)
	}
}
