package transformers

import "github.com/dihedron/snoop/transform"

// Counter holds the state needed to count items across multiple
// invocations of the chain.
type Counter[T any] struct {
	count int64
}

// Add adds 1 to the count of items flowing through the chain. This
// filter does not affect the value flowing through.
func (c *Counter[T]) Add() transform.X[T, T] {
	return func(value T) (T, error) {
		c.count = c.count + 1
		return value, nil
	}
}

// Count returns the count of items.
func (c *Counter[T]) Count() int64 {
	return c.count
}
