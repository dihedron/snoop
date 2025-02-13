package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"log/slog"
	"strings"
	"sync"
	"time"

	"github.com/goccy/go-json"
	amqp091 "github.com/rabbitmq/amqp091-go"

	"github.com/go-playground/validator/v10"
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

// RabbitMQ contains all information about RabbitMQ connection and topology.
type RabbitMQ struct {
	// Client contains info about the RabbitMQ client.
	Client Client `json:"client" yaml:"client" validate:"required"`
	// Servers is the set of RabbitMQ servers in the cluster to connect to.
	Servers []Server `json:"servers" yaml:"servers" validate:"required,dive,required"`
	// Queue is the queue to use.
	Queue Queue `json:"queue" yaml:"queue" validate:"required"`
	// Bindings is the set of bindings to establish the RabbitMQ topology.
	Bindings []Binding `json:"bindings" yaml:"bindings" validate:"required,dive,required"`
	// err is the internal field keeping track of errors.
	err error
}

func (r *RabbitMQ) All(ctx context.Context) iter.Seq[amqp091.Delivery] {
	slog.Debug("starting to collect messages from RabbitMQ")
	r.err = nil
	if ctx == nil {
		slog.Debug("no context provided, allocating default context...")
		ctx = context.Background()
	}

	if err := r.validate(); err != nil {
		slog.Error("invalid configuration", "error", err)
		r.err = err
		return nil
	}
	return func(yield func(amqp091.Delivery) bool) {
		// gather the URLs of the servers
		urls := []string{}
		for _, server := range r.Servers {
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
		for _, binding := range r.Bindings {
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
			"name", r.Queue.Name,
			"declare", r.Queue.Declare,
			"durable", r.Queue.Durable,
			"exclusive", r.Queue.Exclusive,
			"autodelete", r.Queue.AutoDelete,
		)

		options := &rabbit.Options{
			URLs:              urls,
			Mode:              rabbit.Consumer,
			QueueName:         r.Queue.Name,
			QueueDeclare:      r.Queue.Declare,
			QueueDurable:      r.Queue.Durable,
			QueueExclusive:    r.Queue.Exclusive,
			QueueAutoDelete:   r.Queue.AutoDelete,
			Bindings:          binds,
			QosPrefetchCount:  DefaultQosPrefetchCount,
			QosPrefetchSize:   DefaultQosPrefetchSize,
			RetryReconnectSec: DefaultReconnectSec,
			AppID:             DefaultClientID,
			ConsumerTag:       DefaultClientID,
			ConnectionTimeout: r.Client.Timeout,
		}
		if r.Client.ID != "" {
			options.AppID = r.Client.ID
		}
		if r.Client.Tag != "" {
			options.ConsumerTag = r.Client.Tag
		}
		slog.Info("RabbitMQ source ready", "client id", r.Client.ID, "tag", r.Client.Tag)

		queue, err := rabbit.New(options)
		if err != nil {
			slog.Error("unable to instantiate RabbitMQ client", "error", err)
			r.err = err
			return
		}
		slog.Info("RabbitMQ client ready to drain messages")

		experimental := true
		if experimental {
			slog.Debug("EXPERIMENTAL: retrieving events with an interposed channel")

			values := make(chan amqp091.Delivery, 10)
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				queue.Consume(ctx, nil, func(message amqp091.Delivery) error {
					slog.Debug("enqueuing amqp091.Delivery as message", "value", message)
					select {
					case values <- message:
						slog.Debug("message enqueued")
					case <-ctx.Done():
						slog.Debug("context cancelled, no more enqueuing and exiting")
					default:
						slog.Debug("cannot enqueue message")
					}
					return nil
				}, rabbit.DefaultAckPolicy())
				close(values)
				wg.Done()
			}()
		loop:
			for {
				select {
				case message, ok := <-values:
					if !ok || !yield(message) {
						slog.Info("stop generating messages for iterator")
						break loop
					}
				case <-ctx.Done():
					slog.Info("context done")
					break loop
				}
			}
			slog.Debug("waiting for queue consumer to exit...")
			wg.Wait()
			slog.Debug("queue consumer exited")
		} else {
			// TODO: this one seems to work best
			slog.Debug("retrieving events without an interposed channel")
			queue.Consume(ctx, nil, func(message amqp091.Delivery) error {
				slog.Debug("sending amqp091.Delivery as message", "value", message)
				if !yield(message) {
					slog.Info("stop sending messages (cancel context)")
					// TODO: check if thins works in a highly concurrent context
					// where the select() within consume() might not read from
					// the cancel channel and read some more messages from the
					// RabbitMQ input queue; in that case, this will panic. It might
					// be necessary to:
					// 1. implement a queue that sends messages from this inner func
					//    to the iterator
					// 2. have the iterator dequeue messages, check if the yield() func
					//    return false, signal this func() that it's over and exit immediately
					cancel()
					return nil
				}
				return nil
			}, rabbit.DefaultAckPolicy())
		}
	}
}

