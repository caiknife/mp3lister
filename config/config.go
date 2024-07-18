package config

import (
	"gorm.io/gorm"
)

var (
	Config      *AppConfig
	DBMusic     *gorm.DB
	DBWarTankCN *gorm.DB
)

const (
	DB_Music        = "music"
	DB_Music_Read_1 = "music_read_1"
	DB_Music_Read_2 = "music_read_2"
	DB_Wartank      = "wartank"
	DB_Wartank_CN   = "wartank_cn"
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
