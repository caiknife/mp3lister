package test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/random"
)

func TestUUID(t *testing.T) {
	v4, err := random.UUIdV4()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(len(v4), v4)
	s := strings.ReplaceAll(v4, "-", "")
	t.Log(len(s), s)
}

func TestListFiles(t *testing.T) {
	names, err := fileutil.ListFileNames(".")
	if err != nil {
		t.Error(err)
		return
	}
	for _, name := range names {
		t.Log(filepath.Abs(name))
	}
}
