package test

import (
	"sync"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/pkg/errors"

	_ "github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/mp3lister/orm/music/model"
)

func TestMovies_Create(t *testing.T) {
	entries := types.Slice[*model.Movie]{}
	for range 500 {
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

func TestMovies_DeleteAll(t *testing.T) {
	info, err := movie.Where(movie.ID).Delete()
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

func TestMovies_GoRoutine(t *testing.T) {
	wg := &sync.WaitGroup{}
	for range 10 {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err := updateMovie(1)
			if err != nil {
				t.Log(err)
				return
			}
			t.Log(err)
		}(wg)
	}

	wg.Wait()
}

func TestMovieUpdate(t *testing.T) {
	err := updateMovie(1)
	if err != nil {
		t.Error(err)
		return
	}
}

func updateMovie(id uint64) error {
	find, err := movie.Where(
		movie.ID.Eq(id),
	).First()
	if err != nil {
		return err
	}

	m := gofakeit.Movie()
	find.Name = m.Name
	find.Genre = m.Genre

	simple, err := movie.Where(
		movie.ID.Eq(id),
		movie.Version.Eq(find.Version),
	).UpdateSimple(
		movie.Name.Value(find.Name),
		movie.Genre.Value(find.Genre),
		movie.Version.Add(1),
	)
	if err != nil {
		return err
	}

	if simple.RowsAffected > 0 {
		return nil
	}
	return errors.New("update failed")
}
