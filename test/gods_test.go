package test

import (
	"testing"

	"github.com/emirpasic/gods/v2/lists/arraylist"
	"github.com/emirpasic/gods/v2/maps/hashmap"
	"github.com/emirpasic/gods/v2/maps/linkedhashmap"
	"github.com/emirpasic/gods/v2/sets/hashset"
)

func TestGODS_ArrayList(t *testing.T) {
	l := arraylist.New(5, 4, 3, 2, 1)
	t.Log(l.Values())
	l.Sort(func(x, y int) int {
		return x - y
	})
	t.Log(l.Values())

	iterator := l.Iterator()
	for iterator.End(); iterator.Prev(); {
		t.Log(iterator.Value())
	}
	l.Each(func(_ int, value int) {
		t.Log(value)
	})
}

func TestGODS_HashSet(t *testing.T) {
	s := hashset.New(1, 2, 3, 4, 5)
	s.Add(1, 2, 3, 4, 5)
	t.Log(s.Values(), s.Empty(), s.Size())

	s2 := hashset.New(4, 5, 6, 7, 8)
	s = s.Union(s2)
	t.Log(s)
}

func TestGODS_HashMap(t *testing.T) {
	m := hashmap.New[string, int]()
	m.Put("name", 1)
	m.Put("age", 2)
	m.Put("email", 3)
	t.Log(m.Values(), m.Keys())
}

func TestGODS_LinkedHashMap(t *testing.T) {
	m := linkedhashmap.New[string, int]()
	m.Put("name", 1)
	m.Put("age", 2)
	m.Put("email", 3)
	t.Log(m.Values(), m.Keys())
}
