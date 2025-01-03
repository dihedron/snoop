package filters

import (
	"context"
	"log/slog"

	"github.com/dihedron/snoop/message"
	"github.com/dihedron/snoop/pipeline"
)

// OpenStackMessageUnwrapper unwraps an Oslo message and returns
// an OpenStack notification, with a reference to the original
// delivery so that it can be acknowledged.
type OpenStackMessageUnwrapper struct{}

// NewOpenStackMessageUnwrapper allocates a new OsloMessageUnwrapper.
func NewOpenStackMessageUnwrapper() *OpenStackMessageUnwrapper {
	return &OpenStackMessageUnwrapper{}
}

// Name returns the name of the current filter.
func (f *OpenStackMessageUnwrapper) Name() string {
	return "github.com/dihedron/snoop/message/OpenStackMessageUnwrapper"
}

// Process checks if the input message is of type message.OsloMessage,
// then unwraps it into an Oslo message and again into an OpenStack
// notification; the returned object will retain a reference to the
// original amqp.Delivery so it can be acknowledged.
func (f *OpenStackMessageUnwrapper) Process(ctx context.Context, msg pipeline.Message) (context.Context, pipeline.Message, error) {
	if m, ok := msg.(*message.Oslo); ok {
		notification, err := message.NewNotificationFromOslo(m, true)
		if err != nil {
			slog.Error("error reading Notification message", "error", err)
			return ctx, nil, ErrUnsupportedMessageType
		}
		return ctx, notification, nil

	}
	slog.Error("not a valid Oslo message", "error", ErrUnsupportedMessageType)
	return ctx, nil, ErrUnsupportedMessageType
}
