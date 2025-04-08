# Gluetun Exporter
A Promtheus Exporter for the [VPN Client Gluetun](https://github.com/qdm12/gluetun)

---

## Features

- Current Public IP Information (IP, Country, City)
- Forwarded Ports
- Connection Status

## Installation

### Standalone
The Exporter can be deployed Standalone in this Variation it will not Collect the Troughput of the Tunnel, in this Installation only the VPN Data from the API Endpoint is available.
```bash
docker build -t gluetun-exporter:standalone --file Dockerfile-standalone .
```

### Bundled
The Exporter can be deployed Bundled with Gluetun self, the docker Image get builded around the Gluetun Image and will run in the Background. </br>
In this Installation the Troughput of the Tunnel is available, it get's collected with the Usage of [netlink](github.com/vishvananda/netlink) and the Link Statistics.
```bash
docker build -t gluetun-exporter:bundled --file Dockerfile-bundled .
```

## Configuration
The Configuration is currently only available via the Environment Vars:
```env
GLUETUN_HOST=localhost
GLUETUN_PORT=8000
EXPORTER_PORT=8001
EXPORTER_ROLE=gluetun
EXPORTER_INTERVAL=30
```

The Authentication is either via Usernername/Password Combi or API-Key:
```env
GLUETUN_USERNAME=username
GLUETUN_PASSWORD=password
# OR
GLUETUN_APIKEY=apikey
```

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
