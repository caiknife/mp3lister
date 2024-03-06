package model

import (
	"github.com/caiknife/mp3lister/lib/fjson"
)

func (e *Entry) String() string {
	toString, err := fjson.MarshalToString(e)
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
