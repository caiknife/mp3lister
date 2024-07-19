package config

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config *AppConfig

	DBMusic     *gorm.DB
	DBWarTankCN *gorm.DB

	RedisDefault *redis.Client
)

const (
	dbMusic       = "music"
	dbMusicRead_1 = "music_read_1"
	dbMusicRead_2 = "music_read_2"
	dbWarTank     = "wartank"
	dbWarTankCN   = "wartank_cn"
)

const (
	Redis_Default = "default"
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
