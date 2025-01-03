package integer

import (
	"context"
	"log/slog"

	"github.com/dihedron/snoop/pipeline"
	"github.com/dihedron/snoop/pointer"
)

// Option is the type for functional options.
type Option func(*Source)

// New creates a new Source, applying all the provided functional options.
func New(options ...Option) *Source {
	p := &Source{
		start: 0,
		step:  1,
		end:   100,
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
func From(start int64) Option {
	return func(s *Source) {
		s.start = start
	}
}

// Step allows to specify the step of the sequence; if
// it is positive the number in the sequence will grow, if
// negative will decrease; if 0, the sequence will be an infinite
// generator of the same, initial value. If not specified,
// the sequence will assume the default value of +1.
func Step(step int64) Option {
	return func(s *Source) {
		s.step = step
	}
}

// Until allows to specify the end value (exclusive); in order to
// generate an infinite sequence don't set this value.
func Until(end int64) Option {
	return func(s *Source) {
		s.end = end
	}
}

// Source is a mock Source that emits the value
// series one integer at a time.
type Source struct {
	start int64
	step  int64
	end   int64
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
				message := pointer.To(Message(value))
				slog.Debug("sending sequence number as message", "value", int64(*message))
				messages <- message
				value += s.step
				if value >= s.end {
					slog.Debug("end of sequence reached")
					return
				}
			}
		}
	}(ctx)
	slog.Info("returning sequence channel")
	return messages, nil
}
