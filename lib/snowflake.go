package lib

import (
	"net"

	"github.com/bwmarrin/snowflake"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/pkg/errors"
)

const snowFlakeNodeMax = 1024

func SnowflakeMachineID() uint16 {
	ip := netutil.GetInternalIp()
	parseIP := net.ParseIP(ip).To4()
	seed := uint16(parseIP[2])<<8 + uint16(parseIP[3])
	return seed % snowFlakeNodeMax
}

func SnowflakeNode() (*snowflake.Node, error) {
	seed := SnowflakeMachineID()
	node, err := snowflake.NewNode(int64(seed))
	if err != nil {
		err = errors.WithMessage(err, "snowflake.NewNode failed")
		return nil, err
	}
	return node, nil
}
