package experiment

import (
	"testing"
)

func TestChainTogether(t *testing.T) {
	ChainTogether()
}

// func TestWithThen(t *testing.T) {
// 	// Define functions with different types
// 	StringToInt := func(s string) (int, error) {
// 		var i int
// 		_, err := fmt.Sscanf(s, "%d", &i)
// 		return i, err
// 	}

// 	IntToDouble := func(i int) (float64, error) {
// 		return float64(i) * 2, nil
// 	}

// 	DoubleToString := func(f float64) (string, error) {
// 		return fmt.Sprintf("%f", f), nil
// 	}

// 	// Chain the functions together
// 	chain := Transformer[string, int](StringToInt).
// 		FollowedBy(Transformer[int, float64](IntToDouble)).
// 		FollowedBy(Transformer[float64, string](DoubleToString))

// 	// Use the chained function
// 	result, _ := chain("10")
// 	fmt.Println(result) // Output: true
// 	result, _ = chain("11")
// 	fmt.Println(result) // Output: true
// }
