package main

import (
	"testing"

	_ "github.com/caiknife/mp3lister/test"
)

func Test_loadDB(t *testing.T) {
	orders, err := loadDB()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(orders.Len())
}

func Test_getPriceByProductID(t *testing.T) {
	price := getPriceByProductID("diamonds0")
	t.Log(price)
}

func Test_chargeRefund(t *testing.T) {
	orders, err := loadDB()
	if err != nil {
		t.Error(err)
		return
	}
	playerTotalCharges, err := transformOrder(orders)
	if err != nil {
		t.Error(err)
		return
	}

	refund, err := chargeRefund(playerTotalCharges)
	if err != nil {
		t.Error(err)
		return
	}
	refund.ForEach(func(s string, refund *ChargeRefund) {
		t.Log(s, refund)
	})
}

func Test_transformOrder(t *testing.T) {
	orders, err := loadDB()
	if err != nil {
		t.Error(err)
		return
	}
	playerTotalCharges, err := transformOrder(orders)
	if err != nil {
		return
	}
	playerTotalCharges.ForEach(func(i int64, f float64) {
		t.Log(i, f)
	})
}
