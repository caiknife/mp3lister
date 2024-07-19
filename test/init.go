package test

import (
	"github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/orm/music"
)

var (
	book   = music.Book
	car    = music.Car
	movie  = music.Movie
	player = music.Player
	song   = music.Song
)

func init() {
	e := &config.Environment{
		Env: "test",
	}
	err := config.LoadConfigFile(e,
		config.InitDBMusic,
		config.InitDBWarTankCN,
		config.InitRedisDefault,
	)
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
	initTables()
}

func initTables() {
	book = music.Book
	car = music.Car
	movie = music.Movie
	player = music.Player
	song = music.Song
}

func snowflakeMachineID() uint16 {
	return lib.SnowflakeMachineID()
}
