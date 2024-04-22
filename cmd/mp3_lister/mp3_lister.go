package main

import (
	"os"

	"github.com/caiknife/mp3lister/lib/logger"
)

// 列出当前文件夹下的所有mp3文件
func main() {
	if err := newApp().Run(os.Args); err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
}
