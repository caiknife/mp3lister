package types

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gcottom/audiometa/v2"
)

var searchPaths = Slice[string]{
	"/Users/caiknife/Music/网易云音乐",
	"/Users/caiknife/Music/QQ音乐",
	"/Users/caiknife/Music/虾米音乐",
	"/Users/caiknife/百度云同步盘/LINDY HOP MIX",
}

var m4aFile = "/Users/caiknife/Music/虾米音乐/Jonathan Stout and his Campus Five - Moppin' and Boppin' - 18 I Can't Believe You're in Love with Me.m4a"

func TestSearchM4aFiles(t *testing.T) {
	result := Slice[string]{}
	searchPaths.ForEach(func(path string, i int) {
		err := filepath.WalkDir(path, func(filePath string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !strings.HasSuffix(d.Name(), "m4a") {
				return nil
			}
			result = append(result, filePath)
			return nil
		})
		if err != nil {
			t.Fatal(err)
			return
		}
	})

	result.ForEach(func(file string, i int) {
		t.Log(file)
	})
}

func TestAudioMeta(t *testing.T) {
	open, err := os.Open(m4aFile)
	if err != nil {
		t.Error(err)
		return
	}
	path, err := audiometa.Open(open, audiometa.ParseOptions{Format: audiometa.M4A})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(path.BPM())
}
