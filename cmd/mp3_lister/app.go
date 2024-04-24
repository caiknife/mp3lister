package main

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2/v2"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/lib"
	"github.com/caiknife/mp3lister/lib/logger"
	"github.com/caiknife/mp3lister/lib/types"
)

func action() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		inputPath := ctx.StringSlice("input")
		outputPath := ctx.String("output")

		if !strings.HasSuffix(outputPath, ".csv") {
			outputPath += ".csv"
		}

		mp3Files := types.Slice[*types.MP3]{}
		for _, s := range inputPath {
			logger.ConsoleLogger.Infoln("查询路径：", s)
			if !fileutil.IsExist(s) {
				return errors.New("该路径不存在！")
			}
			files, err := collectFiles(s)
			if err != nil {
				return err
			}
			if files.IsEmpty() {
				continue
			}
			logger.ConsoleLogger.Infoln("找到MP3文件数量：", files.Len())
			mp3Files = append(mp3Files, files...)
		}

		if mp3Files.IsEmpty() {
			logger.ConsoleLogger.Warnln("没有找到MP3文件")
			return nil
		}

		err := lib.WriteCSV(mp3Files, outputPath)
		if err != nil {
			return err
		}

		logger.ConsoleLogger.Infoln("输出文件：", outputPath)

		return nil
	}
}

func collectFiles(inputPath string) (types.Slice[*types.MP3], error) {
	mp3files := types.Slice[*types.MP3]{}

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
		return nil, err
	}

	return mp3files, nil
}

func newApp() *cli.App {
	app := &cli.App{
		Name:  "MP3文件列表展示",
		Usage: "将路径下的MP3文件导出为CSV文件",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "输入路径，搜索该路径下的所有MP3文件",
				Value:   cli.NewStringSlice("."),
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "输出文件名，输出CSV统计文件到该文件，不需要带.csv扩展名",
				Value:   "output",
			},
		},
		Action: action(),
	}

	return app
}
