package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	if err := newApp().Run(os.Args); err != nil {
		_, _ = fmt.Fprintf(
			color.Output,
			"Run %s failed: %s\n",
			color.CyanString("%s", Name),
			color.RedString("%v", err),
		)
		os.Exit(1)
	}
}
