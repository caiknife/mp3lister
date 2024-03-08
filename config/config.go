package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/orm/music"
)

var (
	Config *AppConfig
	DB     *gorm.DB
)

func init() {
	initConfig()
	initORM()
}

func initConfig() {
	Config = &AppConfig{}
	InitYAMLConfig(Config, "config.yml")
}

func initORM() {
	DB, err := gorm.Open(mysql.Open(Config.MySQL[DB_Music]))
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
	music.SetDefault(DB)
}

const (
	DB_Music = "music"
)
