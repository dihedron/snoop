package merge

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/pipeline/source/integer"
)

func TestMerge(t *testing.T) {
	stopNumber := int64(1_000)
	source1 := integer.New(
		integer.Until(500),
	)
	source2 := integer.New(
		integer.Until(500),
	)
	source := New(source1, source2)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	messages, _ := source.Emit(ctx)
	defer source.Close()
loop:
	for message := range messages {
		if message, ok := message.(*integer.Message); ok {
			if int64(*message) >= stopNumber {
				slog.Info("upper bound reached, cancelling...")
				cancel()
				break loop
			}
			slog.Debug("message received", "value", int64(*message))
			message.Ack(false)
		}
	}
	slog.Debug("no more messages, test complete")
}
