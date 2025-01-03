package rabbitmq

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/dihedron/snoop/pipeline"
	amqp091 "github.com/rabbitmq/amqp091-go"
	"github.com/streamdal/rabbit"
)

const (
	// DefaultClientID is the constant used to identify the client at the
	// server, both as a producer and a consumer.
	DefaultClientID = "snoop-v1.0.0"
	// DefaultReconnectSec is the delay in seconds between successive
	// attempts to reconnect to the server after a failure.
	DefaultReconnectSec = 5
	// DefaultResendDelay is the delay between attempts to resend messages
	// the server didn't confirm.
	DefaultResendDelay = 5 * time.Second
	// DefaultQosPrefetchCount is the default value of messages that the
	// server should prefetch.
	DefaultQosPrefetchCount = 0
	// DefaultQosPrefetchSize is the default number of bytes that should be
	// prefetched without being acknowledged.
	DefaultQosPrefetchSize = 0
)

// Source is the concrete suorce that reads deliveries
// from a RabbitMQ topology.
type Source struct {
	options *rabbit.Options
	cancel  context.CancelFunc
}

// New creates a new Source.
func New(configuration *RabbitMQ) (*Source, error) {
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
		Log:               nil,
	}
	if configuration.Client.ID != "" {
		options.AppID = configuration.Client.ID
	}
	if configuration.Client.Tag != "" {
		options.ConsumerTag = configuration.Client.Tag
	}
	slog.Info("configuring source to present as client ID", "client id", configuration.Client.ID, "tag", configuration.Client.Tag)
	source := &Source{
		options: options,
	}
	slog.Info("RabbitMQ source ready")
	return source, nil
}

// Emit emits messages one at a time on the return channel.
func (s *Source) Emit(ctx context.Context) (<-chan pipeline.Message, error) {
	var cancellable context.Context
	// wrap the context in a WithCancel() so it can be used by Close()
	cancellable, s.cancel = context.WithCancel(ctx)
	// receive one message at a time from RabbitMQ (because a message must
	// be acknowledged before moving to the next one in a distributed setup)
	messages := make(chan pipeline.Message, 1)

	go func(ctx context.Context) {
		defer func() {
			slog.Info("closing output message channel")
			close(messages)
		}()

		// // An operation that may fail
		// operation := func() error {
		// 	return nil // or an error
		// }

		// err := Retry(operation, NewExponentialBackOff())
		// if err != nil {
		// 	// Handle error.
		// 	return
		// }
		queue, err := rabbit.New(s.options)
		if err != nil {
			slog.Error("unable to instantiate RabbitMQ client", "error", err)
			return
		}
		slog.Info("RabbitMQ client ready to drain messages")
		queue.Consume(ctx, nil, func(message amqp091.Delivery) error {
			//slog.Debug("sending amqp091.Delivery as message", "value", message)
			messages <- &message
			return nil
		})
	}(cancellable)
	return messages, nil
}

// Close closes the underlying queue, if open.
func (s *Source) Close() error {
	s.cancel()
	return nil
}
