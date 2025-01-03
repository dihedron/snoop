package integer

import (
	"context"
	"log/slog"
	"testing"
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
