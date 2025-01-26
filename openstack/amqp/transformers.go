package amqp

import (
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/rabbitmq/amqp091-go"
)

// DeliveryToMessage is a transformer that creates a Message from a
// RabbitMQ amqp091.Delivery message.
func DeliveryToMessage(includeBackRef bool) func(*amqp091.Delivery) (*Message, error) {
	return func(delivery *amqp091.Delivery) (*Message, error) {
		if delivery == nil {
			slog.Error("input must not be nil")
			return nil, errors.New("invalid input") // was: ErrInvalidInput
		}
		message := &Message{
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
		if includeBackRef {
			slog.Debug("adding back-reference to original AMQP delivery", "reference", delivery.DeliveryTag)
			message.backref = delivery
		}
		return message, nil
	}
}

// JSONToMessage is a transformer that creates a new Message by
// parsing an input []byte containing a JSON-serialised version
// of the original message.
func JSONToMessage() func([]byte) (*Message, error) {
	return func(data []byte) (*Message, error) {
		if len(data) == 0 {
			slog.Error("input must not be empty")
			return nil, errors.New("invalid input")
		}
		message := &Message{}
		if err := json.Unmarshal(data, message); err != nil {
			slog.Error("error parsing AMQP delivery from JSON", "error", err)
			return nil, err
		}
		return message, nil
	}
}
