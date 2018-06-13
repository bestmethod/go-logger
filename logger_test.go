package Logger

import (
	"testing"
	"os"
)

func TestLoggerMissingLevel(t *testing.T) {
	logger := new(Logger)
	err := logger.Init("TEST", "TEST", LEVEL_WARN, LEVEL_ERROR, LEVEL_NONE)
	if err != nil {
		t.FailNow()
	}
	logger.Info("InfoTest")
	logger.Debug("InfoDebug")
	logger.Warn("InfoWarn")
	logger.Error("InfoError")
	logger.Critical("InfoCritical")
}

func TestLoggerConsoleError(t *testing.T) {
	logger := new(Logger)
	err := logger.Init("TEST", "TEST", LEVEL_DEBUG|LEVEL_INFO|LEVEL_WARN, LEVEL_ERROR|LEVEL_CRITICAL, LEVEL_NONE)
	if err != nil {
		t.FailNow()
	}
	logger.Info("InfoTest")
	logger.Debug("InfoDebug")
	logger.Warn("InfoWarn")
	logger.Error("InfoError")
	logger.Critical("InfoCritical")
}

func TestLoggerConsole(t *testing.T) {
	logger := new(Logger)
	err := logger.Init("TEST", "TEST", LEVEL_DEBUG|LEVEL_INFO|LEVEL_WARN|LEVEL_ERROR|LEVEL_CRITICAL, LEVEL_NONE, LEVEL_NONE)
	if err != nil {
		t.FailNow()
	}
	logger.Info("InfoTest")
	logger.Debug("InfoDebug")
	logger.Warn("InfoWarn")
	logger.Error("InfoError")
	logger.Critical("InfoCritical")
}

func TestLoggerError(t *testing.T) {
	logger := new(Logger)
	err := logger.Init("TEST", "TEST", LEVEL_NONE, LEVEL_DEBUG|LEVEL_INFO|LEVEL_WARN|LEVEL_ERROR|LEVEL_CRITICAL, LEVEL_NONE)
	if err != nil {
		t.FailNow()
	}
	logger.Info("InfoTest")
	logger.Debug("InfoDebug")
	logger.Warn("InfoWarn")
	logger.Error("InfoError")
	logger.Critical("InfoCritical")
	logger.Destroy()
}

func TestLoggerDevlog(t *testing.T) {
	if _, err := os.Stat("/dev/log"); err == nil {
	logger := new(Logger)
	err := logger.Init("TEST", "TEST", LEVEL_NONE, LEVEL_NONE, LEVEL_DEBUG|LEVEL_INFO|LEVEL_WARN|LEVEL_ERROR|LEVEL_CRITICAL)
	if err != nil {
		t.FailNow()
	}
	logger.Info("InfoTest")
	logger.Debug("InfoDebug")
	logger.Warn("InfoWarn")
	logger.Error("InfoError")
	logger.Critical("InfoCritical")
	logger.Destroy()
	}
}

func TestLoggerAsync(t *testing.T) {
	logger := new(Logger)
	err := logger.Init("TEST", "TEST", LEVEL_NONE, LEVEL_DEBUG|LEVEL_INFO|LEVEL_WARN|LEVEL_ERROR|LEVEL_CRITICAL, LEVEL_NONE)
	logger.Async = true
	if err != nil {
		t.FailNow()
	}
	logger.Info("InfoTest")
	logger.Debug("InfoDebug")
	logger.Warn("InfoWarn")
	logger.Error("InfoError")
	logger.Critical("InfoCritical")
	logger.Destroy()
}

func mockOsExit(code int) {
	return
}

func TestLoggerFatal(t *testing.T) {
	logger := new(Logger)
	err := logger.Init("TEST", "TEST", LEVEL_CRITICAL, LEVEL_NONE, LEVEL_NONE)
	if err != nil {
		t.FailNow()
	}
	logger.osExit = mockOsExit
	logger.Fatal("FatalTest", 0)
}
