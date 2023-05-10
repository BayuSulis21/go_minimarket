package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// function to create signature, kalau nama func huruf besar=public
func CreateSignature(secretKey string, data string) string {
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	return signature
}

// function to validate signature
func ValidateSignature(secretKey string, data string, signature string) bool {
	key := []byte(secretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}
