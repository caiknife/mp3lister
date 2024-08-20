package types

import (
	"io"
	"os"

	"github.com/bogem/id3v2/v2"
	"github.com/pkg/errors"
	"github.com/tcolgate/mp3"

	"github.com/caiknife/mp3lister/lib/fjson"
)

const (
	TagQuodLibetMemo = "QuodLibet::memo"
)

type MP3 struct {
	OriginFile string  `json:"origin_file"`
	BPM        string  `json:"bpm"`
	Title      string  `json:"title"`
	Artist     string  `json:"artist"`
	Album      string  `json:"album"`
	Memo       string  `json:"memo"`
	Length     float64 `json:"length"`
}

func (m *MP3) String() string {
	toString, err := fjson.MarshalToString(m)
	if err != nil {
		return ""
	}
	return toString
}

func NewMP3(name string) (*MP3, error) {
	m := &MP3{OriginFile: name}
	return m.Init()
}

func (m *MP3) LoadLength() error {
	open, err := os.Open(m.OriginFile)
	if err != nil {
		return errors.WithMessage(err, "mp3 file open error")
	}
	defer open.Close()

	v := 0.0
	d := mp3.NewDecoder(open)
	var f mp3.Frame
	skipped := 0

	for {
		if err := d.Decode(&f, &skipped); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return errors.WithMessage(err, "mp3 decode error")
		}

		v += f.Duration().Seconds()
	}

	m.Length = v

	return nil
}

func (m *MP3) loadMemo(tag *id3v2.Tag) error {
	frames := tag.GetFrames("TXXX")
	for _, frame := range frames {
		switch frame := frame.(type) {
		case id3v2.UserDefinedTextFrame:
			if frame.Description == TagQuodLibetMemo {
				m.Memo = CutInvisibleSeparator(frame.Value)
			}
		}
	}
	return nil
}

func (m *MP3) Init() (*MP3, error) {
	tag, err := id3v2.Open(m.OriginFile, id3v2.Options{Parse: true})
	if err != nil {
		return nil, errors.WithMessage(err, "id3v2 open error")
	}
	defer tag.Close()

	m.BPM = tag.GetTextFrame(tag.CommonID("BPM")).Text
	m.Title = tag.Title()
	m.Artist = m.transform(tag.Artist())
	m.Album = tag.Album()
	_ = m.loadMemo(tag)
	// _ = m.LoadLength()

	return m, nil
}

func (m *MP3) transform(input string) string {
	return CutInvisibleSeparator(input)
}
