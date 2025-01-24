package experiment

import "fmt"

// Define a generic type for the function in the chain
type Handler[S, T any] func(S) (T, error)

// Chain creates a chain of functions
func Chain[S any, T any, U any](first Handler[S, T], second Handler[T, U]) Handler[S, U] {
	return func(s S) (U, error) {
		t, err := first(s)
		if err != nil {
			var u U
			return u, err
		}
		return second(t)
	}
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
	chainedFunc := Chain(StringToInt, Chain(IntToDouble, DoubleToString))

	// Use the chained function
	result, _ := chainedFunc("10")
	fmt.Println(result) // Output: true
	result, _ = chainedFunc("11")
	fmt.Println(result) // Output: true
}
