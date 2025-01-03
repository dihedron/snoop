package message

import "fmt"

var (
	// ErrInvalidInput is the error returned when the input is nil or invalid.
	ErrInvalidInput = fmt.Errorf("invalid input to function")
	// ErrInvalidPayload is the error returned when the Oslo message body could not be parsed.
	ErrInvalidPayload = fmt.Errorf("invalid message payload or format")
)
