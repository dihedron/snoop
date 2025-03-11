package oslo

import (
	"log/slog"

	"github.com/dihedron/snoop/format"
	"github.com/rabbitmq/amqp091-go"
)

// Oslo is an OpenStack Oslo v2 message; the Payload is the actual JSON-like
// informational message sent from the OpenStack services. It may contain a
// reference to the original RabbitMQ amqp.Delivery, so it allows to
// acknowledge it if needed.
type Oslo struct {
	// Version contains the version of the Oslo message.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	// Payload contains the original, unparsed Oslo message payload as a string;
	// in order for it to be parsed as valid JSON, it requires unescaping the
	// quotes.
	Payload string `json:"payload,omitempty" yaml:"payload,omitempty"`
	// backref is a reference to the underlying RabbitMQ Delivery
	backref *amqp091.Delivery
}

// String converts the message into its JSON one-liner representation.
func (o *Oslo) String() string {
	return format.ToJSON(o)
}

// Ack allows to acknowledge the Oslo's underlying amqp091.Delivery, if set.
func (o *Oslo) Ack(multiple bool) error {
	slog.Debug("acknowledging Oslo message...", "type", format.TypeAsString(o))
	if o.backref != nil {
		slog.Debug("acknowledging message", "correlation id", o.backref.CorrelationId)
		if err := o.backref.Ack(multiple); err != nil {
			slog.Error("error acnowledging message", "correlation id", o.backref.CorrelationId, "error", err)
			return err
		}
		o.backref = nil
	}
	return nil
}

// BackRef returns the reference to the original AMQP delivery.
func (o *Oslo) BackRef() *amqp091.Delivery {
	return o.backref
}
