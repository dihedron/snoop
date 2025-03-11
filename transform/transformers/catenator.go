package transformers

import (
	"fmt"
	"strings"

	"github.com/dihedron/snoop/transform/chain"
)

// Catenator accumulates the values flowing through the transformer
// in their string representation, and then provides a way to join
// them with an optional separator.
type Catenator[T any] struct {
	Join   string
	values []string
}

// Add adds the value flowing into the transformer to an internal
// buffer after converting it to a string (using fmt.Sprintf("%v")).
// This filter does not affect the value flowing through.
func (c *Catenator[T]) Add() chain.X[T, T] {
	return func(value T) (T, error) {
		c.values = append(c.values, fmt.Sprintf("%v", value))
		return value, nil
	}
}

// Values returns the accumulated values, joined with the optional string.
func (c *Catenator[T]) Value() string {
	return strings.Join(c.values, c.Join)
}
