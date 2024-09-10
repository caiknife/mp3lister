package types

import (
	"maps"

	"github.com/duke-git/lancet/v2/maputil"
	"github.com/samber/lo"

	"github.com/caiknife/mp3lister/lib/fjson"
)

var _ IMap[string, int] = Map[int](nil)

type Map[V any] map[string]V

type H = Map[any]

func (m Map[V]) UnmarshalBinary(data []byte) error {
	return fjson.Unmarshal(data, &m)
}

func (m Map[V]) MarshalBinary() (data []byte, err error) {
	return fjson.Marshal(m)
}

func (m Map[V]) String() string {
	toString, err := fjson.MarshalToString(m)
	if err != nil {
		return ""
	}
	return toString
}

func (m Map[V]) Len() int {
	return len(m)
}

func (m Map[V]) Keys() []string {
	return lo.Keys(m)
}

func (m Map[V]) Values() []V {
	return lo.Values(m)
}

func (m Map[V]) IsEmpty() bool {
	return m.Len() == 0
}

func (m Map[V]) HasKey(k string) bool {
	return maputil.HasKey(m, k)
}

func (m Map[V]) Get(key string) (V, bool) {
	v, ok := m[key]
	return v, ok
}

func (m Map[V]) Set(k string, v V) {
	m[k] = v
}

func (m Map[V]) ForEach(f func(string, V)) {
	maputil.ForEach(m, f)
}

func (m Map[V]) ParallelForEach(f func(string, V)) {
	parallelForEach(m, f)
}

func (m Map[V]) Clone() Map[V] {
	return maps.Clone(m)
}
