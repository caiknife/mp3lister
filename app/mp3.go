package app

import "github.com/bogem/id3v2/v2"

type MP3 struct {
	OriginFile string
	BPM        string
	Title      string
	Artist     string
	Album      string
}

func NewMP3(name string) (*MP3, error) {
	mp3 := &MP3{OriginFile: name}
	return mp3.Init()
}

func (m *MP3) Init() (*MP3, error) {
	tag, err := id3v2.Open(m.OriginFile, id3v2.Options{Parse: true})
	if err != nil {
		return nil, err
	}
	defer tag.Close()

	m.BPM = tag.GetTextFrame(tag.CommonID("BPM")).Text
	m.Title = tag.Title()
	m.Artist = tag.Artist()
	m.Album = tag.Album()
	return m, nil
}
