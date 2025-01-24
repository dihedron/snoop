package processor

import (
	"errors"
	"fmt"
	"io"
	"sync/atomic"
	"time"
)

// Handler is the type of a value processor.
type Handler[T any] func(value T) (T, error)

// type Error struct {
// 	message string
// 	wrapped error
// }

// func (e *Error) Error() string {
// 	return fmt.Sprintf(e.message, e.wrapped)
// }

var (
	//lint:ignore ST1012 skip is not a real error, more a wrapper of one.
	skip = errors.New("handler requests to skip value and continue")
	//lint:ignore ST1012 abort is not a real error, more a wrapper of one.
	abort = errors.New("handler requests to abort processing and exit")
)

// Chain creates a chain of Handlers that will be executed one after the
// other, until one handler returns and error or the chain is completely
// applied.
func Chain[T any](handlers ...Handler[T]) Handler[T] {
	return func(value T) (T, error) {
		var (
			err   error
			nihil T
		)
		for _, handler := range handlers {
			value, err = handler(value)
			if err != nil {
				if errors.Is(err, skip) {
					// slog.Warn("handler requested to skip further processing", "error", err)
					return value, nil
				} else if errors.Is(err, abort) {
					// slog.Error("handler requested to abort processing", "error", err)
					return nihil, err
				} else {
					// slog.Error("unexpected error, this is probably a bug", "error", err)
					panic(fmt.Errorf("unexpected error, this is probably a bug: %w", err))
				}
			}
		}
		return value, nil
	}
}

// Profile computes the time elapsed before the first and the second
// occurrence in a Processor chain; the caller must allocate the input
// values: the first call must have a valid reference to the start
// variable and a nil reference to the elapsed; the second call must
// have both references pointing to valid variables; the elapsed variable
// will be populated with the time between the two calls.
func Profile[T any](start *time.Time, elapsed *time.Duration) Handler[T] {
	return func(value T) (T, error) {
		if elapsed == nil {
			*start = time.Now()
		} else {
			*elapsed = time.Since(*start)
		}
		//t.Logf("value flowing through: %v (type: %T)\n", value, value)
		return value, nil
	}
}

// Delay inserts a configurable delay inside the chain.
func Delay[T any](delay time.Duration) Handler[T] {
	return func(value T) (T, error) {
		time.Sleep(delay)
		return value, nil
	}
}

// Count counts the items that flow through the chain.
func Count[T any](count *int64) Handler[T] {
	return func(value T) (T, error) {
		atomic.AddInt64(count, 1)
		return value, nil
	}
}

// Record records the messages to the given writer.
func Record[T any](writer io.Writer, format string, lenient bool) Handler[T] {
	return func(value T) (T, error) {
		_, err := writer.Write([]byte(fmt.Sprintf(format, value)))
		if err != nil {
			if !lenient {
				// slog.Error("error writing value", "error", err)
				return value, err
			}
			// slog.Warn("ignored error writing value", "error", err)
		}
		return value, nil
	}
}

// Skip skips the value if the condition is true.
func Skip[T any](cond func(value T) bool) Handler[T] {
	return func(value T) (T, error) {
		var nihil T
		if cond(value) {
			return nihil, skip
		}
		return value, nil
	}
}

// Accept accepts the value if the condition is true.
func Accept[T any](cond func(value T) bool) Handler[T] {
	return func(value T) (T, error) {
		var nihil T
		if cond(value) {
			return value, nil
		}
		return nihil, skip
	}
}

// Accumulate adds the value to the given buffer.
func Accumulate[T any](buffer *[]T) Handler[T] {
	return func(value T) (T, error) {
		*buffer = append(*buffer, value)
		return value, nil
	}
}
