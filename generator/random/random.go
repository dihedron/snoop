package random

import (
	"context"
	"crypto/rand"
	"iter"
	"log/slog"
	"math/big"
)

// Sequence uses the new Go 1.23 style generator to generate a sequence
// of random integers.
func Sequence(min int64, max int64) iter.Seq[int64] {
	return SequenceContext(context.Background(), min, max)
}

// SequenceContext uses the new Go 1.23 style generator to generate a sequence
// of random integers.
func SequenceContext(ctx context.Context, min int64, max int64) iter.Seq[int64] {
	return func(yield func(int64) bool) {
		for {
			select {
			case <-ctx.Done():
				slog.Info("context cancelled")
				return
			default:
				value, err := rand.Int(rand.Reader, big.NewInt(max-min))
				if err != nil {
					slog.Error("error retrieving random value", "error", err)
					return
				}
				slog.Debug("sending sequence number as message", "value", value)
				if !yield(value.Int64()) {
					return
				}
			}
		}
	}
}
