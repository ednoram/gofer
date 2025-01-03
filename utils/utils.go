package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashAPIKey(apiKey string) string {
	hash := sha256.New()
	hash.Write([]byte(apiKey))
	return hex.EncodeToString(hash.Sum(nil))
}
