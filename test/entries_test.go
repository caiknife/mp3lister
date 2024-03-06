package test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	_ "github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/orm/music"
	"github.com/caiknife/mp3lister/orm/music/model"
)

var (
	entry = music.Entry
)

func TestEntries_Create(t *testing.T) {
	e := &model.Entry{
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Title:      gofakeit.BookTitle(),
		Artist:     gofakeit.Name(),
		Album:      gofakeit.MonthString(),
		Bpm:        int32(gofakeit.IntRange(40, 300)),
		OriginFile: "",
	}
	err := entry.Create(e)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(e)
}

func TestEntries_Restore(t *testing.T) {
	simple, err := entry.Unscoped().Where(entry.ID).UpdateSimple(
		entry.DeletedAt.Value(nil),
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(simple)
}

func TestEntries_Update(t *testing.T) {
	simple, err := entry.Where(
		entry.ID.Eq(3),
	).UpdateSimple(
		entry.Title.Value("test"),
		entry.Artist.Value("test"),
		entry.Album.Value("test"),
		entry.Bpm.Value(100),
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(simple)
}

func TestEntries_Get(t *testing.T) {
	find, err := entry.Where(
		entry.ID.Eq(3),
	).First()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(find)
}

func TestEntries_GetAll(t *testing.T) {
	find, err := entry.Order(entry.ID.Desc()).Find()
	if err != nil {
		t.Error(err)
		return
	}
	for _, m := range find {
		t.Log(m)
	}
}

func TestEntries_Delete(t *testing.T) {
	info, err := entry.Where(
		entry.ID.Eq(3),
	).Delete()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}
