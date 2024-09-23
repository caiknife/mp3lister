package lib

import (
	"testing"
)

func TestGetInputPath(t *testing.T) {
	t.Log(GetInputPath(""))
	t.Log(GetInputPath("."))
	t.Log(GetInputPath(".."))
	t.Log(GetInputPath("../.."))
}
