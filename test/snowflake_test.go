package test

import (
	"net"
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/duke-git/lancet/v2/netutil"
)

func TestSnowflake(t *testing.T) {
	seed := snowflakeMachineID()
	t.Log(seed)
	node, err := snowflake.NewNode(int64(seed))
	if err != nil {
		t.Error(err)
		return
	}

	for range 10 {
		id := node.Generate()
		t.Log(id.Int64(), id.String())
	}
}

func TestInternalIP(t *testing.T) {
	ip := netutil.GetInternalIp()
	t.Log(ip)
	parsedIP := net.ParseIP(ip).To4()
	t.Log(parsedIP[0], parsedIP[1], parsedIP[2], parsedIP[3])

	b := int(parsedIP[0])<<24 + int(parsedIP[1])<<16 + int(parsedIP[2])<<8 + int(parsedIP[3])
	t.Log(b)
	a := int(parsedIP[2])<<8 + int(parsedIP[3])
	t.Log(a)
}
