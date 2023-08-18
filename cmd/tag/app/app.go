package app

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const (
	Name    = "mp3tagger"
	Version = "v0.0.1"
)

func init() {
	cli.VersionPrinter = func(c *cli.Context) {
		blue := color.New(color.FgBlue)
		cyan := color.New(color.FgCyan)
		fmt.Fprintf(
			color.Output,
			"\n%s: version %s, Just a mp3 files tagger.\n\n",
			cyan.Sprintf(Name),
			blue.Sprintf(c.App.Version),
		)
	}
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
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"V"},
				Usage:   "verbose mode",
			},
			&cli.StringFlag{
				Name: "input",
				// Value:   "",
				Aliases: []string{"i"},
				Usage:   "input path for mp3 lister, default value is current dir",
			},
		},
		Action: action,
	}

	return app
}
