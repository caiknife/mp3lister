package config

import (
	"testing"
)

func TestPrint(t *testing.T) {
	t.Log(AuthSuccess)
	t.Log(AuthNotYet)
	t.Log(AuthFailed)
}
