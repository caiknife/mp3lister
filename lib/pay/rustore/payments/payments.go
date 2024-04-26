package payments

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cast"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/pay/rustore/client"
)

type InvoiceStatus string

func (i InvoiceStatus) Success() bool {
	return i == Confirmed || i == Paid
}

const (
	Created   InvoiceStatus = "created"
	Executed                = "executed"
	Cancelled               = "cancelled"
	Paid                    = "paid"
	Confirmed               = "confirmed"
	Reversed                = "reversed"
	Refunded                = "refunded"
)

const (
	PublicAPIURL             = "https://public-api.rustore.ru/public/"
	GetSandboxPaymentInfoURL = PublicAPIURL + "sandbox/purchase/%s"
	GetPaymentInfoURL        = PublicAPIURL + "purchase/%s"
	GetSubscriptionInfoURL   = PublicAPIURL + "subscription/%s"
	GetSubscriptionInfoV2URL = PublicAPIURL + "glike/subscription/%s/%s/%s"
	GetSubscriptionInfoV3URL = PublicAPIURL + "v3/subscription/%s/%s/%s"
	GetSubscriptionStateURL  = PublicAPIURL + "subscription/%s/state"
	SubmitSubscriptionURL    = PublicAPIURL + "glike/subscription/%s/%s/%s:acknowledge"
)

type Payment struct {
	client      *client.Client
	packageName string
}

func New(c *client.Client, packageName string) *Payment {
	return &Payment{
		client:      c,
		packageName: packageName,
	}
}

type GetTokenPaymentResponse struct {
	Code      string       `json:"code"`
	Message   string       `json:"message"`
	Body      TokenPayment `json:"body"`
	Timestamp time.Time    `json:"timestamp"`
}

func (t GetTokenPaymentResponse) String() string {
	toString, err := fjson.MarshalToString(t)
	if err != nil {
		return ""
	}
	return toString
}

type TokenPayment struct {
	InvoiceID       string          `json:"invoice_id"`
	InvoiceDate     string          `json:"invoice_date"`
	InvoiceStatus   InvoiceStatus   `json:"invoice_status"`
	ApplicationCode string          `json:"application_code"`
	ApplicationName string          `json:"application_name"`
	OwnerCode       string          `json:"owner_code"`
	OwnerName       string          `json:"owner_name"`
	PaymentInfo     PaymentInfo     `json:"payment_info"`
	PaymentMethods  *PaymentMethods `json:"payment_methods"`
	Error           Error           `json:"error"`
	Invoice         Invoice         `json:"invoice"`
	Image           string          `json:"image"`
}

type PaymentMethods struct {
	UserMessage *string   `json:"user_message"`
	Methods     []Methods `json:"methods"`
}

type Methods struct {
	Method string `json:"method"`
	Action string `json:"action"`
}

type Error struct {
	UserMessage      string `json:"user_message"`
	ErrorDescription string `json:"error_description"`
	ErrorCode        int    `json:"error_code"`
}

type Invoice struct {
	DeliveryInfo  DeliveryInfo     `json:"delivery_info"`
	InvoiceParams *[]InvoiceParams `json:"invoice_params"`
	Purchaser     Purchaser        `json:"purchaser"`
	Order         Order            `json:"order"`
}

type Purchaser struct {
	Email   string  `json:"email"`
	Phone   *string `json:"phone"`
	Contact *string `json:"contact"`
}

type InvoiceParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Order struct {
	OrderID        string        `json:"order_id"`
	OrderNumber    string        `json:"order_number"`
	OrderDate      string        `json:"order_date"`
	ServiceID      string        `json:"service_id"`
	ExpirationDate string        `json:"expiration_date"`
	TaxSystem      int           `json:"tax_system"`
	TradeName      *string       `json:"trade_name"`
	VisualName     string        `json:"visual_name"`
	OrgName        string        `json:"org_name"`
	OrgInn         string        `json:"org_inn"`
	VisualAmount   string        `json:"visual_amount"`
	OrderBundle    []OrderBundle `json:"order_bundle"`
	Amount         int           `json:"amount"`
	Currency       string        `json:"currency"`
	Purpose        string        `json:"purpose"`
	Description    string        `json:"description"`
	Language       string        `json:"language"`
}

