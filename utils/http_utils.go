package utils

import (
	"errors"
	"net/http"
	"strings"
)

func ExtractIDFromURL(r *http.Request) (string, error) {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		return "", errors.New("Invalid URL")
	}
	return parts[2], nil
}
