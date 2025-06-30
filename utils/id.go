package utils

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

func GenerateCode(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	s := base64.URLEncoding.EncodeToString(b)
	return strings.TrimRight(s, "=")[:n], nil
}