type OrderBundle struct {
	PositionID    int          `json:"position_id"`
	ItemParams    []ItemParams `json:"item_params"`
	ItemAmount    int          `json:"item_amount"`
	ItemCode      string       `json:"item_code"`
	ItemPrice     int          `json:"item_price"`
	DiscountType  *string      `json:"discount_type"`
	DiscountValue *float64     `json:"discount_value"`
	InterestType  *string      `json:"interest_type"`
	InterestValue *float64     `json:"interest_value"`
	TaxType       int          `json:"tax_type"`
	TaxSum        *int         `json:"tax_sum"`
	Name          string       `json:"name"`
	Quantity      Quantity     `json:"quantity"`
	Currency      string       `json:"currency"`
	Image         string       `json:"image"`
}

type ItemParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Quantity struct {
	Value   int    `json:"value"`
	Measure string `json:"measure"`
}

type DeliveryInfo struct {
	DeliveryType *string `json:"delivery_type"`
	Address      Address `json:"address"`
	Description  *string `json:"description"`
}

type Address struct {
	Country *string `json:"country"`
	City    *string `json:"city"`
	Address *string `json:"address"`
}

type PaymentInfo struct {
	PaymentDate        string         `json:"payment_date"`
	PaymentID          string         `json:"payment_id"`
	PaymentParams      *PaymentParams `json:"payment_params"`
	LoyaltyInfo        *LoyaltyInfo   `json:"loyalty_info"`
	CardID             string         `json:"card_id"`
	PaysysCode         string         `json:"paysys_code"`
	MaskedPan          string         `json:"masked_pan"`
	ExpiryDate         string         `json:"expiry_date"`
	PaymentSystem      string         `json:"payment_system"`
	PaymentSystemImage string         `json:"payment_system_image"`
	PaysysImage        *string        `json:"paysys_image"`
	PaymentWay         string         `json:"payment_way"`
	PaymentWayCode     string         `json:"payment_way_code"`
	PaymentWayLogo     string         `json:"payment_way_logo"`
	BankInfo           BankInfo       `json:"bank_info"`
	DeviceInfo         *DeviceInfo    `json:"device_info"`
	Name               string         `json:"name"`
	Cardholder         string         `json:"cardholder"`
	Image              string         `json:"image"`
	Paysys             string         `json:"paysys"`
}

type LoyaltyInfo struct {
	ServiceCode  string `json:"service_code"`
	ServiceName  string `json:"service_name"`
	ChangeRate   int    `json:"change_rate"`
	PaymentBonus int    `json:"payment_bonus"`
	AwardBonus   int    `json:"award_bonus"`
	Image        string `json:"image"`
}

type DeviceInfo struct {
	DevicePlatformType    string `json:"device_platform_type"`
	DevicePlatformVersion string `json:"device_platform_version"`
	DeviceModel           string `json:"device_model"`
	DeviceManufacturer    string `json:"device_manufacturer"`
	DeviceId              string `json:"device_id"`
	Surface               string `json:"surface"`
	SurfaceVersion        string `json:"surface_version"`
}

type PaymentParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type BankInfo struct {
	BankName        string `json:"bank_name"`
	BankCountryCode string `json:"bank_country_code"`
	BankCountryName string `json:"bank_country_name"`
	BankImage       string `json:"bank_image"`
}

func (p *Payment) IsTestPayment(purchaseToken string) bool {
	s := strings.Split(purchaseToken, ".")[0]
	return cast.ToInt(s) >= 1000000000
}

func (p *Payment) doPayment(ctx context.Context, url string) (GetTokenPaymentResponse, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return GetTokenPaymentResponse{}, err
	}
	response, err := p.client.Do(request, client.RequestOpts{})
	if err != nil {
		return GetTokenPaymentResponse{}, err
	}

	defer response.Body.Close()

	tokenPaymentResponse := GetTokenPaymentResponse{}
	err = json.NewDecoder(response.Body).Decode(&tokenPaymentResponse)

	return tokenPaymentResponse, err
}

