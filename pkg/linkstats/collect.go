package linkstats

import (
	"github.com/thecfu/gluetun-exporter/pkg/linkstats/types"
	"github.com/thecfu/gluetun-exporter/pkg/promexporter"
	"github.com/vishvananda/netlink"
)

func Scrape(linkName string) {
	link, err := netlink.LinkByName(linkName)
	if err != nil {
		return
	}
	stats := link.Attrs().Statistics
	if stats == nil {
		return
	}

	newStats := types.Statistics{
		Interface:  linkName,
		RxBytes:    stats.RxBytes,
		TxBytes:    stats.TxBytes,
		RxPackets:  stats.RxPackets,
		TxPackets:  stats.TxPackets,
		RxErrors:   stats.RxErrors,
		TxErrors:   stats.TxErrors,
		RxDropped:  stats.RxDropped,
		TxDropped:  stats.TxDropped,
		Collisions: stats.Collisions,
	}

	// Update the Prometheus gauges with the new statistics
	promexporter.UpdateLinkStats(newStats)
}
