package integer

import (
	"context"
	"iter"
)

// Sequence uses the new Go 1.23 style generator to generate a sequence
// of integers; it may be configured with functional options. In order
// to create an infinite sequence provide no end value; in order to create
// a Sequence repeating the same value over and over, set Step to 0 and
// Start as the desired value.
func Sequence(from int64, to int64, step int64) iter.Seq[int64] {
	return SequenceContext(context.Background(), from, to, step)
}

// Sequence uses the new Go 1.23 style generator to generate a sequence
// of integers; it may be configured with functional options. In order
// to create an infinite sequence provide no end value; in order to create
// a Sequence repeating the same value over and over, set Step to 0 and
// Start as the desired value. When the given context is cancelled, the
// generator stops.
func SequenceContext(ctx context.Context, from int64, to int64, step int64) iter.Seq[int64] {
	return func(yield func(int64) bool) {
		value := from
		for {
			select {
			case <-ctx.Done():
				// slog.Info("cancelling...", "from", settings.start, "to", settings.end, "step", settings.step)
				return
			default:
				if !yield(value) {
					return
				}
				value += step
				if to > 0 && value >= to {
					return
				}
			}
		}
	}
}
