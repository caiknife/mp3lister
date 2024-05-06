package harmonyos

import (
	"context"
	"fmt"
	"time"

	"github.com/duke-git/lancet/v2/cryptor"
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
	c.token.Header["kid"] = "this"
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
	return nil
}

func (c *Client) QueryOrder(purchaseToken, purchaseOrderId string) (err error) {
	body := types.Map[string]{
		"purchaseOrderId": purchaseOrderId,
		"purchaseToken":   purchaseToken,
	}
	token, err := c.GetToken(body)
	if err != nil {
		return errors.WithMessage(err, "get token failed")
	}
	client := resty.New()
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	post, err := client.R().SetBody(body).SetContext(timeout).SetHeaders(types.Map[string]{
		"Authorization": fmt.Sprintf("Bearer %s", token),
		"Content-Type":  "application/json;charset=UTF-8",
	}).Post(QueryOrder)
	if err != nil {
		return errors.WithMessage(err, "query order failed")
	}
	fmt.Println(post.String())
	return nil
}
