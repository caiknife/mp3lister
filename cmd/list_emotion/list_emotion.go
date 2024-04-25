package main

import (
	"os"

	"github.com/caiknife/mp3lister/lib/logger"
)

func main() {
	if err := newApp().Run(os.Args); err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
}
