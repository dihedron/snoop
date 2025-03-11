package openstack

import (
	"log/slog"
)

// Acknowledge is a transformer that acknowledges a RabbitMQ delivery.
func Acknowledge(multiple bool) func(Acknowledgeable) (Acknowledgeable, error) {
	return func(ack Acknowledgeable) (Acknowledgeable, error) {
		slog.Info("acknowledging AMQP delivery", "reference", ack.BackRef().DeliveryTag, "multiple", multiple)
		return ack, ack.Ack(multiple)
	}
}
