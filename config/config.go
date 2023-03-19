package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	Config *viper.Viper
)

func init() {
	initConfig()
}

func initConfig() {
	Config = viper.New()
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	Config.AddConfigPath(dir)
	Config.AddConfigPath("../")

	Config.SetConfigName("config")
	Config.SetConfigType("yml")

	if err := Config.ReadInConfig(); err != nil {
		panic(err)
	}
}
