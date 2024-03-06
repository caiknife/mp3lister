package test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	_ "github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/mp3lister/orm/music"
	"github.com/caiknife/mp3lister/orm/music/model"
)

var (
	movie = music.Movie
)

func TestMovies_Create(t *testing.T) {
	entries := types.Slice[*model.Movie]{}
	for range 100 {
		b := gofakeit.Movie()
		e := &model.Movie{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      b.Name,
			Genre:     b.Genre,
		}
		entries = append(entries, e)
	}

	err := movie.CreateInBatches(entries, 100)
	if err != nil {
		t.Error(err)
		return
	}
	entries.ForEach(func(m *model.Movie, i int) {
		t.Log(m)
	})
}

func TestMovies_Restore(t *testing.T) {
	simple, err := movie.Unscoped().Where(movie.ID).UpdateSimple(
		movie.DeletedAt.Value(nil),
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(simple)
}

func TestMovies_Update(t *testing.T) {
	b := gofakeit.Movie()
	simple, err := movie.Where(
		movie.ID.Eq(3),
	).UpdateSimple(
		movie.Name.Value(b.Name),
		movie.Genre.Value(b.Genre),
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(simple)
}

func TestMovies_Get(t *testing.T) {
	find, err := movie.Where(
		movie.ID.Eq(3),
	).First()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(find)
}

func TestMovies_GetAll(t *testing.T) {
	find, err := movie.Order(movie.ID.Desc()).Find()
	if err != nil {
		t.Error(err)
		return
	}
	for _, m := range find {
		t.Log(m)
	}
}

func TestMovies_Truncate(t *testing.T) {
	info, err := movie.Unscoped().Where(movie.ID).Delete()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}

func TestMovies_Delete(t *testing.T) {
	info, err := movie.Where(
		movie.ID.Eq(3),
	).Delete()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}
