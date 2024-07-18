package test

import (
	"testing"

	"github.com/duke-git/lancet/v2/random"
)

func TestUpdateLength(t *testing.T) {
	float := random.RandFloat(2*60, 5*60, 10)
	update, err := song.Where(song.ID).UpdateSimple(song.Length.Value(float))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(update)
}
