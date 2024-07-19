package rediskey

import (
	"context"
	"testing"

	"github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib/types"
	_ "github.com/caiknife/mp3lister/test"
)

func TestChargeDiamondPool(t *testing.T) {
	result, err := config.RedisDefault.HGetAll(context.TODO(), ChargeDiamondPool()).Result()
	if err != nil {
		t.Error(err)
		return
	}
	types.Map[string](result).ForEach(func(key string, value string) {
		t.Log(key, value, len(value))
	})
}

func TestFirstChargePool(t *testing.T) {
	result, err := config.RedisDefault.HGetAll(context.TODO(), FirstChargePool()).Result()
	if err != nil {
		t.Error(err)
		return
	}
	types.Map[string](result).ForEach(func(key string, value string) {
		t.Log(key, value, len(value))
	})
}

func TestHighestQualityPool(t *testing.T) {
	result, err := config.RedisDefault.HGetAll(context.TODO(), HighestQualityPool()).Result()
	if err != nil {
		t.Error(err)
		return
	}
	types.Map[string](result).ForEach(func(key string, value string) {
		t.Log(key, value, len(value))
	})
}

func TestPlayerProficiencyExp(t *testing.T) {
	result, err := config.RedisDefault.HGetAll(context.TODO(), PlayerProficiencyExp()).Result()
	if err != nil {
		t.Error(err)
		return
	}
	types.Map[string](result).ForEach(func(key string, value string) {
		t.Log(key, value, len(value))
	})
}

func TestResetTrophyPool(t *testing.T) {
	result, err := config.RedisDefault.SMembers(context.TODO(), ResetTrophyPool()).Result()
	if err != nil {
		t.Error(err)
		return
	}
	types.Slice[string](result).ForEach(func(s string, i int) {
		t.Log(i, s, len(s))
	})
}

func TestSettlePlayerRewards(t *testing.T) {
	result, err := config.RedisDefault.Keys(context.TODO(), keySettlePlayerRewards+"*").Result()
	if err != nil {
		t.Error(err)
		return
	}
	types.Slice[string](result).ForEach(func(s string, i int) {
		t.Log(s, splitPlayerIDFromKey(s), SettlePlayerRewards(splitPlayerIDFromKey(s)))
	})
}

func TestShopDailyChestPool(t *testing.T) {
	result, err := config.RedisDefault.HGetAll(context.TODO(), ShopDailyChestPool()).Result()
	if err != nil {
		t.Error(err)
		return
	}
	types.Map[string](result).ForEach(func(key string, value string) {
		t.Log(key, value, len(value))
	})
}
