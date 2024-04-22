package client

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	BaseURL        = "https://public-api.rustore.ru/public/v1/application"
	AuthURL        = "https://public-api.rustore.ru/public/auth"
	TimeOutSeconds = 5
)

type JWETokenResponse struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Body    JWEToken `json:"body"`
}

type JWEToken struct {
	JWE string `json:"jwe"`
}

type Client struct {
	companyID          string
	privateKeyRaw      string
	privateKeyComplete string
	keyID              string
	jwe                string
	httpClient         *http.Client
}

func New(keyID, privateKeyRaw, companyID string) *Client {
	c := &Client{
		companyID:     companyID,
		privateKeyRaw: privateKeyRaw,
		httpClient:    &http.Client{},
		keyID:         keyID,
	}

	return c
} // создание нового клиента

func (c *Client) Auth() error {
	privateKeyComplete, err := c.GetEncodedSignature(c.privateKeyRaw)
	if err != nil {
		return err
	}

	c.privateKeyComplete = privateKeyComplete // получили зашифрованный приватный ключ
	c.jwe, err = c.GetJWEToken(c.privateKeyComplete)

	return err
}

func (c *Client) GetEncodedSignature(privateKeyRaw string) (string, error) {
	completePrivateKey := fmt.Sprintf(
		"-----BEGIN PRIVATE KEY-----\n%s\n-----END PRIVATE KEY-----",
		privateKeyRaw,
	)

	timestamp := time.Now().Format(time.RFC3339Nano)

	block, _ := pem.Decode([]byte(completePrivateKey))
	if block == nil {
		return "", fmt.Errorf("не удалось расшифровать ключ: проверьте правильность приватного ключа")
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", fmt.Errorf("не удалось расшифровать ключ: неверный формат ключа")
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("не удалось прочитать приватный ключ")
	}

	messageToSign := c.companyID + timestamp // Отсюда возьмём хеш
	if c.keyID != "" {
		messageToSign = c.keyID + timestamp
	}

	hashed := sha512.Sum512([]byte(messageToSign))

	signatureBytes, err := rsa.SignPKCS1v15(
		rand.Reader,
		rsaPrivateKey,
		crypto.SHA512,
		hashed[:],
	)
	if err != nil {
		return "", fmt.Errorf("не удалось зашифровать ключ: неверный хеш")
	}

	signatureValue := base64.StdEncoding.EncodeToString(signatureBytes)

	resultMap := map[string]string{
		"signature": signatureValue,
		"timestamp": timestamp,
	}
	if c.keyID != "" {
		resultMap["keyId"] = c.keyID
	} else {
		resultMap["companyId"] = c.companyID
	}

	ResultJSON, err := json.Marshal(resultMap)

	return string(ResultJSON), err
}

func (c *Client) GetJWEToken(privateKeyComplete string) (string, error) {
	request, err := http.NewRequest(
		http.MethodPost,
		AuthURL,
		bytes.NewBuffer([]byte(privateKeyComplete)),
	)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	response, err := c.httpClient.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	jweToken := JWETokenResponse{}

	err = json.Unmarshal(body, &jweToken)

	return jweToken.Body.JWE, err
}

type RequestOpts struct {
	CustomContentType string
}

func (c *Client) Do(request *http.Request, opts RequestOpts) (*http.Response, error) {
	contentType := opts.CustomContentType
	if contentType == "" {
		contentType = "application/json; charset=UTF-8"
	}

	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Public-Token", c.jwe)
	response, err := c.httpClient.Do(request)

	return response, err
}
