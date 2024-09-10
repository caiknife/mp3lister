package types

import (
	"maps"
	"sort"

	"github.com/duke-git/lancet/v2/maputil"
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"

	"github.com/caiknife/mp3lister/lib/fjson"
)

var (
	_ IMap[int, string] = ValueMessage[int](nil)
)

type ValueMessage[K constraints.Ordered] map[K]string

func (msg ValueMessage[K]) HasKey(k K) bool {
	return maputil.HasKey(msg, k)
}

func (msg ValueMessage[K]) Get(key K) (string, bool) {
	v, ok := msg[key]
	return v, ok
}

func (msg ValueMessage[K]) Set(k K, v string) {
	msg[k] = v
}

func (msg ValueMessage[K]) ForEach(f func(K, string)) {
	maputil.ForEach(msg, f)
}

func (msg ValueMessage[K]) ParallelForEach(f func(K, string)) {
	parallelForEach(msg, f)
}

func (msg ValueMessage[K]) Keys() []K {
	return maputil.Keys(msg)
}

func (msg ValueMessage[K]) Values() []string {
	return maputil.Values(msg)
}

func (msg ValueMessage[K]) Len() int {
	return len(msg)
}

func (msg ValueMessage[K]) IsEmpty() bool {
	return msg.Len() == 0
}

func (msg ValueMessage[K]) Clone() ValueMessage[K] {
	return maps.Clone(msg)
}

func (msg ValueMessage[K]) ToSortedSlice() Slice[*Entry[K]] {
	entries := lo.Entries(msg)
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Key < entries[j].Key
	})

	codes := lo.Map[lo.Entry[K, string], *Entry[K]](entries, func(item lo.Entry[K, string], index int) *Entry[K] {
		return &Entry[K]{
			Value: item.Key,
			Msg:   item.Value,
		}
	})

	return codes
}

type Entry[V comparable] struct {
	Value V      `json:"value"`
	Msg   string `json:"msg"`
}

func (e *Entry[V]) String() string {
	toString, _ := fjson.MarshalToString(e)
	return toString
}
