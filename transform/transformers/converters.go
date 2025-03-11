package transformers

import (
	"fmt"

	"github.com/dihedron/snoop/transform/chain"
	"github.com/goccy/go-json"

	"gopkg.in/yaml.v3"
)

// StringToByteArray converts a string to a []byte.
func StringToByteArray() chain.X[string, []byte] {
	return func(value string) ([]byte, error) {
		return []byte(value), nil
	}
}

// ByteArrayToString converts a []byte to a string.
func ByteArrayToString() chain.X[[]byte, string] {
	return func(value []byte) (string, error) {
		return string(value), nil
	}
}

// ToString transforms any input type into a string representation.
func ToString[T any]() chain.X[T, string] {
	return ToStringf[T]("%v")
}

// ToStringf transforms any input type into a string representation
// according to the specified format.
func ToStringf[T any](format string) chain.X[T, string] {
	return func(value T) (string, error) {
		return fmt.Sprintf(format, value), nil
	}
}

// ToJSON transforms any input value into its JSON representation.
func ToJSON[T any]() chain.X[T, []byte] {
	return func(value T) ([]byte, error) {
		return json.Marshal(value)
		//		return format.ToJSON(value), nil
	}
}

// ToPrettyJSON transforms any input value into its JSON representation.
func ToPrettyJSON[T any]() chain.X[T, []byte] {
	return func(value T) ([]byte, error) {
		return json.MarshalIndent(value, "", "  ")
	}
}

// ToYAML transforms any input value into its YAML representation.
func ToYAML[T any]() chain.X[T, []byte] {
	return func(value T) ([]byte, error) {
		return yaml.Marshal(value)
	}
}
