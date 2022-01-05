package log

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	*zap.Logger
	dynamicFields func(context.Context) []Field
}

func NewZapLogger(parameter *Parameter) *ZapLogger {
	encoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	writer := zapcore.AddSync(parameter.Writer)

	level := newZapLogLevel(parameter.LogLevel)
	core := zapcore.NewCore(encoder, writer, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	}))

	logger := &ZapLogger{
		Logger:        zap.New(core, zap.AddCaller(), zap.AddCallerSkip(2)),
		dynamicFields: parameter.DynamicFields,
	}

	if len(parameter.StaticFields) > 0 {
		logger.Logger = logger.Logger.With(logger.parseFields(parameter.StaticFields)...)
	}

	return logger
}

func (l ZapLogger) Debug(ctx context.Context, message string, fields ...Field) {
	if l.dynamicFields != nil {
		fields = append(fields, l.dynamicFields(ctx)...)
	}

	l.Logger.Debug(message, l.parseFields(fields)...)
}

func (l ZapLogger) Info(ctx context.Context, message string, fields ...Field) {
	if l.dynamicFields != nil {
		fields = append(fields, l.dynamicFields(ctx)...)
	}

	l.Logger.Info(message, l.parseFields(fields)...)
}

func (l ZapLogger) Warn(ctx context.Context, message string, fields ...Field) {
	if l.dynamicFields != nil {
		fields = append(fields, l.dynamicFields(ctx)...)
	}

	l.Logger.Warn(message, l.parseFields(fields)...)
}

func (l ZapLogger) Error(ctx context.Context, message string, fields ...Field) {
	if l.dynamicFields != nil {
		fields = append(fields, l.dynamicFields(ctx)...)
	}

	l.Logger.Error(message, l.parseFields(fields)...)
}

func (l ZapLogger) Panic(ctx context.Context, message string, fields ...Field) {
	if l.dynamicFields != nil {
		fields = append(fields, l.dynamicFields(ctx)...)
	}

	l.Logger.Panic(message, l.parseFields(fields)...)
}

func (l ZapLogger) Fatal(ctx context.Context, message string, fields ...Field) {
	if l.dynamicFields != nil {
		fields = append(fields, l.dynamicFields(ctx)...)
	}

	l.Logger.Fatal(message, l.parseFields(fields)...)
}

func (l ZapLogger) parseFields(fields []Field) []zap.Field {
	zfields := make([]zap.Field, len(fields))
	for index, field := range fields {
		switch field.Type {
		case BinaryType:
			zfields[index] = zap.Binary(field.Key, field.Value.([]byte))
		case BoolType:
			zfields[index] = zap.Bool(field.Key, field.Value.(bool))
		case ByteStringType:
			zfields[index] = zap.ByteString(field.Key, field.Value.([]byte))
		case Complex128Type:
			zfields[index] = zap.Complex128(field.Key, field.Value.(complex128))
		case Complex64Type:
			zfields[index] = zap.Complex64(field.Key, field.Value.(complex64))
		case DurationType:
			zfields[index] = zap.Duration(field.Key, field.Value.(time.Duration))
		case Float64Type:
			zfields[index] = zap.Float64(field.Key, field.Value.(float64))
		case Float32Type:
			zfields[index] = zap.Float32(field.Key, field.Value.(float32))
		case Int64Type:
			zfields[index] = zap.Int64(field.Key, field.Value.(int64))
		case Int32Type:
			zfields[index] = zap.Int32(field.Key, field.Value.(int32))
		case Int16Type:
			zfields[index] = zap.Int16(field.Key, field.Value.(int16))
		case Int8Type:
			zfields[index] = zap.Int8(field.Key, field.Value.(int8))
		case IntType:
			zfields[index] = zap.Int(field.Key, field.Value.(int))
		case StringType:
			zfields[index] = zap.String(field.Key, field.Value.(string))
		case TimeType:
			zfields[index] = zap.Time(field.Key, field.Value.(time.Time))
		case Uint64Type:
			zfields[index] = zap.Uint64(field.Key, field.Value.(uint64))
		case Uint32Type:
			zfields[index] = zap.Uint32(field.Key, field.Value.(uint32))
		case Uint16Type:
			zfields[index] = zap.Uint16(field.Key, field.Value.(uint16))
		case Uint8Type:
			zfields[index] = zap.Uint8(field.Key, field.Value.(uint8))
		case UintType:
			zfields[index] = zap.Uint(field.Key, field.Value.(uint))
		case UintptrType:
			zfields[index] = zap.Uintptr(field.Key, field.Value.(uintptr))
		case ReflectType:
			zfields[index] = zap.Reflect(field.Key, field.Value)
		case NamespaceType:
			zfields[index] = zap.Namespace(field.Key)
		case StringerType:
			zfields[index] = zap.Stringer(field.Key, field.Value.(fmt.Stringer))
		case ErrorType:
			zfields[index] = zap.NamedError(field.Key, field.Value.(error))
		case SkipType:
			zfields[index] = zap.Skip()
		case BoolsType:
			zfields[index] = zap.Bools(field.Key, field.Value.([]bool))
		case ByteStringsType:
			zfields[index] = zap.ByteStrings(field.Key, field.Value.([][]byte))
		case Complex128sType:
			zfields[index] = zap.Complex128s(field.Key, field.Value.([]complex128))
		case Complex64sType:
			zfields[index] = zap.Complex64s(field.Key, field.Value.([]complex64))
		case DurationsType:
			zfields[index] = zap.Durations(field.Key, field.Value.([]time.Duration))
		case Float64sType:
			zfields[index] = zap.Float64s(field.Key, field.Value.([]float64))
		case Float32sType:
			zfields[index] = zap.Float32s(field.Key, field.Value.([]float32))
		case Int64sType:
			zfields[index] = zap.Int64s(field.Key, field.Value.([]int64))
		case Int32sType:
			zfields[index] = zap.Int32s(field.Key, field.Value.([]int32))
		case Int16sType:
			zfields[index] = zap.Int16s(field.Key, field.Value.([]int16))
		case Int8sType:
			zfields[index] = zap.Int8s(field.Key, field.Value.([]int8))
		case IntsType:
			zfields[index] = zap.Ints(field.Key, field.Value.([]int))
		case StringsType:
			zfields[index] = zap.Strings(field.Key, field.Value.([]string))
		case TimesType:
			zfields[index] = zap.Times(field.Key, field.Value.([]time.Time))
		case Uint64sType:
			zfields[index] = zap.Uint64s(field.Key, field.Value.([]uint64))
		case Uint32sType:
			zfields[index] = zap.Uint32s(field.Key, field.Value.([]uint32))
		case Uint16sType:
			zfields[index] = zap.Uint16s(field.Key, field.Value.([]uint16))
		case Uint8sType:
			zfields[index] = zap.Uint8s(field.Key, field.Value.([]uint8))
		case UintsType:
			zfields[index] = zap.Uints(field.Key, field.Value.([]uint))
		case UintptrsType:
			zfields[index] = zap.Uintptrs(field.Key, field.Value.([]uintptr))
		default:
			zfields[index] = zap.Any(field.Key, field.Value)
		}
	}

	return zfields
}

func newZapLogLevel(level LogLevel) zapcore.Level {
	switch level {
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelError:
		return zapcore.ErrorLevel
	case LevelFatal:
		return zapcore.FatalLevel
	default:
		return zapcore.InfoLevel
	}
}
