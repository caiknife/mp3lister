package test

import (
	"net"

	"github.com/duke-git/lancet/v2/netutil"

	_ "github.com/caiknife/mp3lister/config" // load module
)

func snowflakeMachineID() uint16 {
	ip := netutil.GetInternalIp()
	parseIP := net.ParseIP(ip).To4()
	seed := uint16(parseIP[2])<<8 + uint16(parseIP[3])
	return seed
}
