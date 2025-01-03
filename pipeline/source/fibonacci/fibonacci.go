package fibonacci

import (
	"context"
	"log/slog"
	"math/big"

	"github.com/dihedron/snoop/pipeline"
)

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
)

// Source is a mock Source that emits the Fibonacci
// series one integer at a time.
type Source struct{}

// New creates a new Source.
func New() (*Source, error) {
	source := &Source{}
	return source, nil
}

// Emit opens the underlying sequence generator and returns a channel
// on which it will emit messages one at a time; the Fibonacci is an
// infinite series generator and will therefore never post a message
// onto the quit channel.
func (s *Source) Emit(ctx context.Context) (<-chan pipeline.Message, error) {
	messages := make(chan pipeline.Message, 1)
	go func(ctx context.Context) {
		defer func() {
			slog.Info("closing output message channel")
			close(messages)
		}()
		a := zero
		b := one
		n := big.NewInt(0)
		var message *Message
		for {
			select {
			case <-ctx.Done():
				slog.Info("context cancelled")
				return
			default:
				if n.Cmp(zero) == 0 {
					message = &Message{Value: zero}
				} else if n.Cmp(one) == 0 {
					message = &Message{Value: one}
				} else {
					message = &Message{Value: big.NewInt(0).Add(a, b)}
					a, b = b, message.Value
				}
				slog.Debug("sending Fibonacci number as message", "value", message.Value)
				messages <- message
			}
			n.Add(n, one)
		}
	}(ctx)
	return messages, nil
}
