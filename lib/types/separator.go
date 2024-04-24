package types

import (
	"strings"
)

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
