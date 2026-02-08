# 依赖与运行环境

本文档说明 Teamgram Server 的核心依赖、推荐版本及 Docker 环境栈。旧版组件列表见 [high-performance-components.md](high-performance-components.md)（重定向）。

## 核心依赖（必须）

| 组件 | 用途 | 推荐版本 |
|------|------|----------|
| **MySQL** | 主数据存储 | 5.7 或 8.0（Docker 栈使用 8.0） |
| **Redis** | 缓存、会话、去重 | 6.x / 7.x |
| **etcd** | 服务发现与配置 | v3.5.x |
| **Kafka** | 消息与事件管道 | 3.x（KRaft，无需 Zookeeper） |
| **MinIO** | 对象存储（文档、图片、视频等） | 当前 stable |
| **FFmpeg** | 媒体转码（服务端安装） | 见各安装文档 |

- 创建数据库 `teamgram`，并按顺序执行 **`teamgramd/deploy/sql/`** 下所有初始化与 migrate 脚本。
- MinIO：创建桶 `documents`、`encryptedfiles`、`photos`、`videos`（使用仓库提供的 Docker 栈时由 minio-mc 自动创建）。

## 可选：监控与日志

`docker-compose-env.yaml` 中可选的组件包括：

| 组件 | 用途 |
|------|------|
| **Jaeger** | 分布式追踪 |
| **Prometheus** | 指标采集 |
| **Grafana** | 仪表盘 |
| **Node Exporter** | 主机指标 |
| **Elasticsearch** | 日志存储与搜索 |
| **Kibana** | 日志/数据可视化 |
| **Filebeat** | 日志采集（如 teamgram 日志 → Kafka） |
| **go-stash** | 日志管道（Kafka → Elasticsearch） |

详见主 README 的「日志、监控与追踪」及 `teamgramd/deploy/` 配置。

## 版本与 Docker 镜像（参考）

与 `docker-compose-env.yaml` 及 [README-env-cn.md](../README-env-cn.md) / [README-env-en.md](../README-env-en.md) 一致：

- **Kafka**: bitnamilegacy/kafka:3.5.1（KRaft）
- **etcd**: quay.io/coreos/etcd:v3.5.11
- **Redis**: redis:7-alpine
- **MySQL**: mysql:8.0
- **MinIO**: minio/minio:latest；minio-mc 用于 bucket 初始化
- **Jaeger**: jaegertracing/all-in-one:1.52
- **Prometheus**: prom/prometheus:v2.47.2
- **Grafana**: grafana/grafana:10.2.3
- **Node Exporter**: prom/node-exporter:v1.7.0
- **Elasticsearch / Kibana / Filebeat**: 8.11.x
- **go-stash**: kevinwan/go-stash:1.1.1

## Docker 环境栈

- **Compose 文件**：**`docker-compose-env.yaml`**（部分旧文档中称为「docker-compose-env2」）。
- **使用**：复制 `.env.example` 为 `.env`，执行 `docker compose -f docker-compose-env.yaml up -d`。详见 README-env-cn.md / README-env-en.md。
- **网络**：`teamgram_net`；数据在项目 `data/` 目录下。

## 无 Docker 时的安装

自行安装组件时参考：

- [手动安装（Linux）](../docs/install-manual-linux-zh.md)
- [手动安装（macOS）](../docs/install-manual-macos-zh.md)

确保 MySQL、Redis、etcd、Kafka、MinIO、FFmpeg 可用后，执行 `make` 与 `teamgramd/bin/runall2.sh`。
