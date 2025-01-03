package filters

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/dihedron/snoop/message"
	"github.com/dihedron/snoop/pipeline"
	"github.com/dihedron/snoop/pipeline/source/file"
	amqp091 "github.com/rabbitmq/amqp091-go"
)

// ErrUnsupportedMessageType is returned when the input message
// was not recognised as a valid message.
var ErrUnsupportedMessageType = errors.New("invalid message")

// AMQPMessageUnwrapper unwraps an AMQP delivery and returns
// an OpenStack notification, with a reference to the original
// delivery so that it can be acknowledged.
type AMQPMessageUnwrapper struct{}

// NewAMQPMessageUnwrapper allocates a new AMQPMessageUnwrapper.
func NewAMQPMessageUnwrapper() *AMQPMessageUnwrapper {
	return &AMQPMessageUnwrapper{}
}

// Name returns the name of the current filter.
func (f *AMQPMessageUnwrapper) Name() string {
	return "github.com/dihedron/snoop/message/AMQPMessageUnwrapper"
}

// Process checks if the input message is of type amqp091.Delivery,
// then unwraps it into an Oslo message and again into an OpenStack
// notification; the returned object will retain a reference to the
// original amqp091.Delivery so it can be acknowledged.
func (f *AMQPMessageUnwrapper) Process(ctx context.Context, msg pipeline.Message) (context.Context, pipeline.Message, error) {
	var (
		result *message.AMQPMessage
		err    error
	)
	switch m := msg.(type) {
	case *file.Message:
		slog.Debug("input value is file.Message")
		result = &message.AMQPMessage{}
		if err = json.Unmarshal([]byte(m.Value), result); err != nil {
			slog.Error("error parsing AMQP delivery from JSON", "error", err)
			return ctx, nil, pipeline.ErrDone
			// return ctx, nil, err
		}
	case *amqp091.Delivery:
		slog.Debug("input value is amqp091.Delivery")
		result, err = message.NewAMQPMessage(m, true)
		if err != nil {
			slog.Error("error reading AMQP message", "error", err)
			return ctx, nil, err
		}
	default:
		slog.Error("unknown input message type", "type", fmt.Sprintf("%T", msg))
		return ctx, nil, ErrUnsupportedMessageType
	}
	return ctx, result, nil
}
