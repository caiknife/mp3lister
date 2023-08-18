package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/fatih/color"
	"github.com/golang-module/carbon/v2"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/lib"
)

var (
	inputPath  string
	outputName string
)

const (
	OutputNameTemplate = "mp3lister_%s"
)

func action(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		cli.VersionPrinter(ctx)
	}
	// 输入路径
	result, err := lib.GetInputPath(ctx.String("input"))
	if err != nil {
		return err
	}
	inputPath = result

	lib.ColorPrintf("working in %s\n", color.CyanString("%s", inputPath))

	if !fileutil.IsDir(inputPath) {
		return lib.ErrInputIsNotDir
	}

	// 输出名称
	name, err := getOutputName(ctx.String("output"))
	if err != nil {
		return err
	}
	outputName = name
	lib.ColorPrintf("output name is %s\n", color.YellowString("%s", outputName))

	lister := lib.NewMP3Lister(
		lib.WithInputPath(inputPath),
		lib.WithOutputName(outputName),
		lib.WithOutputExt("csv"),
	)
	err = lister.Do()
	if err != nil {
		return err
	}

	if ctx.Bool("verbose") {
		lister.Print()
	}

	if ctx.String("savetodb") != "" {
		err := lister.SaveToDB(ctx.String("savetodb"))
		if err != nil {
			return err
		}
	}

	return nil
}

func getInputPath(path string) (result string, err error) {
	if path == "" {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = dir
	}
	result, err = filepath.Abs(path)
	if err != nil {
		return path, err
	}
	return result, nil
}

func getOutputName(name string) (result string, err error) {
	if name == "" {
		result = fmt.Sprintf(OutputNameTemplate, carbon.Now().Format("Y_m_d_H_i_s"))
		return
	}
	return name, nil
}
