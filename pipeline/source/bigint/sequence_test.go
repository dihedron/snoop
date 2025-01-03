package bigint

import (
	"context"
	"log/slog"
	"math/big"
	"testing"
)

func TestSequence(t *testing.T) {
	stopNumber := big.NewInt(1_000)
	source := New(
		From(big.NewInt(0)),
		Step(big.NewInt(3)),
		Until(big.NewInt(500)),
	)
	messages, _ := source.Emit(context.Background())
loop:
	for message := range messages {
		if message, ok := message.(*Message); ok {
			if message.Value.Cmp(stopNumber) >= 0 {
				slog.Info("upper bound reached, cancelling...")
				break loop
			}
			slog.Debug("message received", "value", message.Value)
			message.Ack(false)
		}
	}
	slog.Info("channel closed (no more messages), test complete")
}