func (p *Payment) GetPaymentInfo(ctx context.Context, purchaseToken string) (GetTokenPaymentResponse, error) {
	if p.IsTestPayment(purchaseToken) {
		return p.GetSandboxPaymentInfo(ctx, purchaseToken)
	}
	return p.GetProductionPaymentInfo(ctx, purchaseToken)
}

func (p *Payment) GetSandboxPaymentInfo(ctx context.Context, purchaseToken string) (GetTokenPaymentResponse, error) {
	url := fmt.Sprintf(GetSandboxPaymentInfoURL, purchaseToken)
	return p.doPayment(ctx, url)
}

func (p *Payment) GetProductionPaymentInfo(ctx context.Context, purchaseToken string) (GetTokenPaymentResponse, error) {
	url := fmt.Sprintf(GetPaymentInfoURL, purchaseToken)
	return p.doPayment(ctx, url)
}

type GetTokenSubscriptionResponse struct {
	Code      string            `json:"code"`
	Message   string            `json:"message"`
	Timestamp time.Time         `json:"timestamp"`
	Body      TokenSubscription `json:"body"`
}

type TokenSubscription struct {
	Code    int                   `json:"code"`
	Success bool                  `json:"success"`
	Message string                `json:"message"`
	Body    TokenSubscriptionBody `json:"body"`
}

type TokenSubscriptionBody struct {
	ServiceName       string    `json:"serviceName"`
	SubscriptionID    int       `json:"subscriptionId"`
	AddParameters     string    `json:"addParameters"`
	ProductType       string    `json:"productType"`
	ProductName       string    `json:"productName"`
	ProductCode       string    `json:"productCode"`
	Recurrent         bool      `json:"recurrent"`
	CountOfDay        int       `json:"countOfDay"`
	PeriodType        string    `json:"periodType"`
	PeriodDuration    int       `json:"periodDuration"`
	NextPaymentDate   string    `json:"nextPaymentDate"`
	Price             int       `json:"price"`
	Currency          string    `json:"currency"`
	ImageURL          string    `json:"imageUrl"`
	State             string    `json:"state"`
	CurrentPeriod     string    `json:"currentPeriod"`
	DebtPaymentPeriod string    `json:"debtPaymentPeriod"`
	Description       string    `json:"description"`
	TariffID          int       `json:"tariffId"`
	Periods           []Periods `json:"periods"`
}

type Periods struct {
	PeriodName     string `json:"periodName"`
	PeriodType     string `json:"periodType"`
	PeriodDuration int    `json:"periodDuration"`
	PeriodPrice    int    `json:"periodPrice"`
	NextPeriod     string `json:"nextPeriod"`
}

func (p *Payment) GetSubscriptionInfo(ctx context.Context, subscriptionToken string) (GetTokenSubscriptionResponse, error) {
	url := fmt.Sprintf(GetSubscriptionInfoURL, subscriptionToken)

	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return GetTokenSubscriptionResponse{}, err
	}

	response, err := p.client.Do(request, client.RequestOpts{})
	if err != nil {
		return GetTokenSubscriptionResponse{}, err
	}

	defer response.Body.Close()

	tokenSubscriptionResponse := GetTokenSubscriptionResponse{}
	err = json.NewDecoder(response.Body).Decode(&tokenSubscriptionResponse)

	return tokenSubscriptionResponse, err
}

type GetTokenSubscriptionV2Response struct {
	StartTimeMillis       string                `json:"startTimeMillis"`
	ExpiryTimeMillis      string                `json:"expiryTimeMillis"`
	AutoRenewing          bool                  `json:"autoRenewing"`
	PriceCurrencyCode     string                `json:"priceCurrencyCode"`
	PriceAmountMicros     string                `json:"priceAmountMicros"`
	CountryCode           string                `json:"countryCode"`
	PaymentState          int                   `json:"paymentState"`
	OrderID               string                `json:"orderId"`
	AcknowledgementState  int                   `json:"acknowledgementState"`
	Kind                  string                `json:"kind"`
	PurchaseType          int                   `json:"purchaseType"`
	IntroductoryPriceInfo IntroductoryPriceInfo `json:"introductoryPriceInfo"`
}

