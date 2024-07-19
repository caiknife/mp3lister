package rediskey

import (
	"strings"
)

func splitPlayerIDFromKey(key string) string {
	return strings.TrimPrefix(key, keySettlePlayerRewards)
}
