package main

import (
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/lib"
	"github.com/caiknife/mp3lister/lib/logger"
)

func action() cli.ActionFunc {
	return func(c *cli.Context) error {
		inputFiles := c.StringSlice("input")
		if len(inputFiles) == 0 {
			return errors.New("no input files")
		}

		for _, inputFile := range inputFiles {
			logger.ConsoleLogger.Infoln("Reading input file:", inputFile)
			u, err := lib.ReadM3U(inputFile)
			if err != nil {
				return errors.WithMessage(err, "read input file")
			}
			outputFile := lib.GetOutputFileWithExt(inputFile, "csv")
			err = lib.WriteCSV(u, outputFile)
			if err != nil {
				return errors.WithMessage(err, "write output file")
			}
			logger.ConsoleLogger.Infoln("Wrote output file:", outputFile)
		}
		return nil
	}
}

func newApp() *cli.App {
	app := &cli.App{
		Name:  "将m3u歌单文件内容导出成csv文件",
		Usage: "将m3u歌单文件内容导出成csv文件",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "input",
				Aliases: []string{"i"},
				Usage:   "需要导出的m3u歌单文件",
				Value:   cli.NewStringSlice(),
			}},
		Action: action(),
	}
	return app
}
