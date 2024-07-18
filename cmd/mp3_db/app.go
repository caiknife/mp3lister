package main

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
	"github.com/caiknife/mp3lister/orm/music"
	"github.com/caiknife/mp3lister/orm/music/model"
)

func action() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		inputPath := ctx.StringSlice("input")
		dsn := ctx.String("dsn")

		logger.ConsoleLogger.Infoln("dsn:", dsn, config.Config.MySQL[dsn])
		mp3Files := types.Slice[*types.MP3]{}
		for _, s := range inputPath {
			logger.ConsoleLogger.Infoln("input path:", s)
			if !fileutil.IsExist(s) {
				return errors.New(s + " input path does not exist")
			}

			files := collectFiles(s)
			if files.IsEmpty() {
				continue
			}
			logger.ConsoleLogger.Infoln("mp3 files total", files.Len())
			mp3Files = append(mp3Files, files...)
		}

		if mp3Files.IsEmpty() {
			logger.ConsoleLogger.Warnln("no mp3 files")
			return nil
		}

		err := saveToDB(config.Config.MySQL[dsn], mp3Files)
		if err != nil {
			return err
		}
		return nil
	}
}

func saveToDB(dsn string, mp3Files types.Slice[*types.MP3]) error {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}

	if !db.Migrator().HasTable(&model.Song{}) {
		err := db.AutoMigrate(&model.Song{})
		if err != nil {
			return err
		}
	}

	music.SetDefault(db)
	// 清空表
	_, err = music.Song.Unscoped().Where(music.Song.ID).Delete()
	if err != nil {
		return err
	}

	// 插入数据
	songs := slice.Map[*types.MP3, *model.Song](mp3Files, func(_ int, item *types.MP3) *model.Song {
		song := &model.Song{
			Title:      item.Title,
			Artist:     item.Artist,
			Album:      item.Album,
			Bpm:        cast.ToInt32(item.BPM),
			OriginFile: item.OriginFile,
			Length:     item.Length,
		}
		return song
	})
	err = music.Song.CreateInBatches(songs, 100)
	if err != nil {
		return errors.WithMessage(err, "create in batches")
	}

	return nil
}

func collectFiles(inputPath string) types.Slice[*types.MP3] {
	var mp3files = types.Slice[*types.MP3]{}
	err := filepath.WalkDir(inputPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".mp3") {
			return nil
		}
		mp3, err := types.NewMP3(path)
		if err != nil {
			if errors.Is(err, id3v2.ErrUnsupportedVersion) {
				return nil
			}
			return err
		}

		mp3files = append(mp3files, mp3)
		return nil
	})
	if err != nil {
		logger.ConsoleLogger.Errorln(err)
	}
	return mp3files
}

func newApp() *cli.App {
	app := &cli.App{
		Name:  "MP3文件列表展示",
		Usage: "将路径下的MP3文件导出为CSV文件",
		Flags: []cli.Flag{
			config.EnvFlag,
			&cli.StringSliceFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "输入路径，搜索该路径下的所有MP3文件",
				Value:   cli.NewStringSlice("."),
			},
			&cli.StringFlag{
				Name:    "dsn",
				Aliases: []string{"d"},
				Usage:   "数据库连接的名称",
				Value:   "music",
			},
		},
		Action: action(),
	}

	return app
}
