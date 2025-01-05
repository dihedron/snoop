package merge

import (
	"context"
	"io"
	"iter"
	"log/slog"
	"sync"

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

func Merge[T any](ctx context.Context, sequences ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		var wg sync.WaitGroup
		channels := []<-chan T{}
		for _, sequence := range sequences {
			wg.Add(1)
			c := make(chan T)
			go func(c chan<- T) {
				for value := range sequence {
					c <- value
				}
				close(c)
			}(c)
			channels = append(channels, c)
		}
		out := merge(channels...)
		defer func() {
			wg.Wait()
		}()
		for {
			select {
			case value := <-out:
				slog.Info("forwarding value received from channel", "value", value)
				if !yield(value) {
					return
				}
			case <-ctx.Done():
				slog.Info("context closed")
				return
			}
		}
	}
}

func Merge2[K any, V any](ctx context.Context, sequences ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	type pair struct {
		k K
		v V
	}
	return func(yield func(K, V) bool) {
		var wg sync.WaitGroup
		channels := []<-chan pair{}
		for _, sequence := range sequences {
			wg.Add(1)
			c := make(chan pair)
			go func(c chan<- pair) {
				for k, v := range sequence {
					c <- pair{k: k, v: v}
				}
				close(c)
			}(c)
			channels = append(channels, c)
		}
		out := merge(channels...)
		defer func() {
			wg.Wait()
		}()
		for {
			select {
			case value := <-out:
				slog.Info("forwarding value received from channel", "key", value.k, "value", value.v)
				if !yield(value.k, value.v) {
					return
				}
			case <-ctx.Done():
				slog.Info("context closed")
				return
			}
		}
	}
}

func merge[T any](cs ...<-chan T) <-chan T {
	out := make(chan T)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan T) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
