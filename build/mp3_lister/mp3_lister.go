package main

import (
	"os"

	"github.com/caiknife/mp3lister/lib"
)

func main() {
	if err := newApp().Run(os.Args); err != nil {
		lib.ConsoleLogger.Fatalln(err)
		return
	}
}
