package test

import (
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {
	dir, err := os.ReadDir(".")
	if err != nil {
		t.Error(err)
		return
	}
	for _, entry := range dir {
		t.Log(entry.Name())
	}
}
