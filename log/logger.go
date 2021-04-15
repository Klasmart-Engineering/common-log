package log

import (
	"context"
	"os"
)

// A Logger provides fast, leveled, structured logging
type Logger interface {
	Debug(string, ...Field)
	Info(string, ...Field)
	Warn(string, ...Field)
	Error(string, ...Field)
	Panic(string, ...Field)
	Fatal(string, ...Field)
}

var (
	globalLogger Logger = New()
)

func New(options ...Option) Logger {
	parameter := &Parameter{
		Writer: os.Stdout,
	}

	for _, option := range options {
		option(parameter)
	}

	return NewDefaultZapLogger(parameter)
}

// ReplaceGlobals replace package level logger
func ReplaceGlobals(logger Logger) {
	globalLogger = logger
}

func Debug(ctx context.Context, message string, fields ...Field) {
	fields = appendFields(ctx, fields)
	globalLogger.Debug(message, fields...)
}

func Info(ctx context.Context, message string, fields ...Field) {
	fields = appendFields(ctx, fields)
	globalLogger.Info(message, fields...)
}

func Warn(ctx context.Context, message string, fields ...Field) {
	fields = appendFields(ctx, fields)
	globalLogger.Warn(message, fields...)
}

func Error(ctx context.Context, message string, fields ...Field) {
	fields = appendFields(ctx, fields)
	globalLogger.Error(message, fields...)
}

func Panic(ctx context.Context, message string, fields ...Field) {
	fields = appendFields(ctx, fields)
	globalLogger.Panic(message, fields...)
}

func Fatal(ctx context.Context, message string, fields ...Field) {
	fields = appendFields(ctx, fields)
	globalLogger.Fatal(message, fields...)
}

func appendFields(ctx context.Context, fields []Field) []Field {
	return append(fields, TraceContext(ctx)...)
}
