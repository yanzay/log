// Package log is a logging package with levels and pluggable writers
package log

import (
	"flag"
	"fmt"
	"io"
)

type logLevel int

const (
	// LevelTrace is most detailed logging
	LevelTrace logLevel = 1 + iota
	// LevelDebug is level for debugging logs
	LevelDebug
	// LevelInfo is level for info logs
	LevelInfo
	// LevelWarning is level for warning logs
	LevelWarning
	// LevelError is level for error logs
	LevelError
	// LevelCritical is logging only for fatal errors
	LevelCritical
)

var (
	// Level is current log level for logger
	Level logLevel
	// Writer for writing logs to. You can change it for your own writer
	Writer io.Writer = DefaultWriter{}
)

var logFlag = flag.String("log-level", "info", "Log level: trace|debug|info|warning|error|critical")

func levelFromString(str string) logLevel {
	switch str {
	case "trace":
		return LevelTrace
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warning":
		return LevelWarning
	case "error":
		return LevelError
	case "critical":
		return LevelCritical
	}
	return LevelInfo // Default Level
}

func printString(s string) {
	_, err := Writer.Write([]byte(s))
	if err != nil {
		fmt.Printf("Error write log: %s\n", err)
	}
}

func lprint(level logLevel, value interface{}) {
	if Level == 0 {
		Level = levelFromString(*logFlag)
	}
	if level >= Level {
		printString(fmt.Sprint(value))
	}
}

func lprintf(level logLevel, format string, params ...interface{}) {
	lprint(level, fmt.Sprintf(format, params...))
}

// Println is unconditional log
func Println(value interface{}) {
	printString(fmt.Sprint(value))
}

// Printf is unconditional formatted log
func Printf(format string, params ...interface{}) {
	printString(fmt.Sprintf(format, params...))
}

// Trace logging. Use it for most detailed logs
func Trace(value interface{}) { lprint(LevelTrace, value) }

// Tracef is formatted trace logging
func Tracef(format string, params ...interface{}) { lprintf(LevelTrace, format, params...) }

// Debug logging
func Debug(value interface{}) { lprint(LevelDebug, value) }

// Debugf is formatted debug logging
func Debugf(format string, params ...interface{}) { lprintf(LevelDebug, format, params...) }

// Info logging
func Info(value interface{}) { lprint(LevelInfo, value) }

// Infof is formatted info logging
func Infof(format string, params ...interface{}) { lprintf(LevelInfo, format, params...) }

// Warning logging
func Warning(value interface{}) { lprint(LevelWarning, value) }

// Warningf is formatted warning logging
func Warningf(format string, params ...interface{}) { lprintf(LevelWarning, format, params...) }

// Error logging
func Error(value interface{}) { lprint(LevelError, value) }

// Errorf is formatted error logging
func Errorf(format string, params ...interface{}) { lprintf(LevelError, format, params...) }

// Fatal logs fatal error and panic
func Fatal(value interface{}) {
	str := fmt.Sprint(value)
	printString(str)
	panic(str)
}

// Fatalf logs fatal error with format and panic
func Fatalf(format string, params ...interface{}) {
	str := fmt.Sprintf(format, params...)
	printString(str)
	panic(str)
}
