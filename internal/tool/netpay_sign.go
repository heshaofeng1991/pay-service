package tool

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// Sign Base64_Encode(HmacSHA256(appId + timestamp + nonce + SHA256_HEX(content), AppKey))
func Sign(appId, timestamp, nonce, content, appKey string) string {
	return HmacSha256(appId, timestamp, nonce, content, appKey)
}

// HmacSha256 HmacSHA256(appId + timestamp + nonce + SHA256_HEX(content)
func HmacSha256(appId, timestamp, nonce, content, appKey string) string {
	message := appId + timestamp + nonce + SHA256HEX(content)
	key := []byte(appKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	fmt.Println(h.Sum(nil))
	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sha)
	return base64.StdEncoding.EncodeToString([]byte(sha))
}

func SHA256HEX(content string) string {
	h := sha256.New()
	h.Write([]byte(content))
	fmt.Println(h.Sum(nil))
	sha := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sha)
	return sha
}
