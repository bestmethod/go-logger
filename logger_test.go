package Logger

import (
	"testing"
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
}

func TestLoggerDevlog(t *testing.T) {
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
