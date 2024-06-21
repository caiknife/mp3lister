package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"

	"github.com/caiknife/mp3lister/lib"
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
	lib.InitYAMLConfig(Config, "config.yml")
}

func initORM() {
	newLogger := gLogger.New(
		log.New(os.Stdout, "", log.LstdFlags), // io writer
		gLogger.Config{
			SlowThreshold:             time.Second * 2, // Slow SQL threshold
			LogLevel:                  gLogger.Info,    // Log level
			IgnoreRecordNotFoundError: true,            // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,            // Don't include params in the SQL log
			Colorful:                  true,            // Disable color
		},
	)

	var err error
	DB, err = gorm.Open(mysql.Open(Config.MySQL[DB_Music]), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
	// 读写分离
	err = DB.Use(dbresolver.Register(dbresolver.Config{
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
	music.SetDefault(DB)
}

const (
	DB_Music        = "music"
	DB_Music_Read_1 = "music_read_1"
	DB_Music_Read_2 = "music_read_2"
)

type RoundRobinPolicy struct {
	index int
}

func (r *RoundRobinPolicy) Resolve(connPools []gorm.ConnPool) gorm.ConnPool {
	if r.index >= len(connPools) {
		r.index = 0
	}
	result := connPools[r.index]
	r.index++
	return result
}
