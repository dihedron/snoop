package transformers

import (
	"fmt"

	"github.com/dihedron/snoop/transform/chain"
)

// MultiCounter holds the state needed to group items and count them
// across multiple invocations of the chain. If Every is specified and
// more than 0, it will print a dot to STDOUT every total%Every values
type MultiCounter[T any, K comparable] struct {
	Every  int64
	counts map[K]int64
	total  int64
}

// Add adds 1 to the count of items flowing through the chain. This
// filter does not affect the value flowing through.
func (c *MultiCounter[T, K]) Add(keyer func(value T) K) chain.X[T, T] {
	return c.AddIf(keyer, func(value T) bool { return true })
}

// AddIf adds 1 to the count of items flowing through the chain, if
// the given condition returns true. This filter does not affect the
// value flowing through.
func (c *MultiCounter[T, K]) AddIf(keyer func(value T) K, condition func(value T) bool) chain.X[T, T] {
	c.counts = map[K]int64{}
	return func(value T) (T, error) {
		if condition(value) {
			k := keyer(value)
			if _, ok := c.counts[k]; !ok {
				c.counts[k] = 1
			} else {
				c.counts[k] = c.counts[k] + 1
			}
			c.total = c.total + 1
			if c.Every > 0 && c.total%c.Every == 0 {
				fmt.Printf(". ")
			}
		}
		return value, nil
	}
}

// AddUnless adds 1 to the count of items flowing through the chain, unless
// the given condition returns true. This filter does not affect the
// value flowing through.
func (c *MultiCounter[T, K]) AddUnless(keyer func(value T) K, condition func(value T) bool) chain.X[T, T] {
	c.counts = map[K]int64{}
	return func(value T) (T, error) {
		if !condition(value) {
			k := keyer(value)
			if _, ok := c.counts[k]; !ok {
				c.counts[k] = 1
			} else {
				c.counts[k] = c.counts[k] + 1
			}
			c.total = c.total + 1
			if c.Every > 0 && c.total%c.Every == 0 {
				fmt.Printf(". ")
			}
		}
		return value, nil
	}
}

// Count returns the count of items.
func (c *MultiCounter[T, K]) Count() (map[K]int64, int64) {
	return c.counts, c.total
}
