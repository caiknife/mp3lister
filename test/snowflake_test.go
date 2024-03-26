package test

import (
	"testing"

	"github.com/bwmarrin/snowflake"
)

func TestSnowflake(t *testing.T) {
	seed := snowflakeMachineID()
	t.Log(seed)
	node, err := snowflake.NewNode(int64(seed))
	if err != nil {
		t.Error(err)
		return
	}
	id := node.Generate()
	t.Log(id.Int64())
}
