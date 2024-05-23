package harmonyos

import (
	"context"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"
	"time"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/slice"
	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

type Client struct {
	PrivateKey string `json:"private_key"`
	KeyID      string `json:"key_id"`
	IssuerID   string `json:"issuer_id"`
	AppID      string `json:"app_id"`

	token *jwt.Token
}

func NewClient(privateKey, keyID, issuerID, appID string) *Client {
	c := &Client{
		PrivateKey: privateKey,
		KeyID:      keyID,
		IssuerID:   issuerID,
		AppID:      appID,
	}
	c.token = jwt.New(jwt.SigningMethodES256)
	return c
}

func (c *Client) GetToken(body any) (s string, err error) {
	c.token.Header["kid"] = c.KeyID
	now := time.Now().Unix()
	hash, err := c.Hash(body)
	if err != nil {
		return "", errors.WithMessage(err, "hash failed")
	}
	c.token.Claims = jwt.MapClaims{
		"iss":    c.IssuerID,
		"aud":    "iap-v1",
		"iat":    now,
		"exp":    now + 60*60,
		"aid":    c.AppID,
		"digest": hash,
	}
	pem, err := jwt.ParseECPrivateKeyFromPEM([]byte(c.PrivateKey))
	if err != nil {
		return "", errors.WithMessage(err, "parse private key failed")
	}
	s, err = c.token.SignedString(pem)
	if err != nil {
		return "", errors.WithMessage(err, "sign token failed")
	}
	return s, nil
}

func (c *Client) Hash(body any) (s string, err error) {
	toString, err := fjson.MarshalToString(body)
	if err != nil {
		return "", errors.WithMessage(err, "json marshal failed")
	}
	s = cryptor.Sha256(toString)
	return s, nil
}

func (c *Client) Verify(token string) (err error) {
	split := strings.Split(token, ".")
	if len(split) != 3 {
		return errors.New("invalid purchase order token")
	}

	header, err := GetHeader(split[0])
	if err != nil {
		return errors.WithMessage(err, "get header failed")
	}

	if len(header.X5C) != 3 {
		return errors.New("invalid x5c header")
	}

	// 证书链验证
	leafCert, err := LoadCertificate(header.X5C[0])
	if err != nil {
		return errors.WithMessage(err, "load leaf certificate failed")
	}
	caCert, err := LoadCertificate(header.X5C[1])
	if err != nil {
		return errors.WithMessage(err, "load ca certificate failed")
	}
	rootCert, err := LoadCertificate(header.X5C[2])
	if err != nil {
		return errors.WithMessage(err, "load root certificate failed")
	}

	err = leafCert.CheckSignatureFrom(caCert)
	if err != nil {
		return errors.WithMessage(err, "check leaf from ca signature failed")
	}

	err = caCert.CheckSignatureFrom(rootCert)
	if err != nil {
		return errors.WithMessage(err, "check ca from root certificate failed")
	}

	// 子证书oid验证
	extensions := types.Slice[*pkix.Extension](slice.Map[pkix.Extension, *pkix.Extension](leafCert.Extensions, func(index int, item pkix.Extension) *pkix.Extension {
		return &item
	}))

	_, find := extensions.Find(func(extension *pkix.Extension) bool {
		return extension.Id.String() == oID
	})
	if !find {
		return errors.New("leaf certificate oid not found in extensions")
	}

	// 完整验证
	_, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return leafCert.PublicKey, nil
	})
	if err != nil {
		return errors.WithMessage(err, "parse order failed")
	}
	return nil
}

func (c *Client) QueryOrder(purchaseToken, purchaseOrderID string) (r *Response, err error) {
	body := types.Map[string]{
		"purchaseOrderId": purchaseOrderID,
		"purchaseToken":   purchaseToken,
	}
	token, err := c.GetToken(body)
	if err != nil {
		return nil, errors.WithMessage(err, "get token failed")
	}
	client := resty.New()
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()
	post, err := client.R().SetBody(body).SetContext(timeout).SetHeaders(types.Map[string]{
		"Authorization": fmt.Sprintf("Bearer %s", token),
		"Content-Type":  "application/json;charset=UTF-8",
	}).Post(QueryOrder)
	if err != nil {
		return nil, errors.WithMessage(err, "http client query failed")
	}
	r = &Response{}
	err = fjson.UnmarshalFromString(post.String(), r)
	if err != nil {
		return nil, errors.WithMessage(err, "json unmarshal failed")
	}
	return r, nil
}

func (c *Client) Check(purchaseToken, purchaseOrderID string) (err error) {
	order, err := c.QueryOrder(purchaseToken, purchaseOrderID)
	if err != nil {
		return errors.WithMessage(err, "query order failed")
	}
	if !order.ResponseOK() {
		return errors.New("query order response is not ok")
	}

	return c.Verify(order.JWSPurchaseOrder)
}

func GetHeader(s string) (*Header, error) {
	s = Base64RawStdDecode(s)
	h := &Header{}
	err := fjson.UnmarshalFromString(s, h)
	if err != nil {
		return nil, errors.WithMessage(err, "json unmarshal failed")
	}
	return h, nil
}

func GetPayload(s string) (*Payload, error) {
	s = Base64RawStdDecode(s)
	h := &Payload{}
	err := fjson.UnmarshalFromString(s, h)
	if err != nil {
		return nil, errors.WithMessage(err, "json unmarshal failed")
	}
	return h, nil
}

func Base64RawStdDecode(s string) string {
	decodeString, _ := base64.RawStdEncoding.DecodeString(s)
	return string(decodeString)
}

func Base64RawStdEncode(s string) string {
	return base64.RawStdEncoding.EncodeToString([]byte(s))
}

func LoadCertificate(key string) (*x509.Certificate, error) {
	if !strings.HasPrefix(key, "-----BEGIN") {
		key = fmt.Sprintf(`-----BEGIN CERTIFICATE-----
%s
-----END CERTIFICATE-----`, key)
	}
	pemBlock, _ := pem.Decode([]byte(key))
	if pemBlock == nil {
		return nil, errors.New("pem decode failed")
	}
	cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return nil, errors.New("failed to parse certificate")
	}
	return cert, nil
}
