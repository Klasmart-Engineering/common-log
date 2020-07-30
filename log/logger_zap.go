package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	*zap.Logger
}

func NewDefaultZapLogger() *ZapLogger {
	lc := zap.NewDevelopmentConfig()
	lc.Encoding = "json"

	logger, _ := lc.Build()

	return NewZapLogger(logger)
}

func NewZapLogger(logger *zap.Logger) *ZapLogger {
	return &ZapLogger{Logger: logger}
}

func (l ZapLogger) Debug(message string, fields ...Field) {
	l.Logger.Debug(message, l.parseFields(fields)...)
}

func (l ZapLogger) Info(message string, fields ...Field) {
	l.Logger.Info(message, l.parseFields(fields)...)
}

func (l ZapLogger) Warn(message string, fields ...Field) {
	l.Logger.Warn(message, l.parseFields(fields)...)
}

func (l ZapLogger) Error(message string, fields ...Field) {
	l.Logger.Error(message, l.parseFields(fields)...)
}

func (l ZapLogger) Panic(message string, fields ...Field) {
	l.Logger.Panic(message, l.parseFields(fields)...)
}

func (l ZapLogger) Fatal(message string, fields ...Field) {
	l.Logger.Fatal(message, l.parseFields(fields)...)
}

func (l ZapLogger) parseFields(fields []Field) []zap.Field {
	zfields := make([]zap.Field, len(fields))
	for index, field := range fields {
		zfields[index] = zap.Field{
			Key:       field.Key,
			Type:      zapcore.FieldType(uint8(field.Type)),
			Integer:   field.Integer,
			String:    field.String,
			Interface: field.Interface,
		}
	}

	return zfields
}
