package openstack

import "github.com/rabbitmq/amqp091-go"

type Acknowledgeable interface {
	Ack(multiple bool) error
	BackRef() *amqp091.Delivery
}
