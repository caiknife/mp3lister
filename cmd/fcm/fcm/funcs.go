package fcm

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

func AES128GCMWithBase64(originText, key string) string {
	decodeString, err := hex.DecodeString(key)
	if err != nil {
		return ""
	}
	bytes, err := encrypt([]byte(originText), decodeString)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(bytes)
}

func encrypt(originText, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	data := append(nonce, gcm.Seal(nil, nonce, originText, nil)...)
	return data, nil
}
