# Architecture and data flow

This document describes Teamgram Server’s service topology, request path, and how it maps to MTProto/API. For a detailed port table and Mermaid diagram, see [Service topology and configuration](../docs/service-topology.md).

## Architecture diagram

High-level diagram in the project README:

![Architecture](../docs/image/architecture-001.png)

## Service list and roles

| Service | Default RPC port | Client-facing (example) | Role |
|---------|------------------|--------------------------|------|
| **gnetway** | 20110 | 10443 (TCP MTProto), 11443 (WebSocket), 5222 (TCP) | Protocol gateway: accepts client MTProto (TCP/WebSocket), forwards to session |
| **session** | 20120 | — | Session layer: connection state, auth, routes RPC to BFF |
| **bff** | 20010 | — | Backend-for-frontend: MTProto RPC → backend gRPC calls |
| **authsession** | 20450 | — | Auth session: login state, session validation |
| **biz** | 20020 | — | Core business: user, chat, dialog, message, updates gRPC |
| **msg** | 20030 | — | Message service: storage, delivery, inbox |
| **sync** | 20420 | — | Sync: multi-device updates |
| **dfs** | 20640 | 11701 (HTTP optional) | Distributed file: upload/download routing, storage |
| **media** | 20650 | — | Media: image/document/video metadata and processing |
| **idgen** | 20660 | — | Distributed ID generation |
| **status** | 20670 | — | User online status |
| **httpserver** | 8801 | 8801 (HTTP) | Optional: HTTP API or web callbacks |

Discovery is via **etcd**; config files live under `teamgramd/etc/` (or `teamgramd/etc2/` for Docker). Startup order: see `teamgramd/bin/runall2.sh` or `teamgramd/bin/runall-docker.sh`.

## Request path (summary)

Client → gnetway (MTProto) → session → BFF → biz / msg / dfs / media / sync; messaging and events use Kafka.

## Data and storage

| Component | Purpose |
|-----------|---------|
| **MySQL** | Business data; run init and migrate scripts under `teamgramd/deploy/sql/` |
| **Redis** | Cache, session, deduplication |
| **Kafka** | Message and event pipeline (msg, sync, inbox) |
| **MinIO** | Object storage: buckets `documents`, `encryptedfiles`, `photos`, `videos` |
| **etcd** | Service discovery and config |

## Ports and config

- Client-facing ports are those exposed by **gnetway** (default 10443, 11443, 5222). Other ports are internal RPC/HTTP.
- YAML configs: `teamgramd/etc/` (or `teamgramd/etc2/` for Docker). Binaries: `make` → `teamgramd/bin/`.
