package pay

import (
	"github.com/caiknife/mp3lister/lib"
)

var (
	Pay *pay
)

func init() {
	initPayTest()
}

func initPayTest() {
	Pay = &pay{}
	lib.InitJSONConfig(Pay, "test.json")
}

type wxpayConf struct {
	AppID      string `json:"app_id"`
	MchID      string `json:"mch_id"`
	SerialNo   string `json:"serial_no"`
	APIKey     string `json:"api_key"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	OrderID    string `json:"order_id"`
}

type rustoreConf struct {
	KeyID         string `json:"key_id"`
	CompanyID     string `json:"company_id"`
	PrivateKey    string `json:"private_key"`
	PackageName   string `json:"package_name"`
	PurchaseToken string `json:"purchase_token"`
}

type alipayConf struct {
	AppID      string `json:"app_id"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	OrderID    string `json:"order_id"`
}

type pay struct {
	Alipay  alipayConf  `json:"alipay"`
	Wxpay   wxpayConf   `json:"wxpay"`
	RuStore rustoreConf `json:"rustore"`
}