func (r *RabbitMQ) Err() error {
	return r.err
}

func (r *RabbitMQ) validate() error {
	validate := validator.New()
	return validate.Struct(*r)
}

// Client contains information about the client connecting to RabbitMQ.
type Client struct {
	// ID is the id of the client connecting to RabbitMQ.
	ID string `json:"id" yaml:"id" validate:"required"`
	// Tag is the tag used by the client connecting to RabbitMQ.
	Tag string `json:"tag" yaml:"tag" validate:"required"`
	// Timeout is the timeout for connections to the server.
	Timeout time.Duration `json:"timeout" yaml:"timeout"`
}

// Server contains all the information needed to connect to a RabbitMQ server.
type Server struct {
	// Address is the name or IP address of the RabbitMQ host.
	Address string `json:"address" yaml:"address" validate:"required"`
	// Port is the port on which RabbitMQ is listening.
	Port uint16 `json:"port" yaml:"port" validate:"required,gte=0,lte=65535"`
	// Username is the username to use to connect to the RabbitMQ server.
	Username *string `json:"username,omitempty" yaml:"username,omitempty"`
	// Password is the password to use to connect to the RabbitMQ server.
	Password *string `json:"password,omitempty" yaml:"password,omitempty"`
	// TLSInfo contains the information to configure TLS on the connection
	// to the RabbitMQ server.
	TLSInfo *TLSInfo `json:"tlsinfo,omitempty" yaml:"tlsinfo,omitempty"`
}

// TLSInfo contains the information needed to set-up a TLS endpoint
// or connection, such as a private key/certificate pair; it should
// be embedded as a pointer into any relevan configuration struct, so
// that whin nil it implies that TLS is not enabled.
type TLSInfo struct {
	// Enables specifies whether TLS support should be enabled.
	EnableTLS bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	// SkipVerify specifies whether the certificate verification should
	// be skipped (allow invalid server- and client-side certificates).
	SkipVerify bool `json:"skipverify,omitempty" yaml:"skipverify,omitempty"`
	// CATrustAnchor is the path to the cacert.pem file containing the CA
	// certificates to use as implied trusted certificates for TLS connections.
	CATrustAnchor string `json:"cacert,omitempty" yaml:"cacert,omitempty"`
	// PrivateKey is the path to the key.pem file containing the private
	// key to use for connections.
	//PrivateKey string `json:"privatekey,omitempty" yaml:"privatekey,omitempty" mapstructure:"privatekey,omitempty" validate:"required,file"`
	PrivateKey string `json:"privatekey,omitempty" yaml:"privatekey,omitempty"`
	// Certificate is the path to the cert.pem file containing the certificate
	// to use for TLS connections..
	//Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty" mapstructure:"certificate,omitempty" validate:"required,file"`
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty"`
}

// Validate checks if the struct is valid.
func (t *TLSInfo) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

// ExchangeType is the type of exchange.
type ExchangeType int8

const (
	// ExchangeTypeFanout identifies the "fanout" exchange.
	ExchangeTypeFanout ExchangeType = iota
	// ExchangeTypeTopic identifies the "fanout" exchange.
	ExchangeTypeTopic
	// ExchangeTypeDirect identifies the "direct" exchange.
	ExchangeTypeDirect
	// ExchangeTypeHeaders identifies the "headers" exchange.
	ExchangeTypeHeaders
)

