package chain

import (
	"errors"
)

// X stands for and exomorphic transformation; it is the generic
// type for a function in the transformer chain, which accepts a
// type and returns a potentially different one.
type X[S any, T any] func(S) (T, error)

// F is a special type of transformation, aka endomorphism or filter,
// which does not mutate the type (domain) of the value but it may
// change the value.
type F[T any] = X[T, T]

var (
	//lint:ignore ST1012 Drop allows to notify upstream that we want to drop this value.
	Drop = errors.New("transformation wants to drop this value and continue")
	//lint:ignore ST1012 Quit allows to notify upstream that we want to abort processing.
	Quit = errors.New("transformation wants to abort processing and exit")
)

// Of returns a function that applies a single transformation.
func Of[A any, B any](xform X[A, B]) X[A, B] {
	return func(a A) (B, error) {
		return xform(a)
	}
}

// Then is an alias for Of, to allow a more fluent API.
func Then[A any, B any](xform X[A, B]) X[A, B] {
	return Of(xform)
}

// Of2 returns a composite transformation that chains the two given
// transformations, executing one after the other and bailing out as
// soon as an error is encountered.
func Of2[A any, B any, C any](first X[A, B], second X[B, C]) X[A, C] {
	return func(s A) (C, error) {
		t, err := first(s)
		if err != nil {
			var c C
			return c, err
		}
		return second(t)
	}
}

// Then2 is an alias for Of2, to allow a more fluent API.
func Then2[A any, B any, C any](first X[A, B], second X[B, C]) X[A, C] {
	return Of2(first, second)
}

// // ThenElse concatenates the first transform with one of the two
// // transforms ("either" and "or") based on a condition on the
// // output value of the first transform. The types of the "either"
// // and "or" transforms must be the same. This allows to apply
// // conditional processing based on a test.
// func ThenElse[A any, B any, C any](first X[A, B], condition func(value B) bool, either X[B, C], or X[B, C]) X[A, C] {
// 	return func(s A) (C, error) {
// 		t, err := first(s)
// 		if err != nil {
// 			var c C
// 			return c, err
// 		}
// 		if condition(t) {
// 			return either(t)
// 		}
// 		return or(t)
// 	}
// }

// Of3 returns a transformation that chains the three given
// transformations, executing one after the other and bailing
// out as soon as an error is encountered.
func Of3[A any, B any, C any, D any](first X[A, B], second X[B, C], third X[C, D]) X[A, D] {
	return func(s A) (D, error) {
		t, err := first(s)
		if err != nil {
			var d D
			return d, err
		}
		u, err := second(t)
		if err != nil {
			var d D
			return d, err
		}
		return third(u)
	}
}

// Then3 is an alias for Of3, to allow a more fluent API.
func Then3[A any, B any, C any, D any](first X[A, B], second X[B, C], third X[C, D]) X[A, D] {
	return Of3(first, second, third)
}

// Of4 returns a transformation that chains the four given
// transformations, executing one after the other and bailing
// out as soon as an error is encountered.
func Of4[A any, B any, C any, D any, E any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E]) X[A, E] {
	return func(s A) (E, error) {
		t, err := first(s)
		if err != nil {
			var e E
			return e, err
		}
		u, err := second(t)
		if err != nil {
			var e E
			return e, err
		}
		v, err := third(u)
		if err != nil {
			var e E
			return e, err
		}
		return fourth(v)
	}
}

// Then4 is an alias for Of4, to allow a more fluent API.
func Then4[A any, B any, C any, D any, E any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E]) X[A, E] {
	return Of4(first, second, third, fourth)
}

// Of5 returns a transformation that chains the five given
// transformations, executing one after the other and bailing
// out as soon as an error is encountered.
func Of5[A any, B any, C any, D any, E any, F any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F]) X[A, F] {
	return func(s A) (F, error) {
		t, err := first(s)
		if err != nil {
			var f F
			return f, err
		}
		u, err := second(t)
		if err != nil {
			var f F
			return f, err
		}
		v, err := third(u)
		if err != nil {
			var f F
			return f, err
		}
		w, err := fourth(v)
		if err != nil {
			var f F
			return f, err
		}
		return fifth(w)
	}
}

// Then5 is an alias for Of5, to allow a more fluent API.
func Then5[A any, B any, C any, D any, E any, F any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F]) X[A, F] {
	return Of5(first, second, third, fourth, fifth)
}

