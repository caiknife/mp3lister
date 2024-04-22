package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var s = Slice[int]{9, 8, 7, 6, 5}

func TestSlice_Len(t *testing.T) {
	assert.Equal(t, s.Len(), 5)
}

func TestSlice_Clone(t *testing.T) {
	v := s.Clone()
	assert.Equal(t, v, s)
}

func TestSlice_Contain(t *testing.T) {
	assert.True(t, s.Contain(5))
	assert.False(t, s.Contain(1))
}

func TestSlice_Count(t *testing.T) {
	assert.Equal(t, s.Count(1), 0)
	assert.Equal(t, s.Count(5), 1)
}

func TestSlice_Find(t *testing.T) {
	find, b := s.Find(func(i int) bool {
		return i == 1
	})
	assert.Equal(t, find, 0)
	assert.Equal(t, b, false)

	find, b = s.Find(func(i int) bool {
		return i == 5
	})
	assert.Equal(t, find, 5)
	assert.Equal(t, b, true)
}

func TestSlice_ForEach(t *testing.T) {
	s.ForEach(func(_ int, i int) {
		s[i] *= 2
	})
	t.Log(s)
}

func TestSlice_IsEmpty(t *testing.T) {
	assert.False(t, s.IsEmpty())
}

func TestSlice_ParallelForEach(t *testing.T) {
	s.ParallelForEach(func(_ int, i int) {
		s[i] *= 2
	})
	t.Log(s)
}

func TestSlice_Random(t *testing.T) {
	assert.True(t, s.Contain(s.Random()))
}

func TestSlice_Shuffle(t *testing.T) {
	s.Shuffle()
	t.Log(s)
}

func TestSlice_Sort(t *testing.T) {
	s.Sort(func(i, j int) bool {
		return s[i] < s[j]
	})
	t.Log(s)
}
