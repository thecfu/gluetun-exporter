package gluetun

import (
	"github.com/qdm12/log"
)

var (
	logger = log.New(log.SetLevel(log.LevelInfo), log.SetComponent("gluetun-exporter"))
)

type Server struct {
	username string
	password string
	apikey   string
	url      string
}
