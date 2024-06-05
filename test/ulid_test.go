package test

import (
	"testing"

	"github.com/oklog/ulid/v2"
)

func TestULID(t *testing.T) {
	for range 10 {
		id := ulid.Make()
		t.Log(id.String(), id.Time())
	}
}
