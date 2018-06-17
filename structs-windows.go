// +build windows

package Logger

import (
"os"
)

const (
	LEVEL_DEBUG    = 16
	LEVEL_INFO     = 8
	LEVEL_WARN     = 4
	LEVEL_ERROR    = 2
	LEVEL_CRITICAL = 1
	LEVEL_NONE     = 0
)

type Logger struct {
	Header      string
	ServiceName string
	pid         int
	Stdout      *os.File
	StdoutLevel int
	Stderr      *os.File
	StderrLevel int
	DevlogLevel int
	osExit      func(code int)
	Async       bool
}

