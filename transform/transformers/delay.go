package transformers

import (
	"time"

	"github.com/dihedron/snoop/transform/chain"
)

// Delay inserts a configurable delay inside the chain. This
// filter does not affect the value flowing through.
func Delay[T any](delay time.Duration) chain.X[T, T] {
	return func(value T) (T, error) {
		time.Sleep(delay)
		return value, nil
	}
}
