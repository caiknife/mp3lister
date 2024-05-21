package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"

	"github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib/logger"
)

func main() {
	dsn, b := config.Config.MySQL.Get(config.DB_Music)
	if !b {
		logger.ConsoleLogger.Fatalln("数据库连接不存在！")
		return
	}

	genORM(dsn, "music")
}

var (
	dataMap = map[string]func(columnType gorm.ColumnType) (dataType string){
		"json": func(columnType gorm.ColumnType) (dataType string) {
			return "datatypes.JSON"
		},
		"longtext": func(columnType gorm.ColumnType) (dataType string) {
			return "datatypes.JSON"
		},
		"tinyint": func(columnType gorm.ColumnType) (dataType string) {
			return "bool"
		},
	}
)

func genORM(dsn, ormName string) {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}

	g := gen.NewGenerator(gen.Config{
		OutPath: "./orm/" + ormName,
		// OutFile:        "",
		ModelPkgPath:  "./orm/" + ormName + "/model",
		WithUnitTest:  true,
		FieldSignable: true,
		// FieldNullable:     true,
		// FieldCoverable:    true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		Mode:              gen.WithDefaultQuery | gen.WithoutContext,
	})

	g.UseDB(db)
	g.WithDataTypeMap(dataMap)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
