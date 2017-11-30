package Logger

import (
	"fmt"
	"time"
)

func (l *Logger) Debug(m string) {
	if l.Async == false {
		l.dispatch(LEVEL_DEBUG, m)
	} else {
		go l.dispatch(LEVEL_DEBUG, m)
	}
}

func (l *Logger) Info(m string) {
	if l.Async == false {
		l.dispatch(LEVEL_INFO, m)
	} else {
		go l.dispatch(LEVEL_INFO, m)
	}
}

func (l *Logger) Warn(m string) {
	if l.Async == false {
		l.dispatch(LEVEL_WARN, m)
	} else {
		go l.dispatch(LEVEL_WARN, m)
	}
}

func (l *Logger) Error(m string) {
	if l.Async == false {
		l.dispatch(LEVEL_ERROR, m)
	} else {
		go l.dispatch(LEVEL_ERROR, m)
	}
}

func (l *Logger) Critical(m string) {
	if l.Async == false {
		l.dispatch(LEVEL_CRITICAL, m)
	} else {
		go l.dispatch(LEVEL_CRITICAL, m)
	}
}

func (l *Logger) Fatal(m string, exitCode int) {
	l.dispatch(LEVEL_CRITICAL, m)
	l.osExit(exitCode)
}

func (l *Logger) dispatch(logLevel int, m string) {
	var mm string
	if logLevel == LEVEL_DEBUG {
		mm = fmt.Sprintf("DEBUG    %s %s", l.Header, m)
	} else if logLevel == LEVEL_INFO {
		mm = fmt.Sprintf("INFO     %s %s", l.Header, m)
	} else if logLevel == LEVEL_WARN {
		mm = fmt.Sprintf("WARN     %s %s", l.Header, m)
	} else if logLevel == LEVEL_ERROR {
		mm = fmt.Sprintf("ERROR    %s %s", l.Header, m)
	} else if logLevel == LEVEL_CRITICAL {
		mm = fmt.Sprintf("CRITICAL %s %s", l.Header, m)
	}
	mm = fmt.Sprintf("%s %s[%d]: %s\n", time.Now().UTC().Format("Jan 02 15:04:05-0700"), l.ServiceName, l.pid, mm)
	if (l.StdoutLevel & logLevel) != 0 {
		l.Stdout.WriteString(mm)
	}
	if (l.StderrLevel & logLevel) != 0 {
		l.Stderr.WriteString(mm)
	}
	if (l.DevlogLevel & logLevel) != 0 {
		switch logLevel {
		case LEVEL_DEBUG:
			l.Devlog.Debug(mm)
		case LEVEL_INFO:
			l.Devlog.Info(mm)
		case LEVEL_WARN:
			l.Devlog.Warning(mm)
		case LEVEL_ERROR:
			l.Devlog.Err(mm)
		case LEVEL_CRITICAL:
			l.Devlog.Crit(mm)
		}
	}
}
