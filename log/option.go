package log

import (
	"context"
	"io"
)

type Parameter struct {
	Writer        io.Writer
	LogLevel      LogLevel
	StaticFields  []Field
	DynamicFields func(context.Context) []Field
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

func WithStaticFields(fields []Field) Option {
	return func(c *Parameter) {
		c.StaticFields = fields
	}
}

func WithDynamicFields(fn func(context.Context) []Field) Option {
	return func(c *Parameter) {
		c.DynamicFields = fn
	}
}
