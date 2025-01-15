package random

import (
	"context"
	"log/slog"
	"testing"
	"time"
)

func TestSequenceGenerator(t *testing.T) {
	count := 0
	for n := range Sequence(0, 1000) {
		slog.Info("received item", "value", n)
		if count++; count >= 10 {
			break
		}
	}
}

func TestSequenceContextGenerator(t *testing.T) {
	t.Log("test with random sequence (from: 0, to: 1000, buffering: 10) and cancellation after ~10 items")
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		for n := range SequenceContext(ctx, 0, 1000) {
			slog.Info("received item", "value", n)
			time.Sleep(10 * time.Millisecond)
		}
	}()
	t.Log("test with constant sequence (from: 0, to: 1000, buffering: 10) and cancellation after ~10 items")
	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		for n := range SequenceContext(ctx, 0, 1000) {
			slog.Info("received item", "value", n)
			time.Sleep(10 * time.Millisecond)
		}
	}()
}
