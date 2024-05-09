package test

import (
	"testing"

	"github.com/teris-io/shortid"
	"golang.org/x/exp/rand"
)

func TestShortID(t *testing.T) {
	generate, err := shortid.Generate()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(generate)
}

func TestShortID_generator(t *testing.T) {
	s, err := shortid.New(0, shortid.DefaultABC, rand.Uint64())
	if err != nil {
		t.Error(err)
		return
	}
	generate, err := s.Generate()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(generate)
}
