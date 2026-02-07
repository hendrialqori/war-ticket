package util

import (
	"crypto/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandomString membuat random string aman dengan panjang n
func GenerateRandomString(n int) (string, error) {
	bytes := make([]byte, n)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	for i := range bytes {
		bytes[i] = charset[int(bytes[i])%len(charset)]
	}

	return string(bytes), nil
}
