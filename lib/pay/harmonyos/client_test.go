package harmonyos

import (
	"encoding/base64"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/cryptor"
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
