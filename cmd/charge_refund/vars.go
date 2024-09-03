package main

import (
	"github.com/duke-git/lancet/v2/strutil"

	"github.com/caiknife/mp3lister/lib"
	"github.com/caiknife/mp3lister/lib/types"
)

var ShopCfg types.Hash[string, *Product]

func init() {
	lib.InitJSONConfig(&ShopCfg, "shop.json")
}

func getProductID(completeProductID string) string {
	return strutil.AfterLast(completeProductID, ".")
}
