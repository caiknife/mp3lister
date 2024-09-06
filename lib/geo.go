package lib

import (
	"fmt"
	"net"

	"github.com/oschwald/geoip2-golang"
	"github.com/pkg/errors"
)

const geoFile = "GeoLite2-City.mmdb"

type city = geoip2.City

type City struct {
	*city
}

func IP2City(ipv4 string) (*City, error) {
	configFile, err := SearchConfigFile(geoFile)
	if err != nil {
		err = errors.WithMessage(err, "failed to search config geoFile")
		return nil, err
	}
	open, err := geoip2.Open(configFile)
	if err != nil {
		err = errors.WithMessage(err, fmt.Sprintf("failed to open %s", geoFile))
		return nil, err
	}
	defer open.Close()

	ip := net.ParseIP(ipv4)
	city, err := open.City(ip)
	if err != nil {
		err = errors.WithMessage(err, "find city by ip")
		return nil, err
	}

	return &City{city: city}, nil
}

func (c *City) ContinentCode() string {
	return c.Continent.Code
}

func (c *City) ContinentName() string {
	if v, ok := c.Continent.Names["zh-CN"]; ok {
		return v
	}
	return c.Continent.Names["en"]
}

func (c *City) CountryCode() string {
	return c.Country.IsoCode
}

func (c *City) CountryName() string {
	if v, ok := c.Country.Names["zh-CN"]; ok {
		return v
	}
	return c.Country.Names["en"]
}

func (c *City) Name() string {
	if v, ok := c.City.Names["zh-CN"]; ok {
		return v
	}
	return c.City.Names["en"]
}

func (c *City) TimeZone() string {
	return c.Location.TimeZone
}
