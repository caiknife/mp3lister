package main

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/spf13/cast"

	"github.com/caiknife/mp3lister/cmd/migrate_tankcn/rediskey"
	"github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib/types"
)

func modifyRedis() error {
	if err := clearRedisKeys(); err != nil {
		return err
	}
	if err := modifyChargeDiamondPool(); err != nil {
		return err
	}
	if err := modifyFirstChargePool(); err != nil {
		return err
	}
	if err := modifyHighestQualityPool(); err != nil {
		return err
	}
	if err := modifyPlayerProficiencyExp(); err != nil {
		return err
	}
	if err := modifyResetTrophyPool(); err != nil {
		return err
	}
	if err := modifyShopDailyChestPool(); err != nil {
		return err
	}
	if err := modifyVeteranChargePool(); err != nil {
		return err
	}
	if err := modifyWeeklyChestPlayer(); err != nil {
		return err
	}
	if err := modifyChestKey(); err != nil {
		return err
	}
	if err := modifySignIn(); err != nil {
		return err
	}
	if err := modifySVIP(); err != nil {
		return err
	}
	if err := modifyCompetitivePlayerAward(); err != nil {
		return err
	}
	if err := modifyShopMissionDiamond(); err != nil {
		return err
	}
	if err := modifyLegionWarPlayerMedal(); err != nil {
		return err
	}
	return nil
}

func clearRedisKeys() error {
	result := types.Slice[string]{}
	result, err := config.RedisDefault.Keys(context.TODO(), "*").Result()
	if err != nil {
		err = errors.WithMessage(err, "redis keys error")
		return err
	}

	result = lo.Without(result, rediskey.ReservedKeys()...)
	err = config.RedisDefault.Del(context.TODO(), result...).Err()
	if err != nil {
		err = errors.WithMessage(err, "redis del error")
		return err
	}
	return nil
}

func modifyRedisHash(key string) error {
	result := types.Map[string]{}
	result, err := config.RedisDefault.HGetAll(context.TODO(), key).Result()
	if err != nil {
		err = errors.WithMessage(err, fmt.Sprintf("%s redis hgetall error", key))
		return err
	}
	if result.IsEmpty() {
		return nil
	}

	newResult := map[string]string{}
	result.ForEach(func(key string, value string) {
		newKey := cast.ToString(cast.ToInt(key) + playerIDIncrement)
		newResult[newKey] = value
	})
	// 添加新数据
	err = config.RedisDefault.HSet(context.TODO(), key, newResult).Err()
	if err != nil {
		err = errors.WithMessage(err, fmt.Sprintf("%s redis hset error", key))
		return err
	}
	// 删除旧数据
	err = config.RedisDefault.HDel(context.TODO(), key, result.Keys()...).Err()
	if err != nil {
		err = errors.WithMessage(err, fmt.Sprintf("%s redis hdel error", key))
		return err
	}
	return nil
}

func modifyRedisSet(key string) error {
	result := types.Slice[string]{}
	result, err := config.RedisDefault.SMembers(context.TODO(), key).Result()
	if err != nil {
		err = errors.WithMessage(err, fmt.Sprintf("%s redis smembers error", key))
		return err
	}
	newResult := make(types.Slice[string], result.Len())
	result.ForEach(func(s string, i int) {
		newResult[i] = cast.ToString(cast.ToInt(s) + playerIDIncrement)
	})
	// 删除旧数据
	err = config.RedisDefault.SRem(context.TODO(), key, lo.Map[string, any](result, func(item string, index int) any {
		return item
	})...).Err()
	if err != nil {
		err = errors.WithMessage(err, fmt.Sprintf("%s redis srem error", key))
		return err
	}
	// 添加新数据
	err = config.RedisDefault.SAdd(context.TODO(), key, lo.Map[string, any](newResult, func(item string, index int) any {
		return item
	})...).Err()
	if err != nil {
		err = errors.WithMessage(err, fmt.Sprintf("%s redis sadd error", key))
		return err
	}
	return nil
}

func modifyChargeDiamondPool() error {
	err := modifyRedisHash(rediskey.ChargeDiamondPool())
	if err != nil {
		err = errors.WithMessage(err, "redis charge diamond pool error")
		return err
	}
	return nil
}

func modifyFirstChargePool() error {
	err := modifyRedisHash(rediskey.FirstChargePool())
	if err != nil {
		err = errors.WithMessage(err, "redis first charge pool error")
		return err
	}
	return nil
}

func modifyHighestQualityPool() error {
	err := modifyRedisHash(rediskey.HighestQualityPool())
	if err != nil {
		err = errors.WithMessage(err, "redis highest quality pool error")
		return err
	}
	return nil
}

func modifyPlayerProficiencyExp() error {
	err := modifyRedisHash(rediskey.PlayerProficiencyExp())
	if err != nil {
		err = errors.WithMessage(err, "redis player proficiency exp error")
		return err
	}
	return nil
}

func modifyResetTrophyPool() error {
	err := modifyRedisSet(rediskey.ResetTrophyPool())
	if err != nil {
		err = errors.WithMessage(err, "redis reset trophy pool error")
		return err
	}
	return nil
}

func modifyShopDailyChestPool() error {
	err := modifyRedisHash(rediskey.ShopDailyChestPool())
	if err != nil {
		err = errors.WithMessage(err, "redis shop daily chest pool error")
		return err
	}
	return nil
}

func modifyVeteranChargePool() error {
	err := modifyRedisHash(rediskey.VeteranChargePool())
	if err != nil {
		err = errors.WithMessage(err, "redis veteran charge pool error")
		return err
	}
	return nil
}

func modifyWeeklyChestPlayer() error {
	err := modifyRedisHash(rediskey.WeeklyChestPlayer())
	if err != nil {
		err = errors.WithMessage(err, "redis weekly chest player error")
		return err
	}
	return nil
}

func modifyChestKey() error {
	err := modifyRedisHash(rediskey.ChestKey())
	if err != nil {
		err = errors.WithMessage(err, "redis chest key error")
		return err
	}
	return nil
}

func modifySignIn() error {
	err := modifyRedisHash(rediskey.SignIn())
	if err != nil {
		err = errors.WithMessage(err, "redis sign in error")
		return err
	}
	return nil
}

func modifySVIP() error {
	err := modifyRedisSet(rediskey.SVIP())
	if err != nil {
		err = errors.WithMessage(err, "redis svip error")
		return err
	}
	return nil
}

func modifyCompetitivePlayerAward() error {
	err := modifyRedisHash(rediskey.CompetitivePlayerAward())
	if err != nil {
		err = errors.WithMessage(err, "redis competitive player award error")
		return err
	}
	return nil
}

func modifyShopMissionDiamond() error {
	err := modifyRedisHash(rediskey.ShopMissionDiamond())
	if err != nil {
		err = errors.WithMessage(err, "redis shop mission diamond error")
		return err
	}
	return nil
}

func modifyLegionWarPlayerMedal() error {
	err := modifyRedisHash(rediskey.LegionWarPlayerMedal())
	if err != nil {
		err = errors.WithMessage(err, "redis legion war medal error")
		return err
	}
	return nil
}
