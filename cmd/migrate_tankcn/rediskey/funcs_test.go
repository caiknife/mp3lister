package rediskey

import (
	"testing"

	"github.com/caiknife/mp3lister/config"
	_ "github.com/caiknife/mp3lister/test"
)

func TestChargeDiamondPool(t *testing.T) {
	t.Log(config.RedisDefault)
}
