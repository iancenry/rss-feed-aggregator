package auth

import (
	"errors"
	"net/http"
	"strings"
)


// GetAPIKey returns the API key from the request header
// Example: Authorization: ApiKey <api-key>
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("authorization header not found")
	}

	values := strings.Split(authHeader, " ")
	if len(values) != 2 {
		return "", errors.New("invalid authorization header")
	}

	if values[0] != "ApiKey" {
		return "", errors.New("malformed first element in authorization header")
	}

	return values[1], nil
} 