package random

import (
	"context"
	"crypto/rand"
	"iter"
	"log/slog"
	"math/big"

	"github.com/dihedron/snoop/pipeline"
)

// Source is a mock Source that emits a sequence of
// random integers one at a time.
type Source struct {
	min       int64
	max       int64
	buffering int
	//cancel    context.CancelFunc
}

// New creates a new Source that will generate a sequence
// of integers in the [min,max) range.
func New(min int64, max int64, buffering int) (*Source, error) {
	source := &Source{
		buffering: 1,
		min:       min,
		max:       max,
	}
	return source, nil
}

// Emit emits messages one at a time on the return channel; Random is
// a generator of an infinite series of random integers, thus it will
// only post a message onto the quit message when an error in the RNG
// occurred.
func (s *Source) Emit(ctx context.Context) (<-chan pipeline.Message, error) {
	messages := make(chan pipeline.Message, s.buffering)
	go func(ctx context.Context) {
		defer func() {
			slog.Info("closing output message channel")
			close(messages)
		}()
		for {
			select {
			case <-ctx.Done():
				slog.Info("context cancelled")
				return
			default:
				number, err := rand.Int(rand.Reader, big.NewInt(s.max-s.min))
				if err != nil {
					slog.Error("error retrieving random value", "error", err)
					return
				}
				message := &Message{Value: number.Int64() + s.min}
				slog.Debug("sending random number as message", "value", message.Value)
				messages <- message
			}
		}
	}(ctx)
	return messages, nil
}

// Random uses the new Go 1.23 style generator to generate a sequence
// of random integers.
func Random(min int64, max int64) iter.Seq[int64] {
	return func(yield func(int64) bool) {
		for {
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

// Random uses the new Go 1.23 style generator to generate a sequence
// of random integers.
func RandomContext(ctx context.Context, min int64, max int64) iter.Seq[int64] {
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
