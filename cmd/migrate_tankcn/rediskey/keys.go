package rediskey

const (
	keySettlePlayerRewards = "TK:SETTLE:PLAYER:REWARDS:"
)

func ChargeDiamondPool() string {
	return "TK:CHARGE:DIAMOND:POOL"
}

func FirstChargePool() string {
	return "TK:FIRSTCHARGE:POOL"
}

func HighestQualityPool() string {
	return "TK:HIGHESTQUALITY:POOL"
}

func PlayerProficiencyExp() string {
	return "TK:PROFICENCYEXP:PLAYER"
}

func ResetTrophyPool() string {
	return "TK:RESETTROPHY:POOL"
}

func SettlePlayerRewards(playerID string) string {
	return keySettlePlayerRewards + playerID
}

func ShopDailyChestPool() string {
	return "TK:SHOP:DAILYCHEST:POOL"
}
