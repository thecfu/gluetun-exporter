# Gluetun Exporter
A Promtheus Exporter for the [VPN Client Gluetun](https://github.com/qdm12/gluetun)

---

## Features

- Current Public IP Information (IP, Country, City)
- Forwarded Ports
- Connection Status

### Grafana Dashboards
For example Dashboards please have a look at [Dashboards](docs/dashboards).

## Installation

### Standalone
The Exporter can be deployed Standalone in this Variation it will not Collect the Troughput of the Tunnel, in this Installation only the VPN Data from the API Endpoint is available.
```bash
docker build -t gluetun-exporter:standalone --file Dockerfile-standalone .
docker run -it -p 8001:8001\
  -e GLUETUN_URL=http://localhost:8000\
  -e EXPORTER_INTERVAL=30\
  -e GLUETUN_APIKEY=apikey\
  gluetun-exporter:standalone
```

### Bundled
The Exporter can be deployed Bundled with Gluetun self, the docker Image get builded around the Gluetun Image and will run in the Background. </br>
In this Installation the Troughput of the Tunnel is available, it get's collected with the Usage of [netlink](github.com/vishvananda/netlink) and the Link Statistics.
```bash
docker build -t gluetun-exporter:bundled --file Dockerfile-bundled .
docker run -it -p 8001:8001\
  -e GLUETUN_URL=http://localhost:8000\
  -e EXPORTER_INTERVAL=30\
  -e GLUETUN_APIKEY=apikey\
  --cap-add net_admin\
  gluetun-exporter:bundled
```

## Configuration
The Configuration is currently only available via the Environment Vars:
```env
GLUETUN_URL=http://localhost:8000 # The Url of the Gluetun API Endpoint
EXPORTER_PORT=8001 # Port of the Exporter
EXPORTER_ROLE=gluetun # Not implemented
EXPORTER_INTERVAL=30 # Interval of the Metrics Scrape
EXPORTER_DEBUG=false # Activate the Debuging Logs
```

The Authentication is either via Usernername/Password Combi or API-Key:
```env
GLUETUN_USERNAME=username # Username of the Role
GLUETUN_PASSWORD=password # Password of the Role
# OR
GLUETUN_APIKEY=apikey # ApiKey of the Role
```
the following Routes are needed in the role
- GET /v1/vpn/status
- GET /v1/publicip/ip
- GET /v1/openvpn/portforwarded

## Usage

Once the service is running, it will expose metrics at http://localhost:8001/metrics (default port). You can configure Prometheus to scrape these metrics by adding the following job to your Prometheus configuration:

```yaml
scrape_configs:
  - job_name: 'gluetun'
    static_configs:
      - targets: ['localhost:8001']
```

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request. Make sure to follow the contribution guidelines

> [!IMPORTANT]
> Please fork from the `dev` branch to include any un-released changes.

## License

This project is licensed under the MIT license. See the [LICENSE](LICENSE) file for details.

---
```
___________.__           _________   _____ ____ ___ 
\__    ___/|  |__   ____ \_   ___ \_/ ____\    |   \
  |    |   |  |  \_/ __ \/    \  \/\   __\|    |   /
  |    |   |   Y  \  ___/\     \____|  |  |    |  / 
  |____|   |___|  /\___  >\______  /|__|  |______/  
                \/     \/        \/                 
```
