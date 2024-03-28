package lib

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/fatih/color"
)

func GetInputPath(path string) (string, error) {
	if path == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return "", err
		}
		path = cwd
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return abs, nil
}

func ColorPrintf(format string, a ...any) {
	fmt.Fprintf(color.Output, format, a...)
}

const (
	NullSeparator   = "\u0000"
	NBSPSeparator   = "\u00A0"
	ZWNBSPSeparator = "\uFEFF"
)

func CutInvisibleSeparator(input string) string {
	if strings.Contains(input, NullSeparator) {
		input = strings.ReplaceAll(input, NullSeparator, ",")
	}
	if strings.Contains(input, NBSPSeparator) {
		input = strings.ReplaceAll(input, NBSPSeparator, " ")
	}
	if strings.Contains(input, ZWNBSPSeparator) {
		input = strings.ReplaceAll(input, ZWNBSPSeparator, "")
	}

	return input
}

func SnowflakeMachineID() uint16 {
	ip := netutil.GetInternalIp()
	parseIP := net.ParseIP(ip).To4()
	seed := uint16(parseIP[2])<<8 + uint16(parseIP[3])
	return seed
}
