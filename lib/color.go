package lib

import (
	"fmt"

	"github.com/fatih/color"
)

func ColorPrintf(format string, a ...any) {
	_, _ = fmt.Fprintf(color.Output, format, a...)
}
