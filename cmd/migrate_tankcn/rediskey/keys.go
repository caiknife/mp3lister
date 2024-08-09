package rediskey

import (
	"github.com/caiknife/mp3lister/lib/types"
)

const (
	keySettlePlayerRewards = "TK:SETTLE:PLAYER:REWARDS:"
	defaultLimit           = 500
)

func ReservedKeys() types.Slice[string] {
	return types.Slice[string]{
		ChargeDiamondPool(),
		FirstChargePool(),
		HighestQualityPool(),
		PlayerProficiencyExp(),
		ResetTrophyPool(),
		ShopDailyChestPool(),
		VeteranChargePool(),
		WeeklyChestPlayer(),
		ChestKey(),
		SignIn(),
		SVIP(),
		CompetitivePlayerAward(),
		ShopMissionDiamond(),
		LegionWarPlayerMedal(),
		LegionWarLegionChest(),
		LegionWarPlayerChest(),
	}
}

func LegionWarPlayerMedal() string {
	return "TK:LEGION_WAR:PLAYER_MEDAL"
}

func ShopMissionDiamond() string {
	return "TK:SHOP:MISSION:DIAMOND"
}

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

func CompetitivePlayerAward() string {
	return "TK:COMPETITIVE:PLAYER_AWARD"
}

func LegionWarLegionChest() string {
	return "TK:LEGION_WAR:LEGION_CHEST"
}

func LegionWarPlayerChest() string {
	return "TK:LEGION_WAR:PLAYER_CHEST"
}
