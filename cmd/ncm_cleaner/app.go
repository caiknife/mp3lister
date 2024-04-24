package main

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/lib/logger"
)

func newApp() *cli.App {
	app := &cli.App{
		Name:  "清除NCM文件",
		Usage: "清除路径下的NCM文件",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "输入路径，搜索该路径下的所有NCM文件",
				Value:   cli.NewStringSlice("."),
			}},
		Action: action(),
	}
	return app
}

func action() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		inputPaths := ctx.StringSlice("input")
		for _, inputPath := range inputPaths {
			logger.ConsoleLogger.Infoln("查询路径：", inputPath)
			if !fileutil.IsExist(inputPath) {
				return errors.New("该路径不存在！")
			}
			err := cleanNCM(inputPath)
			if err != nil {
				return errors.WithMessage(err, "app action error")
			}
		}
		return nil
	}
}

func cleanNCM(inputPath string) error {
	abs, err := filepath.Abs(inputPath)
	if err != nil {
		return errors.WithMessage(err, "clean ncm error")
	}
	err = filepath.WalkDir(abs, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return errors.WithMessage(err, "walk dir error")
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".ncm") {
			return nil
		}

		suffix := filepath.Ext(path)
		fileName := strings.TrimSuffix(path, suffix) + ".mp3"

		if fileutil.IsExist(fileName) {
			logger.ConsoleLogger.Warnln(fileName, "ncm转换后的mp3文件已经存在，原ncm文件要被删除")
			logger.ConsoleLogger.Warnln("删除原ncm文件", path)
			err := fileutil.RemoveFile(path)
			if err != nil {
				logger.ConsoleLogger.Fatalln(err)
				return err
			}
		}
		return nil
	})

	return err
}
