package oslo

import (
	"errors"
	"log/slog"

	"github.com/goccy/go-json"

	"github.com/dihedron/snoop/openstack/amqp"
)

// MessageToOslo extracts an Oslo message from a Message payload.
func MessageToOslo(includeBackRef bool) func(*amqp.Message) (*Oslo, error) {
	return func(message *amqp.Message) (*Oslo, error) {
		if message == nil {
			slog.Error("input must not be nil")
			return nil, errors.New("invalid input") //ErrInvalidInput
		}
		oslo, err := JSONToOslo()(message.Body)
		if err == nil && oslo != nil && includeBackRef && message.BackRef() != nil {
			slog.Debug("adding back-reference to original AMQP delivery", "reference", message.BackRef().DeliveryTag)
			oslo.backref = message.BackRef()
		}
		return oslo, err
	}
}

// JSONToOslo is a transformer that extracts an Oslo message from
// a JSON representation of the same (e.g. read from a text file).
func JSONToOslo() func(data []byte) (*Oslo, error) {
	return func(data []byte) (*Oslo, error) {
		oslo := struct {
			Version string `json:"oslo.version" yaml:"oslo.version"`
			Payload string `json:"oslo.message" yaml:"oslo.message"`
		}{}
		if err := json.Unmarshal(data, &oslo); err != nil {
			slog.Error("error parsing Oslo message", "error", err)
			return nil, err
		}
		slog.Debug("oslo message parsed", "version", oslo.Version)
		return &Oslo{
			Version: oslo.Version,
			Payload: oslo.Payload,
		}, nil
	}
}
