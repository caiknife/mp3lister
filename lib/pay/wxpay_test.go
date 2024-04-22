package pay

import (
	"testing"

	"github.com/duke-git/lancet/v2/random"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/stretchr/testify/assert"

	"github.com/caiknife/mp3lister/config"
)

func TestPay(t *testing.T) {
	t.Log(config.Pay.Wxpay.PrivateKey)
	t.Log(config.Pay.Wxpay.PublicKey)
	t.Log(config.Pay.RuStore)
}

func TestWxPay(t *testing.T) {
	wxpay, err := NewWxpay(config.Pay.Wxpay.AppID, config.Pay.Wxpay.MchID, config.Pay.Wxpay.SerialNo,
		config.Pay.Wxpay.APIKey, config.Pay.Wxpay.PrivateKey)
	if err != nil {
		t.Error(err)
		return
	}
	wxpay.SetNotifyURL("http://test.tank.seabirdtech.com.cn:10001/notify/wxpay")
	app, err := wxpay.TransactionApp("test gopay wxpay", random.RandString(32), random.RandFloat(10, 100, 2))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(app)
}

func TestWxpay_Query(t *testing.T) {
	wxpay, err := NewWxpay(config.Pay.Wxpay.AppID, config.Pay.Wxpay.MchID, config.Pay.Wxpay.SerialNo,
		config.Pay.Wxpay.APIKey, config.Pay.Wxpay.PrivateKey)
	if err != nil {
		t.Error(err)
		return
	}
	query, err := wxpay.Query(config.Pay.Wxpay.OrderID)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(query.Code, query.Error)
	t.Log(query.SignInfo)
	t.Log(query.Response)
	assert.Equal(t, wechat.TradeStateSuccess, query.Response.TradeState)
}
