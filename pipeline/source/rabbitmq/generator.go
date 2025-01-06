package rabbitmq

import (
	"context"
	"fmt"
	"iter"
	"log/slog"

	amqp091 "github.com/rabbitmq/amqp091-go"
	"github.com/streamdal/rabbit"
)

func RabbitMQContext(ctx context.Context, configuration *RabbitMQ) iter.Seq[amqp091.Delivery] {
	return func(yield func(amqp091.Delivery) bool) {

		// gather the URLs of the servers
		urls := []string{}
		for _, server := range configuration.Servers {
			proto := ""
			if server.TLSInfo != nil && server.TLSInfo.EnableTLS {
				proto = "amqps"
			} else {
				proto = "amqp"
			}
			url := ""
			if server.Username != nil && server.Password != nil {
				url = fmt.Sprintf("%s://%s:%s@%s:%d/", proto, *server.Username, *server.Password, server.Address, server.Port)
			} else {
				url = fmt.Sprintf("%s://%s:%d/", proto, server.Address, server.Port)
			}
			urls = append(urls, url)
			slog.Info("RabbitMQ server url", "value", url)
		}

		slog.Debug("connecting to RabbitMQ server URLs", "urls", urls)

		binds := []rabbit.Binding{}
		for _, binding := range configuration.Bindings {
			slog.Info("adding exchange with routing keys", "exchange name", binding.Exchange.Name, "routing keys", binding.RoutingKeys)
			binds = append(binds, rabbit.Binding{
				ExchangeName:       binding.Exchange.Name,
				ExchangeType:       binding.Exchange.Type.String(),
				ExchangeDurable:    binding.Exchange.Durable,
				ExchangeDeclare:    binding.Exchange.Declare,
				ExchangeAutoDelete: binding.Exchange.AutoDelete,
				BindingKeys:        binding.RoutingKeys,
			})
		}

		slog.Info(
			"binding to queue",
			"queue name", configuration.Queue.Name,
			"declare", configuration.Queue.Declare,
			"durable", configuration.Queue.Durable,
			"exclusive", configuration.Queue.Exclusive,
			"autodelete", configuration.Queue.AutoDelete,
		)

		options := &rabbit.Options{
			URLs:              urls,
			Mode:              rabbit.Consumer,
			QueueName:         configuration.Queue.Name,
			QueueDeclare:      configuration.Queue.Declare,
			QueueDurable:      configuration.Queue.Durable,
			QueueExclusive:    configuration.Queue.Exclusive,
			QueueAutoDelete:   configuration.Queue.AutoDelete,
			Bindings:          binds,
			QosPrefetchCount:  DefaultQosPrefetchCount,
			QosPrefetchSize:   DefaultQosPrefetchSize,
			RetryReconnectSec: DefaultReconnectSec,
			AppID:             DefaultClientID,
			ConsumerTag:       DefaultClientID,
		}
		if configuration.Client.ID != "" {
			options.AppID = configuration.Client.ID
		}
		if configuration.Client.Tag != "" {
			options.ConsumerTag = configuration.Client.Tag
		}
		slog.Info("configuring source to present as client ID", "client id", configuration.Client.ID, "tag", configuration.Client.Tag)

		slog.Info("RabbitMQ source ready")

		queue, err := rabbit.New(options)
		if err != nil {
			slog.Error("unable to instantiate RabbitMQ client", "error", err)
			return
		}
		slog.Info("RabbitMQ client ready to drain messages")

		queue.Consume(ctx, nil, func(message amqp091.Delivery) error {
			slog.Debug("sending amqp091.Delivery as message", "value", message)
			if !yield(message) {
				slog.Info("stop sending messages")
				return nil
			}
			return nil
		}, rabbit.DefaultAckPolicy())
	}
}
