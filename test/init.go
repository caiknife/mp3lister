package test

import (
	"github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib"
	"github.com/caiknife/mp3lister/lib/logger"
)

func init() {
	e := &config.Environment{
		Env: "lan",
	}
	err := config.LoadConfigFile(e,
		config.InitDBMusic,
		config.InitDBWarTankCN,
	)
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
}

func snowflakeMachineID() uint16 {
	return lib.SnowflakeMachineID()
}
