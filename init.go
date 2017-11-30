package Logger

import (
	"log/syslog"
	"os"
)

func (l *Logger) Init(header string, serviceName string, stdoutLevel int, stderrLevel int, devlogLevel int) error {
	l.pid = os.Getpid()
	l.Header = header
	l.ServiceName = serviceName
	l.StdoutLevel = stdoutLevel
	l.StderrLevel = stderrLevel
	l.DevlogLevel = devlogLevel
	l.osExit = os.Exit
	l.Async = false
	var err error
	err = nil
	if stdoutLevel != 0 {
		//l.Stdout = log.New(os.Stdout, "", 0)
		l.Stdout = os.Stdout
	} else {
		l.Stdout = nil
	}
	if stderrLevel != 0 {
		//l.Stderr = log.New(os.Stderr, "", 0)
		l.Stderr = os.Stderr
	} else {
		l.Stderr = nil
	}
	if devlogLevel != 0 {
		l.Devlog, err = syslog.Dial("", "", syslog.LOG_DAEMON, l.ServiceName)
	} else {
		l.Devlog = nil
	}
	return err
}

func (l *Logger) Destroy() error {
	if l.Devlog != nil {
		l.DevlogLevel = 0
		return l.Devlog.Close()
	} else {
		return nil
	}
}
