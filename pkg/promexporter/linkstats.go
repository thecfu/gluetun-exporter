package promexporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/thecfu/gluetun-exporter/pkg/types"
)

var (
	rxBytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gluetun_rx_bytes",
		Help: "Number of bytes received",
	})
	txBytes = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gluetun_tx_bytes",
		Help: "Number of bytes transmitted",
	})
	rxPackets = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gluetun_rx_packets",
		Help: "Number of packets received",
	})
	txPackets = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gluetun_tx_packets",
		Help: "Number of packets transmitted",
	})
	rxErrors = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gluetun_rx_errors",
		Help: "Number of receive errors",
	})
	txErrors = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gluetun_tx_errors",
		Help: "Number of transmit errors",
	})
	rxDropped = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gluetun_rx_dropped",
		Help: "Number of packets dropped",
	})
	txDropped = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gluetun_tx_dropped",
		Help: "Number of packets dropped",
	})
	Collisions = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gluetun_collisions",
		Help: "Number of collisions",
	})
)

func UpdateLinkStats(stats types.Statistics) {
	rxBytes.Set(float64(stats.RxBytes))
	txBytes.Set(float64(stats.TxBytes))
	rxPackets.Set(float64(stats.RxPackets))
	txPackets.Set(float64(stats.TxPackets))
	rxErrors.Set(float64(stats.RxErrors))
	txErrors.Set(float64(stats.TxErrors))
	rxDropped.Set(float64(stats.RxDropped))
	txDropped.Set(float64(stats.TxDropped))
	Collisions.Set(float64(stats.Collisions))
}

func RegisterLinkStats() {
	prometheus.MustRegister(rxBytes)
	prometheus.MustRegister(txBytes)
	prometheus.MustRegister(rxPackets)
	prometheus.MustRegister(txPackets)
	prometheus.MustRegister(rxErrors)
	prometheus.MustRegister(txErrors)
	prometheus.MustRegister(rxDropped)
	prometheus.MustRegister(txDropped)
	prometheus.MustRegister(Collisions)
}
