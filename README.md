# Teamgram Server

[中文](README-zh.md) | **English**

---

Unofficial open-source [MTProto](https://core.telegram.org/mtproto) server implementation in Go. Compatible with Telegram clients; supports self-hosted deployment.

## Features

- **MTProto 2.0**
  - **Abridged**
  - **Intermediate**
  - **Padded intermediate**
  - **Full**
- **API Layer: 222**
- **Core features**
  - **private chat**
  - **basic group**
  - **contacts**
  - **web**

## Architecture

![Architecture](docs/image/architecture-001.png)

- [Architecture (specs)](specs/architecture.md) — service topology, data flow, and ports  
- [Service topology and configuration](docs/service-topology.md) — ports, infrastructure dependencies, call graph, and Mermaid diagram ([中文](docs/service-topology-zh.md))

## Prerequisites

| Component | Purpose |
|-----------|---------|
| [MySQL](https://www.mysql.com/) 5.7+ / 8.0 | Primary data store |
| [Redis](https://redis.io/) | Cache, session, deduplication |
| [etcd](https://etcd.io/) | Service discovery & config |
| [Kafka](https://kafka.apache.org/) | Message & event pipeline |
| [MinIO](https://minio.io/) | Object storage |
| [FFmpeg](https://ffmpeg.org/) | Media transcoding (on server) |

Detailed versions and optional monitoring stack: [Dependencies and runtime (specs)](specs/dependencies-and-runtime.md).

---

## Manual installation

For running the server from source (Go build), follow the step-by-step guides:

- **[Manual installation (Linux)](docs/install-manual-linux.md)** — CentOS, Fedora, Ubuntu/Debian
- **[Manual installation (macOS)](docs/install-manual-macos.md)** — Intel and Apple Silicon

Requires Go 1.21+. You must install and configure dependencies (MySQL, Redis, etcd, Kafka, MinIO, FFmpeg), initialize the database and MinIO, then build and run.

---

## Docker installation

For running the full stack with Docker. **No manual data initialization:** the dependency stack initializes the database (via mounted SQL) and MinIO buckets (via `minio-mc`) on first start.

### 1. Clone

```bash
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

### 2. Start dependency stack

This starts MySQL, Redis, etcd, Kafka, MinIO (and optional monitoring). The database and MinIO buckets are initialized automatically.

```bash
docker compose -f docker-compose-env.yaml up -d
```

### 3. Start application

```bash
docker compose up -d
```

---

## Compatible clients

**Default sign-in verification code:** `12345` (change for production.)

| Platform | Repository | Patch Link |
|----------|------------|------------|
| Android | [https://github.com/teamgram/teamgram-android](https://github.com/teamgram/teamgram-android) | [teamgram-android](clients/teamgram-android.md) |
| iOS | [https://github.com/teamgram/teamgram-ios](https://github.com/teamgram/teamgram-ios) | [teamgram-ios](clients/teamgram-ios.md) |
| Desktop (TDesktop) | [https://github.com/teamgram/teamgram-tdesktop](https://github.com/teamgram/teamgram-tdesktop) | [teamgram-tdesktop](clients/teamgram-tdesktop.md) |

---

## Documentation

- [Project specs](specs/README.md) — Architecture, protocol, dependencies, contributing, security, roadmap
- [Service topology and configuration](docs/service-topology.md) — Ports, infrastructure, call graph ([中文](docs/service-topology-zh.md))
- [CONTRIBUTING](CONTRIBUTING.md) · [SECURITY](SECURITY.md) · [CHANGELOG](CHANGELOG.md)

---

## Community & feedback

- **Issues:** bugs and feature requests
- **Telegram:** [Teamgram group](https://t.me/+TjD5LZJ5XLRlCYLF)

---

## Enterprise edition

The following are available in the enterprise edition (contact the [author](https://t.me/benqi)):

- sticker/theme/chat_theme/wallpaper/reactions/secretchat/2fa/sms/push(apns/web/fcm)/web/scheduled/autodelete/... 
- channels/megagroups
- audio/video/group/conferenceCall
- bots
- miniapp

See [specs/roadmap.md](specs/roadmap.md) for community vs. enterprise scope.

---

## License

[Apache License 2.0](LICENSE).

---

## Give a Star! ⭐

If this project helps you, consider giving it a star.
