package app

import (
	"github.com/urfave/cli/v2"
)

func action(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		cli.VersionPrinter(ctx)
	}
	
	return nil
}
