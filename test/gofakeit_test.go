package test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"

	"github.com/caiknife/mp3lister/lib/types"
)

func TestGoFakeIt(t *testing.T) {
	mp3 := types.MP3{}
	err := gofakeit.Struct(&mp3)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(mp3.String())
}
