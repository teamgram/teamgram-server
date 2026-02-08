# Dependencies and runtime

This document covers core dependencies, recommended versions, and the Docker stack for Teamgram Server. The legacy component list is in [high-performance-components.md](high-performance-components.md) (redirect).

## Core dependencies (required)

| Component | Purpose | Recommended version |
|-----------|---------|---------------------|
| **MySQL** | Primary data store | 5.7 or 8.0 (Docker stack uses 8.0) |
| **Redis** | Cache, session, deduplication | 6.x / 7.x |
| **etcd** | Service discovery and config | v3.5.x |
| **Kafka** | Message and event pipeline | 3.x (KRaft, no Zookeeper) |
| **MinIO** | Object storage (documents, photos, videos) | Current stable |
| **FFmpeg** | Media transcoding (install on server) | Per install docs |

- Create database `teamgram` and run all init and migrate scripts under **`teamgramd/deploy/sql/`** in order.
- MinIO: create buckets `documents`, `encryptedfiles`, `photos`, `videos` (auto-created by minio-mc when using the provided Docker stack).

## Optional: monitoring and logging

The Docker stack in `docker-compose-env.yaml` can include:

| Component | Purpose |
|-----------|---------|
| **Jaeger** | Distributed tracing |
| **Prometheus** | Metrics |
| **Grafana** | Dashboards |
| **Node Exporter** | Host metrics |
| **Elasticsearch** | Log storage and search |
| **Kibana** | Log/data UI |
| **Filebeat** | Log collection (e.g. teamgram logs → Kafka) |
| **go-stash** | Log pipeline (Kafka → Elasticsearch) |

See the main README section “Logging, monitoring & tracing” and `teamgramd/deploy/` for config.

## Versions and Docker images (reference)

Aligned with `docker-compose-env.yaml` and [README-env-cn.md](../README-env-cn.md) / [README-env-en.md](../README-env-en.md):

- **Kafka**: bitnamilegacy/kafka:3.5.1 (KRaft)
- **etcd**: quay.io/coreos/etcd:v3.5.11
- **Redis**: redis:7-alpine
- **MySQL**: mysql:8.0
- **MinIO**: minio/minio:latest; minio-mc for bucket init
- **Jaeger**: jaegertracing/all-in-one:1.52
- **Prometheus**: prom/prometheus:v2.47.2
- **Grafana**: grafana/grafana:10.2.3
- **Node Exporter**: prom/node-exporter:v1.7.0
- **Elasticsearch / Kibana / Filebeat**: 8.11.x
- **go-stash**: kevinwan/go-stash:1.1.1

## Docker stack

- **Compose file**: **`docker-compose-env.yaml`** (sometimes referred to as “docker-compose-env2” in older docs).
- **Usage**: Copy `.env.example` to `.env`, then `docker compose -f docker-compose-env.yaml up -d`. See README-env-cn.md / README-env-en.md.
- **Network**: `teamgram_net`; data under project `data/` directory.

## Installation without Docker

If you install components yourself, see:

- [Manual installation (Linux)](../docs/install-manual-linux.md)
- [Manual installation (macOS)](../docs/install-manual-macos.md)

Then ensure MySQL, Redis, etcd, Kafka, MinIO, and FFmpeg are available and run `make` and `teamgramd/bin/runall2.sh`.
