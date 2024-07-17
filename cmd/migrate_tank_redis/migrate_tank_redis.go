package main

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/lib/logger"
)

func main() {
	if err := newApp().Run(os.Args); err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
}

func newApp() *cli.App {
	app := &cli.App{
		Name:   "迁移 tank redis",
		Usage:  "迁移 tank redis",
		Flags:  []cli.Flag{},
		Action: action(),
	}
	return app
}

func action() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		return nil
	}
}
