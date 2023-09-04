package lib

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/spf13/cast"
)

var (
	testDirs = []string{
		"/Users/caiknife/Music/虾米音乐",
		"/Users/caiknife/Music/QQ音乐",
	}

	musicDir = "/Users/caiknife/Music/网易云音乐"
)

func TestCleanNCM(t *testing.T) {
	err := filepath.WalkDir(musicDir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".ncm") {
			return nil
		}

		suffix := filepath.Ext(path)
		fileName := strings.TrimSuffix(path, suffix) + ".mp3"

		if fileutil.IsExist(fileName) {
			t.Log(fileName, "ncm转换后的mp3文件已经存在，原ncm文件要被删除")
			t.Log("删除原ncm文件", path)
			err := fileutil.RemoveFile(path)
			if err != nil {
				t.Error(err)
				return err
			}
		}

		return nil
	})

	if err != nil {
		t.Error(err)
		return
	}
}

func TestMP3Lister(t *testing.T) {
	for i, dir := range testDirs {
		t.Run(fmt.Sprintf("mp3lister test %d", i+1), func(t *testing.T) {
			l := NewMP3Lister(WithInputPath(dir), WithOutputName(cast.ToString(i+1)))
			err := l.Do()
			if err != nil {
				t.Error(err)
			}
		})
	}
}
