package rabbitmq

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
)

// RabbitMQ contains all information about RabbitMQ connection and topology.
type RabbitMQ struct {
	// Client contains info about the RabbitMQ client.
	Client Client `json:"client" yaml:"client" mapstructure:"client" validate:"required,dive,required"`
	// Servers is the set of RabbitMQ servers in the cluster to connect to.
	Servers []Server `json:"servers" yaml:"servers" mapstructure:"servers" validate:"required,dive,required"`
	// Queue is the queue to use.
	Queue Queue `json:"queue" yaml:"queue" mapstructure:"queue" validate:"required,dive,required"`
	// Bindings is the set of bindings to establish the RabbitMQ topology.
	Bindings []Binding `json:"bindings" yaml:"bindings" mapstructure:"bindings" validate:"required,dive,required"`
}

func (r *RabbitMQ) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

// Client contains information about the client connecting to RabbitMQ.
type Client struct {
	// ID is the id of the client connecting to RabbitMQ.
	ID string `json:"id" yaml:"id" mapstructure:"id" validate:"required"`
	// Tag is the tag used by the client connecting to RabbitMQ.
	Tag string `json:"tag" yaml:"tag" mapstructure:"tag" validate:"required"`
}

// Server contains all the information needed to connect to a RabbitMQ server.
type Server struct {
	// Address is the name or IP address of the RabbitMQ host.
	Address string `json:"address" yaml:"address" mapstructure:"address" validate:"required"`
	// Port is the port on which RabbitMQ is listening.
	Port uint16 `json:"port" yaml:"port" mapstructure:"port" validate:"required,gte=0,lte=65535"`
	// Username is the username to use to connect to the RabbitMQ server.
	Username *string `json:"username,omitempty" yaml:"username,omitempty" mapstructure:"username"`
	// Password is the password to use to connect to the RabbitMQ server.
	Password *string `json:"password,omitempty" yaml:"password,omitempty" mapstructure:"password,omitempty"`
	// TLSInfo contains the information to configure TLS on the connection
	// to the RabbitMQ server.
	TLSInfo *TLSInfo `json:"tlsinfo,omitempty" yaml:"tlsinfo,omitempty" mapstructure:"tlsinfo,omitempty"`
}

// TLSInfo contains the information needed to set-up a TLS endpoint
// or connection, such as a private key/certificate pair; it should
// be embedded as a pointer into any relevan configuration struct, so
// that whin nil it implies that TLS is not enabled.
type TLSInfo struct {
	// Enables specifies whether TLS support should be enabled.
	EnableTLS bool `json:"enabled,omitempty" yaml:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// SkipVerify specifies whether the certificate verification should
	// be skipped (allow invalid server- and client-side certificates).
	SkipVerify bool `json:"skipverify,omitempty" yaml:"skipverify,omitempty" mapstructure:"skipverify,omitempty"`
	// CATrustAnchor is the path to the cacert.pem file containing the CA
	// certificates to use as implied trusted certificates for TLS connections.
	CATrustAnchor string `json:"cacert,omitempty" yaml:"cacert,omitempty" mapstructure:"cacert,omitempty"`
	// PrivateKey is the path to the key.pem file containing the private
	// key to use for connections.
	//PrivateKey string `json:"privatekey,omitempty" yaml:"privatekey,omitempty" mapstructure:"privatekey,omitempty" validate:"required,file"`
	PrivateKey string `json:"privatekey,omitempty" yaml:"privatekey,omitempty" mapstructure:"privatekey,omitempty"`
	// Certificate is the path to the cert.pem file containing the certificate
	// to use for TLS connections..
	//Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty" mapstructure:"certificate,omitempty" validate:"required,file"`
	Certificate string `json:"certificate,omitempty" yaml:"certificate,omitempty" mapstructure:"certificate,omitempty"`
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

// StringToExchangeTypeHookFunc is used to parse and ExchangeType from its
// string representation when using mapstructure.
func StringToExchangeTypeHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(ExchangeTypeFanout) {
			return data, nil
		}
		var e ExchangeType
		err := e.Parse(data.(string))
		if err != nil {
			return nil, err
		}
		return e, nil
	}
}

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

// String is a utility function that converts a string value into
// an optional string by returning it as a pointer (thus nillable).
func String(value string) *string {
	return &value
}
