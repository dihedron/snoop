package transformers

import (
	"github.com/dihedron/snoop/transform"
)

type Accumulator[T any] struct {
	values []T
}

// Add adds the value flowing into the transformer to an internal
// buffer. This filter does not affect the value flowing through.
func (a *Accumulator[T]) Add() transform.X[T, T] {
	return func(value T) (T, error) {
		a.values = append(a.values, value)
		return value, nil
	}
}

// Values returns the accumulated values
func (a *Accumulator[T]) Values() []T {
	return a.values
}
