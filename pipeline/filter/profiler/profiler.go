package profiler

import (
	"context"
	"log/slog"
	"time"

	"github.com/dihedron/snoop/pipeline"
)

//const TimeKey string = "github.com/dihedron/snoop/pipeline/filter/Profiler::"

// Profiler is a filter that keeps track of the duration between its
// first invocation in the pipeline and its successive invocation(s).
type Profiler struct {
	set     bool
	start   time.Time
	elapsed time.Duration
}

// NewStartTime creates a new StartTime; id is the unique identifier
// of the start time in the context.
func New() *Profiler {
	return &Profiler{}
}

func (*Profiler) Name() string {
	return "github.com/dihedron/snoop/pipeline/filter/profiler/Profiler"
}

func (p *Profiler) Elapsed() time.Duration {
	return p.elapsed
}

func (p *Profiler) Process(ctx context.Context, message pipeline.Message) (context.Context, pipeline.Message, error) {
	select {
	case <-ctx.Done():
		slog.Debug("context cancelled")
		return ctx, message, pipeline.ErrAbort
	default:
		if !p.set {
			slog.Debug("recording start time")
			p.start = time.Now()
			p.set = true
		} else {
			p.elapsed = time.Since(p.start)
			slog.Debug("elapsed time", "ms", p.elapsed.Milliseconds())
			p.set = false
		}
	}
	return ctx, message, nil
}
