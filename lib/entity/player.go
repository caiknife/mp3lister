package entity

import (
	"github.com/brianvoe/gofakeit/v6"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/orm/music/model"
)

type Extra struct {
	Address    gofakeit.AddressInfo    `json:"address"`
	CreditCard gofakeit.CreditCardInfo `json:"credit_card"`
}

type Player struct {
	*model.Player
	Extra Extra `json:"extra"`
}

func NewPlayer(player *model.Player) *Player {
	p := &Player{}
	p.Scan(player)
	return p
}

func (p *Player) Scan(player *model.Player) {
	p.Player = player
	_ = fjson.Unmarshal(player.Extra, &p.Extra)
}

func (p *Player) Value() *model.Player {
	p.Player.Extra, _ = fjson.Marshal(p.Extra)
	return p.Player
}
