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
