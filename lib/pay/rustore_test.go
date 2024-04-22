package pay

import (
	"context"
	"testing"
	"time"

	"github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib/pay/rustore/client"
	"github.com/caiknife/mp3lister/lib/pay/rustore/payments"
)

func TestRuStoreAPI(t *testing.T) {
	c := client.New("", config.Pay.RuStore.PrivateKey, config.Pay.RuStore.CompanyID)
	err := c.Auth()
	if err != nil {
		t.Error(err)
		return
	}

	p := payments.New(c, config.Pay.RuStore.PackageName)
	ctx, cancel := context.WithTimeout(context.Background(), client.TimeOutSeconds*10*time.Second)
	defer cancel()
	info, err := p.GetPaymentInfo(ctx, config.Pay.RuStore.PurchaseToken)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(info)
}
