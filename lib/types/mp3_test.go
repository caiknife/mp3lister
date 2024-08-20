package types

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/dhowden/tag"
	"github.com/spf13/cast"
)

var (
	testsMP3Files = Slice[string]{
		"/Users/caiknife/Music/网易云音乐/Joe Eckert/UNITED STATES AIR FORCE AIRMEN OF NOTE： Let's Dance/United States Air Force Band - Airmen of Note - Rockin' in Rhythm.mp3",
		"/Users/caiknife/Music/虾米音乐/Rasmus Faber-銀河鉄道999 ~はじめてのチュウ.mp3",
		"/Users/caiknife/Music/虾米音乐/Joseph Williams,Jason Weaver,Ernie Sabella - Hakuna Matata.mp3",
	}
)

func TestUserTag(t *testing.T) {
	testFiles := Slice[string]{
		"/Users/caiknife/Music/网易云音乐/The Bourbon Street Stompers/Caution Hot!/The Bourbon Street Stompers - Amazing Grace.mp3",
		"/Users/caiknife/Music/网易云音乐/Cats and Dinosaurs/Kapitalismen är en dröm/Cats and Dinosaurs - Ojämlikheten skördar människoliv.mp3",
	}

	testFiles.ForEach(func(testFile string, i int) {
		mp3, err := NewMP3(testFile)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(mp3.Memo)
	})
}

func TestM4aFile(t *testing.T) {
	testFile := "/Users/caiknife/Music/虾米音乐/Jonathan Stout and his Campus Five - Moppin' and Boppin' - 18 I Can't Believe You're in Love with Me.m4a"
	mp3, err := NewMP3(testFile)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(mp3)

	// testFile = "/Users/caiknife/Music/网易云音乐/Cats and Dinosaurs/Kapitalismen är en dröm/Cats and Dinosaurs - Ojämlikheten skördar människoliv.mp3"
	open, err := os.Open(testFile)
	if err != nil {
		t.Error(err)
		return
	}
	from, err := tag.ReadFrom(open)
	if err != nil {
		t.Error()
		return
	}
	t.Log(from.Format())
	for key, val := range from.Raw() {
		t.Log(key, val)
	}
}

func TestFromStringToInt(t *testing.T) {
	toInt := cast.ToInt("")
	t.Log(toInt)
}

func TestMP3Length(t *testing.T) {
	dir := "/Users/caiknife/Music/网易云音乐"
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".mp3") {
			return nil
		}

		m := &MP3{OriginFile: path}
		err = m.LoadLength()
		if err != nil {
			t.Error(m)
			return err
		}
		t.Log(m)
		return nil
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestBPM(t *testing.T) {
	testFile := "/Users/caiknife/Music/网易云音乐/Stockholm Swing All Stars/In the Spirit of/Stockholm Swing All Stars - Isfahan.mp3"
	mp3, err := NewMP3(testFile)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(mp3)
}

func TestNewMP3(t *testing.T) {
	testsMP3Files.ForEach(func(file string, i int) {
		t.Run(fmt.Sprintf("mp3 test %d", i+1), func(t *testing.T) {
			mp3, err := NewMP3(file)
			if err != nil {
				t.Error(err)
			}
			t.Log(mp3)
		})
	})
}
