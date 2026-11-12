package utils

import (
	"errors"
	"net/http"
	"strings"
)

func GetToken(r *http.Request) (string, error) {
	header := r.Header["Authorization"]
	if len(header) == 0 {
		return "", errors.New("empty header")
	}
	token := strings.TrimPrefix(header[0], "Bearer ")
	if len(token) == 0 {
		return "", errors.New("empty token")
	}

	return token, nil
}
