package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("missing Authorization header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid Authorization header format")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid Authorization header type")
	}

	return vals[1], nil
}
