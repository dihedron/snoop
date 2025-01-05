package random

import (
	"bytes"
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/pipeline/filter/recorder"
	"github.com/dihedron/snoop/pipeline/filter/throttler"
	"github.com/dihedron/snoop/pipeline/flow"
	"github.com/dihedron/snoop/pipeline/sink"
)

const (
	MinValue int64 = 0
	MaxValue int64 = 1_000
)

func TestRandom(t *testing.T) {
	source, err := New(MinValue, MaxValue, 10)
	if err != nil {
		t.Fail()
	}
	var buffer bytes.Buffer
	f := flow.New(
		flow.From(source),
		flow.Through(
			throttler.New(100*time.Millisecond),
			recorder.New(&buffer, true),
		),
		flow.Into(&sink.Null{}),
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = f.Execute(ctx)
	if err != nil {
		slog.Error("error received", "error", err)
	}
	defer f.Close()
	// loop:
	// 	for {
	// 		select {
	// 		case result, ok := <-results:
	// 			if ok {
	// 				slog.Debug("result retrieved", "result", result)
	// 				result.Ack(false)
	// 			} else {
	// 				slog.Debug("pipeline cancelled")
	// 				break loop
	// 			}
	// 		case <-ctx.Done():
	// 			slog.Debug("pipeline context cancelled")
	// 			break loop
	// 		}
	// 	}
	slog.Debug("result", "value", buffer.String())
	slog.Info("test complete")
}

func TestRandomGenerator(t *testing.T) {
	count := 0
	for n := range Random(0, 1000) {
		slog.Info("received item", "value", n)
		if count++; count >= 10 {
			break
		}
	}
}

func TestRandomContextGenerator(t *testing.T) {
	t.Log("test with random sequence (from: 0, to: 1000, buffering: 10) and cancellation after ~10 items")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for n := range RandomContext(ctx, 0, 1000) {
		slog.Info("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}
	t.Log("test with constant sequence (from: 0, to: 1000, buffering: 10) and cancellation after ~10 items")
	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for n := range RandomContext(ctx, 0, 1000) {
		slog.Info("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}

}
