package amqp

import (
	"log/slog"
	"time"

	"github.com/dihedron/snoop/format"
	"github.com/rabbitmq/amqp091-go"
)

// Message is an almost exact replica of amqp.Delivery; it is necessary to
// be able to introspect the message, be able to print out its contents
// in multiple formats (YAML, JSON), perform field-by-field comparisons etc.
// It may optionally contain a reference to the original amqp.Delivery so it
// can be acknowledged.
type Message struct {
	// Headers is the application or header exchange table
	Headers map[string]interface{} `json:"headers,omitempty" yaml:"headers,omitempty"`
	// ContentType is the MIME content type.
	ContentType string `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	// ContentEncoding is the MIME content encoding.
	ContentEncoding string `json:"contentEncoding,omitempty" yaml:"contentEncoding,omitempty"`
	// DeliveryMode indicates whether it is non-persistent (1) or persistent (2).
	DeliveryMode uint8 `json:"deliveryMode,omitempty" yaml:"deliveryMode,omitempty"`
	// DeliveryTag is...
	DeliveryTag uint64 `json:"deliveryTag,omitempty" yaml:"deliveryTag,omitempty"`
	// Priority indicates the message priority (0 to 9).
	Priority uint8 `json:"priority,omitempty" yaml:"priority,omitempty"`
	// CorrelationID is the message correlation id.
	CorrelationID string `json:"correlationId,omitempty" yaml:"correlationId,omitempty"`
	// ReplyTo is the address to reply to (e.g. RPC).
	ReplyTo string `json:"replyTo,omitempty" yaml:"replyTo,omitempty"`
	// MessageID is the message identifier.
	MessageID string `json:"messageId,omitempty" yaml:"messageId,omitempty"`
	// Expiration is the message expiration spec.
	Expiration string `json:"expiration,omitempty" yaml:"expiration,omitempty"`
	// Timestamp is the message timestamp.
	Timestamp time.Time `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
	// Type is the message type name.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	// UserID is the (authenticated) user who created the message.
	UserID string `json:"userId,omitempty" yaml:"userId,omitempty"`
	// ApplicationID is the ID of the application that created the message.
	ApplicationID string `json:"appId,omitempty" yaml:"appplicationId,omitempty"`
	// ConsumerTag is valid only with Channel.Consume
	ConsumerTag string `json:"consumerTag,omitempty" yaml:"consumerTag,omitempty"`
	// MessageCount is valid only with Channel.Get
	MessageCount uint32 `json:"messageCount,omitempty" yaml:"messageCount,omitempty"`
	// Redelivered is true if the message is being redelivered.
	Redelivered bool `json:"redelivered,omitempty" yaml:"redelivered,omitempty"`
	// Exchange is the exchange from which the message is coming.
	Exchange string `json:"exchange,omitempty" yaml:"exchange,omitempty"` // basic.publish exchange
	// RoutingKey is the key used to route the message to this queue.
	RoutingKey string `json:"routingKey,omitempty" yaml:"routingKey,omitempty"` // basic.publish routing key
	// Body is the actual message body.
	Body []byte `json:"body,omitempty" yaml:"body,omitempty"`
	// backref is a reference to the underlying RabbitMQ Delivery
	backref *amqp091.Delivery
}

// String converts the Message into its JSON one-liner representation.
func (m *Message) String() string {
	return format.ToJSON(m)
}

// Ack allows to acknowledge the Message's underlying amqp091.Delivery, if set.
func (m *Message) Ack(multiple bool) error {
	slog.Debug("acknowledging AMQP message...", "type", format.TypeAsString(m))
	if m.backref != nil {
		slog.Debug("acknowledging message", "correlation id", m.backref.CorrelationId)
		if err := m.backref.Ack(multiple); err != nil {
			slog.Error("error acnowledging message", "correlation id", m.backref.CorrelationId, "error", err)
			return err
		}
		m.backref = nil
	}
	return nil
}

// BackRef returns the reference to the original AMQP delivery.
func (m *Message) BackRef() *amqp091.Delivery {
	return m.backref
}
