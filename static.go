package logging

import (
	"fmt"
	"os"
)

var (
	// DefaultLog is the *Log used in static functions and is returned from contexts when no *Log is present
	DefaultLog = &Log{writer: Discard}
)

// Debug log message
func Debug(msg string, args ...interface{}) {
	DefaultLog.Debug(msg, args...)
}

// Info log message
func Info(msg string, args ...interface{}) {
	DefaultLog.Info(msg, args...)
}

// Warning log message
func Warning(msg string, args ...interface{}) {
	DefaultLog.Warning(msg, args...)
}

// Error log message
func Error(msg string, args ...interface{}) {
	DefaultLog.Error(msg, args...)
}

// Fatal log message, and exit -1
func Fatal(msg string, args ...interface{}) {
	DefaultLog.Error(msg, args...)
	os.Exit(-1)
}

// Panic log message - calls panic
func Panic(msg string, args ...interface{}) {
	DefaultLog.Error(msg, args...)

	if len(args) > 0 {
		panic(fmt.Sprintf(msg, args...))
	} else {
		panic(msg)
	}
}
