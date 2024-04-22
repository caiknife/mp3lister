package types

import (
	"github.com/duke-git/lancet/v2/maputil"
	"github.com/samber/lo"

	"github.com/caiknife/mp3lister/lib/fjson"
)

var _ IMap[string, int] = Hash[string, int](nil)

type Hash[K comparable, V any] map[K]V

func (h Hash[K, V]) UnmarshalBinary(data []byte) error {
	return fjson.Unmarshal(data, &h)
}

func (h Hash[K, V]) MarshalBinary() (data []byte, err error) {
	return fjson.Marshal(h)
}

func (h Hash[K, V]) String() string {
	toString, err := fjson.MarshalToString(h)
	if err != nil {
		return ""
	}
	return toString
}

func (h Hash[K, V]) Len() int {
	return len(h)
}

func (h Hash[K, V]) Keys() []K {
	return lo.Keys(h)
}

func (h Hash[K, V]) Values() []V {
	return lo.Values(h)
}

func (h Hash[K, V]) IsEmpty() bool {
	return h.Len() == 0
}

func (h Hash[K, V]) HasKey(k K) bool {
	return maputil.HasKey(h, k)
}

func (h Hash[K, V]) Get(key K) (V, bool) {
	v, ok := h[key]
	return v, ok
}

func (h Hash[K, V]) Set(k K, v V) {
	h[k] = v
}

func (h Hash[K, V]) ForEach(f func(K, V)) {
	maputil.ForEach(h, f)
}

func (h Hash[K, V]) ParallelForEach(f func(K, V)) {
	parallelForEach(h, f)
}
