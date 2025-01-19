package pipeline

import (
	"bytes"
	"testing"
	"time"

	"github.com/dihedron/snoop/pipeline/filter/counter"
	"github.com/dihedron/snoop/pipeline/filter/profiler"
	"github.com/dihedron/snoop/pipeline/filter/recorder"
	"github.com/dihedron/snoop/pipeline/filter/throttler"
	"github.com/dihedron/snoop/pipeline/source/file"
)

/*
type TestLogFilter[T any] struct {
	t *testing.T
}

func (f *TestLogFilter[T]) Apply(value T) (T, error) {
	f.t.Logf("value flowing through: %v (type: %T)\n", value, value)
	return value, nil
}

func TestFibonacciFlow(t *testing.T) {
	var buffer bytes.Buffer
	profiler := profiler.New[int64]()
	counter := counter.New[int64]()
	pipeline := New[int64](
		profiler,
		&TestLogFilter[int64]{t: t},
		throttler.New[int64](50*time.Millisecond),
		recorder.New[int64](&buffer, "", true),
		profiler,
		counter,
	)
	defer pipeline.Close()

	for value := range fibonacci.Series(1_000_000) {
		value, err := pipeline.Apply(value)
		t.Logf("result: %v (took %s, with err %v)", value, profiler.Elapsed(), err)
	}
	t.Logf("final result: items %d, value:\n%s", counter.Count(), buffer.String())
}

func TestRandomFlow(t *testing.T) {
	var buffer bytes.Buffer
	profiler := profiler.New[int64]()
	counter := counter.New[int64]()
	pipeline := New[int64](
		profiler,
		&TestLogFilter[int64]{t: t},
		throttler.New[int64](50*time.Millisecond),
		recorder.New[int64](&buffer, "", true),
		profiler,
		counter,
	)
	defer pipeline.Close()

	for value := range random.Sequence(0, 1_000) {
		value, err := pipeline.Apply(value)
		t.Logf("result: %v (took %s, with err %v)", value, profiler.Elapsed(), err)
		if counter.Count() >= 10 {
			break
		}
	}
	t.Logf("final result: items %d, value:\n%s", counter.Count(), buffer.String())
}

/*
type AcknowledgeableImpl[T any] struct {
	value T
}

func (a AcknowledgeableImpl[T]) Ack(multiple bool) error {
	return nil
}

func Wrap[T any](seq iter.Seq[T]) iter.Seq[AcknowledgeableImpl[T]] {
	return func(yield func(AcknowledgeableImpl[T]) bool) {
		for v := range seq {
			if !yield(AcknowledgeableImpl[T]{value: v}) {
				return
			}
		}
	}
}
*/

func TestFileFlow(t *testing.T) {
	var buffer bytes.Buffer
	profiler := profiler.New[string]()
	counter := counter.New[string]()
	pipeline := New[string](
		profiler,
		&TestLogFilter[string]{t: t},
		throttler.New[string](50*time.Millisecond),
		recorder.New[string](&buffer, "", true),

		profiler,
		counter,
	)
	defer pipeline.Close()

	for value := range file.File("./source/file/test.txt") {
		value, err := pipeline.Apply(value)
		t.Logf("result: %v (took %s, with err %v)", value, profiler.Elapsed(), err)
		if counter.Count() >= 10 {
			break
		}
	}
	t.Logf("final result: items %d, value:\n%s", counter.Count(), buffer.String())

}

/*
func TestIntegerSequenceWithSkippedMessages(t *testing.T) {
	ctx := context.Background()
	accumulator := &Int64Accumulator{}
	pipeline := New(
		From(integer.SequenceContext(ctx, 0, 100, 1)),
		Through[int64](
			accumulator,
		),
		Into[int64](&sink.Null{}),
	)
	defer pipeline.Close()
	pipeline.Execute()
	slog.Debug("results received", "count", len(accumulator.values), "value", accumulator.values)
}

type Int64Accumulator struct {
	values []int64
}

func (a *Int64Accumulator) Name() string {
	return "Int64Accumulator"
}

func (a *Int64Accumulator) Process(message any) (any, error) {
	slog.Debug("processing message")
	if value, ok := message.(int64); ok {
		if value%2 == 0 {
			slog.Info("even value, forwarding...", "value", value)
			return value, nil
		} else {
			slog.Info("odd value: adding to skipped values in accumulator...", "value", value)
			a.values = append(a.values, value)
			return nil, pipeline.ErrSkip
		}
	}
	return nil, fmt.Errorf("invalid message type: expected int64, got %T", message)
}

func TestPipelineWithCounterSource(t *testing.T) {
	source := integer.New(integer.From(0), integer.Step(1), integer.Until(100))
	var buffer bytes.Buffer
	counter := counter.New()
	accumulator := &ListAccumulatorSink{
		values: []pipeline.Message{},
	}
	pipeline := New(
		From(source),
		Through[pipeline.Message](
			&TestLogFilter{t: t},
			throttler.New(10*time.Millisecond),
			recorder.New(&buffer, true),
			counter,
		),
		Into(accumulator),
	)
	ctx := context.Background()
	err := pipeline.Execute(ctx)
	if err != nil {
		t.Fatalf("error executing pipeline: %v", err)
	} else {
		slog.Debug("results received", "count", counter.Count(), "value", buffer.String())
		values := accumulator.GetValues()
		slog.Debug("results accumulated", "count", len(values), "values", values)
	}
}

type ListAccumulatorSink struct {
	values []pipeline.Message
}

func (s *ListAccumulatorSink) Collect(ctx context.Context, message pipeline.Message) error {
	slog.Debug("absorbing message", "message", message)
	if s.values == nil {
		s.values = []pipeline.Message{}
	}
	s.values = append(s.values, message)
	message.Ack(true)
	return nil
}

func (s *ListAccumulatorSink) GetValues() []pipeline.Message {
	return s.values
}
*/
