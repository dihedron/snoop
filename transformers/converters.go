package transformers

import (
	"fmt"

	"github.com/dihedron/snoop/format"
	"github.com/dihedron/snoop/transform"
)

// ToString transforms any input type into a string representation.
func ToString[T any]() transform.X[T, string] {
	return func(value T) (string, error) {
		return fmt.Sprintf("%v", value), nil
	}
}

// ToJSON transforms any input value into its JSON representation.
func ToJSON[T any]() transform.X[T, string] {
	return func(value T) (string, error) {
		return format.ToJSON(value), nil
	}
}

// ToPrettyJSON transforms any input value into its JSON representation.
func ToPrettyJSON[T any]() transform.X[T, string] {
	return func(value T) (string, error) {
		return format.ToPrettyJSON(value), nil
	}
}

// ToYAML transforms any input value into its YAML representation.
func ToYAML[T any]() transform.X[T, string] {
	return func(value T) (string, error) {
		return format.ToYAML(value), nil
	}
}
