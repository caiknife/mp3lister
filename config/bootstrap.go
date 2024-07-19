package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
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

func InitRedisDefault() {
	red, b := Config.Redis.Get(Redis_Default)
	if !b {
		logger.ConsoleLogger.Fatalln(RedisDefault, "redis config not exist")
		return
	}
	RedisDefault = redis.NewClient(&redis.Options{
		Addr:     red.Addr,
		Password: red.Password, // 没有密码，默认值
		DB:       red.DB,       // 默认DB 0
	})
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
	db, b := Config.MySQL.Get(dbWarTankCN)
	if !b {
		logger.ConsoleLogger.Fatalln(dbWarTankCN, "mysql config not exist")
		return
	}
	DBWarTankCN, err = gorm.Open(mysql.Open(db), &gorm.Config{
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

	dbM, b := Config.MySQL.Get(dbMusic)
	if !b {
		logger.ConsoleLogger.Fatalln(dbMusic, "mysql config not exist")
		return
	}
	dbMR1, b := Config.MySQL.Get(dbMusicRead_1)
	if !b {
		logger.ConsoleLogger.Fatalln(dbMusicRead_1, "mysql config not exist")
		return
	}
	dbMR2, b := Config.MySQL.Get(dbMusicRead_2)
	if !b {
		logger.ConsoleLogger.Fatalln(dbMusicRead_2, "mysql config not exist")
		return
	}

	DBMusic, err = gorm.Open(mysql.Open(dbM), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
	// 读写分离
	err = DBMusic.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{
			mysql.Open(dbM),
		},
		Replicas: []gorm.Dialector{
			mysql.Open(dbMR1),
			mysql.Open(dbMR2),
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
