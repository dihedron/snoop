package pipeline

import (
	"context"
	"errors"
)

// Source describes the behaviour of input sources; the concrete
// implementations are expected to be able to output messages one at
// a time into the output channel, in order for them to be piped into
// the actual pipeline and its Processors. When the input Context is
// cancelled, the Source should poison the pipeline and then return
// immediately.
type Source interface {
	// Emit starts the input source which will emit messages one at a time
	// on the returned channel; if an error occurs when opening the underlying
	// source, an error could be returned. The input Context can be used to stop
	// the Source: concrete implementations will have to check on the Done() channel
	// before lengthy operations in order to exit as soon as it is signalled; once
	// the source is done producing messages, it will close the output message so
	// the pipline Engine can detect that there are no more messages available.
	// Note that Open starts a new goroutine to asynchronously produce and output
	// messages onto the output channel.
	Emit(ctx context.Context) (<-chan Message, error)
}

// Filter implements a pipeline filter; it processes one message
// at a time and then passes the bucket to the next filter in the
// chain.
type Filter interface {
	Name() string
	// Process processes a message; it should check if the context
	// has been cancelled and return immediately if so; it can tell
	// the pipeline to drop the current message without further processing
	// by the next filters in the chain by returning a ErrSkip, or to abort
	// the whole pipeline by returning ErrAbort. The output Context should
	// be the same that was given as input or a derived Context; this
	// allows to add objects to the context and have them passed over to
	// the following filters.
	Process(ctx context.Context, message Message) (context.Context, Message, error)
}

// Sink implements a pipeline sink; it absorbs all messages as they
// emerge from the pipeline filters; if the Sink has a Close method,
// the Sink is closed when the pipeline is.
type Sink interface {
	// Collect acquires a message from the pipeline and acknowledges it.
	Collect(ctx context.Context, message Message) error
}

// Message is the interface that delivered messages must comply with.
type Message interface {
	// Ack is used to notify the source that the Message has been processed;
	// this is relevant in those cases where Messages have to be explicitly
	// removed from the source, depending on whether they have been acquired
	// by the client (see RabbitMQ deliveries). When this is not the case,
	// the function can be a no-op.
	Ack(multiple bool) error
}

// MessageWrpper is a generic mechanism to convey the original message
// alongside the modified, or wrapping message; this allows to keep track
// of the acknowledging logic, so that if the original message must
type MessageWrapper struct {
	wrapped Message
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
