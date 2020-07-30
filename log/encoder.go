package log

import "time"

// ArrayMarshaler allows user-defined types to efficiently add themselves to the
// logging context, and to selectively omit information which shouldn't be
// included in logs (e.g., passwords).
type ArrayMarshaler interface {
	MarshalLogArray(ArrayEncoder) error
}

// ArrayEncoder is a strongly-typed, encoding-agnostic interface for adding
// array-like objects to the logging context. Of note, it supports mixed-type
// arrays even though they aren't typical in Go. Like slices, ArrayEncoders
// aren't safe for concurrent use (though typical use shouldn't require locks).
type ArrayEncoder interface {
	// Built-in types.
	PrimitiveArrayEncoder

	// Time-related types.
	AppendDuration(time.Duration)
	AppendTime(time.Time)

	// Logging-specific marshalers.
	AppendArray(ArrayMarshaler) error
	// AppendObject(ObjectMarshaler) error

	// AppendReflected uses reflection to serialize arbitrary objects, so it's
	// slow and allocation-heavy.
	AppendReflected(value interface{}) error
}

// PrimitiveArrayEncoder is the subset of the ArrayEncoder interface that deals
// only in Go's built-in types. It's included only so that Duration- and
// TimeEncoders cannot trigger infinite recursion.
type PrimitiveArrayEncoder interface {
	// Built-in types.
	AppendBool(bool)
	AppendByteString([]byte) // for UTF-8 encoded bytes
	AppendComplex128(complex128)
	AppendComplex64(complex64)
	AppendFloat64(float64)
	AppendFloat32(float32)
	AppendInt(int)
	AppendInt64(int64)
	AppendInt32(int32)
	AppendInt16(int16)
	AppendInt8(int8)
	AppendString(string)
	AppendUint(uint)
	AppendUint64(uint64)
	AppendUint32(uint32)
	AppendUint16(uint16)
	AppendUint8(uint8)
	AppendUintptr(uintptr)
}
