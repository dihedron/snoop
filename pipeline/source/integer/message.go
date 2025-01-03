package integer

import (
	"fmt"
	"log/slog"
)

// Message is an implementation of the Message interface
// for messages taken from a Fibonacci series, which require no Ack().
type Message int64

// String returns a string representation of the Message.
func (m *Message) String() string {
	return fmt.Sprintf("%d", int64(*m))
}

// Ack is a no-ops implementation of the Ack() function.
func (m *Message) Ack(multiple bool) error {
	slog.Debug("message acknowledged", "value", m)
	return nil
}
