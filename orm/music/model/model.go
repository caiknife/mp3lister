package model

import (
	"github.com/caiknife/mp3lister/lib/fjson"
)

func (b *Book) String() string {
	toString, err := fjson.MarshalToString(b)
	if err != nil {
		return ""
	}
	return toString
}

func (s *Song) String() string {
	toString, err := fjson.MarshalToString(s)
	if err != nil {
		return ""
	}
	return toString
}

func (m *Movie) String() string {
	toString, err := fjson.MarshalToString(m)
	if err != nil {
		return ""
	}
	return toString
}

func (c *Car) String() string {
	toString, err := fjson.MarshalToString(c)
	if err != nil {
		return ""
	}
	return toString
}

func (p *Player) String() string {
	toString, err := fjson.MarshalToString(p)
	if err != nil {
		return ""
	}
	return toString
}
