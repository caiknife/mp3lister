// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameWtGamecenter = "wt_gamecenter"

// WtGamecenter 玩家和GameCenter关系表
type WtGamecenter struct {
	ID             string    `gorm:"column:id;type:varchar(255);primaryKey" json:"id"`
	BundleID       string    `gorm:"column:bundle_id;type:varchar(64);not null;index:bundle_id,priority:1;comment:BUNDLE ID" json:"bundle_id"` // BUNDLE ID
	PlayerID       int64     `gorm:"column:player_id;type:bigint(20);not null;index:player_id,priority:1;comment:玩家ID" json:"player_id"`       // 玩家ID
	GcDisplayName  string    `gorm:"column:gc_display_name;type:varchar(128);comment:GC displayName" json:"gc_display_name"`                   // GC displayName
	GcGamePlayerID string    `gorm:"column:gc_game_player_id;type:varchar(255)" json:"gc_game_player_id"`
	GcTeamPlayerID string    `gorm:"column:gc_team_player_id;type:varchar(255)" json:"gc_team_player_id"`
	GcPlayerID     string    `gorm:"column:gc_player_id;type:varchar(255)" json:"gc_player_id"`
	CreateTime     time.Time `gorm:"column:create_time;type:timestamp;not null;default:current_timestamp();comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime     time.Time `gorm:"column:update_time;type:timestamp;not null;default:current_timestamp();comment:更新时间" json:"update_time"` // 更新时间
}

// TableName WtGamecenter's table name
func (*WtGamecenter) TableName() string {
	return TableNameWtGamecenter
}
