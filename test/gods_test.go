package test

import (
	"testing"

	"github.com/emirpasic/gods/v2/lists/arraylist"
	"github.com/emirpasic/gods/v2/maps/hashmap"
	"github.com/emirpasic/gods/v2/maps/linkedhashmap"
	"github.com/emirpasic/gods/v2/queues/arrayqueue"
	"github.com/emirpasic/gods/v2/sets/hashset"
	"github.com/emirpasic/gods/v2/sets/linkedhashset"
	"github.com/emirpasic/gods/v2/sets/treeset"
	"github.com/emirpasic/gods/v2/stacks/arraystack"
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
	m.Each(func(key string, value int) {
		t.Log(key, value)
		m.Put(key, value*2)
	})
	t.Log(m)
}

func TestGODS_LinkedHashSet(t *testing.T) {
	s := linkedhashset.New(1, 2, 3, 4, 5, 1, 2, 3, 4, 5)
	t.Log(s.Values())
	iterator := s.Iterator()
	for iterator.Next() {
		t.Log(iterator.Value())
	}
}

func TestGODS_TreeSet(t *testing.T) {
	s := treeset.New(1, 2, 3, 4, 5, 1, 2, 3, 4, 5)
	t.Log(s.Values())
	iterator := s.Iterator()
	for iterator.Next() {
		t.Log(iterator.Value())
	}
}

func TestGODS_Stack(t *testing.T) {
	s := arraystack.New[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	t.Log(s)

	for {
		value, ok := s.Pop()
		if !ok {
			break
		}
		t.Log(value)
	}
}

func TestGODS_Queue(t *testing.T) {
	q := arrayqueue.New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Enqueue(5)
	t.Log(q)

	for {
		value, ok := q.Dequeue()
		if !ok {
			break
		}
		t.Log(value)
	}
}
