package file

import (
	"context"
	"log/slog"
	"testing"
	"time"
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

func TestFileGenerator(t *testing.T) {
	t.Log("test with no file")
	for n, err := range File() {
		slog.Info("received item", "value", n, "error", err)
	}
	t.Log("test with one file")
	for n, err := range File("../../flow/test.txt") {
		slog.Info("received item", "value", n, "error", err)
	}
	t.Log("test with two files")
	for n, err := range File("../../flow/test.txt", "../../flow/test.txt") {
		slog.Info("received item", "value", n, "error", err)
	}
	t.Log("test with non-existing file")
	for n, err := range File("../../flow/non_existing.txt") {
		slog.Info("received item", "value", n, "error", err)
	}
}

func TestFileContextGenerator(t *testing.T) {
	t.Log("test with cancellation after 10 items")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for n, err := range FileContext(ctx, "../../flow/test.txt") {
		slog.Info("received item", "value", n, "error", err)
		time.Sleep(10 * time.Millisecond)
	}
}
