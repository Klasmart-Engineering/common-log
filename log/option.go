package log

import "io"

type Parameter struct {
	Writer   io.Writer
	LogLevel LogLevel
}

// Option logger option
type Option func(*Parameter)

func WithWriter(w io.Writer) Option {
	return func(c *Parameter) {
		c.Writer = w
	}
}

func WithLogLevel(level LogLevel) Option {
	return func(c *Parameter) {
		c.LogLevel = level
	}
}
