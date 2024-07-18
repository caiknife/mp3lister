package test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gen"

	"github.com/caiknife/mp3lister/lib/entity"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/mp3lister/orm/music/model"
)

func TestPlayers_Create(t *testing.T) {
	entries := types.Slice[*model.Player]{}
	for range 500 {
		b := gofakeit.Person()
		e := &model.Player{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      b.FirstName + " " + b.LastName,
			Phone:     b.Contact.Phone,
			Email:     b.Contact.Email,
		}
		p := entity.NewPlayer(e)
		p.Extra.Address = *b.Address
		p.Extra.CreditCard = *b.CreditCard
		entries = append(entries, p.Value())
	}

	err := player.CreateInBatches(entries, 100)
	if err != nil {
		t.Error(err)
		return
	}
	entries.ForEach(func(m *model.Player, i int) {
		t.Log(m)
	})
}

func TestPlayers_Restore(t *testing.T) {
	simple, err := player.Unscoped().Where(player.ID).UpdateSimple(
		player.DeletedAt.Value(nil),
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(simple)
}

func TestPlayer_Get(t *testing.T) {
	find, err := player.Where(
		player.ID,
	).First()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(find)
	p := entity.NewPlayer(find)
	t.Log(p)
}

func TestPlayers_GetAll(t *testing.T) {
	find, err := player.Order(player.ID.Desc()).Find()
	if err != nil {
		t.Error(err)
		return
	}
	for _, m := range find {
		t.Log(m)
	}
}

func TestPlayers_FindInBatch(t *testing.T) {
	batch, err := player.FindInBatch(10, func(tx gen.Dao, batch int) error {
		return nil
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(len(batch))
}

func TestPlayers_DeleteAll(t *testing.T) {
	info, err := player.Where(player.ID).Delete()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}

func TestPlayers_Truncate(t *testing.T) {
	info, err := player.Unscoped().Where(player.ID).Delete()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}
