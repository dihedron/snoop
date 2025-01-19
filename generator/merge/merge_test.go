package merge

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/generator/file"
	"github.com/dihedron/snoop/generator/integer"
)

func TestMergeContextGenerator(t *testing.T) {
	t.Log("test with 3 alternating sequences of 0, 1 and 2 and cancellation after ~20 items")
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		for n := range Merge(ctx, integer.SequenceContext(ctx, 0, 0, 0), integer.SequenceContext(ctx, 1, 0, 0), integer.SequenceContext(ctx, 2, 0, 0)) {
			slog.Info("received item", "value", n)
			time.Sleep(10 * time.Millisecond)
		}
	}()
	t.Log("test with 2 files read line by line and cancellation after ~20 items")
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		for value := range Merge(ctx, file.LinesContext(ctx, "../file/a2m.txt"), file.LinesContext(ctx, "../file/n2z.txt")) {
			slog.Info("received item", "value", value)
			time.Sleep(10 * time.Millisecond)
		}
	}()
}

func TestMerge2ContextGenerator(t *testing.T) {
	// TODO
}
