package file

import (
	"log/slog"
)

// Message is an implementation of the Message interface
// for messages taken from a file, which requires no Ack().
type Message struct {
	Value string
}

// String returns a string representation of the Message.
func (m *Message) String() string {
	return m.Value
}

// Ack is a no-ops implementation of the Ack() function.
func (m *Message) Ack(multiple bool) error {
	slog.Debug("message acknowledged", "value", m.Value)
	return nil
}
