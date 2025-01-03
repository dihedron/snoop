package recorder

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"github.com/dihedron/snoop/pipeline"
)

// Recorder is a filter that records all messages to a writer
// before moving on to the next filter in the pipeline.
type Recorder struct {
	writer       io.Writer
	abortOnError bool
}

func New(writer io.Writer, abortOnError bool) *Recorder {
	return &Recorder{
		writer:       writer,
		abortOnError: abortOnError,
	}
}

func (d *Recorder) Name() string {
	return "github.com/dihedron/snoop/pipeline/filter/recorder/Recorder"
}

func (r *Recorder) Process(ctx context.Context, message pipeline.Message) (context.Context, pipeline.Message, error) {
	select {
	case <-ctx.Done():
		slog.Debug("context cancelled")
		return ctx, message, pipeline.ErrAbort
	default:
		slog.Debug("writing message to output", "type", fmt.Sprintf("%T", message))
		_, err := r.writer.Write([]byte(fmt.Sprintf("%s\n", message)))
		if err != nil {
			slog.Error("error writing message", "type", fmt.Sprintf("%T", message), "error", err)
			if r.abortOnError {
				// TODO: wrap the real error somehow???
				return ctx, message, pipeline.ErrAbort
			}
		}
	}
	return ctx, message, nil
}
