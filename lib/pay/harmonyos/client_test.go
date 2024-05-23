package harmonyos

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/golang-jwt/jwt/v5"

	"github.com/caiknife/mp3lister/lib/fjson"
	"github.com/caiknife/mp3lister/lib/types"
)

func TestJWT(t *testing.T) {
	token := jwt.New(jwt.SigningMethodES256)
	token.Header["kid"] = "this"
	now := time.Now().Unix()
	token.Claims = jwt.MapClaims{
		"iss": "issuer_id",
		"aud": "iap-v1",
		"iat": now,
		"exp": now + 60*60,
		"aid": "app_id",
	}
	key := "12345678"
	pem, err := jwt.ParseECPrivateKeyFromPEM([]byte(key))
	if err != nil {
		t.Error(err)
		return
	}
	signedString, err := token.SignedString(pem)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(signedString)
}

func TestHashBody(t *testing.T) {
	purchaseToken := "123"
	purchaseOrderId := "321"
	toString, err := fjson.MarshalToString(types.Map[string]{
		"purchaseOrderId": purchaseOrderId,
		"purchaseToken":   purchaseToken,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(toString)
	t.Log(cryptor.Sha256(toString))
}

type MyCustomClaims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

func (m *MyCustomClaims) String() string {
	toString, err := fjson.MarshalToString(m)
	if err != nil {
		return ""
	}
	return toString
}

func TestSign(t *testing.T) {
	claim := &MyCustomClaims{
		Foo: "bar",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:   "test",
			Audience: jwt.ClaimStrings{"single"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedString, err := token.SignedString([]byte("AllYourBase"))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(signedString)
}

func TestDecode(t *testing.T) {
	payload := "eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiYXVkIjpbInNpbmdsZSJdfQ"
	decodeString, err := base64.RawStdEncoding.DecodeString(payload)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(decodeString))

	a := &MyCustomClaims{}
	err = fjson.UnmarshalFromString(string(decodeString), a)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(a)
}

func TestParse(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJpc3MiOiJ0ZXN0IiwiYXVkIjoic2luZ2xlIn0.QAWg1vGvnqRuCFTMcPkjZljXHh8U3L_qUjszOtQbeaA"
	a := &MyCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, a, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		log.Fatal(err)
	} else if claims, ok := token.Claims.(*MyCustomClaims); ok {
		fmt.Println(claims.Foo, claims.RegisteredClaims.Issuer)
		t.Log(a)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}
}

const (
	privateKeyFile = "IAPKey_7dade06d-0dc8-4d2a-8d87-7b95ebef6fd7.p8"
	rootCertFile   = "RootCaG2Ecdsa.cer"
	appID          = "5765880207854356261"
	issuerID       = "6537cc5f-1a50-41e3-9bb1-fdebd25964ec"
	keyID          = "7dade06d-0dc8-4d2a-8d87-7b95ebef6fd7"

	purchaseOrderID = "202405161036449949c73488c1.5765880207854356261"
	purchaseToken   = "0000018f7f4279e9d43bd3b8eb415ce7be42d7b2701540560d53a9af00979ed10995fe5f42c1752a.1.5765880207854356261"
)

const (
	jws = "eyJ4NWMiOlsiTUlJQ3d6Q0NBa21nQXdJQkFnSU9DZnF2WmxNY2JOQU0ySlp6S1g0d0NnWUlLb1pJemowRUF3TXdaekVMTUFrR0ExVUVCZ3dDUTA0eER6QU5CZ05WQkFvTUJraDFZWGRsYVRFVE1CRUdBMVVFQ3d3S1NIVmhkMlZwSUVOQ1J6RXlNREFHQTFVRUF3d3BTSFZoZDJWcElFTkNSeUJCY0hCc2FXTmhkR2x2YmlCSmJuUmxjbWR5WVhScGIyNGdRMEVnUnpNd0hoY05NalF3TXpBeE1EWTBNVEEyV2hjTk1qWXdNekF4TURZME1UQTJXakJ2TVFzd0NRWURWUVFHRXdKRFRqRVBNQTBHQTFVRUNnd0dTSFZoZDJWcE1Ta3dKd1lEVlFRTERDQklkV0YzWldrZ1EwSkhJRU5zYjNWa0lGTmxZM1Z5YVhSNUlGTnBaMjVsY2pFa01DSUdBMVVFQXd3YlNIVmhkMlZwSUVOQ1J5QkpiaTFCY0hBZ1VIVnlZMmhoYzJWek1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRStjak5YRGczK0RTbXpXUC8rbG5xYVNJOENjMEVTUFE5R25DYkR1cDR2SEdaZ3NiOHk0dm1YMWYyVEQrd2ZDVzBPWjRDcHFHMlpWaXpHK3Job3IrQTI2T0IwRENCelRBTUJnTlZIUk1CQWY4RUFqQUFNRmtHQTFVZEh3UlNNRkF3VHFCTW9FcUdTR2gwZEhBNkx5OW9OV2h2YzNScGJtY3RaSEpqYmk1a1ltRnVhMk5rYmk1amJpOWpZMmcxTDJOeWJDOW9ZV2xqWVdjekwwaDFZWGRsYVVOQ1IwaEJTVWN6WTNKc0xtTnliREFmQmdOVkhTTUVHREFXZ0JSdEpsSFd2TW8zaGZSODlIdzYzZnNJVHM3MjZEQVNCZ3dyQmdFRUFZOWJBb01mQVFFRUFnVUFNQjBHQTFVZERnUVdCQlRwVkJrM1JVRjVZRSs5Rm1lN3N0QU8xVkNuNGpBT0JnTlZIUThCQWY4RUJBTUNBK2d3Q2dZSUtvWkl6ajBFQXdNRGFBQXdaUUl3U1ZhbFhiQ2s2NW9UaTNNQzJjYlp1NkFZQk1yUHJIRDM2YU9Ra0dRSEZXL1NxcGEwV2wxOWY5a1MwNEo3Nnp2L0FqRUF2Z0RLK0VsWFI2alJqekl2K05zd2FsYi9HVDQvd1UxK0o5NzZpaXdnbEVoTWg1UHB1bnJiT2xTU09DaU1YT3RkIiwiTUlJQzVEQ0NBbXVnQXdJQkFnSUljSHhrbUoyNlp1TXdDZ1lJS29aSXpqMEVBd013VXpFTE1Ba0dBMVVFQmhNQ1EwNHhEekFOQmdOVkJBb01Ca2gxWVhkbGFURVRNQkVHQTFVRUN3d0tTSFZoZDJWcElFTkNSekVlTUJ3R0ExVUVBd3dWU0hWaGQyVnBJRU5DUnlCU2IyOTBJRU5CSUVjeU1CNFhEVEl5TURVeU5EQXhOVEl4TjFvWERUUXlNRFV5TkRBeE5USXhOMW93WnpFTE1Ba0dBMVVFQmd3Q1EwNHhEekFOQmdOVkJBb01Ca2gxWVhkbGFURVRNQkVHQTFVRUN3d0tTSFZoZDJWcElFTkNSekV5TURBR0ExVUVBd3dwU0hWaGQyVnBJRU5DUnlCQmNIQnNhV05oZEdsdmJpQkpiblJsY21keVlYUnBiMjRnUTBFZ1J6TXdkakFRQmdjcWhrak9QUUlCQmdVcmdRUUFJZ05pQUFTVitMTmplcHlaVm5rdmxvakRSWXFYZGdJRWh3WERtZW5QR2VhZUlXUkdSR29rL1B5ZFVIbWI3d1h3dTZsUUlRWHJVMGNwRk5JckQyN3NXVzR4SllSZi95RUVDbFc0Qjg3QXlVaUZoc2hRU0ZaM1BVdFc3Y2RKaUdmQ0tUSmROQ0NqZ2Zjd2dmUXdId1lEVlIwakJCZ3dGb0FVbzQ1YTlWcThjWXdxYWlWeWZraVM0cExjSUFBd0hRWURWUjBPQkJZRUZHMG1VZGE4eWplRjlIejBmRHJkK3doT3p2Ym9NRVlHQTFVZElBUS9NRDB3T3dZRVZSMGdBREF6TURFR0NDc0dBUVVGQndJQkZpVm9kSFJ3T2k4dmNHdHBMbU52Ym5OMWJXVnlMbWgxWVhkbGFTNWpiMjB2WTJFdlkzQnpNQklHQTFVZEV3RUIvd1FJTUFZQkFmOENBUUF3RGdZRFZSMFBBUUgvQkFRREFnRUdNRVlHQTFVZEh3US9NRDB3TzZBNW9EZUdOV2gwZEhBNkx5OXdhMmt1WTI5dWMzVnRaWEl1YUhWaGQyVnBMbU52YlM5allTOWpjbXd2Y205dmRGOW5NbDlqY213dVkzSnNNQW9HQ0NxR1NNNDlCQU1EQTJjQU1HUUNNSE9HbmNrWlkwNkR2aFl1QmNRQXB1K1ZmTkgvZFpSZzNOTzlZWm1hRkVuZG52dlRkblR4anBkenRnbjZrT0ZsaXdJd1BLZHZHeHVYdlJuV1VsWHRNTUY0cjFzRDlHQ1RsQ1hWZVJQY1RFSThtR0U1eVBNY3hBVmhMNHF1a1paQnp6SmIiLCJNSUlDR2pDQ0FhR2dBd0lCQWdJSVNoaHBuNTE5ak5Bd0NnWUlLb1pJemowRUF3TXdVekVMTUFrR0ExVUVCaE1DUTA0eER6QU5CZ05WQkFvTUJraDFZWGRsYVRFVE1CRUdBMVVFQ3d3S1NIVmhkMlZwSUVOQ1J6RWVNQndHQTFVRUF3d1ZTSFZoZDJWcElFTkNSeUJTYjI5MElFTkJJRWN5TUI0WERUSXdNRE14TmpBek1EUXpPVm9YRFRRNU1ETXhOakF6TURRek9Wb3dVekVMTUFrR0ExVUVCaE1DUTA0eER6QU5CZ05WQkFvTUJraDFZWGRsYVRFVE1CRUdBMVVFQ3d3S1NIVmhkMlZwSUVOQ1J6RWVNQndHQTFVRUF3d1ZTSFZoZDJWcElFTkNSeUJTYjI5MElFTkJJRWN5TUhZd0VBWUhLb1pJemowQ0FRWUZLNEVFQUNJRFlnQUVXaWRrR25EU093My9IRTJ5MkdIbCtmcFdCSWE1UytJbG5OcnNHVXZ3QzFJMlFXdnRxQ0hXbXdGbEZLOTV6S1hpTThzOXlWM1ZWWGg3aXZOOFpKTzNTQzVOMVRDcnZCMmxwSE1Cd2N6NERBMGtnSENNbS93RGVjNmtPSHgxeHZDUm8wSXdRREFPQmdOVkhROEJBZjhFQkFNQ0FRWXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVvNDVhOVZxOGNZd3FhaVZ5ZmtpUzRwTGNJQUF3Q2dZSUtvWkl6ajBFQXdNRFp3QXdaQUl3TXlwZUI3UDBJYlk3YzZncFdjQ2xoUnpuT0pGajh1YXZyTnUyUElvejlLSXFyM2puQmxCSEpzMG15STdudFlwRUFqQmJtOGVETVpZNXpxNWlNWlVDNkg3VXpZU2l4NFV5MVlsc0xWVjczOFB0S1A5aEZUamdESGN0WEpsQzVMNytaRFk9Il0sImFsZyI6IkVTMjU2IiwidHlwIjoiSldUIn0.eyJwdXJjaGFzZU9yZGVySWQiOiIyMDI0MDUxNjEwMzY0NDk5NDljNzM0ODhjMS41NzY1ODgwMjA3ODU0MzU2MjYxIiwicHVyY2hhc2VUb2tlbiI6IjAwMDAwMThmN2Y0Mjc5ZTlkNDNiZDNiOGViNDE1Y2U3YmU0MmQ3YjI3MDE1NDA1NjBkNTNhOWFmMDA5NzllZDEwOTk1ZmU1ZjQyYzE3NTJhLjEuNTc2NTg4MDIwNzg1NDM1NjI2MSIsImFwcGxpY2F0aW9uSWQiOiI1NzY1ODgwMjA3ODU0MzU2MjYxIiwicHJvZHVjdElkIjoiY29tLnJhdmVuLnRhbmsuaG1jbi50ZXN0IiwicHVyY2hhc2VUaW1lIjoxNzE1ODI3MDEzMDAwLCJwcm9kdWN0VHlwZSI6IjAiLCJkZXZlbG9wZXJQYXlsb2FkIjoiIiwic2lnbmVkVGltZSI6MTcxNjQzNDE2MjU2MiwiZW52aXJvbm1lbnQiOiJOT1JNQUwiLCJjb3VudHJ5Q29kZSI6IkNOIiwicHJpY2UiOjEsImN1cnJlbmN5IjoiQ05ZIiwiZmluaXNoU3RhdHVzIjoiMSIsIm5lZWRGaW5pc2giOnRydWV9.x64w2Bb7utHl9729uvZG0dsY689FCXMYD0SOsb9qr_dcgXX9U14-n3ZGuBYiCDz04VQX1UnWVwd4KBR8THNwUw"
)

func TestClient(t *testing.T) {
	privateKey, err := fileutil.ReadFileToString(privateKeyFile)
	if err != nil {
		t.Error(err)
		return
	}

	client := NewClient(privateKey, keyID, issuerID, appID)
	err = client.Check(purchaseToken, purchaseOrderID)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(client.Header)
	t.Log(client.Payload)
}

func TestJWSPurchaseOrder(t *testing.T) {
	split := strings.Split(jws, ".")
	header := &Header{}
	err := fjson.UnmarshalFromString(Base64RawStdDecode(split[0]), header)
	if err != nil {
		t.Error(err)
		return
	}

	x := fmt.Sprintf(`-----BEGIN CERTIFICATE-----
%s
-----END CERTIFICATE-----`, header.X5C[0])
	p, _ := pem.Decode([]byte(x))

	certificate, err := x509.ParseCertificate(p.Bytes)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(certificate.Subject)
	t.Log(certificate.Extensions)
}

func TestPrivateKey(t *testing.T) {
	toString, err := fileutil.ReadFileToString(privateKeyFile)
	if err != nil {
		t.Error(err)
		return
	}
	privateKey, err := jwt.ParseECPrivateKeyFromPEM([]byte(toString))
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(privateKey)
}
