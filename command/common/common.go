package common

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/go-playground/validator/v10"
)

// Validate validates an object using the tags in the strcut.
func Validate(v any) error {
	slog.Debug("validating object", "type", fmt.Sprintf("%T", v))
	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(v); err != nil {
		slog.Error("error validating object", "error", err)
		return err
	}
	slog.Debug("validation successful")
	return nil
}

// GetWriter returns amn io.Writer (possibly an io.WriteCloser)
// where messages can be recorded,
func GetWriter(path string, truncate *bool) (io.Writer, error) {
	if path == "" {
		slog.Error("invalid output path")
		return nil, errors.New("invalid output path")
	}

	if path == "-" {
		slog.Info("STDOUT writer ready")
		return os.Stdout, nil
	}

	// zipped := false
	// if filepath.Ext(path) == ".gz" {
	// 	slog.Debug("enabling GZIP compression support for output stream")
	// 	zipped = true
	// 	path = strings.TrimSuffix(path, ".gz")
	// }

	slog.Info("writing to file", "path", path)
	flags := 0
	if truncate != nil && *truncate {
		slog.Debug("opening output file in truncate mode", "path", path)
		flags = os.O_TRUNC | os.O_CREATE | os.O_WRONLY
	} else {
		slog.Debug("opening output file in append mode", "path", path)
		flags = os.O_APPEND | os.O_CREATE | os.O_WRONLY
	}
	file, err := os.OpenFile(path, flags, 0600)
	if err != nil {
		slog.Error("error opening recorder output file in append mode", "path", path, "truncate", truncate, "flags", flags, "error", err)
		return nil, errors.New("error opening output file")
	}

	// if zipped {
	// 	// 2. Create a gzip writer.
	// 	writer := gzip.NewWriter(file)
	// 	slog.Debug("GZIP writer is ready")
	// 	defer writer.Close()
	// TODO: need to wrap the underlying file in order to flush it
	// 	return writer, nil
	// }

	slog.Debug("writer is ready")
	return file, nil
}
