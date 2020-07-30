package log

import (
	"context"
	"testing"
	"time"

	"gitlab.badanamu.com.cn/calmisland/common-cn/helper"
)

var (
	testContext context.Context
	testFields  []Field
)

func TestMain(m *testing.M) {
	badaCtx := &helper.BadaCtx{
		CurrTid:  "tid-1234567",
		PrevTid:  "prev-2222222",
		EntryTid: "entry-333333",
	}
	testContext = context.WithValue(context.TODO(), helper.CtxKeyBadaCtx, badaCtx)

	testFields = []Field{
		Int("int", 1),
		Int8("int8", 8),
		Int16("int16", 16),
		Int32("int32", 32),
		Int64("int64", 64),
		Float32("float32", 32.3232323),
		Float64("float64", 64.6464646464646),
		String("string", "abc"),
		Time("time", time.Now()),
		Duration("duration", time.Second),
		Any("any", badaCtx),
	}
}

func TestNewLogger(t *testing.T) {
	Debug(testContext, "Debug test", testFields...)
	Info(testContext, "INFO test", testFields...)
	Warn(testContext, "Warn test", testFields...)
	Error(testContext, "Error test", testFields...)
}

func TestPanic(t *testing.T) {
	Panic(testContext, "Panic test", testFields...)
}

func TestFatal(t *testing.T) {
	Fatal(testContext, "Fatal test", testFields...)
}
