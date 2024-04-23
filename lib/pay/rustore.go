package pay

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/caiknife/mp3lister/config"
	"github.com/caiknife/mp3lister/lib/pay/rustore/client"
	"github.com/caiknife/mp3lister/lib/pay/rustore/payments"
)

type RuStore struct {
	KeyID       string `json:"key_id"`
	CompanyID   string `json:"company_id"`
	PrivateKey  string `json:"private_key"`
	PackageName string `json:"package_name"`

	client  *client.Client
	payment *payments.Payment
}

func NewRuStore(keyID, companyID, privateKey, packageName string) *RuStore {
	p := &RuStore{
		KeyID:       keyID,
		CompanyID:   companyID,
		PrivateKey:  privateKey,
		PackageName: packageName,
	}
	p.client = client.New(keyID, privateKey, companyID)
	p.payment = payments.New(p.client, packageName)
	return p
}

func (r *RuStore) GetPurchaseInfo(purchaseToken string) (p payments.GetTokenPaymentResponse, err error) {
	err = r.client.Auth()
	if err != nil {
		return p, errors.WithMessage(err, "rustore auth error")
	}
	ctx, cancel := context.WithTimeout(context.Background(), client.TimeOutSeconds*10*time.Second)
	defer cancel()
	p, err = r.payment.GetPaymentInfo(ctx, config.Pay.RuStore.PurchaseToken)
	if err != nil {
		return p, errors.WithMessage(err, "rustore get purchase info error")
	}
	return p, nil
}
