package main

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/duke-git/lancet/v2/fileutil"
)

func main() {
	abs, err := filepath.Abs(".")
	if err != nil {
		log.Fatalln(err)
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
			log.Println(fileName, "ncm转换后的mp3文件已经存在，原ncm文件要被删除")
			log.Println("删除原ncm文件", path)
			err := fileutil.RemoveFile(path)
			if err != nil {
				log.Fatalln(err)
				return err
			}
		}

		return nil
	})

	if err != nil {
		log.Fatalln(err)
		return
	}
}
