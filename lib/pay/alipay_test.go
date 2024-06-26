package pay

import (
	"testing"

	"github.com/duke-git/lancet/v2/random"
)

func TestAliPay(t *testing.T) {
	alipay, err := NewAlipay(Pay.Alipay.AppID, Pay.Alipay.PrivateKey, true)
	if err != nil {
		t.Error(err)
		return
	}
	alipay.SetNotifyURL("http://test.tank.seabirdtech.com.cn:10001/notify/alipay")

	pay, err := alipay.TransactionApp("test gopay alipay", random.RandString(32), random.RandFloat(10, 100, 2))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(pay)
}

func TestAlipay_Query(t *testing.T) {
	alipay, err := NewAlipay(Pay.Alipay.AppID, Pay.Alipay.PrivateKey, true)
	if err != nil {
		t.Error(err)
		return
	}

	query, err := alipay.Query(Pay.Alipay.OrderID)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(query)
}
