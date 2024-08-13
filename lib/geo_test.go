package lib

import (
	"testing"

	"github.com/caiknife/mp3lister/lib/types"
)

var testIPs = types.Slice[string]{
	"175.0.225.74",
	"43.139.149.180",
}

func TestGeoIP(t *testing.T) {
	testIPs.ForEach(func(s string, i int) {
		city, err := IP2City(s)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log("城市名称", city.Name())
		t.Log("国家名称", city.CountryCode(), city.CountryName())
		t.Log("大陆名称", city.ContinentCode(), city.ContinentName())
	})
}
