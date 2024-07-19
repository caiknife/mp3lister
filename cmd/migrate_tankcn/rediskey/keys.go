package rediskey

const (
	keySettlePlayerRewards = "TK:SETTLE:PLAYER:REWARDS:"
	defaultLimit           = 500
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

func VeteranChargePool() string {
	return "TK:VETERANCHARGE:POOL"
}

func WeeklyChestPlayer() string {
	return "TK:WEEKLY:CHESTPLAYER"
}

func ChestKey() string {
	return "TK:CHEST_KEY"
}

func SignIn() string {
	return "TK:SIGNIN"
}

func SVIP() string {
	return "TK:SVIP"
}
