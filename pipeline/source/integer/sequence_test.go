package integer

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func TestSequence(t *testing.T) {
	stopNumber := int64(1_000)
	source := New(
		From(0),
		Step(3),
		Until(500),
	)
	messages, _ := source.Emit(context.Background())
loop:
	for message := range messages {
		if message, ok := message.(*Message); ok {
			if int64(*message) >= stopNumber {
				slog.Info("upper bound reached, cancelling...")
				break loop
			}
			slog.Debug("message received", "value", int64(*message))
			message.Ack(false)
		}
	}
	slog.Info("channel closed (no more messages), test complete")
}

func TestSequenceGenerator(t *testing.T) {
	for n := range Sequence(0, 100, 2) {
		slog.Info("received item", "value", n)
	}
}

func TestSequenceContextGenerator(t *testing.T) {
	t.Log("test with increasing sequence (from: 0, to: 100, step: 2) and cancellation after ~10 items")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for n := range SequenceContext(ctx, 0, 100, 2) {
		slog.Info("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}
	t.Log("test with constant sequence (from: 5, to: -, step: 0) and cancellation after ~10 items")
	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for n := range SequenceContext(ctx, 5, 0, 0) {
		slog.Info("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}

}
