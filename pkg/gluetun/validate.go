package gluetun

import (
	"net/url"
	"strings"
)

func IsValidURL(u string) bool {
	parsed, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}
	return strings.HasPrefix(parsed.Scheme, "http")
}
