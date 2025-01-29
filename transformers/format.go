package transformers

import (
	"bytes"
	"log/slog"

	"text/template"

	"github.com/Masterminds/sprig/v3"
	f "github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/templating"
	"github.com/dihedron/snoop/transform"
)

func Format[T any](format string) transform.X[T, string] {
	var buffer bytes.Buffer

	// populate the functions map
	functions := template.FuncMap{}
	for k, v := range templating.FuncMap() {
		functions[k] = v
	}
	for k, v := range sprig.FuncMap() {
		functions[k] = v
	}

	template, err := template.New(format).Funcs(functions).Parse(format)
	if err != nil {
		slog.Error("invalid template", "error", err)
		return nil
	}
	return func(value T) (string, error) {
		err = template.Execute(&buffer, value)
		if err != nil {
			slog.Error("error applying template", "object", value, "type", f.TypeAsString(value), "error", err)
			return "", err
		}
		return buffer.String(), nil
	}
}
