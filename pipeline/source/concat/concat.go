package concat

import (
	"context"
	"io"
	"iter"
	"log/slog"

	"github.com/dihedron/snoop/pipeline"
)

// Source is a synthetic Source that emits the messages
// from multiple sources, one after the other.
type Source struct {
	sources []pipeline.Source
}

// New creates a new Source.
func New(sources ...pipeline.Source) *Source {
	return &Source{
		sources: sources,
	}
}

// Emit opens the underlying sources and returns a channel
// on which it will emit messages from both, one at a time.
func (s *Source) Emit(ctx context.Context) (<-chan pipeline.Message, error) {
	messages := make(chan pipeline.Message, 1)
	go func(ctx context.Context) {
		defer func() {
			slog.Debug("closing output message channel")
			close(messages)
		}()
		// one source at a time, open it and drain it into
		// the output channel
	outer:
		for i, source := range s.sources {
			ch, err := source.Emit(ctx)
			if err != nil {
				slog.Error("error opening source", "source id", i, "error", err)
				continue outer
			}
		inner:
			for {
				select {
				case m, ok := <-ch:
					if ok {
						slog.Debug("forwarding message", "message", m, "source id", i)
						messages <- m
					} else {
						slog.Debug("no more messages available", "source id", i)
						break inner
					}
				case <-ctx.Done():
					slog.Info("context cancelled")
					return
				}
			}
		}
	}(ctx)
	return messages, nil
}

// Close closes the underlying sources.
func (s *Source) Close() error {
	slog.Info("closing the source")
	if s != nil && len(s.sources) > 0 {
		for i, source := range s.sources {
			if closeable, ok := source.(io.Closer); ok {
				slog.Info("closing source", "source id", i)
				closeable.Close()
			}
		}
	}
	return nil
}

// Concat concatenates multiple single-valued sequences.
func Concat[T any](sequences ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, sequence := range sequences {
			for value := range sequence {
				if !yield(value) {
					return
				}
			}
		}
	}
}

// Concat2 concatenates multiple double-valued sequences.
func Concat2[T any, S any](sequences ...iter.Seq2[T, S]) iter.Seq2[T, S] {
	return func(yield func(T, S) bool) {
		for _, sequence := range sequences {
			for v1, v2 := range sequence {
				if !yield(v1, v2) {
					return
				}
			}
		}
	}
}
