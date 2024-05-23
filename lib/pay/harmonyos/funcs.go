package harmonyos

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/caiknife/mp3lister/lib/fjson"
)

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
