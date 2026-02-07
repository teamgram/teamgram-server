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
- **API Layer: 220**
- **Core features**
  - **private chat**
  - **basic group**
  - **contacts**
  - **web**

## Architecture

![Architecture](docs/image/architecture-001.png)

For service topology, data flow, and ports, see [Architecture (specs)](specs/architecture.md).

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

**Installation guides (no Docker):**

- [CentOS 9 Stream](docs/install-centos-9.md)
- [CentOS 7](docs/install-centos-7.md)
- [Fedora 40](docs/install-fedora.md)

**One-shot environment (Docker):** use [docker-compose-env.yaml](docker-compose-env.yaml). See [README-env-cn.md](README-env-cn.md) / [README-env-en.md](README-env-en.md).

---

## Manual installation

For running the server from source (Go build). You must install and configure dependencies, then initialize the database and MinIO yourself.

**Requires Go 1.21+.**

### 1. Clone

```bash
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

### 2. Dependencies

Install MySQL, Redis, etcd, Kafka, MinIO, and FFmpeg (see installation guides in Prerequisites above), or start only the dependency stack with Docker:

```bash
docker compose -f docker-compose-env.yaml up -d
```

### 3. Initialize data

**Database**

1. Create database: `teamgram`
2. Run SQL in order (from repo root):

```bash
for f in teamgramd/deploy/sql/1_teamgram.sql teamgramd/deploy/sql/migrate-*.sql teamgramd/deploy/sql/z_init.sql; do
  mysql -uroot teamgram < "$f"
done
```

Or run each file manually; see [teamgramd/deploy/sql/README.md](teamgramd/deploy/sql/README.md).

**MinIO**

Create buckets: `documents`, `encryptedfiles`, `photos`, `videos` (e.g. via MinIO Console at `http://<host>:9001`). If you started the env stack with Docker, the `minio-mc` service creates them automatically.

### 4. Build & run

```bash
make
cd teamgramd/bin
./runall2.sh
```

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

## Star & visitors

If this project helps you, consider giving it a star.

[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fteamgram%2Fteamgram-server&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=visitors&edge_flat=false)](https://hits.seeyoufarm.com)
