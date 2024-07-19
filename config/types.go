package config

import (
	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type File = types.Map[*AppConfig]

type AppConfig struct {
	MySQL types.Map[string]       `yaml:"mysql"`
	Redis types.Map[*RedisConfig] `yaml:"redis"`
}

func (a *AppConfig) String() string {
	toString, _ := fjson.MarshalToString(a)
	return toString
}

var (
	EnvFlag = &cli.StringFlag{
		Name:    "env",
		Aliases: []string{"e"},
		Usage:   "设置程序使用的环境，根据config文件配置决定",
		Value:   "lan",
	}
)

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func (r *RedisConfig) String() string {
	toString, _ := fjson.MarshalToString(r)
	return toString
}
