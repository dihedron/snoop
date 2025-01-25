package experiment

import "fmt"

// Define a generic type for the function in the chain
type Transformer[S any, T any] func(S) (T, error)

type Filter[T any] = Transformer[T, T]

// FollowedBy does not work, but it would be the best solution.
func (before Transformer[S, T]) FollowedBy(after Transformer[T, any]) Transformer[S, any] {
	return func(s S) (any, error) {
		t, err := before(s)
		if err != nil {
			return nil, err
		}
		return after(t)
	}
}

// Apply creates a chain of functions
func Apply[S any, T any, U any](first Transformer[S, T], second Transformer[T, U]) Transformer[S, U] {
	return func(s S) (U, error) {
		t, err := first(s)
		if err != nil {
			var u U
			return u, err
		}
		return second(t)
	}
}

func Then[S any, T any, U any](first Transformer[S, T], second Transformer[T, U]) Transformer[S, U] {
	return Apply(first, second)
}

func ChainTogether() {
	// Define functions with different types
	StringToInt := func(s string) (int, error) {
		var i int
		_, err := fmt.Sscanf(s, "%d", &i)
		return i, err
	}

	IntToDouble := func(i int) (float64, error) {
		return float64(i) * 2, nil
	}

	DoubleToString := func(f float64) (string, error) {
		return fmt.Sprintf("%f", f), nil
	}

	// Chain the functions together
	chainedFunc := Apply(
		StringToInt,
		Then(
			IntToDouble,
			DoubleToString,
		),
	)

	// Use the chained function
	result, _ := chainedFunc("10")
	fmt.Println(result) // Output: true
	result, _ = chainedFunc("11")
	fmt.Println(result) // Output: true
}
