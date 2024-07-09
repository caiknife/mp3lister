package pay

import (
	"context"

	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/pkg/errors"
)

type Wxpay struct {
	AppID      string `json:"app_id"`
	MchID      string `json:"mch_id"`
	SerialNo   string `json:"serial_no"`
	APIv3Key   string `json:"api_v3_key"`
	PrivateKey string `json:"private_key"`
	NotifyURL  string `json:"notify_url"`

	client *wechat.ClientV3
}

func NewWxpay(appID, mchID, serialNo, apiv3Key, privateKey string) (p *Wxpay, err error) {
	p = &Wxpay{
		AppID:      appID,
		MchID:      mchID,
		SerialNo:   serialNo,
		APIv3Key:   apiv3Key,
		PrivateKey: privateKey,
	}

	p.client, err = wechat.NewClientV3(mchID, serialNo, apiv3Key, privateKey)
	if err != nil {
		return nil, errors.WithMessage(err, "new wxpay client error")
	}
	return p, nil
}

func (w *Wxpay) SetNotifyURL(notifyURL string) {
	w.NotifyURL = notifyURL
}

func (w *Wxpay) TransactionApp(description string, orderID string, amount float64) (*wechat.AppPayParams, error) {
	bm := make(gopay.BodyMap)
	bm.Set("appid", w.AppID).
		Set("description", description).
		Set("out_trade_no", orderID).
		Set("notify_url", w.NotifyURL).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("total", int64(amount*100)).Set("currency", "CNY")
		})

	timeout, cancelFunc := context.WithTimeout(context.TODO(), defaultTimeOut)
	defer cancelFunc()

	app, err := w.client.V3TransactionApp(timeout, bm)
	if err != nil {
		return nil, errors.WithMessage(err, "wxpay transaction app error")
	}
	if app.Code != wechat.Success {
		return nil, errors.New("wxpay code error")
	}

	ofApp, err := w.client.PaySignOfApp(w.AppID, app.Response.PrepayId)
	if err != nil {
		return nil, errors.WithMessage(err, "wxpay sign of app error")
	}
	return ofApp, nil
}

func (w *Wxpay) Query(orderID string) (*wechat.QueryOrderRsp, error) {
	timeout, cancelFunc := context.WithTimeout(context.TODO(), defaultTimeOut)
	defer cancelFunc()

	order, err := w.client.V3TransactionQueryOrder(timeout, wechat.OutTradeNo, orderID)
	if err != nil {
		return nil, errors.WithMessage(err, "wxpay query error")
	}
	return order, nil
}
