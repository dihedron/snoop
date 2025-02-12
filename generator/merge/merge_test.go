package merge

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/generator/file"
	"github.com/dihedron/snoop/generator/integer"
	"github.com/dihedron/snoop/test"
)

func TestMergeContextGenerator(t *testing.T) {
	test.Setup(t)
	slog.Info("test with 3 alternating sequences of 0, 1 and 2 and cancellation after ~20 items")
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		for n := range Merge(ctx, integer.SequenceContext(ctx, 0, 0, 0), integer.SequenceContext(ctx, 1, 0, 0), integer.SequenceContext(ctx, 2, 0, 0)) {
			slog.Debug("received item", "value", n)
			time.Sleep(10 * time.Millisecond)
		}
	}()
	slog.Info("test with 2 files read line by line and cancellation after ~20 items")
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		files := file.New()
		for value := range Merge(ctx, files.AllLinesContext(ctx, "../file/a2m.txt"), files.AllLinesContext(ctx, "../file/n2z.txt")) {
			slog.Debug("received item", "value", value)
			time.Sleep(10 * time.Millisecond)
		}
	}()
}

func TestMerge2ContextGenerator(t *testing.T) {
	// TODO
}
