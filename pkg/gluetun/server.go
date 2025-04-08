package gluetun

import (
	"github.com/qdm12/log"
	"os"
)

var (
	logger = log.New(log.SetLevel(log.LevelInfo), log.SetComponent("gluetun-exporter"))
)

type Server struct {
	username string
	password string
	apikey   string
	host     string
	port     string
}

func New() *Server {
	host := os.Getenv("GLUETUN_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("GLUETUN_PORT")
	if port == "" {
		port = "8000"
	}
	username := os.Getenv("GLUETUN_USERNAME")
	password := os.Getenv("GLUETUN_PASSWORD")
	apikey := os.Getenv("GLUETUN_APIKEY")

	return &Server{
		username: username,
		password: password,
		apikey:   apikey,
		host:     host,
		port:     port,
	}
}
