package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"

	"github.com/caiknife/mp3lister/lib"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/orm/music"
	"github.com/caiknife/mp3lister/orm/wartankcn"
)

func InitCliContext(ctx *cli.Context, callbacks ...func()) error {
	env := ctx.String("env")
	e := &Environment{
		Env: env,
	}
	return LoadConfigFile(e, callbacks...)
}

func LoadConfigFile(e *Environment, callbacks ...func()) error {
	f := &File{}
	lib.InitYAMLConfig(f, "config.yml")
	get, b := f.Get(e.Env)
	if !b {
		return errors.New(fmt.Sprintf("config file %s section not found", e.Env))
	}
	Config = get
	for _, callback := range callbacks {
		callback()
	}
	return nil
}

func InitDBWarTankCN() {
	newLogger := gLogger.New(
		log.New(os.Stdout, "", log.LstdFlags), // io writer
		gLogger.Config{
			SlowThreshold:             time.Second * 2, // Slow SQL threshold
			LogLevel:                  gLogger.Info,    // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,           // Don't include params in the SQL log
			Colorful:                  true,            // Disable color
		},
	)

	var err error
	DBWarTankCN, err = gorm.Open(mysql.Open(Config.MySQL[DB_Wartank_CN]), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
	wartankcn.SetDefault(DBWarTankCN)
}

func InitDBMusic() {
	newLogger := gLogger.New(
		log.New(os.Stdout, "", log.LstdFlags), // io writer
		gLogger.Config{
			SlowThreshold:             time.Second * 2, // Slow SQL threshold
			LogLevel:                  gLogger.Info,    // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,           // Don't include params in the SQL log
			Colorful:                  true,            // Disable color
		},
	)

	var err error
	DBMusic, err = gorm.Open(mysql.Open(Config.MySQL[DB_Music]), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
	// 读写分离
	err = DBMusic.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{
			mysql.Open(Config.MySQL[DB_Music]),
		},
		Replicas: []gorm.Dialector{
			mysql.Open(Config.MySQL[DB_Music_Read_1]),
			mysql.Open(Config.MySQL[DB_Music_Read_2]),
		},
		Policy:            &RoundRobinPolicy{},
		TraceResolverMode: true,
	}))
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
	music.SetDefault(DBMusic)
}
