package filters

import (
	"context"
	"log/slog"

	"github.com/dihedron/snoop/message"
	"github.com/dihedron/snoop/pipeline"
)

// OsloMessageUnwrapper unwraps an AMQP delivery and returns
// an Oslo message, with a reference to the original delivery
// so that it can be acknowledged.
type OsloMessageUnwrapper struct{}

// NewOsloMessageUnwrapper allocates a new OsloMessageUnwrapper.
func NewOsloMessageUnwrapper() *OsloMessageUnwrapper {
	return &OsloMessageUnwrapper{}
}

// Name returns the name of the current filter.
func (f *OsloMessageUnwrapper) Name() string {
	return "github.com/dihedron/snoop/message/OsloMessageUnwrapper"
}

// Process checks if the input message is of type message.AMQPMessage,
// then unwraps it into an Oslo message and again into an OpenStack
// notification; the returned object will retain a reference to the
// original amqp.Delivery so it can be acknowledged.
func (f *OsloMessageUnwrapper) Process(ctx context.Context, msg pipeline.Message) (context.Context, pipeline.Message, error) {
	if m, ok := msg.(*message.AMQPMessage); ok {
		oslo, err := message.NewOsloFromAMQPMessage(m, true)
		if err != nil {
			slog.Error("error reading Oslo message", "error", err)
			return ctx, nil, ErrUnsupportedMessageType
		}
		return ctx, oslo, nil
	}
	slog.Error("not a valid AMQP message", "error", ErrUnsupportedMessageType)
	return ctx, nil, ErrUnsupportedMessageType
}

// UnwrapOslo unwraps a message received from RabbitMQ into an Oslo
// notification.
func UnwrapOslo(value T) (T, error) {
	var nihil pipeline.Acknowledgeable
	if m, ok := pipeline.Acknowledgeable(value).(*message.AMQPMessage); ok {
		oslo, err := message.NewOsloFromAMQPMessage(m, true)
		if err != nil {
			slog.Error("error reading Oslo message", "error", err)
			r, _ := any(nihil).(pipeline.Acknowledgeable)
			return r.(T), ErrUnsupportedMessageType
		}
		return oslo.(T), nil
	}
	slog.Error("not a valid AMQP message", "error", ErrUnsupportedMessageType)
	return nil, ErrUnsupportedMessageType
}
