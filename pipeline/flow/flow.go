package flow

import (
	"errors"
	"fmt"
	"io"
	"iter"
	"log/slog"

	"github.com/dihedron/snoop/pipeline"
)

// Flow represents a data flow pipeline; it allows to specify a single
// Source (which is an iter.Seq[Message] generator) and an ordered list of
// Filters (which must comply with the pipeline.Filter interface).
// The pipeline also accepts a set of callbacks which are called whenever
// an error occurs and have a chance to abort the whole pipeline.
type Flow[T any] struct {
	source  iter.Seq[T]
	filters []pipeline.Filter
	sink    pipeline.Sink
}

// Option is the type for functional options.
type Option[T any] func(*Flow[T])

// New creates a new Flow, applying all the provided functional options.
func New[T any](options ...Option[T]) *Flow[T] {
	f := &Flow[T]{
		filters: []pipeline.Filter{},
	}
	for _, option := range options {
		option(f)
	}
	return f
}

// From allows to specify the Flow messages generator.
func From[T any](source iter.Seq[T]) Option[T] {
	return func(f *Flow[T]) {
		if source != nil {
			f.source = source
		}
	}
}

// Through allows to add one or more Filters to the Flow.
func Through[T any](filters ...pipeline.Filter) Option[T] {
	return func(f *Flow[T]) {
		for _, filter := range filters {
			if filter != nil {
				f.filters = append(f.filters, filter)
			}
		}
	}
}

// Into allows to specify the Flow sink.
func Into[T any](sink pipeline.Sink) Option[T] {
	return func(f *Flow[T]) {
		if sink != nil {
			f.sink = sink
		}
	}
}

// Execute starts the pipeline by retrieving messages from the underlying Source and
// piping the Messages into the filters and eventually the pipeline sink.
func (f *Flow[T]) Execute() error {
	if f == nil || f.source == nil || f.sink == nil {
		return errors.New("invalid flow")
	}

messages:
	for value := range f.source {
		slog.Debug("processing new incoming message...", "type", fmt.Sprintf("%T", value))
		message := any(value)
		var err error
	filtering:
		for _, filter := range f.filters {
			backup := message
			message, err = filter.Process(message.(any))
			switch err {
			case nil:
				slog.Debug("filter processing successful", "filter", filter.Name())
				continue filtering
			case pipeline.ErrSkip:
				slog.Warn("filter requested skipping of message", "filter", filter.Name(), "message", message)
				continue messages
			case pipeline.ErrDone:
				slog.Warn("filter requested skipping of further processing for message", "filter", filter.Name(), "message", message)
				break filtering
			case pipeline.ErrAbort:
				slog.Error("filter requested aborting the pipeline", "filter", filter.Name())
				break messages
			default:
				// TODO: check if sending the original message to the sink for collection is ok
				slog.Warn("filter returned an error on message ", "filter", filter.Name(), "message", backup, "type", fmt.Sprintf("%T", backup), "error", err)
				message = backup
				break filtering
			}
		}
		// we emit the message; there must be a sink that acknowledges it;
		// if the message has been transformed, the filter must make sure
		// that a reference to the original message's Ack() method is
		// available and invoked when acknowledging the final incarnation
		// of the message, or there will be no ack at all.
		err = f.sink.Collect(message)
		if err != nil {
			slog.Error("error collecting message into sink", "error", err)
			return err
		}

	}
	slog.Debug("terminating pipeline message pump...")
	return nil
}

// Close stops the flow and releases associated resources (e.g.
// by closing the Source).
func (f *Flow[T]) Close() error {
	if f == nil || f.source == nil || f.sink == nil {
		return errors.New("invalid flow")
	}
	for _, filter := range f.filters {
		if closeable, ok := filter.(io.Closer); ok {
			slog.Info("closing filter", "filter", filter.Name())
			closeable.Close()
		}
	}
	if closeable, ok := f.sink.(io.Closer); ok {
		slog.Info("closing the sink")
		closeable.Close()
	}

	return nil
}
