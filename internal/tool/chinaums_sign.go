package tool

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

// Sign Base64_Encode(HmacSHA256(appId + timestamp + nonce + SHA256_HEX(content), AppKey))
func Sign(appId, timestamp, nonce, content, appKey string) string {
	return HmacSha256(appId+timestamp+nonce+SHA256HEX(content), appKey)
}

// HmacSha256 HmacSHA256(appId + timestamp + nonce + SHA256_HEX(content)
func HmacSha256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// SHA256HEX SHA256_HEX(content)
func SHA256HEX(content string) string {
	h := sha256.New()
	h.Write([]byte(content))

	return hex.EncodeToString(h.Sum(nil))
}
