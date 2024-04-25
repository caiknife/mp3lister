package lib

import (
	"github.com/caiknife/mp3lister/lib/types"
)

const (
	ErrInputIsNotDir     = types.Error("input is not a dir")
	ErrInputPathIsEmpty  = types.Error("input path is empty")
	ErrOutputNameIsEmpty = types.Error("output name is empty")
	ErrOutputExtIsEmpty  = types.Error("output ext is empty")
	ErrDataIsEmpty       = types.Error("data is empty")
)
