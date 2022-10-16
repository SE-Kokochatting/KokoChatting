package util

import (
	"crypto/sha256"
	"encoding/hex"
)

// Encryption 加密函数
func Encryption(encryptionObject string) (encryptedObject string) {
	encryptedObject = ""
	sha256 := sha256.New()
	sha256.Write([]byte(encryptionObject))
	return hex.EncodeToString(sha256.Sum([]byte("")))
}