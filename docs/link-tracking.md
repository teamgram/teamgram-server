# Link Tracking (Distributed Tracing)

Teamgram uses go-zero’s built-in support for distributed tracing. You can use [Jaeger](https://www.jaegertracing.io/) or Zipkin to collect traces and analyze request flows across services.

## How it works

Each service reports spans to the tracing backend (Jaeger or Zipkin). The backend correlates spans by trace ID so you can see the full path of a request (e.g. gateway → bff → biz_service → database).

## Enabling tracing

### 1. Configure the Telemetry block

In each service config under `teamgramd/etc2/` (e.g. `bff.yaml`, `gateway.yaml`), add or adjust the `Telemetry` block:

```yaml
Telemetry:
  Name: bff.bff                    # Service name shown in traces
  Endpoint: http://jaeger:14268/api/traces
  Sampler: 1.0                     # Sampling rate (1.0 = 100%)
  Batcher: jaeger                  # or zipkin
```

- **Name**: Service name in the trace UI.
- **Endpoint**: Jaeger collector (e.g. `http://jaeger:14268/api/traces`) or Zipkin endpoint.
- **Sampler**: `1.0` = sample all requests; use a smaller value (e.g. `0.1`) in production if needed.
- **Batcher**: `jaeger` or `zipkin`.

### 2. Run Jaeger (or Zipkin)

The env stack in `docker-compose-env.yaml` includes Jaeger. Start it with:

```bash
docker compose -f docker-compose-env.yaml up -d
```

Jaeger all-in-one listens for traces (e.g. on 14268) and serves the UI (default port 16686). Open `http://localhost:16686` to search traces by service, operation, or time range.

### 3. Optional: Jaeger with Elasticsearch storage

For production you may want to store spans in Elasticsearch. Use a Jaeger deployment that points to your Elasticsearch cluster instead of the all-in-one in-memory storage. The `Telemetry.Endpoint` in service configs should point to that Jaeger collector.

## Jaeger UI

- **Search**: Select service, operation, time range, and click “Find Traces”.
- **Trace detail**: Click a trace to see the timeline and span tree (which service called which, and how long each step took).

## See also

- [Log collection](log-collection.md)
- [Service monitoring](service-monitoring.md)
