package bigint

import (
	"context"
	"log/slog"
	"math/big"

	"github.com/dihedron/snoop/pipeline"
)

// Option is the type for functional options.
type Option func(*Source)

// New creates a new Source, applying all the provided functional options.
func New(options ...Option) *Source {
	p := &Source{
		start: big.NewInt(0),
		step:  big.NewInt(1),
	}
	for _, option := range options {
		option(p)
	}
	return p
}

// From allows to specify the start value; in order to generate
// an infinite sequence of the same value, set this to the desired
// value and the step to 0. If not specified, the sequence will assume
// the default value of 0.
func From(start *big.Int) Option {
	return func(s *Source) {
		s.start = new(big.Int).Set(start)
	}
}

// Step allows to specify the step of the sequence; if
// it is positive the number in the sequence will grow, if
// negative will decrease; if 0, the sequence will be an infinite
// generator of the same, initial value. If not specified,
// the sequence will assume the default value of +1.
func Step(step *big.Int) Option {
	return func(s *Source) {
		s.step = new(big.Int).Set(step)
	}
}

// WithEnd allows to specify the end value (exclusive); in order to
// generate an infinite sequence don't set this value.
func Until(end *big.Int) Option {
	return func(s *Source) {
		s.end = new(big.Int).Set(end)
	}
}

// Source is a mock Source that emits the value
// series one integer at a time.
type Source struct {
	start *big.Int
	step  *big.Int
	end   *big.Int
}

// Emit opens the underlying sequence generator and returns a channel
// on which it will emit messages one at a time; it is safe to assume
// that it will never emit a message onto the quit message.
func (s *Source) Emit(ctx context.Context) (<-chan pipeline.Message, error) {
	messages := make(chan pipeline.Message, 1)
	go func(ctx context.Context) {
		defer func() {
			slog.Info("closing output message channel")
			close(messages)
		}()
		value := s.start
		for {
			select {
			case <-ctx.Done():
				slog.Info("context cancelled")
				return
			default:
				message := &Message{Value: new(big.Int).Set(value)}
				slog.Debug("sending sequence number as message", "value", message.Value)
				messages <- message
				value.Add(value, s.step)
				if s.end != nil && value.Cmp(s.end) >= 0 {
					slog.Debug("end of sequence reached")
					return
				}
			}
		}
	}(ctx)
	slog.Info("returning sequence channel")
	return messages, nil
}
