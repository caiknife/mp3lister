package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib/logger"
)

func main() {
	dsn, b := config.Config.MySQL.Get("music")
	if !b {
		logger.ConsoleLogger.Fatalln("数据库连接不存在！")
		return
	}

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logger.ConsoleLogger.Errorln(err)
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:           "./orm/dal",
		WithUnitTest:      true,
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		Mode:              gen.WithDefaultQuery | gen.WithoutContext,
	})

	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
