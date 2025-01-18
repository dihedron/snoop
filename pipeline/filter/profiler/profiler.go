package profiler

import (
	"log/slog"
	"time"
)

// Profiler is a filter that keeps track of the duration between its
// first invocation in the pipeline and its successive invocation(s).
type Profiler[_ any] struct {
	set     bool
	start   time.Time
	elapsed time.Duration
}

// New creates a new Profiler.
func New[T any]() *Profiler[T] {
	return &Profiler[T]{}
}

func (p *Profiler[_]) Reset() {
	p.set = false
	p.start = time.Time{}
}

// Elapsed returns the time elapsed between the first and the
// second time it is called; it should be placed at the beginning
// and at the end of a pipeline, and can then be reset at each
// iteration when the pipeline is used in a loop..
func (p *Profiler[_]) Elapsed() time.Duration {
	return p.elapsed
}

// Apply records the start time the first time it is called, and
// computes the elapsed time the second time, resetting the Profiler
// so it can be reused.
func (p *Profiler[T]) Apply(value T) (T, error) {
	if !p.set {
		slog.Debug("recording start time")
		p.start = time.Now()
		p.set = true
	} else {
		p.elapsed = time.Since(p.start)
		slog.Debug("elapsed time", "ms", p.elapsed.Milliseconds())
		p.Reset()
	}
	return value, nil
}
