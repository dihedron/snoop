package transformers

import (
	"bytes"
	"log/slog"

	"text/template"

	"github.com/dihedron/snoop/transform"
)

func Format[T any](format string) transform.X[T, string] {
	var buffer bytes.Buffer
	template, err := template.New(format).Parse(format)
	if err != nil {
		slog.Error("invalid template", "error", err)
		return nil
	}
	return func(value T) (string, error) {
		err = template.Execute(&buffer, value)
		if err != nil {
			return "", err
		}
		return buffer.String(), nil
	}
}
