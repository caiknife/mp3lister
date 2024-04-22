package types

import (
	"sort"
	"time"

	"github.com/samber/lo"
	"github.com/samber/lo/parallel"
	"golang.org/x/exp/rand"

	"github.com/caiknife/mp3lister/lib/fjson"
)

func init() {
	rand.Seed(uint64(time.Now().UnixNano()))
}

var _ ISlice[int] = Slice[int](nil)

type Slice[T comparable] []T

func (s Slice[T]) UnmarshalBinary(data []byte) error {
	return fjson.Unmarshal(data, &s)
}

func (s Slice[T]) MarshalBinary() (data []byte, err error) {
	return fjson.Marshal(s)
}

func (s Slice[T]) String() string {
	toString, err := fjson.MarshalToString(s)
	if err != nil {
		return ""
	}
	return toString
}

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s Slice[T]) Count(v T) int {
	return lo.Count(s, v)
}

func (s Slice[T]) Contain(v T) bool {
	return lo.Contains(s, v)
}

func (s Slice[T]) Random() T {
	var v T
	if s.IsEmpty() {
		return v
	}
	idx := rand.Intn(s.Len())
	return s[idx]
}

func (s Slice[T]) Shuffle() {
	s = lo.Shuffle(s)
}

func (s Slice[T]) Find(f func(T) bool) (T, bool) {
	return lo.Find(s, f)
}

func (s Slice[T]) ForEach(f func(T, int)) {
	lo.ForEach(s, f)
}

func (s Slice[T]) ParallelForEach(f func(T, int)) {
	parallel.ForEach(s, f)
}

func (s Slice[T]) Sort(less func(i, j int) bool) {
	sort.SliceStable(s, less)
}

func (s Slice[T]) Clone() Slice[T] {
	v := make(Slice[T], len(s))
	copy(v, s)
	return v
}

func (s Slice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
