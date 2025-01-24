package processor

import (
	"bytes"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/generator/fibonacci"
	"github.com/dihedron/snoop/generator/file"
	"github.com/dihedron/snoop/generator/integer"
	"github.com/dihedron/snoop/generator/random"
	"github.com/dihedron/snoop/test"
)

func Log[T any](t *testing.T) Handler[T] {
	test.Setup(t)
	return func(value T) (T, error) {
		t.Logf("value flowing through: %v (type: %T)\n", value, value)
		return value, nil
	}
}

func TestFibonacciChain(t *testing.T) {
	var (
		buffer      bytes.Buffer
		start       time.Time
		elapsed     time.Duration
		count       int64
		accumulator []int64 = []int64{}
	)
	test.Setup(t)
	chain := Chain(
		Profile[int64](&start, nil),
		Log[int64](t),
		Delay[int64](50*time.Millisecond),
		Record[int64](&buffer, "%d\n", true),
		Count[int64](&count),
		Accumulate(&accumulator),
		Profile[int64](&start, &elapsed),
	)

	for value := range fibonacci.Series(1_000_000) {
		if _, err := chain(value); err != nil {
			break
		}
	}
	slog.Info("final result", "items", count, "accumulator", accumulator, "buffer", buffer.String())
}

func TestRandomChain(t *testing.T) {
	var (
		buffer      bytes.Buffer
		start       time.Time
		elapsed     time.Duration
		count       int64
		accumulator []int64 = []int64{}
	)
	test.Setup(t)
	chain := Chain(
		Profile[int64](&start, nil),
		Log[int64](t),
		Delay[int64](50*time.Millisecond),
		Record[int64](&buffer, "%d\n", true),
		Count[int64](&count),
		Accumulate(&accumulator),
		Profile[int64](&start, &elapsed),
	)

	for value := range random.Sequence(0, 1_000) {
		if _, err := chain(value); err != nil || count >= 10 {
			break
		}
	}
	slog.Info("final result", "items", count, "accumulator", accumulator, "buffer", buffer.String())
}

func TestFileChain(t *testing.T) {
	var (
		buffer      bytes.Buffer
		start       time.Time
		elapsed     time.Duration
		count       int64
		accumulator []string = []string{}
	)
	test.Setup(t)
	chain := Chain(
		Profile[string](&start, nil),
		Log[string](t),
		Delay[string](50*time.Millisecond),
		Record[string](&buffer, "%s\n", true),
		Count[string](&count),
		Accumulate(&accumulator),
		Profile[string](&start, &elapsed),
	)

	for value := range file.Lines("../generator/file/test.txt") {
		if _, err := chain(value); err != nil || count >= 10 {
			break
		}
	}
	slog.Info("final result", "items", count, "accumulator", accumulator, "buffer", buffer.String())
}

func TestSequenceWithSkipOddChain(t *testing.T) {
	var (
		buffer      bytes.Buffer
		start       time.Time
		elapsed     time.Duration
		count       int64
		accumulator []int64 = []int64{}
	)
	test.Setup(t)
	chain := Chain(
		Profile[int64](&start, nil),
		Log[int64](t),
		Delay[int64](50*time.Millisecond),
		Count[int64](&count),
		Accept[int64](func(value int64) bool { return value%2 == 0 }),
		Record[int64](&buffer, "%d\n", true),
		Accumulate(&accumulator),
		Profile[int64](&start, &elapsed),
	)

	for value := range integer.Sequence(0, 1_000, 1) {
		if _, err := chain(value); err != nil || count >= 10 {
			break
		}
	}
	slog.Info("final result", "items", count, "accumulator", accumulator, "buffer", buffer.String())
}
