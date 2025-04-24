package util

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
