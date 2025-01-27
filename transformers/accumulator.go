package transformers

import (
	"github.com/dihedron/snoop/transform"
)

// Accumulate adds the value to the given buffer. This filter
// does not affect the value flowing through.
func Accumulate[T any](buffer *[]T) transform.X[T, T] {
	return func(value T) (T, error) {
		*buffer = append(*buffer, value)
		return value, nil
	}
}
