package main

import (
	"os"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib/logger"
)

// 1USD=6RMB=400钻石
// 返还300% 6RMB=1200钻石
func main() {
	if err := newApp().Run(os.Args); err != nil {
		logger.ConsoleLogger.Fatalln(err)
		return
	}
}

func newApp() *cli.App {
	app := &cli.App{
		Name:  "tank内测充值返还",
		Usage: "tank内测中指返还 ",
		Flags: []cli.Flag{
			config.EnvFlag,
			&cli.BoolFlag{
				Name:    "full",
				Aliases: []string{"f"},
				Usage:   "完整流程，从order表读取数据处理之后，存入charge_refund表，并存入redis",
				Value:   false,
			},
			&cli.BoolFlag{
				Name:    "redis",
				Aliases: []string{"r"},
				Usage:   "只从charge_refund表读取数据并存入redis",
				Value:   false,
			},
		},
		Action: action(),
	}
	return app
}

func action() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if err := config.InitCliContext(
			ctx,
			config.InitDBWarTankCN,
			config.InitRedisDefault,
		); err != nil {
			err = errors.WithMessage(err, "init cli context")
			return err
		}

		if ctx.Bool("full") {
			if err := doRefund(); err != nil {
				err = errors.WithMessage(err, "do refund")
				return err
			}
		}

		if ctx.Bool("redis") {
			if err := doRedis(); err != nil {
				err = errors.WithMessage(err, "do redis")
				return err
			}
		}

		return nil
	}
}
