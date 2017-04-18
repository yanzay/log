package log

import (
	"io"
	"log/syslog"
)

type SyslogWriter struct {
	*syslog.Writer
}

func NewSyslogWriter(priority int, tag string) (io.Writer, error) {
	writer, err := syslog.New(syslog.Priority(priority), tag)
	if err != nil {
		return nil, err
	}
	return &SyslogWriter{Writer: writer}, nil
}
