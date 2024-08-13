package lib

import (
	"testing"
)

func TestSnowflakeNode(t *testing.T) {
	node, err := SnowflakeNode()
	if err != nil {
		t.Error(err)
		return
	}
	id := node.Generate()
	t.Log(id, len(id.String()))
}
