package fcm

import (
	"github.com/caiknife/mp3lister/lib/fjson"
)

type Check struct {
	Ai    string `json:"ai"`
	Name  string `json:"name"`
	IdNum string `json:"idNum"`
}

func (c *Check) String() string {
	toString, _ := fjson.MarshalToString(c)
	return toString
}

type Query struct {
	Ai string `json:"ai"`
}

func (q *Query) String() string {
	toString, _ := fjson.MarshalToString(q)
	return toString
}

type Behavior struct {
	No int    `json:"no"`
	Si string `json:"si"`
	Bt int    `json:"bt"`
	Ot int64  `json:"ot"`
	Ct int    `json:"ct"`
	Di string `json:"di"`
	Pi string `json:"pi"`
}

func (l *Behavior) String() string {
	toString, _ := fjson.MarshalToString(l)
	return toString
}

type Collections struct {
	Collections *[]Behavior `json:"collections"`
}

func (c *Collections) String() string {
	toString, _ := fjson.MarshalToString(c)
	return toString
}
