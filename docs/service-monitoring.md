# Service Monitoring

Teamgram uses [Prometheus](https://prometheus.io/) for metrics collection and [Grafana](https://grafana.com/) for dashboards. go-zero exposes metrics automatically when the `Prometheus` block is configured in each service.

## Architecture

- **Services**: Each go-zero service exposes a `/metrics` endpoint (host and port set in the service YAML).
- **Prometheus**: Scrapes these endpoints at a regular interval and stores time-series metrics.
- **Grafana**: Connects to Prometheus as a data source and displays dashboards.

## Enabling metrics in services

In each service config under `teamgramd/etc2/` (e.g. `bff.yaml`), ensure the `Prometheus` block is present:

```yaml
Prometheus:
  Host: 0.0.0.0
  Port: 20011        # Unique per service; avoid port conflicts
  Path: /metrics
```

- **Host** / **Port**: Where the metrics HTTP server listens (often the same host as the service, with a dedicated port).
- **Path**: Endpoint path, typically `/metrics`.

The repo’s Prometheus config already defines scrape jobs for these targets; ensure `Host`/`Port` (or the advertised address in Docker) match what Prometheus scrapes.

## Prometheus configuration

- **Path**: `teamgramd/deploy/prometheus/server/prometheus.yml`
- **Docker**: Mounted in `docker-compose-env.yaml` (e.g. `./teamgramd/deploy/prometheus/server/prometheus.yml:/etc/prometheus/prometheus.yml`).

`prometheus.yml` contains:

- **global**: e.g. `scrape_interval: 15s`
- **scrape_configs**: One job per service (e.g. `bff.bff`, `service.authsession`), with `metrics_path: /metrics` and `static_configs` with the target host and port (e.g. `teamgram:20011` for bff).

When you add a new service, add a corresponding scrape job and ensure the service’s `Prometheus.Port` (and network visibility) match.

## Grafana

### 1. Start Grafana

Grafana is included in `docker-compose-env.yaml`. Start the env stack:

```bash
docker compose -f docker-compose-env.yaml up -d
```

Default URL: `http://localhost:3000` (or the port you expose). Default login is often `admin` / `admin` (check env or docs; change on first login if prompted).

### 2. Add Prometheus data source

- **Configuration** → **Data sources** → **Add data source** → **Prometheus**.
- **URL**: `http://prometheus:9090` (when running in the same Docker network) or `http://<host>:9090` if Prometheus is elsewhere.
- Save & test.

### 3. Dashboards

- You can create panels manually (PromQL queries) or import dashboards (e.g. go-zero or generic Go/Prometheus dashboards from Grafana.com).
- Use metrics such as `http_server_requests_total`, `http_server_request_duration_ms`, or any custom metrics your services expose.

## See also

- [Log collection](log-collection.md)
- [Link tracking](link-tracking.md)
