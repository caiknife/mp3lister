package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func GetInputPath(path string) (string, error) {
	if path == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = cwd
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return abs, nil
}

func ColorPrintf(format string, a ...any) {
	fmt.Fprintf(color.Output, format, a...)
}

const (
	NullSeparator   = "\u0000"
	NBSPSeparator   = "\u00A0"
	ZWNBSPSeparator = "\uFEFF"
)

func CutInvisibleSeparator(input string) string {
	if strings.Contains(input, NullSeparator) {
		// split := strings.Split(input, NullSeparator)
		// input = strings.Join(split, "|")
		input = strings.ReplaceAll(input, NullSeparator, ",")
	}
	if strings.Contains(input, NBSPSeparator) {
		// split := strings.Split(input, NullSeperator)
		// input = strings.Join(split, "|")
		input = strings.ReplaceAll(input, NBSPSeparator, " ")
	}
	if strings.Contains(input, ZWNBSPSeparator) {
		// split := strings.Split(input, NullSeperator)
		// input = strings.Join(split, "|")
		input = strings.ReplaceAll(input, ZWNBSPSeparator, "")
	}

	return input
}
