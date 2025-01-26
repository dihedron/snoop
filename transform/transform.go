package transform

import (
	"errors"
)

// X strands for transformation; it is the generic type for a
// function in the transformer chain, which accepts a type and
// returns a potentially different one.
type X[S any, T any] func(S) (T, error)

// F is a special type of transformation, aka filter, which
// does not mutate the type of the value but it may change the
// value.
type F[T any] = X[T, T]

var (
	//lint:ignore ST1012 Drop allows to notify upstream that we want to drop this value.
	Drop = errors.New("transformation wants to drop this value and continue")
	//lint:ignore ST1012 Quit allows to notify upstream that we want to abort processing.
	Quit = errors.New("transformation wants to abort processing and exit")
)

// Apply chains the two transformations, executing one after
// the other and bailing out as soon as an error is encountered.
func Apply[S any, T any, U any](first X[S, T], second X[T, U]) X[S, U] {
	return func(s S) (U, error) {
		// slog.Debug("calling first...")
		t, err := first(s)
		if err != nil {
			var u U
			// slog.Error("error executing first transformer", "error", err)
			return u, err
		}
		// slog.Debug("calling second...")
		return second(t)
	}
}

// Then is an alias for Apply, to allow a more fluent API.
func Then[S any, T any, U any](first X[S, T], second X[T, U]) X[S, U] {
	return Apply(first, second)
}
