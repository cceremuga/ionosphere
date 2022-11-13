// Package log is a service to wrap multiple logger types.
package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var stdoutLog *logrus.Logger
var fileLog *logrus.Logger

func init() {
	stdoutLog = logrus.New()
	stdoutLog.SetOutput(os.Stdout)
	stdoutLog.SetLevel(logrus.DebugLevel)

	fileLog = logrus.New()

	file, err := os.OpenFile(
		"logs/ionosphere.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err == nil {
		fileLog.SetOutput(file)
		fileLog.SetLevel(logrus.DebugLevel)
	} else {
		stdoutLog.Fatal("Failed to initialize log file.")
	}
}

// Debug wraps logger.Debug
func Debug(args ...interface{}) {
	stdoutLog.Debug(args...)
	fileLog.Debug(args...)
}

// Info wraps logger.Info
func Info(args ...interface{}) {
	stdoutLog.Info(args...)
	fileLog.Info(args...)
}

// Error wraps logger.Error
func Error(args ...interface{}) {
	stdoutLog.Error(args...)
	fileLog.Error(args...)
}

// Fatal wraps logger.Fatal
func Fatal(args ...interface{}) {
	stdoutLog.Fatal(args...)
	fileLog.Fatal(args...)
}

// Fatalf wraps logger.Fatalf
func Fatalf(msg string, args ...interface{}) {
	stdoutLog.Fatalf(msg, args...)
	fileLog.Fatalf(msg, args...)
}

// Warn wraps logger.Warn
func Warn(args ...interface{}) {
	stdoutLog.Warn(args...)
	fileLog.Warn(args...)
}

// Printf wraps logger.Printf
func Printf(msg string, args ...interface{}) {
	stdoutLog.Printf(msg, args...)
	fileLog.Printf(msg, args...)
}

type StderrLogger struct{}

func (StderrLogger) Write(b []byte) (int, error) {
	Debug(string(b))
	return len(b), nil
}
