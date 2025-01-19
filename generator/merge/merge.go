package merge

import (
	"context"
	"iter"
	"sync"
)

func Merge[T any](ctx context.Context, sequences ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		var wg sync.WaitGroup
		channels := []<-chan T{}
		// for id, sequence := range sequences {
		for _, sequence := range sequences {
			wg.Add(1)
			c := make(chan T)
			go func(c chan<- T) {
				// slog.Info("Merge: generator starting", "id", id)
			loop:
				for value := range sequence {
					select {
					case c <- value:
						// slog.Info("Merge: message sent", "id", id, "value", value)
					case <-ctx.Done():
						// slog.Info("Merge: context closed", "id", id)
						break loop
					}
				}
				// slog.Info("Merge: closing channel", "id", id)
				wg.Done()
				close(c)
				// slog.Info("Merge: channel closed, exiting...", "id", id)
			}(c)
			channels = append(channels, c)
		}
		out := merge(channels...)
		defer func() {
			// slog.Info("Merge: main waiting for goroutines to close...")
			wg.Wait()
			// slog.Info("Merge: all goroutines closed")
		}()
		for {
			select {
			case value := <-out:
				// slog.Info("Merge: forwarding value received from channel", "value", value)
				if !yield(value) {
					return
				}
			case <-ctx.Done():
				// slog.Info("Merge: context closed")
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
		// for id, sequence := range sequences {
		for _, sequence := range sequences {
			wg.Add(1)
			c := make(chan pair)
			go func(c chan<- pair) {
				// slog.Info("Merge2: generator starting", "id", id)
			loop:
				for k, v := range sequence {
					select {
					case c <- pair{k: k, v: v}:
						// slog.Info("Merge2: message sent", "id", id, "k", k, "v", v)
					case <-ctx.Done():
						// slog.Info("Merge2: context closed", "id", id)
						break loop
					}
				}
				// slog.Info("Merge2: closing channel", "id", id)
				wg.Done()
				close(c)
				// slog.Info("Merge2: channel closed, exiting...", "id", id)
			}(c)
			channels = append(channels, c)
		}
		out := merge(channels...)
		defer func() {
			// slog.Info("Merge2: main waiting for goroutines to close...")
			wg.Wait()
			// slog.Info("Merge2: all goroutines closed")
		}()
		for {
			select {
			case value := <-out:
				// slog.Info("Merge2: forwarding value received from channel", "value", value)
				if !yield(value.k, value.v) {
					return
				}
			case <-ctx.Done():
				// slog.Info("Merge2: context closed")
				return
			}
		}
	}
}

func merge[T any](cs ...<-chan T) <-chan T {
	out := make(chan T, 100)
	var wg sync.WaitGroup
	for _, c := range cs {
		// slog.Info("merge: start new goroutine on generator")
		wg.Add(1)
		go func(c <-chan T) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		// slog.Info("merge: waiting for goroutines to close")
		wg.Wait()
		// slog.Info("merge: all goroutines closed, closing output channel")
		close(out)
	}()
	return out
}
