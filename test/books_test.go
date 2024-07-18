package test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/mp3lister/orm/music/model"
)

func TestBooks_Create(t *testing.T) {
	entries := types.Slice[*model.Book]{}
	for range 500 {
		b := gofakeit.Book()
		e := &model.Book{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title:     b.Title,
			Author:    b.Author,
			Genre:     b.Genre,
		}
		entries = append(entries, e)
	}

	err := book.CreateInBatches(entries, 100)
	if err != nil {
		t.Error(err)
		return
	}
	entries.ForEach(func(m *model.Book, i int) {
		t.Log(m)
	})
}

func TestBooks_Restore(t *testing.T) {
	simple, err := book.Unscoped().Where(book.ID).UpdateSimple(
		book.DeletedAt.Value(nil),
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(simple)
}

func TestBooks_Update(t *testing.T) {
	b := gofakeit.Book()
	simple, err := book.Where(
		book.ID.Eq(3),
	).UpdateSimple(
		book.Title.Value(b.Title),
		book.Author.Value(b.Author),
		book.Genre.Value(b.Genre),
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(simple)
}

func TestBooks_Count(t *testing.T) {
	count, err := book.Count()
	if err != nil {
		t.Error()
		return
	}
	t.Log(count)
}

type query struct {
	Total int
	Count int
}

func TestBooks_Sum(t *testing.T) {
	a := query{}
	err := book.Select(book.ID.Sum().As("total"), book.ID.Count().As("count")).Scan(&a)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(a)
}

func TestBooks_Get(t *testing.T) {
	find, err := book.Where(
		book.ID.Eq(3),
	).First()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(find)
}

func TestBooks_GetAll(t *testing.T) {
	find, err := book.Order(book.ID.Desc()).Find()
	if err != nil {
		t.Error(err)
		return
	}
	for _, m := range find {
		t.Log(m)
	}
}

func TestBooks_Truncate(t *testing.T) {
	info, err := book.Unscoped().Where(book.ID).Delete()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}

func TestBooks_Delete(t *testing.T) {
	info, err := book.Where(
		book.ID.Eq(3),
	).Delete()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}

func TestBooks_DeleteAll(t *testing.T) {
	info, err := book.Where(book.ID).Delete()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}
