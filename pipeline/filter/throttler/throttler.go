package throttler

import (
	"context"
	"log/slog"
	"time"

	"github.com/dihedron/snoop/pipeline"
)

// Throttler is a filter that delays message processing
// by inserting a delay between messages.
type Throttler struct {
	delay time.Duration
}

func New(delay time.Duration) *Throttler {
	return &Throttler{delay: delay}
}

func (d *Throttler) Name() string {
	return "github.com/dihedron/snoop/pipeline/filter/throttler/Throttler"
}

func (d *Throttler) Process(ctx context.Context, message pipeline.Message) (context.Context, pipeline.Message, error) {
	select {
	case <-ctx.Done():
		slog.Debug("context cancelled")
		return ctx, message, pipeline.ErrAbort
	default:
		slog.Debug("throttling before forwarding message")
		time.Sleep(d.delay)
	}
	return ctx, message, nil
}
