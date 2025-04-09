package main

import (
	"github.com/qdm12/log"
	"github.com/thecfu/gluetun-exporter/pkg/gluetun"
	"github.com/thecfu/gluetun-exporter/pkg/linkstats"
	"github.com/thecfu/gluetun-exporter/pkg/promexporter"
	"os"
	"strconv"
	"time"
)

func main() {
	logger := log.New(log.SetLevel(log.LevelInfo), log.SetComponent("gluetun-exporter"))

	// Get the EXPORTER_INTERVAL environment variable
	exporterInterval := os.Getenv("EXPORTER_INTERVAL")
	if exporterInterval == "" {
		exporterInterval = "60" // Default to 60 seconds if not provided
	}

	// Get Interface Name from ENV Variable
	interfaceName := os.Getenv("VPN_INTERFACE")
	if interfaceName == "" {
		interfaceName = "tun0"
	}

	// Convert the interval to an integer
	interval, err := time.ParseDuration(exporterInterval + "s")
	if err != nil {
		logger.Errorf("Invalid EXPORTER_INTERVAL value: %v", err)
		os.Exit(1)
	}

	bundledStr := os.Getenv("EXPORTER_BUNDLED")
	bundled, err := strconv.ParseBool(bundledStr)
	if err != nil {
		bundled = false
		logger.Warnf("Invalid EXPORTER_BUNDLED value, defaulting to false: %v", err)
	}

	// Start the Prometheus exporter server in a background goroutine
	go func() {
		logger.Info("Starting prometheus exporter...")
		promexporter.Serve(bundled)
	}()

	// Start the metric collection loop
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	controlServer := gluetun.New()

	for {
		select {
		case <-ticker.C:
			// Collect metrics from the control server
			logger.Info("Updating Metrics...")
			if bundled {
				linkstats.Scrape(interfaceName)
			}
			controlServer.Collect()
			logger.Info("Updated Metrics")
		}
	}
}
