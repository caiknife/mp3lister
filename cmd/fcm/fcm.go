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
		Name:     "防沉迷系统测试接口",
		Usage:    "防沉迷系统测试接口",
		Flags:    []cli.Flag{},
		Action:   action(),
		Commands: commands,
	}
	return app
}

func action() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		logger.ConsoleLogger.Infoln("请输入help获取更多详细内容")
		return nil
	}
}