// String converts an ExchangeType into its string representation.
func (e ExchangeType) String() string {
	return []string{
		"fanout",
		"topic",
		"direct",
		"headers",
	}[e]
}

// Parse parses a string into an ExchangeType value.
func (e *ExchangeType) Parse(value string) error {
	switch strings.ToLower(value) {
	case "fanout":
		*e = ExchangeTypeFanout
	case "topic":
		*e = ExchangeTypeTopic
	case "direct":
		*e = ExchangeTypeDirect
	case "headers":
		*e = ExchangeTypeHeaders
	default:
		return errors.New("unsupported ExchangeType value")
	}
	return nil
}

// MarshalYAML marshals the ExchangeType value into a YAML string.
func (e ExchangeType) MarshalYAML() (interface{}, error) {
	v := e.String()
	if v != "" {
		return v, nil
	}
	return "", errors.New("unsupported ExchangeType value")
}

// UnmarshalYAML unmarshals a YAML value into an ExchangeType value.
func (e *ExchangeType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	size := ""
	err := unmarshal(&size)
	if err != nil {
		return err
	}
	return e.Parse(size)
}

// MarshalJSON marshals the ExchangeType value into a JSON string.
func (e ExchangeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON unmarshals a JSON value into an ExchangeType value.
func (e *ExchangeType) UnmarshalJSON(b []byte) error {
	var value string
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}
	e.Parse(value)
	return nil
}

// // StringToExchangeTypeHookFunc is used to parse and ExchangeType from its
// // string representation when using mapstructure.
// func StringToExchangeTypeHookFunc() mapstructure.DecodeHookFunc {
// 	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
// 		if f.Kind() != reflect.String {
// 			return data, nil
// 		}
// 		if t != reflect.TypeOf(ExchangeTypeFanout) {
// 			return data, nil
// 		}
// 		var e ExchangeType
// 		err := e.Parse(data.(string))
// 		if err != nil {
// 			return nil, err
// 		}
// 		return e, nil
// 	}
// }

// Exchange contains information about a RabbitMQ exchange.
type Exchange struct {
	Name       string       `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" validate:"required"`
	Type       ExchangeType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty" validate:"gte=0,lte=3"`
	Durable    bool         `json:"durable,omitempty" yaml:"durable,omitempty" mapstructure:"durable,omitempty"`
	Declare    bool         `json:"declare,omitempty" yaml:"declare,omitempty" mapstructure:"declare,omitempty"`
	AutoDelete bool         `json:"autodelete,omitempty" yaml:"autodelete,omitempty" mapstructure:"autodelete,omitempty"`
}

// Queue contains information about a RabbitMQ exchange.
type Queue struct {
	Name       string `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty" validate:"required"`
	Durable    bool   `json:"durable,omitempty" yaml:"durable,omitempty" mapstructure:"durable,omitempty"`
	Declare    bool   `json:"declare,omitempty" yaml:"declare,omitempty" mapstructure:"declare,omitempty"`
	Exclusive  bool   `json:"exclusive,omitempty" yaml:"exclusive,omitempty" mapstructure:"exclusive,omitempty"`
	AutoDelete bool   `json:"autodelete,omitempty" yaml:"autodelete,omitempty" mapstructure:"autodelete,omitempty"`
	// QosPrefetchCount:  rabbitmq.DefaultQosPrefetchCount,
	// QosPrefetchSize:   rabbitmq.DefaultQosPrefetchSize,
	// RetryReconnectSec: rabbitmq.DefaultReconnectSec,
}

// Binding is the exchange and routing key(s) to use for connecting to RabbitMQ
type Binding struct {
	// Exchange is the name of the RabbitMQ exchange to connect to.
	Exchange *Exchange `json:"exchange,omitempty" yaml:"exchange,omitempty" mapstructure:"exchange,omitempty"`
	// RoutingKeys is the set of routing keys to use on the given exchange.
	RoutingKeys []string `json:"routingkeys,omitempty" yaml:"routingkeys,omitempty" mapstructure:"routingkeys,omitempty" validate:"required"`
}
