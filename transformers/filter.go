package transformers

import "github.com/dihedron/snoop/transform"

// Filter lets the value flow if the condition is true. If the
// condition is true,  this filter does not affect the value
// flowing through.
func Filter[T any](cond func(value T) bool) transform.X[T, T] {
	return func(value T) (T, error) {
		if cond(value) {
			return value, nil
		}
		var nihil T
		return nihil, transform.Drop
	}
}
