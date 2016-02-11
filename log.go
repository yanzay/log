package log

import (
	"flag"
	"log"
)

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
	LevelCritical
)

var Level = LevelInfo

func init() {
	logLevel := flag.String("log-level", "info", "Log level: trace|debug|info|warning|error|critical")
	flag.Parse()
	Level = levelFromString(*logLevel)
}

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

func print(level LogLevel, value interface{}) {
	if level >= Level {
		log.Println(value)
	}
}

func printf(level LogLevel, format string, params ...interface{}) {
	if level >= Level {
		log.Printf(format, params...)
	}
}

func Println(value interface{})                   { log.Println(value) }
func Printf(format string, params ...interface{}) { log.Printf(format, params...) }

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

func Fatal(value interface{})                     { log.Fatal(value) }
func Fatalf(format string, params ...interface{}) { log.Fatalf(format, params...) }