// Of6 returns a transformation that chains the six given
// transformations, executing one after the other and bailing
// out as soon as an error is encountered.
func Of6[A any, B any, C any, D any, E any, F any, G any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G]) X[A, G] {
	return func(s A) (G, error) {
		t, err := first(s)
		if err != nil {
			var g G
			return g, err
		}
		u, err := second(t)
		if err != nil {
			var g G
			return g, err
		}
		v, err := third(u)
		if err != nil {
			var g G
			return g, err
		}
		w, err := fourth(v)
		if err != nil {
			var g G
			return g, err
		}
		x, err := fifth(w)
		if err != nil {
			var g G
			return g, err
		}
		return sixth(x)
	}
}

// Then6 is an alias for Of6, to allow a more fluent API.
func Then6[A any, B any, C any, D any, E any, F any, G any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G]) X[A, G] {
	return Of6(first, second, third, fourth, fifth, sixth)
}

// Of7 returns a transformation that chains the seven given
// transformations, executing one after the other and bailing
// out as soon as an error is encountered.
func Of7[A any, B any, C any, D any, E any, F any, G any, H any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G], seventh X[G, H]) X[A, H] {
	return func(s A) (H, error) {
		t, err := first(s)
		if err != nil {
			var h H
			return h, err
		}
		u, err := second(t)
		if err != nil {
			var h H
			return h, err
		}
		v, err := third(u)
		if err != nil {
			var h H
			return h, err
		}
		w, err := fourth(v)
		if err != nil {
			var h H
			return h, err
		}
		x, err := fifth(w)
		if err != nil {
			var h H
			return h, err
		}
		y, err := sixth(x)
		if err != nil {
			var h H
			return h, err
		}
		return seventh(y)
	}
}

// Then7 is an alias for Of7, to allow a more fluent API.
func Then7[A any, B any, C any, D any, E any, F any, G any, H any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G], seventh X[G, H]) X[A, H] {
	return Of7(first, second, third, fourth, fifth, sixth, seventh)
}

// Of8 returns a transformation that chains the eight given
// transformations, executing one after the other and bailing
// out as soon as an Of9 is encountered.
func Of8[A any, B any, C any, D any, E any, F any, G any, H any, I any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G], seventh X[G, H], eighth X[H, I]) X[A, I] {
	return func(s A) (I, error) {
		t, err := first(s)
		if err != nil {
			var i I
			return i, err
		}
		u, err := second(t)
		if err != nil {
			var i I
			return i, err
		}
		v, err := third(u)
		if err != nil {
			var i I
			return i, err
		}
		w, err := fourth(v)
		if err != nil {
			var i I
			return i, err
		}
		x, err := fifth(w)
		if err != nil {
			var i I
			return i, err
		}
		y, err := sixth(x)
		if err != nil {
			var i I
			return i, err
		}
		z, err := seventh(y)
		if err != nil {
			var i I
			return i, err
		}
		return eighth(z)
	}
}

// Then8 is an alias for Of8, to allow a more fluent API.
func Then8[A any, B any, C any, D any, E any, F any, G any, H any, I any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G], seventh X[G, H], eighth X[H, I]) X[A, I] {
	return Of8(first, second, third, fourth, fifth, sixth, seventh, eighth)
}

// Of9 returns a transformation that chains the nine given
// transformations, executing one after the other and bailing
// out as soon as an error is encountered.
func Of9[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G], seventh X[G, H], eighth X[H, I], ninth X[I, J]) X[A, J] {
	return func(s A) (J, error) {
		t, err := first(s)
		if err != nil {
			var j J
			return j, err
		}
		u, err := second(t)
		if err != nil {
			var j J
			return j, err
		}
		v, err := third(u)
		if err != nil {
			var j J
			return j, err
		}
		w, err := fourth(v)
		if err != nil {
			var j J
			return j, err
		}
		x, err := fifth(w)
		if err != nil {
			var j J
			return j, err
		}
		y, err := sixth(x)
		if err != nil {
			var j J
			return j, err
		}
		z, err := seventh(y)
		if err != nil {
			var j J
			return j, err
		}
		a, err := eighth(z)
		if err != nil {
			var j J
			return j, err
		}
		return ninth(a)
	}
}

// Then9 is an alias for Of9, to allow a more fluent API.
func Then9[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G], seventh X[G, H], eighth X[H, I], ninth X[I, J]) X[A, J] {
	return Of9(first, second, third, fourth, fifth, sixth, seventh, eighth, ninth)
}

