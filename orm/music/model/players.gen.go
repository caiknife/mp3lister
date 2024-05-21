// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const TableNamePlayer = "players"

// Player mapped from table <players>
type Player struct {
	ID        uint64         `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;not null;default:current_timestamp()" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;not null;default:current_timestamp()" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;index:deleted_at,priority:1" json:"deleted_at"`
	Name      string         `gorm:"column:name;type:varchar(200);not null;comment:姓名" json:"name"`                            // 姓名
	Phone     string         `gorm:"column:phone;type:varchar(200);not null;index:phone,priority:1;comment:电话" json:"phone"`   // 电话
	Email     string         `gorm:"column:email;type:varchar(200);not null;index:email,priority:1;comment:邮件地址" json:"email"` // 邮件地址
	Gold      int64          `gorm:"column:gold;type:bigint(20);not null;comment:金币数量" json:"gold"`                            // 金币数量
	Extra     datatypes.JSON `gorm:"column:extra;type:longtext;comment:扩展信息" json:"extra"`                                     // 扩展信息
}

// TableName Player's table name
func (*Player) TableName() string {
	return TableNamePlayer
}
