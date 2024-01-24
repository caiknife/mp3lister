package types

import (
	"sync"
)

type ISlice[T comparable] interface {
	Len() int
	IsEmpty() bool
	Count(v T) int
	Contain(v T) bool
	Random() T
	Shuffle()
	Find(f func(T) bool) (T, bool)
	ForEach(f func(T, int))
	ParallelForEach(f func(T, int))
}

type IMap[K comparable, V any] interface {
	Len() int
	Keys() []K
	Values() []V
	IsEmpty() bool
	HasKey(K) bool
	Get(key K) (V, bool)
	Set(k K, v V)
	ForEach(func(K, V))
	ParallelForEach(func(K, V))
}

func parallelForEach[K comparable, V any](m map[K]V, f func(K, V)) {
	var wg sync.WaitGroup
	wg.Add(len(m))

	for s, t := range m {
		go func(_s K, _t V, wg *sync.WaitGroup) {
			f(_s, _t)
			wg.Done()
		}(s, t, &wg)
	}

	wg.Wait()
}
