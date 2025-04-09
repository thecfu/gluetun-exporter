package gluetun

import (
	"encoding/json"
	"github.com/thecfu/gluetun-exporter/pkg/gluetun/types"
	"github.com/thecfu/gluetun-exporter/pkg/promexporter"
	"net/http"
	"os"
)

func (s *Server) Collect() {
	client := &http.Client{}

	urls := map[string]func(*http.Response){
		s.url + "/v1/vpn/status":            handleStatus,
		s.url + "/v1/publicip/ip":           handlePublicIP,
		s.url + "/v1/openvpn/portforwarded": handlePortForward,
	}

	for url, handler := range urls {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			continue
		}

		if s.username != "" && s.password != "" {
			req.SetBasicAuth(s.username, s.password)
		} else if s.apikey != "" {
			req.Header.Set("Authorization", "Bearer "+s.apikey)
		}

		resp, err := client.Do(req)
		if err != nil {
			continue
		}

		handler(resp)
		err = resp.Body.Close()
		if err != nil {
			logger.Errorf("Failed to close response body: %v", err)
			os.Exit(1)
		}
	}
}

func handleStatus(resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		logger.Error("Failed to get VPN status")
		promexporter.UpdateVPNStatus(0)
		return
	}

	var statusMap = map[string]float64{
		"running": 1,
		"stopped": 0,
		"error":   -1,
	}

	type StatusResponse struct {
		Status string `json:"status"`
	}

	var status StatusResponse
	err := json.NewDecoder(resp.Body).Decode(&status)
	if err != nil {
		logger.Errorf("Failed to decode VPN status: %v", err)
		return
	}

	value, ok := statusMap[status.Status]
	if !ok {
		value = -2 // Unknown status
	}

	promexporter.UpdateVPNStatus(value)
}

func handlePortForward(resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		logger.Error("Failed to get port forward status")
		promexporter.UpdateVPNStatus(0)
		return
	}
	type PortForwardResponse struct {
		Ports []int `json:"ports"`
	}

	var portForward PortForwardResponse
	err := json.NewDecoder(resp.Body).Decode(&portForward)
	if err != nil {
		logger.Errorf("Failed to decode port forward response: %v", err)
		return
	}

	promexporter.UpdateForwardedPorts(portForward.Ports)
}

func handlePublicIP(resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		logger.Error("Failed to get public IP")
		promexporter.UpdateVPNStatus(0)
		return
	}

	type PublicIPResponse struct {
		IP      string `json:"public_ip"`
		Country string `json:"country"`
		City    string `json:"city"`
	}

	var publicIP PublicIPResponse
	err := json.NewDecoder(resp.Body).Decode(&publicIP)
	if err != nil {
		logger.Errorf("Failed to decode public IP response: %v", err)
		return
	}

	vpnInfo := types.VPNInfo{
		IP:      publicIP.IP,
		Country: publicIP.Country,
		City:    publicIP.City,
	}

	promexporter.UpdateVPNInfos(vpnInfo)
}
