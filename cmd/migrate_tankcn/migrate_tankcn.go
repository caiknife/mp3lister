package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/config"
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
		Name:  "tankcn 数据迁移",
		Usage: "tankcn 数据迁移 ",
		Flags: []cli.Flag{
			config.EnvFlag,
			&cli.BoolFlag{
				Name:    "db",
				Aliases: []string{"d"},
				Usage:   "修改DB数据",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "redis",
				Aliases: []string{"r"},
				Usage:   "修改Redis数据",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "迁移Redis数据",
				Value:   false,
			},
		},
		Action: action(),
	}
	return app
}

func action() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if ctx.Bool("db") {
			if err := modifyDB(); err != nil {
				err = errors.WithMessage(err, "modify db failed")
				return err
			}
		}

		if ctx.Bool("redis") {
			if err := modifyRedis(); err != nil {
				err = errors.WithMessage(err, "modify redis failed")
				return err
			}
		}

		return nil
	}
}
