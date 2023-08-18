package lib

import (
	"fmt"
	"os"
	"path/filepath"

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
