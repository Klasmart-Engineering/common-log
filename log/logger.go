package log

import (
	"context"
	"os"
)

// A Logger provides fast, leveled, structured logging
type Logger interface {
	Debug(context.Context, string, ...Field)
	Info(context.Context, string, ...Field)
	Warn(context.Context, string, ...Field)
	Error(context.Context, string, ...Field)
	Panic(context.Context, string, ...Field)
	Fatal(context.Context, string, ...Field)
}

var (
	globalLogger Logger = New()
)

func New(options ...Option) Logger {
	parameter := &Parameter{
		Writer:   os.Stdout,
		LogLevel: LevelDebug,
	}

	for _, option := range options {
		option(parameter)
	}

	return NewZapLogger(parameter)
}

// ReplaceGlobals replace package level logger
func ReplaceGlobals(logger Logger) {
	globalLogger = logger
}

func Debug(ctx context.Context, message string, fields ...Field) {
	globalLogger.Debug(ctx, message, fields...)
}

func Info(ctx context.Context, message string, fields ...Field) {
	globalLogger.Info(ctx, message, fields...)
}

func Warn(ctx context.Context, message string, fields ...Field) {
	globalLogger.Warn(ctx, message, fields...)
}

func Error(ctx context.Context, message string, fields ...Field) {
	globalLogger.Error(ctx, message, fields...)
}

func Panic(ctx context.Context, message string, fields ...Field) {
	globalLogger.Panic(ctx, message, fields...)
}

func Fatal(ctx context.Context, message string, fields ...Field) {
	globalLogger.Fatal(ctx, message, fields...)
}
