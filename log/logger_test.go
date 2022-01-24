package log

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var (
	testContext context.Context
	testFields  []Field
)

func TestMain(m *testing.M) {
	testContext = context.Background()

	testFields = []Field{
		Int("int", 1),
		Ints("ints", []int{1, 2}),
		Int8("int8", 8),
		Int8s("int8s", []int8{9, 7}),
		Int16("int16", 16),
		Int16s("int16s", []int16{16, 17}),
		Int32("int32", 32),
		Int32s("int32s", []int32{32, 33}),
		Int64("int64", 64),
		Int64s("int64s", []int64{64, 65}),
		Float32("float32", 32.3232323),
		Float32s("float32s", []float32{32.33, 33.32}),
		Float64("float64", 64.6464646464646),
		Float64s("float64s", []float64{32.33, 33.32}),
		Uint("uint", 4),
		Uints("uints", []uint{5, 6}),
		Uint8("uint8", 8),
		Uint8s("uint8s", []uint8{9, 6}),
		Uint16("uint16", 16),
		Uint16s("uint16s", []uint16{19, 16}),
		Uint32("uint32", 32),
		Uint32s("uint32s", []uint32{33, 46}),
		Uint64("uint64", 64),
		Uint64s("uint64s", []uint64{9, 6}),
		String("string", "abc"),
		Strings("strings", []string{"ccc", "def"}),
		Time("time", time.Now()),
		Times("times", []time.Time{time.Now(), time.Now().Add(time.Hour)}),
		Duration("duration", time.Second),
		Durations("durations", []time.Duration{time.Minute, time.Hour}),
		Err(errors.New("test error")),
		NamedError("my error", errors.New("my error")),
		Stack("stack"),
		Skip(),
	}

	os.Exit(m.Run())
}

func TestNewLogger(t *testing.T) {
	Debug(testContext, "Debug test", testFields...)
	Info(testContext, "Info test", testFields...)
	Warn(testContext, "Warn test", testFields...)
	Error(testContext, "Error test", testFields...)
}

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Logf("panic recovered due to %v", err)
		}
	}()
	Panic(testContext, "Panic test", testFields...)
}

func TestFatal(t *testing.T) {
	Fatal(testContext, "Fatal test", testFields...)
}

func TestNewWriter(t *testing.T) {
	file, _ := ioutil.TempFile(".", "logger*.log")
	defer file.Close()

	ReplaceGlobals(New(
		WithWriter(file),
		WithLogLevel(LevelDebug),
		WithStaticFields([]Field{
			String("service", "test"),
			String("hello", "world"),
		}),
		WithDynamicFields(func(ctx context.Context) []Field {
			value, ok := ctx.Value("aabbccddKK").(string)

			return []Field{
				String("aabbccddKK", value),
				Bool("ok", ok),
			}
		}),
	))

	Debug(testContext, "Debug test", testFields...)
	Info(testContext, "Info test", testFields...)
	Warn(testContext, "Warn test", testFields...)
	Error(testContext, "Error test", testFields...)

	ctx := context.WithValue(testContext, "aabbccddKK", "665544332211")

	Debug(ctx, "Debug test", testFields...)
	Info(ctx, "Info test", testFields...)
	Warn(ctx, "Warn test", testFields...)
	Error(ctx, "Error test", testFields...)
}
