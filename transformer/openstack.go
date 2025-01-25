package transformer

import (
	"encoding/json"
	"errors"
	"log/slog"
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

// AMQPMessage is an almost exact replica of amqp.Delivery; it is necessary to
// be able to introspect the message, be able to print out its contents
// in multiple formats (YAML, JSON), perform field-by-field comparisons etc.
// It may optionally contain a reference to the original amqp.Delivery so it
// can be acknowledged.
type AMQPMessage struct {
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
	// mutex protects the delivery reference from concurrent access
	mutex sync.RWMutex
	// delivery is the underlying RabbitMQ Delivery
	delivery *amqp091.Delivery
}

// NewAMQPMessage creates an AMQPMessage from a RabbitMQ amqp.Delivery message.
func AMQPDeliveryToMessage(includeDelivery bool) func(*amqp091.Delivery) (*AMQPMessage, error) {
	return func(delivery *amqp091.Delivery) (*AMQPMessage, error) {
		if delivery == nil {
			slog.Error("input must not be nil")
			return nil, errors.New("invalid input") // was: ErrInvalidInput
		}
		message := &AMQPMessage{
			ContentType:     delivery.ContentType,
			ContentEncoding: delivery.ContentEncoding,
			DeliveryMode:    delivery.DeliveryMode,
			DeliveryTag:     delivery.DeliveryTag,
			Priority:        delivery.Priority,
			CorrelationID:   delivery.CorrelationId,
			ReplyTo:         delivery.ReplyTo,
			MessageID:       delivery.MessageId,
			Expiration:      delivery.Expiration,
			Timestamp:       delivery.Timestamp,
			Type:            delivery.Type,
			UserID:          delivery.UserId,
			ApplicationID:   delivery.AppId,
			ConsumerTag:     delivery.ConsumerTag,
			MessageCount:    delivery.MessageCount,
			Redelivered:     delivery.Redelivered,
			Exchange:        delivery.Exchange,
			RoutingKey:      delivery.RoutingKey,
			Body:            delivery.Body,
		}
		if includeDelivery {
			slog.Debug("adding original delivery reference to AMQP message", "reference", delivery.DeliveryTag)
			// TODO: check if we still need concurrent access (no channels...)
			message.mutex.Lock()
			defer message.mutex.Unlock()
			message.delivery = delivery
		}
		return message, nil
	}
}

// AMQPMessageToOslo extracts an Oslo message from an AMQP delivery payload.
func AMQPMessageToOslo(includeDelivery bool) func(*AMQPMessage)(*Oslo, error) {
	return func(message *AMQPMessage)(*Oslo, error) {
		if message == nil {
			slog.Error("input must not be nil")
			return nil, errors.New("invalid input")//ErrInvalidInput
		}
		//oslo, err := NewOsloFromJSON(string(message.Body))

		tmp := struct {
			Version string `json:"oslo.version" yaml:"oslo.version"`
			Payload string `json:"oslo.message" yaml:"oslo.message"`
		}{}
		if err := json.Unmarshal([]byte(data), &tmp); err != nil {
			slog.Error("error parsing Oslo message", "error", err)
			return nil, err
		}
		slog.Debug("oslo message parsed", "version", oslo.Version)
		oslo := &Oslo{
			Version: tmp.Version,
			Payload: tmp.Payload,
		}

		if err == nil && oslo != nil && includeDelivery {
			message.mutex.RLock()
			defer message.mutex.RUnlock()
			delivery := message.delivery
			if delivery != nil {
				slog.Debug("valid delivery reference acquired, adding to Oslo message...")
				oslo.mutex.Lock()
				defer oslo.mutex.Unlock()
				oslo.delivery = delivery
				slog.Debug("added reference to original delivery to Oslo message", "tag", delivery.DeliveryTag)
			}
		}
		return oslo, err
	}
}


/*

// Process checks if the input message is of type amqp091.Delivery,
// then unwraps it into an Oslo message and again into an OpenStack
// notification; the returned object will retain a reference to the
// original amqp091.Delivery so it can be acknowledged.
func DeliveryToOslo(message any) (, error) {
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
*/
