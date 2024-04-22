package lib

import (
	"fmt"
	"net"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/fatih/color"
)

func ColorPrintf(format string, a ...any) {
	fmt.Fprintf(color.Output, format, a...)
}

func SnowflakeMachineID() uint16 {
	ip := netutil.GetInternalIp()
	parseIP := net.ParseIP(ip).To4()
	seed := uint16(parseIP[2])<<8 + uint16(parseIP[3])
	return seed
}
