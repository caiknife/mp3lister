package harmonyos

import (
	"github.com/caiknife/mp3lister/lib/fjson"
)

type Response struct {
	ResponseCode     string `json:"responseCode"`
	ResponseMessage  string `json:"responseMessage"`
	JWSPurchaseOrder string `json:"jwsPurchaseOrder"`
}

func (r *Response) String() string {
	toString, _ := fjson.MarshalToString(r)
	return toString
}

func (r *Response) ResponseOK() bool {
	return r.ResponseCode == "0"
}

type Header struct {
	X5C []string `json:"x5c"`
	Alg string   `json:"alg"`
	Typ string   `json:"typ"`
}

func (h *Header) String() string {
	toString, _ := fjson.MarshalToString(h)
	return toString
}

type Payload struct {
	PurchaseOrderId  string `json:"purchaseOrderId"`
	PurchaseToken    string `json:"purchaseToken"`
	ApplicationId    string `json:"applicationId"`
	ProductId        string `json:"productId"`
	PurchaseTime     int64  `json:"purchaseTime"`
	ProductType      string `json:"productType"`
	DeveloperPayload string `json:"developerPayload"`
	SignedTime       int64  `json:"signedTime"`
	Environment      string `json:"environment"`
	CountryCode      string `json:"countryCode"`
	Price            int    `json:"price"`
	Currency         string `json:"currency"`
	FinishStatus     string `json:"finishStatus"`
	NeedFinish       bool   `json:"needFinish"`
}

func (p *Payload) String() string {
	toString, _ := fjson.MarshalToString(p)
	return toString
}
