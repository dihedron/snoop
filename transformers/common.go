package transformers

import (
	"fmt"
	"io"
	"log/slog"
	"sync/atomic"
	"time"

	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/transform"
)

// Profile computes the time elapsed before the first and the second
// occurrence in a Transformer chain; the caller must allocate the input
// values: the first call must have a valid reference to the start
// variable and a nil reference to the elapsed; the second call must
// have both references pointing to valid variables; the elapsed variable
// will be populated with the time between the two calls. This filter
// does not affect the value flowing through.
func Profile[T any](start *time.Time, elapsed *time.Duration) transform.X[T, T] {
	return func(value T) (T, error) {
		if elapsed == nil {
			*start = time.Now()
			slog.Debug("profile start", "value", value, "type", format.TypeAsString(value))
		} else {
			*elapsed = time.Since(*start)
			slog.Debug("profile end", "elapsed", elapsed.String(), "value", value, "type", format.TypeAsString(value))
		}

		return value, nil
	}
}

// Delay inserts a configurable delay inside the chain. This
// filter does not affect the value flowing through.
func Delay[T any](delay time.Duration) transform.X[T, T] {
	return func(value T) (T, error) {
		time.Sleep(delay)
		return value, nil
	}
}

// Count counts the items that flow through the chain. This
// filter does not affect the value flowing through.
func Count[T any](count *int64) transform.X[T, T] {
	return func(value T) (T, error) {
		atomic.AddInt64(count, 1)
		return value, nil
	}
}

// Record records the messages to the given writer. This filter
// does not affect the value flowing through.
func Record[T any](writer io.Writer, format string, lenient bool) transform.X[T, T] {
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

// Skip skips the value if the condition is true. This filter
// does not affect the value flowing through.
func Skip[T any](cond func(value T) bool) transform.X[T, T] {
	return func(value T) (T, error) {
		if cond(value) {
			var nihil T
			return nihil, transform.Drop
		}
		return value, nil
	}
}

// Accept accepts the value if the condition is true. This
// filter does not affect the value flowing through.
func Accept[T any](cond func(value T) bool) transform.X[T, T] {
	return func(value T) (T, error) {
		if cond(value) {
			return value, nil
		}
		var nihil T
		return nihil, transform.Drop
	}
}

// Accumulate adds the value to the given buffer. This filter
// does not affect the value flowing through.
func Accumulate[T any](buffer *[]T) transform.X[T, T] {
	return func(value T) (T, error) {
		*buffer = append(*buffer, value)
		return value, nil
	}
}
