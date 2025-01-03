package file

import (
	"context"
	"log/slog"
	"testing"
)

const MaxIter = 15

func TestFile(t *testing.T) {
	ctx := context.Background()
	source, err := New("../../flow/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	messages, _ := source.Emit(ctx)
	defer source.Close()
	n := 0
loop:
	for message := range messages {
		if message, ok := message.(*Message); ok {
			if n >= MaxIter {
				slog.Debug("cancelling...")
				break loop
			}
			n++
			slog.Debug("received message", "value", message.Value)
			message.Ack(false)
		}
	}
	slog.Info("done reading file, test complete")
}
