package transformers

import (
	"log/slog"
	"time"

	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/transform"
)

// StopWatch allowa a way to measure the time elapsed
type StopWatch[S any, T any] struct {
	start   time.Time
	elapsed time.Duration
}

// Start sets the stopwatch start time, based on which the elapsed
// time is computed. This filter does not affect the value flowing
// through.
func (s *StopWatch[S, T]) Start() transform.X[S, S] {
	return func(value S) (S, error) {
		s.start = time.Now()
		slog.Debug("profile start", "value", value, "type", format.TypeAsString(value))
		return value, nil
	}
}

// Stop sets the stopwatch stop time and computes the elapsed time
// from the start moment. This filter does not affect the value
// flowing through.
func (s *StopWatch[S, T]) Stop() transform.X[T, T] {
	return func(value T) (T, error) {
		s.elapsed = time.Since(s.start)
		slog.Debug("profile stop", "elapsed", s.elapsed.String(), "value", value, "type", format.TypeAsString(value))
		return value, nil
	}
}

// Elapsed returns the time elapsed between the Start and Stop calls.
func (s *StopWatch[S, T]) Elapsed() time.Duration {
	return s.elapsed
}
