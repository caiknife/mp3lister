package types

import (
	"fmt"
	"testing"

	"github.com/spf13/cast"
)

var (
	testsMP3Files = []string{
		"/Users/caiknife/Music/网易云音乐/Joe Eckert/UNITED STATES AIR FORCE AIRMEN OF NOTE： Let's Dance/United States Air Force Band - Airmen of Note - Rockin' in Rhythm.mp3",
		"/Users/caiknife/Music/虾米音乐/Rasmus Faber-銀河鉄道999 ~はじめてのチュウ.mp3",
		"/Users/caiknife/Music/虾米音乐/Joseph Williams,Jason Weaver,Ernie Sabella - Hakuna Matata.mp3",
	}
)

func TestFromStringToInt(t *testing.T) {
	toInt := cast.ToInt("")
	t.Log(toInt)
}

func TestBPM(t *testing.T) {
	testFile := "/Users/caiknife/Music/网易云音乐/Stockholm Swing All Stars/In the Spirit of/Stockholm Swing All Stars - Isfahan.mp3"
	mp3, err := NewMP3(testFile)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(mp3.BPM)
}

func TestNewMP3(t *testing.T) {
	for i, file := range testsMP3Files {
		t.Run(fmt.Sprintf("mp3 test %d", i+1), func(t *testing.T) {
			mp3, err := NewMP3(file)
			if err != nil {
				t.Error(err)
			}
			t.Log(mp3.Staffing, len(mp3.Staffing))
		})
	}
}
