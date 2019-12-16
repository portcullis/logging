package logging

// This is all placeholder stuff

import (
	"fmt"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

// Debug log message
func Debug(msg string, args ...interface{}) {
	output("[DEBUG] "+msg, args...)
}

// Info log message
func Info(msg string, args ...interface{}) {
	output("[INFO] "+msg, args...)
}

// Warning log message
func Warning(msg string, args ...interface{}) {
	output("[WARNING] "+msg, args...)
}

// Error log message
func Error(msg string, args ...interface{}) {
	output("[FATAL] "+msg, args...)
}

// Fatal log message, and exit -1
func Fatal(msg string, args ...interface{}) {
	output("[FATAL] "+msg, args...)
	os.Exit(-1)
}

// Panic log message - calls panic
func Panic(msg string, args ...interface{}) {
	if len(args) > 0 {
		panic(fmt.Sprintf(msg, args...))
	} else {
		panic(msg)
	}
}

func output(msg string, args ...interface{}) {
	if len(args) > 0 {
		log.Printf(msg, args...)
	} else {
		log.Print(msg)
	}
}
