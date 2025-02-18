package transformers

import (
	"fmt"
	"io"
	"log/slog"

	"github.com/dihedron/snoop/transform/chain"
)

// Write records the messages to the given writer. This filter
// does not affect the value flowing through.
func Write[T any](writer io.Writer, lenient bool) chain.X[T, T] {
	return WritefIf[T](writer, "%v", lenient, func(value T) bool { return true })
}

// Writef records the messages to the given writer, applying
// the given format to convert it into a []byte. This filter
// does not affect the value flowing through.
func Writef[T any](writer io.Writer, format string, lenient bool) chain.X[T, T] {
	return WritefIf[T](writer, format, lenient, func(value T) bool { return true })
}

// WriteIf records the messages to the given writer if the given
// condition is true. This filter does not affect the value flowing
// through.
func WriteIf[T any](writer io.Writer, lenient bool, condition func(value T) bool) chain.X[T, T] {
	return WritefIf(writer, "%v", lenient, condition)
}

// WritefIf records the messages to the given writer if the given
// condition is true, applying the given format to convert the value
// to a []byte. This filter does not affect the value flowing through.
func WritefIf[T any](writer io.Writer, format string, lenient bool, condition func(value T) bool) chain.X[T, T] {
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

// WriteUnless records the messages to the given writer unless the given
// condition is true. This filter does not affect the value flowing through.
func WriteUnless[T any](writer io.Writer, lenient bool, condition func(value T) bool) chain.X[T, T] {
	return WritefUnless(writer, "%v", lenient, condition)
}

// WritefUnless records the messages to the given writer unless the given
// condition is true, applying the given format to convert it to a []byte.
// This filter does not affect the value flowing through.
func WritefUnless[T any](writer io.Writer, format string, lenient bool, condition func(value T) bool) chain.X[T, T] {
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
