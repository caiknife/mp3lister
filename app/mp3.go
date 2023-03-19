package app

import (
	"github.com/bogem/id3v2/v2"
)

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

type MP3Collection []*MP3

func (m MP3Collection) Len() int {
	return len(m)
}

func (m MP3Collection) Less(i, j int) bool {
	return m[i].Artist < m[j].Artist &&
		m[i].Album < m[j].Album &&
		m[i].Title < m[j].Title
}

func (m MP3Collection) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
