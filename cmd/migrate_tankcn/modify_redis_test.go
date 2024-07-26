package main

import (
	"testing"

	_ "github.com/caiknife/mp3lister/test"
)

func Test_clearRedisKeys(t *testing.T) {
	err := clearRedisKeys()
	if err != nil {
		t.Error(err)
		return
	}
}
