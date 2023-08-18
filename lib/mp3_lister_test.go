package lib

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

var (
	testDirs = []string{
		"/Users/caiknife/Music/虾米音乐",
		"/Users/caiknife/Music/QQ音乐",
	}
)

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
