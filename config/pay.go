package config

var (
	Pay *pay
)

func init() {
	initPayTest()
}

func initPayTest() {
	Pay = &pay{}
	InitJSONConfig(Pay, "test.json")
}

type wxpay struct {
	AppID      string `json:"app_id"`
	MchID      string `json:"mch_id"`
	SerialNo   string `json:"serial_no"`
	APIKey     string `json:"api_key"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	OrderID    string `json:"order_id"`
}

type rustore struct {
	KeyID         string `json:"key_id"`
	CompanyID     string `json:"company_id"`
	PrivateKey    string `json:"private_key"`
	PackageName   string `json:"package_name"`
	PurchaseToken string `json:"purchase_token"`
}

type alipay struct {
	AppID      string `json:"app_id"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	OrderID    string `json:"order_id"`
}

type pay struct {
	Alipay  alipay  `json:"alipay"`
	Wxpay   wxpay   `json:"wxpay"`
	RuStore rustore `json:"rustore"`
}
