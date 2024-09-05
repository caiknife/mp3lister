package lib

import (
	"fmt"
	"net"

	"github.com/oschwald/geoip2-golang"
	"github.com/pkg/errors"
)

const geoFile = "GeoLite2-City.mmdb"

type City struct {
	city *geoip2.City
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
	return c.city.Continent.Code
}

func (c *City) ContinentName() string {
	if v, ok := c.city.Continent.Names["zh-CN"]; ok {
		return v
	}
	return c.city.Continent.Names["en"]
}

func (c *City) CountryCode() string {
	return c.city.Country.IsoCode
}

func (c *City) CountryName() string {
	if v, ok := c.city.Country.Names["zh-CN"]; ok {
		return v
	}
	return c.city.Country.Names["en"]
}

func (c *City) Name() string {
	if v, ok := c.city.City.Names["zh-CN"]; ok {
		return v
	}
	return c.city.City.Names["en"]
}

func (c *City) TimeZone() string {
	return c.city.Location.TimeZone
}
