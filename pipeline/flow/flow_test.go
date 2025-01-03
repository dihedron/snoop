package flow

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/pipeline"
	"github.com/dihedron/snoop/pipeline/filter/counter"
	"github.com/dihedron/snoop/pipeline/filter/profiler"
	"github.com/dihedron/snoop/pipeline/filter/recorder"
	"github.com/dihedron/snoop/pipeline/filter/throttler"
	"github.com/dihedron/snoop/pipeline/sink"
	"github.com/dihedron/snoop/pipeline/source/fibonacci"
	"github.com/dihedron/snoop/pipeline/source/file"
	"github.com/dihedron/snoop/pipeline/source/integer"
	"github.com/dihedron/snoop/pipeline/source/random"
)

type TestLogFilter struct {
	t *testing.T
}

func (f *TestLogFilter) Name() string {
	return "github.com/dihedron/snoop/pipeline/flow/TestLogFilter"
}

func (f *TestLogFilter) Process(ctx context.Context, message pipeline.Message) (context.Context, pipeline.Message, error) {
	f.t.Logf("message flowing through: %v\n", message)
	return ctx, message, nil
}

func TestFibonacciEngineStartStop(t *testing.T) {
	source, err := fibonacci.New()
	if err != nil {
		slog.Error("error creating fibonacci source", "error", err)
	}
	var buffer bytes.Buffer
	counter := counter.New()
	profiler := profiler.New()
	pipeline := New(
		From(source),
		Through(
			profiler,
			&TestLogFilter{t: t},
			throttler.New(50*time.Millisecond),
			recorder.New(&buffer, true),
			profiler,
			counter,
		),
		Into(&sink.Null{}),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pipeline.Execute(ctx)
	defer pipeline.Close()
	slog.Debug("final result", "items", counter.Count(), "each", profiler.Elapsed(), "value", buffer.String())
}

func TestRandomEngineStartStop(t *testing.T) {
	source, err := random.New(0, 1_000, 10)
	if err != nil {
		slog.Error("error creating random source", "error", err)
	}
	var buffer bytes.Buffer
	counter := counter.New()
	pipeline := New(
		From(source),
		Through(
			&TestLogFilter{t: t},
			throttler.New(50*time.Millisecond),
			recorder.New(&buffer, true),
			counter,
		),
		Into(&sink.Null{}),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pipeline.Execute(ctx)
	defer pipeline.Close()
	slog.Debug("final result", "count", counter.Count(), "value", buffer.String())
}

func TestFileEngineStartStop(t *testing.T) {
	source, err := file.New("./test.txt")
	if err != nil {
		slog.Error("error creating file source", "error", err)
	}
	var buffer bytes.Buffer
	counter := counter.New()
	pipeline := New(
		From(source),
		Through(
			&TestLogFilter{t: t},
			throttler.New(50*time.Millisecond),
			recorder.New(&buffer, true),
			counter,
		),
		Into(&sink.Null{}),
	)
	ctx := context.Background()
	pipeline.Execute(ctx)
	defer pipeline.Close()
	slog.Debug("results received", "count", counter.Count(), "value", buffer.String())
}

func TestEngineStartStopWithSkippedMessages(t *testing.T) {
	source := integer.New(
		integer.From(0),
		integer.Step(1),
		integer.Until(100))
	accumulator := &Int64Accumulator{}
	pipeline := New(
		From(source),
		Through(
			accumulator,
		),
		Into(&sink.Null{}),
	)
	ctx := context.Background()
	pipeline.Execute(ctx)
	defer pipeline.Close()
	slog.Debug("results received", "count", len(accumulator.values), "value", accumulator.values)
}

type Int64Accumulator struct {
	values []int64
}

func (a *Int64Accumulator) Name() string {
	return "Int64Accumulator"
}

func (a *Int64Accumulator) Process(ctx context.Context, message pipeline.Message) (context.Context, pipeline.Message, error) {
	select {
	case <-ctx.Done():
		slog.Debug("context cancelled")
		return ctx, message, pipeline.ErrAbort
	default:
		slog.Debug("processing message")
		if message, ok := message.(*integer.Message); ok {
			value := int64(*message)
			if value%2 == 0 {
				slog.Info("even value, forwarding...", "value", value)
				return ctx, message, nil
			} else {
				slog.Info("odd value: adding to skipped values in accumulator...", "value", value)
				a.values = append(a.values, value)
				message.Ack(true)
				return ctx, nil, pipeline.ErrSkip
			}
		}
	}
	return ctx, nil, fmt.Errorf("invalid message type: expected *int64, got %T", message)
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
		Through(
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
