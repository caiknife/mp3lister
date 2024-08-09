package rediskey

import (
	"encoding"

	"github.com/caiknife/mp3lister/lib/fjson"
)

// RedisValue
//
//	@Description:	凡是存储在redis中的结构体，都需要实现这个接口
type RedisValue interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type PlayerChest struct {
	ChestProgress       int  `json:"chest_progress"`
	ActualChestProgress int  `json:"actual_chest_progress"`
	Acquired            bool `json:"acquired"`
	LegionId            int  `json:"legion_id"`
}

func (p *PlayerChest) NoLegion() bool {
	return p.LegionId == -1
}

func (p *PlayerChest) MarshalBinary() (data []byte, err error) {
	return fjson.Marshal(p)
}

func DefaultPlayerChest() *PlayerChest {
	p := &PlayerChest{
		ChestProgress:       0,
		ActualChestProgress: 0,
		Acquired:            true,
		LegionId:            -1,
	}
	return p
}

func (p *PlayerChest) UnmarshalBinary(data []byte) error {
	return fjson.Unmarshal(data, p)
}

func (p *PlayerChest) String() string {
	toString, _ := fjson.MarshalToString(p)
	return toString
}

var _ RedisValue = (*PlayerChest)(nil)