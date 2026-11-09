package helpers

import (
	"crypto/rand"
	"encoding/base64"
)

func GenToken() (string, error) {
	refreshBytes := make([]byte, 32)
	if _, err := rand.Read(refreshBytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(refreshBytes), nil
}
