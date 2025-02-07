package transformers

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/dihedron/snoop/transform"
)

// Record records the messages to the given writer. This filter
// does not affect the value flowing through.
func Record[T any](writer io.Writer, format string, lenient bool) transform.X[T, T] {
	return RecordIf[T](writer, format, lenient, func(value T) bool { return true })
}

// RecordIf records the messages to the given writer if the given
// condition is true. This filter does not affect the value flowing
// through.
func RecordIf[T any](writer io.Writer, format string, lenient bool, condition func(value T) bool) transform.X[T, T] {
	if format == "" {
		format = "%v"
	}
	return func(value T) (T, error) {
		if condition(value) {
			_, err := writer.Write([]byte(fmt.Sprintf(format, value)))
			if err != nil {
				if !lenient {
					slog.Error("error writing value", "error", err)
					return value, err
				}
				slog.Warn("ignored error writing value", "error", err)
			}
		}
		return value, nil
	}
}

// RecordUnless records the messages to the given writer unless the given
// condition is true. This filter does not affect the value flowing through.
func RecordUnless[T any](writer io.Writer, format string, lenient bool, condition func(value T) bool) transform.X[T, T] {
	return func(value T) (T, error) {
		if !condition(value) {
			_, err := writer.Write([]byte(fmt.Sprintf(format, value)))
			if err != nil {
				if !lenient {
					slog.Error("error writing value", "error", err)
					return value, err
				}
				slog.Warn("ignored error writing value", "error", err)
			}
		}
		return value, nil
	}
}
