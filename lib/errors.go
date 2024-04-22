package lib

import (
	"github.com/pkg/errors"
)

var (
	ErrInputIsNotDir = errors.New("input is not a dir")
)

var (
	ErrInputPathIsEmpty  = errors.New("input path is empty")
	ErrOutputNameIsEmpty = errors.New("output name is empty")
	ErrOutputExtIsEmpty  = errors.New("output ext is empty")
	ErrDataIsEmpty       = errors.New("data is empty")
)