// Of10 returns a transformation that chains the ten given
// transformations, executing one after the other and bailing
// out as soon as an error is encountered.
func Of10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G], seventh X[G, H], eighth X[H, I], ninth X[I, J], tenth X[J, K]) X[A, K] {
	return func(s A) (K, error) {
		t, err := first(s)
		if err != nil {
			var k K
			return k, err
		}
		u, err := second(t)
		if err != nil {
			var k K
			return k, err
		}
		v, err := third(u)
		if err != nil {
			var k K
			return k, err
		}
		w, err := fourth(v)
		if err != nil {
			var k K
			return k, err
		}
		x, err := fifth(w)
		if err != nil {
			var k K
			return k, err
		}
		y, err := sixth(x)
		if err != nil {
			var k K
			return k, err
		}
		z, err := seventh(y)
		if err != nil {
			var k K
			return k, err
		}
		a, err := eighth(z)
		if err != nil {
			var k K
			return k, err
		}
		b, err := ninth(a)
		if err != nil {
			var k K
			return k, err
		}
		return tenth(b)
	}
}

// Then10 is an alias for Of10, to allow a more fluent API.
func Then10[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G], seventh X[G, H], eighth X[H, I], ninth X[I, J], tenth X[J, K]) X[A, K] {
	return Of10(first, second, third, fourth, fifth, sixth, seventh, eighth, ninth, tenth)
}

// Of11 returns a transformation that chains the eleven given
// transformations, executing one after the other and bailing
// out as soon as an error is encountered.
func Of11[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G], seventh X[G, H], eighth X[H, I], ninth X[I, J], tenth X[J, K], eleventh X[K, L]) X[A, L] {
	return func(s A) (L, error) {
		t, err := first(s)
		if err != nil {
			var l L
			return l, err
		}
		u, err := second(t)
		if err != nil {
			var l L
			return l, err
		}
		v, err := third(u)
		if err != nil {
			var l L
			return l, err
		}
		w, err := fourth(v)
		if err != nil {
			var l L
			return l, err
		}
		x, err := fifth(w)
		if err != nil {
			var l L
			return l, err
		}
		y, err := sixth(x)
		if err != nil {
			var l L
			return l, err
		}
		z, err := seventh(y)
		if err != nil {
			var l L
			return l, err
		}
		a, err := eighth(z)
		if err != nil {
			var l L
			return l, err
		}
		b, err := ninth(a)
		if err != nil {
			var l L
			return l, err
		}
		c, err := tenth(b)
		if err != nil {
			var l L
			return l, err
		}
		return eleventh(c)
	}
}

// Then11 is an alias for Of11, to allow a more fluent API.
func Then11[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G], seventh X[G, H], eighth X[H, I], ninth X[I, J], tenth X[J, K], eleventh X[K, L]) X[A, L] {
	return Of11(first, second, third, fourth, fifth, sixth, seventh, eighth, ninth, tenth, eleventh)
}

// Of12 returns a transformation that chains the twelve given
// transformations, executing one after the other and bailing
// out as soon as an error is encountered.
func Of12[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G], seventh X[G, H], eighth X[H, I], ninth X[I, J], tenth X[J, K], eleventh X[K, L], twelfth X[L, M]) X[A, M] {
	return func(s A) (M, error) {
		t, err := first(s)
		if err != nil {
			var m M
			return m, err
		}
		u, err := second(t)
		if err != nil {
			var m M
			return m, err
		}
		v, err := third(u)
		if err != nil {
			var m M
			return m, err
		}
		w, err := fourth(v)
		if err != nil {
			var m M
			return m, err
		}
		x, err := fifth(w)
		if err != nil {
			var m M
			return m, err
		}
		y, err := sixth(x)
		if err != nil {
			var m M
			return m, err
		}
		z, err := seventh(y)
		if err != nil {
			var m M
			return m, err
		}
		a, err := eighth(z)
		if err != nil {
			var m M
			return m, err
		}
		b, err := ninth(a)
		if err != nil {
			var m M
			return m, err
		}
		c, err := tenth(b)
		if err != nil {
			var m M
			return m, err
		}
		d, err := eleventh(c)
		if err != nil {
			var m M
			return m, err
		}
		return twelfth(d)
	}
}

// Then12 is an alias for Of12, to allow a more fluent API.
func Then12[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any, K any, L any, M any](first X[A, B], second X[B, C], third X[C, D], fourth X[D, E], fifth X[E, F], sixth X[F, G], seventh X[G, H], eighth X[H, I], ninth X[I, J], tenth X[J, K], eleventh X[K, L], twelfth X[L, M]) X[A, M] {
	return Of12(first, second, third, fourth, fifth, sixth, seventh, eighth, ninth, tenth, eleventh, twelfth)
}
