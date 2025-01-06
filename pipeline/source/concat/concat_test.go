package concat

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/pipeline/source/file"
	"github.com/dihedron/snoop/pipeline/source/integer"
)

func TestConcatSequences(t *testing.T) {
	source1 := integer.New(
		integer.From(1),
		integer.Step(2),
		integer.Until(11),
	)
	source2 := integer.New(
		integer.From(0),
		integer.Step(2),
		integer.Until(10),
	)
	source := New(source1, source2)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	messages, _ := source.Emit(ctx)
	defer source.Close()
	for message := range messages {
		if message, ok := message.(*integer.Message); ok {
			slog.Debug("message received", "value", int64(*message))
			message.Ack(false)
		}
	}
	slog.Info("channel closed (no more messages), test complete")
}

func TestConcatFiles(t *testing.T) {
	source1, err := file.New("../../flow/a2m.txt")
	if err != nil {
		t.Fatal(err)
	}
	source2, err := file.New("../../flow/n2z.txt")
	if err != nil {
		t.Fatal(err)
	}
	source := New(source1, source2)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	messages, _ := source.Emit(ctx)
	defer source.Close()
	for message := range messages {
		if message, ok := message.(*file.Message); ok {
			slog.Debug("message received", "value", message.Value)
			message.Ack(false)
		}
	}
	slog.Info("channel closed (no more messages), test complete")
}

func TestConcatGenerator(t *testing.T) {
	for n := range Concat(file.File("../../flow/a2m.txt"), file.File("../../flow/n2z.txt")) {
		slog.Info("received item", "value", n)
	}
}

func TestConcatGeneratorContext(t *testing.T) {
	ctx1, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	ctx2, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	for n := range Concat(integer.SequenceContext(ctx1, 0, 0, 0), integer.SequenceContext(ctx2, 1, 0, 0)) {
		slog.Info("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}
}
