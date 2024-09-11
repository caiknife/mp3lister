package main

import (
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

const (
	defaultDiamonds = 400
)

var _ types.RedisValue = (*ChargeRefund)(nil)

type ChargeRefund struct {
	PlayerID     int64   `json:"player_id"`
	GameCenterID string  `json:"game_center_id"`
	TotalCharge  float64 `json:"total_charge"`
	Diamonds     int64   `json:"diamonds"`
	Acquired     bool    `json:"acquired"`
}

func (c *ChargeRefund) MarshalBinary() (data []byte, err error) {
	return fjson.Marshal(c)
}

func (c *ChargeRefund) UnmarshalBinary(data []byte) error {
	return fjson.Unmarshal(data, c)
}

func (c *ChargeRefund) CalcDiamonds() {
	c.Diamonds = int64(c.TotalCharge * defaultDiamonds)
}

func (c *ChargeRefund) String() string {
	toString, _ := fjson.MarshalToString(c)
	return toString
}

type PlayerOrder struct {
	PlayerID  int64   `json:"player_id"`
	BundleID  string  `json:"bundle_id"`
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
}

func (p *PlayerOrder) String() string {
	toString, _ := fjson.MarshalToString(p)
	return toString
}

type Product struct {
	Price             float64 `json:"price"`
	Diamond           int     `json:"diamond"`
	ExtraDiamond      int     `json:"extra_diamond"`
	FirstExtraDiamond int     `json:"first_extra_diamond"`
	Name              string  `json:"name"`
	DisplayName       string  `json:"displayname"`
	Desc              string  `json:"desc"`
	SortIdx           int     `json:"sort_idx"`
}

func (p *Product) String() string {
	toString, _ := fjson.MarshalToString(p)
	return toString
}
