package main

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/duke-git/lancet/v2/fileutil"

	"github.com/caiknife/mp3lister/lib"
)

func main() {
	abs, err := filepath.Abs(".")
	if err != nil {
		lib.ConsoleLogger.Fatalln(err)
		return
	}

	err = filepath.WalkDir(abs, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, ".ncm") {
			return nil
		}

		suffix := filepath.Ext(path)
		fileName := strings.TrimSuffix(path, suffix) + ".mp3"

		if fileutil.IsExist(fileName) {
			lib.ConsoleLogger.Warnln(fileName, "ncm转换后的mp3文件已经存在，原ncm文件要被删除")
			lib.ConsoleLogger.Warnln("删除原ncm文件", path)
			err := fileutil.RemoveFile(path)
			if err != nil {
				lib.ConsoleLogger.Fatalln(err)
				return err
			}
		}

		return nil
	})

	if err != nil {
		lib.ConsoleLogger.Fatalln(err)
		return
	}
}
