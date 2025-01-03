package merge

import (
	"context"
	"io"
	"log/slog"

	"github.com/dihedron/snoop/pipeline"
)

// Source is a synthetic Source that emits the messages coming from
// the underlying sources in a non-deterministic, interleaved way.
type Source struct {
	source1 pipeline.Source
	source2 pipeline.Source
}

// New creates a new Source.
func New(source1, source2 pipeline.Source) *Source {
	return &Source{
		source1: source1,
		source2: source2,
	}
}

// Emit opens the underlying sources and returns a channel
// on which it will emit messages from both, one at a time.
func (s *Source) Emit(ctx context.Context) (<-chan pipeline.Message, error) {
	messages := make(chan pipeline.Message, 1)
	ch1, err := s.source1.Emit(ctx)
	if err != nil {
		slog.Error("error opening first source", "error", err)
		if closeable, ok := s.source1.(io.Closer); ok {
			closeable.Close()
		}
	}
	ch2, err := s.source2.Emit(ctx)
	if err != nil {
		slog.Error("error opening second source", "error", err)
		if closeable, ok := s.source1.(io.Closer); ok {
			closeable.Close()
		}
		if closeable, ok := s.source2.(io.Closer); ok {
			closeable.Close()
		}
	}
	go func(ctx context.Context) {
		defer func() {
			slog.Info("closing output message channel")
			close(messages)
		}()
		closed := []bool{false, false}
	loop:
		for {
			select {
			case m, ok := <-ch1:
				if ok {
					slog.Debug("forwarding message received from source 1", "message", m)
					messages <- m
				} else {
					slog.Debug("source 1 closed")
					closed[0] = true
					if closed[0] && closed[1] {
						slog.Debug("both sources closed, breaking...")
						break loop
					}
				}
			case m, ok := <-ch2:
				if ok {
					slog.Debug("forwarding message received from source 2", "message", m)
					messages <- m
				} else {
					slog.Debug("source 2 closed")
					closed[1] = true
					if closed[0] && closed[1] {
						slog.Debug("both sources closed, breaking...")
						break loop
					}
				}
			case <-ctx.Done():
				slog.Info("context cancelled")
				return
			default:
				slog.Debug("in default case")
				if closed[0] && closed[1] {
					slog.Debug("both sources closed, breaking...")
					break loop
				}
			}
		}
	}(ctx)
	return messages, nil
}

// Close closes the underlying sources.
func (s *Source) Close() error {
	if s != nil && s.source1 != nil {
		if closeable, ok := s.source1.(io.Closer); ok {
			closeable.Close()
		}
	}
	if s != nil && s.source2 != nil {
		if closeable, ok := s.source2.(io.Closer); ok {
			closeable.Close()
		}
	}
	return nil
}
