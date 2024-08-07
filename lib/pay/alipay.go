package pay

import (
	"context"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
)

type Alipay struct {
	AppID        string `json:"app_id"`
	PrivateKey   string `json:"private_key"`
	IsProduction bool   `json:"is_production"`
	NotifyURL    string `json:"notify_url"`

	client *alipay.Client
}

func NewAlipay(appID, privateKey string, isProduction bool) (p *Alipay, err error) {
	p = &Alipay{
		AppID:        appID,
		PrivateKey:   privateKey,
		IsProduction: isProduction,
	}
	p.client, err = alipay.NewClient(appID, privateKey, isProduction)
	if err != nil {
		return nil, errors.WithMessage(err, "new alipay client error")
	}
	return p, nil
}

func (a *Alipay) SetNotifyURL(notifyURL string) {
	a.NotifyURL = notifyURL
	a.client.SetNotifyUrl(notifyURL)
}

func (a *Alipay) TransactionApp(description, orderID string, amount float64) (string, error) {
	bm := make(gopay.BodyMap)
	bm.Set("subject", description).
		Set("out_trade_no", orderID).
		Set("total_amount", cast.ToString(amount))
	timeout, cancelFunc := context.WithTimeout(context.TODO(), defaultTimeOut)
	defer cancelFunc()
	pay, err := a.client.TradeAppPay(timeout, bm)
	if err != nil {
		return "", errors.WithMessage(err, "alipay transaction app error")
	}
	return pay, nil
}

func (a *Alipay) Query(orderID string) (q *alipay.TradeQueryResponse, err error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", orderID)
	timeout, cancelFunc := context.WithTimeout(context.TODO(), defaultTimeOut)
	defer cancelFunc()
	q, err = a.client.TradeQuery(timeout, bm)
	if err != nil {
		return nil, errors.WithMessage(err, "alipay query error")
	}
	return q, nil
}
