package app

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const (
	Name    = "mp3lister"
	Version = "v0.0.1"
)

func init() {
	cli.VersionPrinter = func(c *cli.Context) {
		blue := color.New(color.FgBlue)
		cyan := color.New(color.FgCyan)
		fmt.Fprintf(
			color.Output,
			"\n%s: version %s, Just a mp3 files lister.\n\n",
			cyan.Sprintf(Name),
			blue.Sprintf(c.App.Version),
		)
	}
}

func ColorPrintf(format string, a ...any) {
	fmt.Fprintf(color.Output, format, a...)
}

func New() *cli.App {
	app := &cli.App{
		Name:    Name,
		Usage:   "Just a mp3 files lister.",
		Version: Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "debug mode",
			},
			&cli.StringFlag{
				Name:    "savetodb",
				Aliases: []string{"s"},
				Usage:   "all mp3 files info will be save to a mysql database, please fill a dsn",
			},
			&cli.StringFlag{
				Name: "input",
				// Value:   "",
				Aliases: []string{"i"},
				Usage:   "input path for mp3 lister, default value is current dir",
			},
			&cli.StringFlag{
				Name: "output",
				// Value:   "",
				Aliases: []string{"o"},
				Usage:   "output file name for mp3 lister, just the name without extension. By default output file is mp3lister_yyyy_mm_dd_hh_ii_ss.csv",
			},
		},
		Action: action,
	}

	return app
}
