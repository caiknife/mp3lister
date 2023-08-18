package app

import (
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"

	"github.com/caiknife/mp3lister/lib"
)

func action(ctx *cli.Context) error {
	if ctx.Bool("debug") {
		cli.VersionPrinter(ctx)
	}

	path, err := lib.GetInputPath(ctx.String("input"))
	if err != nil {
		return err
	}

	lib.ColorPrintf("working in %s\n", color.CyanString("%s", path))

	return nil
}
