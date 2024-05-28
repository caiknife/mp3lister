package fcm

import (
	"encoding/hex"
	"testing"

	"github.com/duke-git/lancet/v2/cryptor"
)

const (
	testKey = "b4c932250dbe59ff53b15ee993a9feb5"
)

func TestAES128GCMWithBase64(t *testing.T) {
	base64 := AES128GCMWithBase64(`{"ai":"100000000000000001","name":"某一一","idNum":"110000190101010001"}`, testKey)
	t.Log(base64)
}

func TestSha256(t *testing.T) {
	t.Log(len(testKey))
	decodeString, err := hex.DecodeString(testKey)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(len(decodeString))

	s := `b4c932250dbe59ff53b15ee993a9feb5appIdfdc6688637bc468e9aea874654cbead2bizId1101999999timestamps1705975788903{"data":"kGuV06piX8av9vsZGofHI1viPrHG/IpjsGGu75DYmRyQx6UEvPXrKkAdwWs3SmzEQ5GctOK/N5x5J4Yykw61plWqIL/PytfMZfcnqM43+HmW04agmLU6TJ1ydUnirDl8xGiofmrLLg=="}`
	sha256 := cryptor.Sha256(s)
	t.Log(sha256)
}
