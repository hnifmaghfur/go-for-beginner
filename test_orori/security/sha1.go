package security

import (
	"crypto/sha1"
	"crypto/sha512"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"crypto/hmac"
	"fmt"
)

func ShaOneEncrypt(s string) (string) {
	h := sha1.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}

func Sha256Encrypt(s string) (string) {
	h := sha256.New()
	h.Write([]byte(s))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}

func Sha512Encrypt(s string) (string) {
	h := sha512.New()
	h.Write([]byte(s))
	sha512_hash := hex.EncodeToString(h.Sum(nil))
	
	return sha512_hash
}

func Sha256_HMAC(k string, message string) (string) {
	key := []byte(k)
	sig := hmac.New(sha256.New, key)
	sig.Write([]byte(message))

	hmac := fmt.Sprintf(hex.EncodeToString(sig.Sum(nil)))

	return hmac
}

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Sha1_HMAC(k string, message string) (string) {
	key := []byte(k)
	sig := hmac.New(sha1.New, key)
	sig.Write([]byte(message))

	hmac := fmt.Sprintf(hex.EncodeToString(sig.Sum(nil)))

	return hmac
}
