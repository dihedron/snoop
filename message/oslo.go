package message

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

// Oslo is an OpenStack Oslo v2 message; the Payload is the actual JSON-like
// informational message sent from the OpenStack services. It may contain a
// reference to the original RabbitMQ amqp091.Delivery, so it allows to
// acknowledge it if needed.
type Oslo struct {
	// Version contains the version of the Oslo message.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`
	// Payload contains the original, unparsed Oslo message payload as a string;
	// in order for it to be parsed as valid JSON, it requires unescaping the
	// quotes.
	Payload string `json:"payload,omitempty" yaml:"payload,omitempty"`
	// mutex protects the delivery reference from concurrent access.
	mutex sync.RWMutex
	// delivery is the underlying RabbitMQ Delivery
	delivery *amqp091.Delivery
}

// NewOsloFromAMQPMessage extracts an Oslo message from an AMQPMessage payload.
func NewOsloFromAMQPMessage(message *AMQPMessage, includeDelivery bool) (*Oslo, error) {
	if message == nil {
		slog.Error("input must not be nil")
		return nil, ErrInvalidInput
	}
	oslo, err := NewOsloFromJSON(string(message.Body))
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

// NewOsloFromJSON extracts an Oslo message from a JSON string.
func NewOsloFromJSON(data string) (*Oslo, error) {
	oslo := struct {
		Version string `json:"oslo.version" yaml:"oslo.version"`
		Payload string `json:"oslo.message" yaml:"oslo.message"`
	}{}
	if err := json.Unmarshal([]byte(data), &oslo); err != nil {
		slog.Error("error parsing Oslo message", "error", err)
		return nil, err
	}
	slog.Debug("oslo message parsed", "version", oslo.Version)
	return &Oslo{
		Version: oslo.Version,
		Payload: oslo.Payload,
	}, nil
}

// String converts the message into its JSON one-liner representation.
func (msg *Oslo) String() string {
	bytes, err := json.Marshal(msg)
	if err != nil {
		slog.Error("failure marshaling object to JSON", "error", err)
		return ""
	}
	return string(bytes)
}

// Ack allows to propagate and acknowledge the underlying
// amqp091.Delivery if set.
func (msg *Oslo) Ack(multiple bool) error {
	slog.Debug("acknowledging Oslo message...", "type", fmt.Sprintf("%T", msg))
	if msg != nil {
		msg.mutex.Lock()
		defer msg.mutex.Unlock()
		if msg.delivery != nil {
			slog.Debug("acknowledging message", "correlation id", msg.delivery.CorrelationId)
			err := msg.delivery.Ack(multiple)
			if err == nil {
				msg.delivery = nil
			}
			return err
		}
	}
	return nil
}
