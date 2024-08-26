package pay

import (
	"testing"

	"github.com/caiknife/mp3lister/lib"
	"github.com/caiknife/mp3lister/lib/fjson"
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

func (w *wxpayConf) String() string {
	toString, _ := fjson.MarshalToString(w)
	return toString
}

type rustoreConf struct {
	KeyID         string `json:"key_id"`
	CompanyID     string `json:"company_id"`
	PrivateKey    string `json:"private_key"`
	PackageName   string `json:"package_name"`
	PurchaseToken string `json:"purchase_token"`
}

func (r *rustoreConf) String() string {
	toString, _ := fjson.MarshalToString(r)
	return toString
}

type alipayConf struct {
	AppID      string `json:"app_id"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	OrderID    string `json:"order_id"`
}

func (a *alipayConf) String() string {
	toString, _ := fjson.MarshalToString(a)
	return toString
}

type quickConf struct {
	MD5Key      string `json:"md5_key"`
	CallbackKey string `json:"callback_key"`
	ProductKey  string `json:"product_key"`
	ProductCode string `json:"product_code"`
}

func (q *quickConf) String() string {
	toString, _ := fjson.MarshalToString(q)
	return toString
}

type pay struct {
	Alipay  *alipayConf  `json:"alipay"`
	Wxpay   *wxpayConf   `json:"wxpay"`
	RuStore *rustoreConf `json:"rustore"`
	Quick   *quickConf   `json:"quick"`
}

func TestConf(t *testing.T) {
	t.Log(Pay.Quick)
}
