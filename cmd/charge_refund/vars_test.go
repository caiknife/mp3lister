package main

import (
	"testing"
)

func TestShopCfg(t *testing.T) {
	ShopCfg.ForEach(func(s string, product *Product) {
		t.Log(s, product)
	})
}

func Test_getProductID(t *testing.T) {
	t.Log(getProductID("123"))
	t.Log(getProductID("com.raven.tank.ttcn.diamonds2"))
}
