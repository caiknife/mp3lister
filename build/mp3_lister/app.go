package main

import (
	"encoding/csv"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/bogem/id3v2/v2"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/golang-module/carbon/v2"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/lib"
)

func action() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		var inputPath, outputPath string
		if ctx.String("input") != "" {
			inputPath = ctx.String("input")
		} else {
			inputPath = "."
		}
		inputPath, err := filepath.Abs(inputPath)
		if err != nil {
			return nil
		}

		if ctx.String("output") != "" {
			outputPath = ctx.String("output") + ".csv"
		} else {
			outputPath = time.Now().Format(carbon.ShortDateTimeLayout) + ".csv"
		}

		if !fileutil.IsExist(inputPath) {
			return errors.New("该路径不存在！")
		}

		log.Println("查询路径：", inputPath)

		files, err := collectFiles(inputPath)
		if err != nil {
			return err
		}

		if files.Len() == 0 {
			return errors.New("该路径下没有MP3文件！")
		}

		err = writeFiles(files, outputPath)
		if err != nil {
			return err
		}

		log.Println("输出文件：", outputPath)

		return nil
	}
}

func writeFiles(mp3files lib.MP3Collection, outputPath string) error {
	create, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer create.Close()

	writer := csv.NewWriter(create)
	err = writer.Write([]string{"No.", "BPM", "Title", "Artist", "Album"})
	if err != nil {
		return err
	}
	for i, file := range mp3files {
		err := writer.Write([]string{
			cast.ToString(i + 1),
			file.BPM,
			file.Title,
			file.Artist,
			file.Album,
		})
		if err != nil {
			return err
		}
	}
	writer.Flush()

	return nil
}

func collectFiles(inputPath string) (lib.MP3Collection, error) {
	mp3files := lib.MP3Collection{}

	err := filepath.WalkDir(inputPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(d.Name(), ".mp3") {
			return nil
		}
		mp3, err := lib.NewMP3(path)
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
			&cli.StringFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "输入路径，搜索该路径下的所有MP3文件",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "输出文件名，输出CSV统计文件到该文件，不需要带.csv扩展名",
			},
		},
		Action: action(),
	}

	return app
}
