package transformer

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
)

func Log[T any](t *testing.T) Filter[T] {
	test.Setup(t)
	return func(value T) (T, error) {
		slog.Debug("value flowing through", "value", value, "type", format.TypeAsString(value))
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
	chain := Apply(
		Profile[int64](&start, nil),
		Then(
			Log[int64](t),
			Then(
				Delay[int64](50*time.Millisecond),
				Then(
					Record[int64](&buffer, "%d\n", true),
					Then(
						Count[int64](&count),
						Then(
							Accumulate(&accumulator),
							Profile[int64](&start, &elapsed),
						),
					),
				),
			),
		),
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
	chain := Apply(
		Profile[int64](&start, nil),
		Then(
			Log[int64](t),
			Then(
				Delay[int64](50*time.Millisecond),
				Then(
					Record[int64](&buffer, "%d\n", true),
					Then(
						Count[int64](&count),
						Then(
							Accumulate(&accumulator),
							Profile[int64](&start, &elapsed),
						),
					),
				),
			),
		),
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
	chain := Apply(
		Profile[string](&start, nil),
		Then(
			Log[string](t),
			Then(
				Delay[string](50*time.Millisecond),
				Then(
					Record[string](&buffer, "%s\n", true),
					Then(
						Count[string](&count),
						Then(
							Accumulate(&accumulator),
							Profile[string](&start, &elapsed),
						),
					),
				),
			),
		),
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
	chain := Apply(
		Profile[int64](&start, nil),
		Then(
			Log[int64](t),
			Then(
				Delay[int64](50*time.Millisecond),
				Then(
					Count[int64](&count),
					Then(
						Accept[int64](func(value int64) bool { return value%2 == 0 }),
						Then(
							Record[int64](&buffer, "%d\n", true),
							Then(
								Accumulate(&accumulator),
								Profile[int64](&start, &elapsed),
							),
						),
					),
				),
			),
		),
	)

	for value := range integer.Sequence(0, 1_000, 1) {
		if _, err := chain(value); err != nil || count >= 10 {
			break
		}
	}
	slog.Info("final result", "items", count, "accumulator", accumulator, "buffer", buffer.String())
}
