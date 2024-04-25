package lib

import (
	"testing"

	"github.com/pkg/errors"
)

func inputPath() error {
	return errors.WithStack(ErrInputPathIsEmpty)
}

func outputPath() error {
	err := inputPath()
	return errors.WithStack(err)
}

func TestErrors(t *testing.T) {
	err := outputPath()
	if err != nil {
		t.Errorf("%+v\n", err)
		return
	}
}
