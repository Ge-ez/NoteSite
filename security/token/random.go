package token

import (
	"crypto/rand"
	"encoding/base64"
	mrand "math/rand"
	"time"
)
// GenerateRandomBytes returns securely generated random bytes.
func GenerateRandomBytes(n int) ([]byte, error) {
	mrand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
// GenerateRandomString returns a URL-safe, base64 encoded securely generated random string.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}


