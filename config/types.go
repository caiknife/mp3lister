package config

import (
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type AppConfig struct {
	MySQL types.Map[string] `yaml:"mysql"`
}

func (a *AppConfig) String() string {
	toString, err := fjson.MarshalToString(a)
	if err != nil {
		return ""
	}
	return toString
}
