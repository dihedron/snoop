package transformers

import "github.com/dihedron/snoop/transform/chain"

// AcceptIf lets the value flow if the condition is true. If the
// condition is true, this filter does not affect the value
// flowing through.
func AcceptIf[T any](condition func(value T) bool) chain.X[T, T] {
	return func(value T) (T, error) {
		if condition(value) {
			return value, nil
		}
		var nihil T
		return nihil, chain.Drop
	}
}

// AcceptUnless drops the value flow if the condition is true. If the
// condition is false, this filter does not affect the value flowing
// through.
func AcceptUnless[T any](condition func(value T) bool) chain.X[T, T] {
	return func(value T) (T, error) {
		if !condition(value) {
			return value, nil
		}
		var nihil T
		return nihil, chain.Drop
	}
}

// DropIf drops the value if the condition is true. If the condition is
// false, this filter does not affect the value flowing through.
func DropIf[T any](condition func(value T) bool) chain.X[T, T] {
	return AcceptUnless(condition)
}

// DropUnless drops the value if the condition is false. If the condition
// is true, this filter does not affect the value flowing through.
func DropUnless[T any](condition func(value T) bool) chain.X[T, T] {
	return AcceptIf(condition)
}
