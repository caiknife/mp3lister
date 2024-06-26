package lib

import (
	"os"
	"path/filepath"

	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/logger"
)

var (
	ErrConfigFileMissing = errors.New("config file missing")
	defaultSearchPaths   = []string{
		"./", "./config/", // 当前路径和config
		"../", "../config/", // 上层路径和config
		"../../", "../../config/", // 上两层路径和config
		"../../../", "../../../config/", // 上三层路径和config
	}
)

func SearchConfigFile(fileName string, searchPaths ...string) (string, error) {
	if len(searchPaths) == 0 {
		searchPaths = defaultSearchPaths
	}

	for _, path := range searchPaths {
		p, _ := filepath.Abs(path)
		filePath := filepath.Join(p, fileName)
		if fileutil.IsExist(filePath) {
			return filePath, nil
		}
	}

	return "", ErrConfigFileMissing
}

func readFile(fileName string) ([]byte, error) {
	filePath, err := SearchConfigFile(fileName)
	if err != nil {
		return nil, err
	}
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func InitYAMLConfig(receiver any, fileName string) {
	file, err := readFile(fileName)
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}

	err = yaml.Unmarshal(file, receiver)
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
}

func InitJSONConfig(receiver any, fileName string) {
	file, err := readFile(fileName)
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}

	err = fjson.Unmarshal(file, receiver)
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
}
