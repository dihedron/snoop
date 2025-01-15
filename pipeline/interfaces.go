package pipeline

import (
	"errors"
)

// Filter implements a pipeline filter; it processes one message
// at a time and then passes the bucket to the next filter in the
// chain.
type Filter interface {
	Name() string
	// Process processes a message; it can modify the message provided
	// the resulting value still complies with the Message interface;
	// moreover it can instruct the pipeline to drop the current message
	// without further processing by the next filters in the chain by
	// returning a ErrSkip, or to abort the whole pipeline by returning
	// ErrAbort.
	Process(message any) (any, error)
}

// Sink implements a pipeline sink; it absorbs all messages as they
// emerge from the pipeline filters; if the Sink has a Close method,
// the Sink is closed when the pipeline is.
type Sink interface {
	// Collect acquires a message from the pipeline and acknowledges it.
	Collect(message any) error
}

// Acknowledgeable is the interface that delivered messages must comply with
// if they must be acknowledged.
type Acknowledgeable interface {
	// Ack is used to notify the source that the Message has been processed;
	// this is relevant in those cases where Messages have to be explicitly
	// removed from the source, depending on whether they have been acquired
	// by the client (see RabbitMQ deliveries).
	Ack(multiple bool) error
}

// MessageWrapper is a generic mechanism to convey the original message
// alongside the modified, or wrapping message; this allows to keep track
// of the acknowledging logic, so that if the original message must
type MessageWrapper struct {
	wrapped Acknowledgeable
}

func (w *MessageWrapper) Ack(multiple bool) error {
	return w.wrapped.Ack(multiple)
}

// ErrSkip is an error that, when emitted by a filter, causes the
// flow engine to skip BOTH any further processing by the following
// filters, and the collection by the sink.
var ErrSkip error = errors.New("message skipped")

// ErrDone is an error that, when emitted by a filter, causes the
// flow engine to skip any further processing by the following
// filters, and to send the message to the sink for collection.
var ErrDone error = errors.New("no more processing needed")

// ErrAbort is an error that, when emitted by a filter, causes the
// flow engine to abort any processing of any succeeding messages.
var ErrAbort error = errors.New("pipeline aborted")