type IntroductoryPriceInfo struct {
	IntroductoryPriceCurrencyCode string `json:"introductoryPriceCurrencyCode"`
	IntroductoryPriceAmountMicros string `json:"introductoryPriceAmountMicros"`
	IntroductoryPricePeriod       string `json:"introductoryPricePeriod"`
	IntroductoryPriceCycles       string `json:"introductoryPriceCycles"`
}

func (p *Payment) GetSubscriptionInfoV2(ctx context.Context, packageName, subscriptionID,
	subscriptionToken string) (GetTokenSubscriptionV2Response, error) {
	url := fmt.Sprintf(
		GetSubscriptionInfoV2URL,
		packageName,
		subscriptionID,
		subscriptionToken,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return GetTokenSubscriptionV2Response{}, err
	}

	response, err := p.client.Do(request, client.RequestOpts{})
	if err != nil {
		return GetTokenSubscriptionV2Response{}, err
	}

	defer response.Body.Close()

	tokenSubscriptionResponse := GetTokenSubscriptionV2Response{}
	err = json.NewDecoder(response.Body).Decode(&tokenSubscriptionResponse)

	return tokenSubscriptionResponse, err
}

func (p *Payment) GetSubscriptionInfoV3(ctx context.Context, packageName, subscriptionID,
	subscriptionToken string) (GetTokenSubscriptionV2Response, error) {
	url := fmt.Sprintf(
		GetSubscriptionInfoV3URL,
		packageName,
		subscriptionID,
		subscriptionToken,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return GetTokenSubscriptionV2Response{}, err
	}

	response, err := p.client.Do(request, client.RequestOpts{})
	if err != nil {
		return GetTokenSubscriptionV2Response{}, err
	}

	defer response.Body.Close()

	tokenSubscriptionResponse := GetTokenSubscriptionV2Response{}
	err = json.NewDecoder(response.Body).Decode(&tokenSubscriptionResponse)

	return tokenSubscriptionResponse, err
}

type GetSubscriptionStateResponse struct {
	Code      string    `json:"code"`
	Message   *string   `json:"message"`
	Body      IsActive  `json:"body"`
	Timestamp time.Time `json:"timestamp"`
}

type IsActive struct {
	IsActive bool `json:"is_active"`
}

func (p *Payment) GetSubscriptionState(ctx context.Context, subscriptionToken string) (GetSubscriptionStateResponse, error) {
	url := fmt.Sprintf(GetSubscriptionStateURL, subscriptionToken)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return GetSubscriptionStateResponse{}, err
	}

	response, err := p.client.Do(request, client.RequestOpts{})
	if err != nil {
		return GetSubscriptionStateResponse{}, err
	}

	defer response.Body.Close()

	subscriptionState := GetSubscriptionStateResponse{}
	err = json.NewDecoder(response.Body).Decode(&subscriptionState)

	return subscriptionState, err
}

type SubmitSubscription struct {
	Code      string    `json:"code"`
	Message   string    `json:"message"`
	Body      *string   `json:"body"` // при успешном ответе body пустой
	Timestamp time.Time `json:"timestamp"`
}

func (p *Payment) SubmitSubscription(ctx context.Context, packageName, subscriptionID,
	subscriptionToken string) (SubmitSubscription, error) {
	url := fmt.Sprintf(
		SubmitSubscriptionURL,
		packageName,
		subscriptionID,
		subscriptionToken,
	)

	request, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		url,
		nil,
	)
	if err != nil {
		return SubmitSubscription{}, err
	}

	response, err := p.client.Do(request, client.RequestOpts{})
	if err != nil {
		return SubmitSubscription{}, err
	}

	defer response.Body.Close()

	submitSubscription := SubmitSubscription{}
	err = json.NewDecoder(response.Body).Decode(&submitSubscription)

	return submitSubscription, err
}
