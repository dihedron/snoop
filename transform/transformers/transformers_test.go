package transformers

import (
	"bytes"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/generator/concat"
	"github.com/dihedron/snoop/generator/fibonacci"
	"github.com/dihedron/snoop/generator/file"
	"github.com/dihedron/snoop/generator/integer"
	"github.com/dihedron/snoop/generator/random"
	"github.com/dihedron/snoop/test"
	"github.com/dihedron/snoop/transform/chain"
)

func Log[T any](t *testing.T) chain.F[T] {
	test.Setup(t)
	return func(value T) (T, error) {
		slog.Debug("value flowing through", "value", value, "type", format.TypeAsString(value))
		return value, nil
	}
}

func TestFibonacciChain(t *testing.T) {
	test.Setup(t)
	var buffer bytes.Buffer

	stopwatch := &StopWatch[int64, int64]{}
	counter := &Counter[int64]{}
	accumulator := &Accumulator[int64]{}

	transform := chain.Of7(
		stopwatch.Start(),
		Log[int64](t),
		Delay[int64](50*time.Millisecond),
		Writef[int64](&buffer, "%d\n", true),
		counter.Add(),
		accumulator.Add(),
		stopwatch.Stop(),
	)
	for value := range fibonacci.Series(1_000_000) {
		if _, err := transform(value); err != nil {
			break
		}
	}
	slog.Info("final result", "elapsed", stopwatch.Elapsed().String(), "items", counter.Count(), "accumulator", accumulator, "buffer", buffer.String())
}

func TestRandomChain(t *testing.T) {
	test.Setup(t)

	var buffer bytes.Buffer

	stopwatch := &StopWatch[int64, int64]{}
	counter := &Counter[int64]{}
	accumulator := &Accumulator[int64]{}

	transform := chain.Of7(
		stopwatch.Start(),
		Log[int64](t),
		Delay[int64](50*time.Millisecond),
		Writef[int64](&buffer, "%d\n", true),
		counter.Add(),
		accumulator.Add(),
		stopwatch.Stop(),
	)

	for value := range random.Sequence(0, 1_000) {
		if _, err := transform(value); err != nil || counter.Count() >= 10 {
			break
		}
	}
	slog.Info("final result", "elapsed", stopwatch.Elapsed().String(), "items", counter.Count(), "accumulator", accumulator, "buffer", buffer.String())
}

func TestFileChain(t *testing.T) {
	test.Setup(t)

	var buffer bytes.Buffer

	stopwatch := &StopWatch[string, string]{}
	counter := &Counter[string]{}
	accumulator := &Accumulator[string]{}

	transform := chain.Of7(
		stopwatch.Start(),
		Log[string](t),
		Delay[string](50*time.Millisecond),
		Writef[string](&buffer, "%s\n", true),
		counter.Add(),
		accumulator.Add(),
		stopwatch.Stop(),
	)

	files := file.New()
	for value := range files.AllLines("../generator/file/test.txt") {
		if _, err := transform(value); err != nil || counter.Count() >= 10 {
			break
		}
	}
	slog.Info("final result", "elapsed", stopwatch.Elapsed().String(), "items", counter.Count(), "accumulator", accumulator, "buffer", buffer.String())
}

func TestSequenceWithSkipOddChain(t *testing.T) {
	test.Setup(t)
	var buffer bytes.Buffer
	stopwatch := &StopWatch[int64, int64]{}
	counter := &Counter[int64]{}
	accumulator := &Accumulator[int64]{}

	transform := chain.Of8(
		stopwatch.Start(),
		Log[int64](t),
		Delay[int64](50*time.Millisecond),
		counter.Add(),
		AcceptIf(func(value int64) bool { return value%2 == 0 }),
		Writef[int64](&buffer, "%d\n", true),
		accumulator.Add(),
		stopwatch.Stop(),
	)
	for value := range integer.Sequence(0, 1_000, 1) {
		if _, err := transform(value); err != nil && err != chain.Drop || counter.Count() >= 10 {
			break
		}
	}
	slog.Info("final result", "elapsed", stopwatch.Elapsed().String(), "items", counter.Count(), "accumulator", accumulator, "buffer", buffer.String())
}

func TestCatenate(t *testing.T) {
	test.Setup(t)

	var buffer bytes.Buffer

	stopwatch := &StopWatch[int64, string]{}
	counter := &Counter[int64]{}
	catenator := &Catenator[string]{Join: ", "}

	transform := chain.Of8(
		stopwatch.Start(),
		Log[int64](t),
		Delay[int64](50*time.Millisecond),
		Writef[int64](&buffer, "%d\n", true),
		counter.Add(),
		ToString[int64](),
		catenator.Add(),
		stopwatch.Stop(),
	)

	for value := range random.Sequence(0, 1_000) {
		if _, err := transform(value); err != nil || counter.Count() >= 10 {
			break
		}
	}
	slog.Info("final result", "elapsed", stopwatch.Elapsed().String(), "items", counter.Count(), "catenator", catenator.Value(), "buffer", buffer.String())
}

func TestCacheFromFile(t *testing.T) {
	test.Setup(t)

	var buffer bytes.Buffer

	stopwatch := &StopWatch[string, string]{}
	counter := &Counter[string]{}
	cache := &Cache[string, string]{}

	transform := chain.Of7(
		stopwatch.Start(),
		Log[string](t),
		Delay[string](50*time.Millisecond),
		Writef[string](&buffer, "%s\n", true),
		counter.Add(),
		cache.Set(func(s string) string { return s[:1] }),
		stopwatch.Stop(),
	)
	files := file.New()
	for value := range files.AllLines("../generator/file/test.txt") {
		if _, err := transform(value); err != nil || counter.Count() >= 10 {
			break
		}
	}
	slog.Info("final result", "elapsed", stopwatch.Elapsed().String(), "items", counter.Count(), "cache", cache, "buffer", buffer.String())
}

func TestMultipleCacheFromMultipleFiles(t *testing.T) {
	test.Setup(t)

	var buffer bytes.Buffer

	stopwatch := &StopWatch[string, string]{}
	counter := &Counter[string]{}
	multicache := &MultiCache[string, string]{}

	transform := chain.Of7(
		stopwatch.Start(),
		Log[string](t),
		Delay[string](50*time.Millisecond),
		Writef[string](&buffer, "%s\n", true),
		counter.Add(),
		multicache.Set(func(s string) string { return s[:1] }),
		stopwatch.Stop(),
	)

	files := file.New()
	for value := range concat.Concat(files.AllLines("../generator/file/test.txt"), files.AllLines("../generator/file/test.txt")) {
		if _, err := transform(value); err != nil {
			break
		}
	}
	slog.Info("final result", "elapsed", stopwatch.Elapsed().String(), "items", counter.Count(), "multicache", multicache, "buffer", buffer.String())
}
