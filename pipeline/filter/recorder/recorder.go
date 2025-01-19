package recorder

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/dihedron/snoop/pipeline"
)

// Recorder is a filter that records all messages to a writer
// before moving on to the next filter in the pipeline.
type Recorder[T any] struct {
	writer  io.Writer
	format  string
	lenient bool
}

func New[T any](writer io.Writer, format string, lenient bool) *Recorder[T] {
	r := &Recorder[T]{
		writer:  writer,
		format:  "%v\n",
		lenient: lenient,
	}
	if format != "" {
		r.format = format
	}
	return r
}

func (r *Recorder[T]) Apply(value T) (T, error) {
	slog.Debug("writing to output", "value", fmt.Sprintf(r.format, value), "type", fmt.Sprintf("%T", value))
	_, err := r.writer.Write([]byte(fmt.Sprintf(r.format, value)))
	if err != nil {
		if !r.lenient {
			slog.Error("error writing value", "error", err)
			return value, err
		}
		slog.Warn("ignored error writing value", "error", err)
	}
	return value, nil
}

func RecorderFunc[T any](writer io.Writer, format string, lenient bool) pipeline.Handler[T] {
	return func(value T) (T, error) {
		_, err := writer.Write([]byte(fmt.Sprintf(format, value)))
		if err != nil {
			if !lenient {
				slog.Error("error writing value", "error", err)
				return value, err
			}
			slog.Warn("ignored error writing value", "error", err)
		}
		return value, nil
	}
}
