package lib

import (
	"path/filepath"
)

func GetInputPath(path string) (string, error) {
	if path == "" {
		path = "."
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return abs, nil
}
