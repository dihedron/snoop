package counter

import (
	"context"
	"log/slog"
	"sync/atomic"

	"github.com/dihedron/snoop/pipeline"
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
	return "github.com/diherdon/snoop/pipeline/filter/counter/Counter"
}

func (c *Counter) Count() int64 {
	return atomic.LoadInt64(&c.count)
}

func (c *Counter) Process(ctx context.Context, message pipeline.Message) (context.Context, pipeline.Message, error) {
	select {
	case <-ctx.Done():
		slog.Debug("context cancelled")
		return ctx, message, pipeline.ErrAbort
	default:
		slog.Debug("adding one to message count")
		atomic.AddInt64(&(c.count), 1)
	}
	return ctx, message, nil
}
