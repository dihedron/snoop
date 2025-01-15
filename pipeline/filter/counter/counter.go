package counter

import (
	"log/slog"
	"sync/atomic"
)

// Counter is a filter that counts all messages as they pass.
type Counter struct {
	count int64
}

// NewCounter creates a new counter.
func New() *Counter {
	return &Counter{}
}

func (*Counter) Name() string {
	return "github.com/dihedron/snoop/pipeline/filter/counter/Counter"
}

func (c *Counter) Count() int64 {
	return atomic.LoadInt64(&c.count)
}

func (c *Counter) Process(message any) (any, error) {
	slog.Debug("adding one to message count")
	atomic.AddInt64(&(c.count), 1)
	return message, nil
}
