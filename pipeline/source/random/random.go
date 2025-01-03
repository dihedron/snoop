package random

import (
	"context"
	"crypto/rand"
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
