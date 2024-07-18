package config

import (
	"testing"

	"github.com/caiknife/mp3lister/lib"
)

func TestFile(t *testing.T) {
	f := &File{}
	lib.InitYAMLConfig(f, "config.yml")
	t.Log(f)
}
