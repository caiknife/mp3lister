package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/golang-module/carbon/v2"
	"github.com/urfave/cli/v2"
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
	result, err := getInputPath(ctx.String("input"))
	if err != nil {
		return err
	}
	inputPath = result

	if !fileutil.IsDir(inputPath) {
		return ErrInputIsNotDir
	}

	fmt.Printf("working in %s\n", inputPath)
	// 输出名称
	name, err := getOutputName(ctx.String("output"))
	if err != nil {
		return err
	}
	outputName = name
	fmt.Printf("output name is %s\n", outputName)

	return nil
}

func getInputPath(path string) (result string, err error) {
	if path == "" {
		dir, err := os.Getwd()
		if err != nil {
			return "", err
		}
		result = dir
	}
	result, err = filepath.Abs(result)
	if err != nil {
		return "", err
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
