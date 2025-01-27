package transformers

import (
	"bytes"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/generator/fibonacci"
	"github.com/dihedron/snoop/generator/file"
	"github.com/dihedron/snoop/generator/integer"
	"github.com/dihedron/snoop/generator/random"
	"github.com/dihedron/snoop/test"
	. "github.com/dihedron/snoop/transform"
)

func Log[T any](t *testing.T) F[T] {
	test.Setup(t)
	return func(value T) (T, error) {
		slog.Debug("value flowing through", "value", value, "type", format.TypeAsString(value))
		return value, nil
	}
}

func TestFibonacciChain(t *testing.T) {
	var (
		buffer      bytes.Buffer
		accumulator []int64 = []int64{}
	)
	test.Setup(t)

	stopwatch := &StopWatch[int64, int64]{}
	counter := &Counter[int64]{}

	transform := Apply(
		stopwatch.Start(),
		Then(
			Log[int64](t),
			Then(
				Delay[int64](50*time.Millisecond),
				Then(
					Record[int64](&buffer, "%d\n", true),
					Then(
						counter.Add(),
						Then(
							Accumulate(&accumulator),
							stopwatch.Stop(),
						),
					),
				),
			),
		),
	)
	for value := range fibonacci.Series(1_000_000) {
		if _, err := transform(value); err != nil {
			break
		}
	}
	slog.Info("final result", "elapsed", stopwatch.Elapsed().String(), "items", counter.Count(), "accumulator", accumulator, "buffer", buffer.String())
}

func TestRandomChain(t *testing.T) {
	var (
		buffer      bytes.Buffer
		accumulator []int64 = []int64{}
	)
	test.Setup(t)

	stopwatch := &StopWatch[int64, int64]{}
	counter := &Counter[int64]{}

	transform := Apply(
		stopwatch.Start(),
		Then(
			Log[int64](t),
			Then(
				Delay[int64](50*time.Millisecond),
				Then(
					Record[int64](&buffer, "%d\n", true),
					Then(
						counter.Add(),
						Then(
							Accumulate(&accumulator),
							stopwatch.Stop(),
						),
					),
				),
			),
		),
	)

	for value := range random.Sequence(0, 1_000) {
		if _, err := transform(value); err != nil || counter.Count() >= 10 {
			break
		}
	}
	slog.Info("final result", "elapsed", stopwatch.Elapsed().String(), "items", counter.Count(), "accumulator", accumulator, "buffer", buffer.String())
}

func TestFileChain(t *testing.T) {
	var (
		buffer      bytes.Buffer
		accumulator []string = []string{}
	)
	test.Setup(t)

	stopwatch := &StopWatch[string, string]{}
	counter := &Counter[string]{}

	transform := Apply[string, string](
		stopwatch.Start(),
		Then(
			Log[string](t),
			Then(
				Delay[string](50*time.Millisecond),
				Then(
					Record[string](&buffer, "%s\n", true),
					Then[string, string](
						counter.Add(),
						Then(
							Accumulate(&accumulator),
							stopwatch.Stop(),
						),
					),
				),
			),
		),
	)
	for value := range file.Lines("../generator/file/test.txt") {
		if _, err := transform(value); err != nil || counter.Count() >= 10 {
			break
		}
	}
	slog.Info("final result", "elapsed", stopwatch.Elapsed().String(), "items", counter.Count(), "accumulator", accumulator, "buffer", buffer.String())
}

func TestSequenceWithSkipOddChain(t *testing.T) {
	var (
		buffer      bytes.Buffer
		accumulator []int64 = []int64{}
	)
	test.Setup(t)

	stopwatch := &StopWatch[int64, int64]{}
	counter := &Counter[int64]{}

	transform := Apply(
		stopwatch.Start(),
		Then(
			Log[int64](t),
			Then(
				Delay[int64](50*time.Millisecond),
				Then(
					counter.Add(),
					Then(
						Filter(func(value int64) bool { return value%2 == 0 }),
						Then(
							Record[int64](&buffer, "%d\n", true),
							Then(
								Accumulate(&accumulator),
								stopwatch.Stop(),
							),
						),
					),
				),
			),
		),
	)
	for value := range integer.Sequence(0, 1_000, 1) {
		if _, err := transform(value); err != nil && err != Drop || counter.Count() >= 10 {
			break
		}
	}
	slog.Info("final result", "elapsed", stopwatch.Elapsed().String(), "items", counter.Count(), "accumulator", accumulator, "buffer", buffer.String())
}
