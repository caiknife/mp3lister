package main

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/duke-git/lancet/v2/fileutil"

	"github.com/caiknife/mp3lister/lib/logger"
)

// 清除当前目录下的ncm文件
func main() {
	abs, err := filepath.Abs(".")
	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}

	err = filepath.WalkDir(abs, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
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

	if err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
}
