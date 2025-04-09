package gluetun

import (
	"os"
	"strings"
)

func New() *Server {
	url := os.Getenv("GLUETUN_URL")
	if url == "" {
		url = "http://localhost:8000"
	}

	url = strings.TrimSuffix(url, "/")

	if !IsValidURL(url) {
		logger.Errorf("Invalid Gluetun URL: %s", url)
		os.Exit(1)
	}

	username := os.Getenv("GLUETUN_USERNAME")
	password := os.Getenv("GLUETUN_PASSWORD")
	apikey := os.Getenv("GLUETUN_APIKEY")

	return &Server{
		username: username,
		password: password,
		apikey:   apikey,
		url:      url,
	}
}
