package test

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAESGCM(t *testing.T) {
	secretMessage := []byte("send reinforcements, we're going to advance")
	rng := rand.Reader
	key := &rsa.PublicKey{}
	cipherData, err := rsa.EncryptOAEP(sha256.New(), rng, key, secretMessage, nil)
	if err != nil {
		t.Error("Error from encryption:", err)
		return
	}
	ciphertext := base64.StdEncoding.EncodeToString(cipherData)
	fmt.Printf("Ciphertext: %s\n", ciphertext)
}
