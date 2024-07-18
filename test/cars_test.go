package test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/mp3lister/orm/music/model"
)

func TestCars_Create(t *testing.T) {
	entries := types.Slice[*model.Car]{}
	for range 500 {
		b := gofakeit.Car()
		e := &model.Car{
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			Type:         b.Type,
			Fuel:         b.Fuel,
			Transmission: b.Transmission,
			Brand:        b.Brand,
			Model:        b.Model,
			Year:         int32(b.Year),
		}
		entries = append(entries, e)
	}

	err := car.CreateInBatches(entries, 100)
	if err != nil {
		t.Error(err)
		return
	}
	entries.ForEach(func(m *model.Car, i int) {
		t.Log(m)
	})
}

func TestCars_Restore(t *testing.T) {
	simple, err := car.Unscoped().Where(car.ID).UpdateSimple(
		car.DeletedAt.Value(nil),
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(simple)
}

func TestCars_Update(t *testing.T) {
	b := gofakeit.Car()
	simple, err := car.Where(
		car.ID.Eq(3),
	).UpdateSimple(
		car.Type.Value(b.Type),
		car.Fuel.Value(b.Fuel),
		car.Transmission.Value(b.Transmission),
		car.Brand.Value(b.Brand),
		car.Model.Value(b.Model),
		car.Year.Value(int32(b.Year)),
	)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(simple)
}

func TestCars_Get(t *testing.T) {
	for i := range 10 {
		find, err := car.Where(
			car.ID.Eq(uint64(i + 1)),
		).First()
		if err != nil {
			t.Error(err)
			continue
		}
		t.Log(find)
	}
}

func TestCars_GetAll(t *testing.T) {
	find, err := car.Order(car.ID.Desc()).Find()
	if err != nil {
		t.Error(err)
		return
	}
	for _, m := range find {
		t.Log(m)
	}
}

func TestCars_Truncate(t *testing.T) {
	info, err := car.Unscoped().Where(car.ID).Delete()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}

func TestCars_Delete(t *testing.T) {
	info, err := car.Where(
		car.ID.Eq(3),
	).Delete()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}

func TestCars_DeleteAll(t *testing.T) {
	info, err := car.Where(car.ID).Delete()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}
