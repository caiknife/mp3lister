package lib

import (
	"fmt"
	"testing"
)

var (
	testsMP3Files = []string{
		"/Users/caiknife/Music/虾米音乐/Rasmus Faber-銀河鉄道999 ~はじめてのチュウ.mp3",
		"/Users/caiknife/Music/虾米音乐/Joseph Williams,Jason Weaver,Ernie Sabella - Hakuna Matata.mp3",
	}
)

func TestNewMP3(t *testing.T) {
	for i, file := range testsMP3Files {
		t.Run(fmt.Sprintf("mp3 test %d", i+1), func(t *testing.T) {
			mp3, err := NewMP3(file)
			if err != nil {
				t.Error(err)
			}
			t.Log(mp3)
		})
	}
}
