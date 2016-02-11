package log

import (
	"flag"
	"fmt"
	"io"
)

type LogLevel int

const (
	LevelTrace LogLevel = 1 + iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

var (
	Level  LogLevel  = 0
	Writer io.Writer = DefaultWriter{}
)

var logFlag = flag.String("log-level", "info", "Log level: trace|debug|info|warning|error|critical")

func levelFromString(str string) LogLevel {
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
	Writer.Write([]byte(s))
}

func print(level LogLevel, value interface{}) {
	if Level == 0 {
		Level = levelFromString(*logFlag)
	}
	if level >= Level {
		printString(fmt.Sprint(value))
	}
}

func printf(level LogLevel, format string, params ...interface{}) {
	print(level, fmt.Sprintf(format, params...))
}

func Println(value interface{}) {
	printString(fmt.Sprint(value))
}
func Printf(format string, params ...interface{}) {
	printString(fmt.Sprintf(format, params...))
}

func Trace(value interface{})                     { print(LevelTrace, value) }
func Tracef(format string, params ...interface{}) { printf(LevelTrace, format, params...) }

func Debug(value interface{})                     { print(LevelDebug, value) }
func Debugf(format string, params ...interface{}) { printf(LevelDebug, format, params...) }

func Info(value interface{})                     { print(LevelInfo, value) }
func Infof(format string, params ...interface{}) { printf(LevelInfo, format, params...) }

func Warning(value interface{})                     { print(LevelWarning, value) }
func Warningf(format string, params ...interface{}) { printf(LevelWarning, format, params...) }

func Error(value interface{})                     { print(LevelError, value) }
func Errorf(format string, params ...interface{}) { printf(LevelError, format, params...) }

func Fatal(value interface{}) {
	str := fmt.Sprint(value)
	printString(str)
	panic(str)
}
func Fatalf(format string, params ...interface{}) {
	str := fmt.Sprintf(format, params...)
	printString(str)
	panic(str)
}
