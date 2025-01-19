package pipeline

import (
	"errors"
	"io"
	"log/slog"

	"github.com/dihedron/snoop/format"
)

type Handler[T any] func(value T) (T, error)

// Filter implements a pipeline filter; it processes one value
// at a time and then passes the bucket to the next filter in the
// chain.
type Filter[T any] interface {
	// Process processes a value; it can modify the message provided
	// the resulting value still has the same type as the input value.
	// if the Filter returns and error, the pipeline is interrupted and
	// the returned value and error are returned to the caller.
	Apply(value T) (T, error)
}

// Pipeline represents a data flow pipeline; it implements a chain of
// responsibility where each Filter processes one step of the overall
// treatment; at the end, the pipeline returns a processed item.
type Pipeline[T any] struct {
	// source  iter.Seq[T]
	filters []Filter[T]
}

// New creates a new Pipeline comprising all the provided filters.
func New[T any](filters ...Filter[T]) *Pipeline[T] {
	return &Pipeline[T]{
		filters: filters,
	}
}

// Execute starts the pipeline by retrieving messages from the underlying Source and
// piping the Messages into the filters and eventually the pipeline sink.
func (p *Pipeline[T]) Apply(value T) (T, error) {
	slog.Debug("applying pipeline", "input", value)
	var err error
	for _, filter := range p.filters {
		slog.Debug("applying filter", "filter", format.TypeAsString(filter))
		value, err = filter.Apply(value)
		if err != nil {
			slog.Error("filter returned error", "filter", format.TypeAsString(filter), "value", value, "error", err)
			return value, err
		}
		slog.Debug("filter result", "filter", format.TypeAsString(filter), "value", value)
	}
	slog.Debug("pipeline complete", "output", value)
	return value, nil
}

// Close calls the Close() method on any filter that implements the
// Closeable interface.
func (p *Pipeline[T]) Close() error {
	var err error
	for _, filter := range p.filters {
		if closeable, ok := filter.(io.Closer); ok {
			slog.Debug("closing filter", "filter", format.TypeAsString(filter))
			err = errors.Join(err, closeable.Close())
		}
	}
	return err
}
