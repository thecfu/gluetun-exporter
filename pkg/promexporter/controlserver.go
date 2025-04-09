package promexporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/thecfu/gluetun-exporter/pkg/gluetun/types"
	"strconv"
)

var (
	vpnStatus = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gluetun_vpn_status",
		Help: "Is the VPN connected",
	})
	forwardedPorts = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gluetun_forwarded_ports",
		Help: "Port forwarding status",
	}, []string{"port"})
	totalForwardedPorts = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "gluetun_forwarded_ports_total",
		Help: "Total number of forwarded ports",
	})
	vpnInfos = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "gluetun_vpn_infos",
		Help: "Label containing the VPN information",
	}, []string{"ip", "country", "city"})
)

func RegisterControlServerMetrics() {
	prometheus.MustRegister(vpnStatus)
	prometheus.MustRegister(forwardedPorts)
	prometheus.MustRegister(totalForwardedPorts)
	prometheus.MustRegister(vpnInfos)
}

func UpdateVPNStatus(status float64) {
	vpnStatus.Set(status)
}

func UpdateForwardedPorts(ports []int) {
	forwardedPorts.Reset()                       // Reset the gauge before updating
	totalForwardedPorts.Set(float64(len(ports))) // Update the total count
	for _, port := range ports {
		forwardedPorts.WithLabelValues(strconv.Itoa(port)).Set(1)
	}
}

func UpdateVPNInfos(infos types.VPNInfo) {
	vpnInfos.Reset() // Reset the gauge before updating
	vpnInfos.WithLabelValues(infos.IP, infos.Country, infos.City).Set(1)
}
