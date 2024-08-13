package test

import (
	"context"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/datetime"

	"github.com/caiknife/mp3lister/config"
)

func TestRedisHIncr(t *testing.T) {
	result, err := config.RedisDefault.HIncrBy(context.TODO(), "myhash", "111111", 1).Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
	n := time.Now()
	next := datetime.EndOfDay(n)
	b, err := config.RedisDefault.Expire(context.TODO(), "myhash", next.Sub(n)).Result()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(b)
}
