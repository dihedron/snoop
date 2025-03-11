package fibonacci

import (
	"iter"
)

// Series uses the new Go 1.23 style generator to generate
// a fibonacci series.
func Series(limit int64) iter.Seq[int64] {
	if limit <= 0 {
		limit = 9223372036854775807
	}
	return func(yield func(int64) bool) {
		a, b := int64(0), int64(1)
		for {
			a, b = b, a+b
			if a >= limit {
				break
			}
			if !yield(a) {
				return
			}
		}
	}
}
