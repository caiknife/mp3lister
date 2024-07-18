package config

import (
	"github.com/caiknife/mp3lister/lib/fjson"
)

type Environment struct {
	Env string `yaml:"env"`
}

func (e *Environment) String() string {
	toString, _ := fjson.MarshalToString(e)
	return toString
}
