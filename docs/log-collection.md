# Log Collection

Teamgram uses [Filebeat](https://www.elastic.co/beats/filebeat) to collect logs, [Kafka](https://kafka.apache.org/) as a buffer, [go-stash](https://github.com/kevwan/go-stash) to consume and process logs (instead of resource-heavy Logstash), and [Elasticsearch](https://www.elastic.co/elasticsearch/) + [Kibana](https://www.elastic.co/kibana/) for storage and visualization.

## Architecture

![Log pipeline architecture](image/log-pipeline-architecture.png)

```
┌─────────────┐     ┌─────────┐     ┌──────────┐     ┌─────────────────┐     ┌────────┐
│  App Logs   │────▶│ Filebeat│────▶│  Kafka   │────▶│    go-stash     │────▶│   ES   │
│ (container) │     │         │     │(teamgram-│     │ (filter/transform)│    │        │
└─────────────┘     └─────────┘     │  log)    │     └─────────────────┘     └───┬────┘
                                    └─────────┘                                  │
                                                                                 ▼
                                                                           ┌─────────┐
                                                                           │ Kibana  │
                                                                           └─────────┘
```

- **Filebeat**: Collects container logs (e.g. from Docker) and sends them to Kafka.
- **Kafka**: Topic `teamgram-log` buffers log events for decoupling and back-pressure.
- **go-stash**: Consumes from Kafka, applies filters (e.g. drop debug, clean fields), and writes to Elasticsearch.
- **Elasticsearch**: Stores logs; index pattern `teamgram-{{yyyy-MM-dd}}`.
- **Kibana**: Query and visualize logs.

## Configuration

### Filebeat

- **Path**: `teamgramd/deploy/filebeat/conf/filebeat.yml`
- **Docker**: Mounted in `docker-compose-env.yaml`; ensure the Docker containers log path (e.g. `/var/lib/docker/containers`) is correct for your host.

Main settings:

- `filebeat.inputs`: `type: container`, paths to container logs, JSON parsing.
- `output.kafka`: `hosts`, `topic: teamgram-log`, compression, etc.

### go-stash

- **Path**: `teamgramd/deploy/go-stash/etc/config.yaml`
- **Docker**: Config directory `teamgramd/deploy/go-stash/etc` is mounted into the go-stash container.

Main settings:

- **Input.Kafka**: Brokers, topic `teamgram-log`, consumer group, concurrency.
- **Filters**: e.g. drop debug logs or certain containers, remove unneeded fields, map `message` to `data`.
- **Output.ElasticSearch**: ES hosts, index `teamgram-{{yyyy-MM-dd}}`.

### Elasticsearch & Kibana

Started via `docker-compose-env.yaml`. Default ports (when exposed): Elasticsearch 9200, Kibana 5601.

## Kibana

1. Open Kibana (e.g. `http://localhost:5601`).
2. **Stack Management** → **Index Patterns** (or **Index Management** in newer versions).
3. Create index pattern: `teamgram-*`, time field e.g. `@timestamp`.
4. Use **Discover** to search and filter logs by time, level, service, message, etc.

## Troubleshooting

- **No logs in Kibana**: Check that Filebeat and go-stash containers are running, Kafka topic `teamgram-log` has messages, and Elasticsearch has indices `teamgram-*`. Check go-stash and Filebeat logs for errors.
- **Docker log path**: Filebeat must see container log files; the path in `filebeat.yml` must match the Docker daemon’s `containers` directory (e.g. `/var/lib/docker/containers`). Adjust the volume in `docker-compose-env.yaml` if your Docker root is different.
- **go-stash image**: Use the image specified in `docker-compose-env.yaml` (e.g. `kevinwan/go-stash:1.1.1`). If the config format differs, check [go-stash](https://github.com/kevwan/go-stash) documentation.

## See also

- [Link tracking](link-tracking.md)
- [Service monitoring](service-monitoring.md)
