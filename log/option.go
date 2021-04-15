package log

import "io"

type Parameter struct {
	Writer io.Writer
}

// Option logger option
type Option func(*Parameter)

func WithWriter(w io.Writer) Option {
	return func(c *Parameter) {
		c.Writer = w
	}
}
