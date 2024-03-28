package test

import (
	_ "github.com/caiknife/mp3lister/config" // load module
	"github.com/caiknife/mp3lister/lib"
)

func snowflakeMachineID() uint16 {
	return lib.SnowflakeMachineID()
}
