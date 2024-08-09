package main

import (
	"testing"

	_ "github.com/caiknife/mp3lister/test"
)

func Test_clearRedisKeys(t *testing.T) {
	err := clearRedisKeys()
	if err != nil {
		t.Error(err)
		return
	}
}

func Test_modifyRedis(t *testing.T) {
	err := modifyRedis()
	if err != nil {
		t.Error(err)
		return
	}
}

func Test_modifyChargeDiamondPool(t *testing.T) {
	err := modifyChargeDiamondPool()
	if err != nil {
		t.Error(err)
		return
	}
}

func Test_modifyLegionWarPlayerChest(t *testing.T) {
	err := modifyLegionWarPlayerChest()
	if err != nil {
		t.Error(err)
		return
	}
}
