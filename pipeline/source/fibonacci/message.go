package fibonacci

import (
	"fmt"
	"log/slog"
	"math/big"
)

// Message is an implementation of the Message interface
// for messages taken from a Fibonacci series, which require no Ack().
type Message struct {
	Value *big.Int
}

// String returns a string representation of the Message.
func (m *Message) String() string {
	return fmt.Sprintf("%d", m.Value)
}

// Ack is a no-ops implementation of the Ack() function.
func (m *Message) Ack(multiple bool) error {
	slog.Debug("message acknowledged", "message", m)
	return nil
}
