package test

import (
	"testing"

	"github.com/caiknife/mp3lister/lib/fjson"
)

type person struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (p *person) String() string {
	toString, err := fjson.MarshalToString(p)
	if err != nil {
		return ""
	}
	return toString
}

type driver struct {
	person
	DriveLicence string `json:"drive_licence"`
}

func (d *driver) String() string {
	toString, err := fjson.MarshalToString(d)
	if err != nil {
		return ""
	}
	return toString
}

func TestPerson(t *testing.T) {
	p := &person{
		Name:  "1",
		Phone: "2",
		Email: "3",
	}
	t.Log(p)
	t.Log((*p).String())
	d := &driver{
		person:       *p,
		DriveLicence: "4",
	}
	t.Log(d)
	t.Log((*d).String())
}
