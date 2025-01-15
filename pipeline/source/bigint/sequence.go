package bigint

import (
	"context"
	"iter"
	"math/big"
)

// Sequence uses the new Go 1.23 style generator to generate a sequence
// of integers; it may be configured with functional options. In order
// to create an infinite sequence provide no end value; in order to create
// a Sequence repeating the same value over and over, set Step to 0 and
// Start as the desired value.
func Sequence(from *big.Int, to *big.Int, step *big.Int) iter.Seq[big.Int] {
	return SequenceContext(context.Background(), from, to, step)
}

// Sequence uses the new Go 1.23 style generator to generate a sequence
// of integers; it may be configured with functional options. In order
// to create an infinite sequence provide no end value; in order to create
// a Sequence repeating the same value over and over, set Step to 0 and
// Start as the desired value. When the given context is cancelled, the
// generator stops.
func SequenceContext(ctx context.Context, from *big.Int, to *big.Int, step *big.Int) iter.Seq[big.Int] {
	return func(yield func(big.Int) bool) {
		zero := new(big.Int).SetInt64(0)
		value := from
		for {
			select {
			case <-ctx.Done():
				// slog.Info("cancelling...", "from", settings.start, "to", settings.end, "step", settings.step)
				return
			default:
				if !yield(*value) {
					return
				}
				value.Add(value, step)
				if to.Cmp(zero) > 0 && value.Cmp(to) > 0 {
					return
				}
			}
		}
	}
}
