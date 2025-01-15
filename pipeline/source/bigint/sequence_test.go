package bigint

import (
	"context"
	"log/slog"
	"math/big"
	"testing"
	"time"
)

func TestSequenceContextGenerator(t *testing.T) {
	t.Log("test with increasing sequence (from: 0, to: 100, step: 2) and cancellation after ~10 items")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for n := range SequenceContext(ctx, new(big.Int).SetInt64(0), new(big.Int).SetInt64(100), new(big.Int).SetInt64(2)) {
		slog.Info("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}
	t.Log("test with constant sequence (from: 5, to: -, step: 0) and cancellation after ~10 items")
	ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for n := range SequenceContext(ctx, new(big.Int).SetInt64(5), new(big.Int).SetInt64(0), new(big.Int).SetInt64(0)) {
		slog.Info("received item", "value", n)
		time.Sleep(10 * time.Millisecond)
	}

}
