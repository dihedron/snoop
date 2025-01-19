package concat

import (
	"iter"
)

// Concat concatenates multiple single-valued sequences.
func Concat[T any](sequences ...iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, sequence := range sequences {
			for value := range sequence {
				if !yield(value) {
					return
				}
			}
		}
	}
}

// Concat2 concatenates multiple double-valued sequences.
func Concat2[K any, V any](sequences ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, sequence := range sequences {
			for k, v := range sequence {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
