package pay

import (
	"testing"

	"github.com/caiknife/mp3lister/lib"
	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
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
	KeyID                string `json:"key_id"`
	CompanyID            string `json:"company_id"`
	PrivateKey           string `json:"private_key"`
	PackageName          string `json:"package_name"`
	PurchaseToken        string `json:"purchase_token"`
	SandboxPurchaseToken string `json:"sandbox_purchase_token"`
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

type appleConf struct {
	IssID        string `json:"iss_id"`
	BundleID     string `json:"bundle_id"`
	KeyID        string `json:"key_id"`
	PrivateKey   string `json:"private_key"`
	IsProduction bool   `json:"is_production"`
	TransID      string `json:"trans_id"`
	OrderID      string `json:"order_id"`
}

func (a *appleConf) String() string {
	toString, _ := fjson.MarshalToString(a)
	return toString
}

type pay struct {
	Alipay  types.Slice[*alipayConf]  `json:"alipay"`
	Wxpay   types.Slice[*wxpayConf]   `json:"wxpay"`
	RuStore types.Slice[*rustoreConf] `json:"rustore"`
	Quick   types.Slice[*quickConf]   `json:"quick"`
	Apple   types.Slice[*appleConf]   `json:"apple"`
}

func TestConf(t *testing.T) {
	t.Log(Pay.Quick)
}
