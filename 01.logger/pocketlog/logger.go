package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold        Level
	output           io.Writer
	maxMessageLength int
}

// New returns you a logger, ready to logf at the required threshold.
// Give it a list of configuration functions to tune it at your will
// The default output is Stdout.
// There is no maxMessageLength character limit
func New(threshold Level, opts ...Option) *Logger {
	l := &Logger{threshold: threshold, output: os.Stdout, maxMessageLength: 0}

	for _, configFunc := range opts {
		configFunc(l)
	}

	return l
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	l.Logf(LevelDebug, format, args...)
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	l.Logf(LevelInfo, format, args...)
}

// Errorf formats and prints a message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	l.Logf(LevelError, format, args...)
}

// Logf formats and prints a message if the log level is high enough
func (l *Logger) Logf(lvl Level, format string, args ...any) {
	if l.threshold > lvl {
		return
	}
	l.logf(lvl, format, args...)
}

// logf prints the message to the output
// Add decorations here, if any
// Text longer than maxMessageLength will be trimmed off
func (l *Logger) logf(level Level, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	if l.maxMessageLength != 0 && len([]rune(message)) > l.maxMessageLength {
		message = string([]rune(message)[:l.maxMessageLength])
	}
	_, _ = fmt.Fprintf(l.output, "%s %s\n", level, message)
}
