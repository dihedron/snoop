package transformers

import (
	"fmt"

	"github.com/dihedron/snoop/transform"
)

// Counter holds the state needed to count items across multiple
// invocations of the chain. If Every is specified and more than
// 0, it will print a dot to STDOUT every count%Every values
type Counter[T any] struct {
	Every int64
	count int64
}

// Add adds 1 to the count of items flowing through the chain. This
// filter does not affect the value flowing through.
func (c *Counter[T]) Add() transform.X[T, T] {
	return func(value T) (T, error) {
		c.count = c.count + 1
		if c.Every > 0 && c.count%c.Every == 0 {
			fmt.Printf(". ")
		}
		return value, nil
	}
}

// Count returns the count of items.
func (c *Counter[T]) Count() int64 {
	return c.count
}
