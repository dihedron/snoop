package transformers

import "github.com/dihedron/snoop/transform/chain"

// Cache holds a map into which values can be accumulated under
// keys dynamically computed through a user-provided function.
type Cache[K comparable, T any] struct {
	cache map[K]T
}

// Set adds the item to the cache by first applying a function
// that extracts the key under which the item will be added
// and then adding it to an internal map.
func (c *Cache[K, T]) Set(keyer func(value T) K) chain.X[T, T] {
	return func(value T) (T, error) {
		key := keyer(value)
		if c.cache == nil {
			c.cache = map[K]T{}
		}
		c.cache[key] = value
		return value, nil
	}
}

// Contents returns the contents of the cache.
func (c *Cache[K, T]) Contents() map[K]T {
	return c.cache
}
