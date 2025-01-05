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

func TestMergeContextGenerator(t *testing.T) {
	t.Log("test with 3 alternating sequences of 0, 1 and 2 and cancellation after ~30 items")
	ctx, cancel0 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel0()
	for n := range Merge(ctx, integer.SequenceContext(ctx, integer.From(0), integer.Step(0)), integer.SequenceContext(ctx, integer.From(1), integer.Step(0)), integer.SequenceContext(ctx, integer.From(2), integer.Step(0))) {
		t.Logf("received item %d", n)
		time.Sleep(10 * time.Millisecond)
	}
}

func TestMerge2ContextGenerator(t *testing.T) {
	t.Log("test with 3 alternating sequences of 0, 1 and 2 and cancellation after ~30 items")
	// ctx, cancel0 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	// defer cancel0()
	// for n := range Merge2(ctx, integer.SequenceContext(ctx, integer.From(0), integer.Step(0)), integer.SequenceContext(ctx, integer.From(1), integer.Step(0)), integer.SequenceContext(ctx, integer.From(2), integer.Step(0))) {
	// 	t.Logf("received item %d", n)
	// 	time.Sleep(10 * time.Millisecond)
	// }
	// TODO!!!
}
