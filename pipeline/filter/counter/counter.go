package counter

import (
	"log/slog"
	"sync/atomic"
)

// Counter is a filter that counts all values as they pass.
type Counter[T any] struct {
	count int64
}

// NewCounter creates a new counter.
func New[T any]() *Counter[T] {
	return &Counter[T]{}
}

func (c *Counter[T]) Count() int64 {
	return atomic.LoadInt64(&c.count)
}

func (c *Counter[T]) Apply(value T) (T, error) {
	atomic.AddInt64(&(c.count), 1)
	slog.Debug("adding one to values count", "count", c.count)
	return value, nil
}

// func CounterFunc[T any]() pipeline.Handler[T] {
// 	var count int64 = 0
// 	return func(value T) (T, error) {
// 		atomic.AddInt64(&(count), 1)
// 		return value, nil
// 	}
// }
