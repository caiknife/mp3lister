package test

import (
	"net"
	"net/url"
	"testing"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/go-resty/resty/v2"

	"github.com/caiknife/mp3lister/lib/types"
)

func TestNetUtil(t *testing.T) {
	ip := netutil.GetInternalIp()
	t.Log(ip)
	t.Log(snowflakeMachineID())

	ips := netutil.GetIps()
	t.Log(ips)

	addrs := netutil.GetMacAddrs()
	t.Log(addrs)

	parseIP := net.ParseIP(ip)
	t.Log(netutil.IsInternalIP(parseIP))
}

func TestUrl(t *testing.T) {
	values := url.Values{}
	values.Add("name", "cai")
	values.Add("name", "fan")
	values.Add("age", "18")
	values.Add("age", "28")
	encode := values.Encode()
	t.Log(encode)

	client := resty.New()
	req := client.R()
	req.SetHeaders(types.Map[string]{
		"Content-Type": "application/x-www-form-urlencoded",
	}).SetBody(values.Encode())
	resp, err := req.Post("https://httpbin.org/post")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp.String())
}

func TestParseURL(t *testing.T) {
	uri := "https://httpbin.org/post"
	parse, err := url.Parse(uri)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(parse)
}
