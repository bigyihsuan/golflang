package util

import (
	"fmt"
	"iter"
)

func SliceMap[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func MapMapKeys[K, R comparable, V any](m map[K]V, fn func(K) R) map[R]V {
	result := make(map[R]V, len(m))
	for i, t := range m {
		result[fn(i)] = t
	}
	return result
}

func MapMapValues[K comparable, V, R any](m map[K]V, fn func(V) R) map[K]R {
	result := make(map[K]R, len(m))
	for i, t := range m {
		result[i] = fn(t)
	}
	return result
}

func MapMapKV[K, L comparable, V, W any](m map[K]V, fn func(K, V) (L, W)) map[L]W {
	result := make(map[L]W, len(m))
	for k, v := range m {
		l, w := fn(k, v)
		result[l] = w
	}
	return result
}

func Reversed[T any](s []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(s[i]) {
				return
			}
		}
	}
}

// Compacts together adjacent elements, leaving only the unique ones. Also returns the count of each unique element.
func SliceCompactCount[S ~[]E, E comparable](s S) (S, []int) {
	if len(s) < 2 {
		return s, []int{1}
	}

	elements := []E{s[0]}
	counts := []int{1}
	// start on 1-th element, since there are at least 2 elements
	for k := 1; k < len(s); k++ {
		// if this and previous are the same
		if s[k] == s[k-1] {
			counts[len(counts)-1]++
		} else {
			// this is different from previous
			// add it to elements and couunts
			elements = append(elements, s[k])
			counts = append(counts, 1)
		}
	}
	return elements, counts
}

func SliceCompactCountFunc[S ~[]E, E any](s S, eq func(a, b E) bool) (S, []int) {
	if len(s) < 2 {
		return s, []int{1}
	}

	elements := []E{s[0]}
	counts := []int{1}
	// start on 1-th element, since there are at least 2 elements
	for k := 1; k < len(s); k++ {
		// if this and previous are the same
		if eq(s[k], s[k-1]) {
			counts[len(counts)-1]++
		} else {
			// this is different from previous
			// add it to elements and couunts
			elements = append(elements, s[k])
			counts = append(counts, 1)
		}
	}
	return elements, counts
}

func SliceMap2[A, B, C any](as []A, bs []B, fn func(a A, b B) C) []C {
	if len(as) != len(bs) {
		panic(fmt.Errorf("SliceMap2 lengths do not match: %d, %d", len(as), len(bs)))
	}

	result := make([]C, len(as))
	for i := range len(as) {
		result[i] = fn(as[i], bs[i])
	}
	return result
}

func SliceZip[S ~[]E, E any](a, b S) []E {
	// if len(a) != len(b) {
	// 	panic(fmt.Errorf("SliceZip lengths do not match: %d, %d", len(a), len(b)))
	// }

	minLen := min(len(a), len(b))
	out := []E{}
	for i := 0; i < minLen; i++ {
		out = append(out, a[i], b[i])
	}
	if len(a) < len(b) {
		out = append(out, b[minLen:]...)
	} else if len(a) > len(b) {
		out = append(out, a[minLen:]...)
	}
	return out
}
