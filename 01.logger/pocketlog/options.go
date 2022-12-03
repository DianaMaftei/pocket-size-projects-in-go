package pocketlog

import "io"

// Option defines a functional option to our logger.
type Option func(*Logger)

func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
}

func WithMaxLength(maxLength int) Option {
	return func(l *Logger) {
		l.maxMessageLength = maxLength
	}
}
