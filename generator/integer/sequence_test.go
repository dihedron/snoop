package integer

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/dihedron/snoop/test"
)

func TestSequenceGenerator(t *testing.T) {
	test.Setup(t, test.Text)
	for n := range Sequence(0, 100, 2) {
		slog.Debug("received item", "value", n)
	}
}

func TestSequenceContextGenerator(t *testing.T) {
	test.Setup(t, test.Text)
	slog.Info("test with increasing sequence (from: 0, to: 100, step: 2) and cancellation after ~10 items")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for n := range SequenceContext(ctx, 0, 100, 2) {
		slog.Debug("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}
	slog.Info("test with constant sequence (from: 5, to: -, step: 0) and cancellation after ~10 items")
	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for n := range SequenceContext(ctx, 5, 0, 0) {
		slog.Debug("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}

}
