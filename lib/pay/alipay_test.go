package pay

import (
	"testing"

	"github.com/duke-git/lancet/v2/random"
)

func TestAliPay(t *testing.T) {
	Pay.Alipay.ForEach(func(conf *alipayConf, i int) {
		alipay, err := NewAlipay(conf.AppID, conf.PrivateKey, true)
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
	})

}

func TestAlipay_Query(t *testing.T) {
	Pay.Alipay.ForEach(func(conf *alipayConf, i int) {
		alipay, err := NewAlipay(conf.AppID, conf.PrivateKey, true)
		if err != nil {
			t.Error(err)
			return
		}

		query, err := alipay.Query(conf.OrderID)
		if err != nil {
			t.Error(err)
			return
		}
		t.Log(query)
	})
}
