package main

func Map[T any, M any](vs []T, f func(T) M) []M {
	vsm := make([]M, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func Filter[T any](target []T, f func(T) bool) []T {
	var store []T
	for i, v := range target {
		if f(v) {
			store = append(store, target[i])
		}
	}
	return store
}
