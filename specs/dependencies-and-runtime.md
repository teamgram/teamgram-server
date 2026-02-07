# 依赖与运行环境

本文档说明 Teamgram Server 的核心依赖、推荐版本及 Docker 一键环境栈。原 [high-performance-components.md](high-performance-components.md) 中的组件列表已合并到本文档。

## 核心依赖（必须）

运行 Teamgram 前需具备以下组件：

| 组件 | 用途 | 推荐版本 |
|------|------|----------|
| **MySQL** | 业务数据存储 | 5.7 或 8.0（当前 Docker 栈使用 8.0） |
| **Redis** | 缓存、会话、去重等 | 6.x / 7.x |
| **etcd** | 服务发现与配置 | v3.5.x |
| **Kafka** | 消息与事件管道 | 3.x（KRaft 模式无需 Zookeeper） |
| **MinIO** | 对象存储（文档、图片、视频等） | 当前 stable |
| **ffmpeg** | 媒体转码与处理（服务端需安装） | 系统安装，见各安装文档 |

- 数据库需先创建库 `teamgram`，并执行 `teamgramd/sql/` 下所有初始化与 migrate 脚本。
- MinIO 需创建 bucket：`documents`、`encryptedfiles`、`photos`、`videos`（使用本仓库提供的 docker-compose 环境时可由 minio-mc 自动创建）。

## 可选：监控与日志栈

若需可观测性，可一并部署以下组件（与仓库提供的 Docker 环境栈一致）：

| 组件 | 用途 |
|------|------|
| **Jaeger** | 分布式追踪 |
| **Prometheus** | 指标采集 |
| **Grafana** | 仪表盘与可视化 |
| **Alertmanager** | 告警（可配入 Prometheus） |
| **Node Exporter** | 主机指标 |
| **Elasticsearch** | 日志存储与搜索 |
| **Kibana** | 日志/数据可视化 |
| **Filebeat** | 日志采集 |
| **go-stash** | 日志管道（Kafka → Elasticsearch） |

## 版本与 Docker 镜像（参考）

与当前 `docker-compose-env.yaml` 及 [README-env-cn.md](../README-env-cn.md) / [README-env-en.md](../README-env-en.md) 对齐的版本示例：

- **Kafka**: bitnamilegacy/kafka:3.5.1（KRaft）
- **etcd**: quay.io/coreos/etcd:v3.5.11
- **Redis**: redis:7-alpine
- **MySQL**: mysql:8.0
- **MinIO**: minio/minio:latest；minio-mc 用于初始化 bucket
- **Jaeger**: jaegertracing/all-in-one:1.52
- **Prometheus**: prom/prometheus:v2.47.2
- **Grafana**: grafana/grafana:10.2.3
- **Node Exporter**: prom/node-exporter:v1.7.0
- **Elasticsearch / Kibana / Filebeat**: 8.11.x
- **go-stash**: kevinwan/go-stash:1.1.1

## Docker 一键环境

- **环境栈文件**：仓库内文件名为 **`docker-compose-env.yaml`**。部分文档中出现的「docker-compose-env2」即指该环境栈（历史命名）。
- **使用方式**：
  - 复制 `.env.example` 为 `.env`，按需修改数据库、MinIO、Grafana 等密码与配置。
  - 启动依赖：`docker compose -f docker-compose-env.yaml up -d`
  - 详见 [README-env-cn.md](../README-env-cn.md)（中文）、[README-env-en.md](../README-env-en.md)（英文）。
- **网络与数据**：网络名为 `teamgram_net`；数据持久化在项目下 `data/` 目录，各服务子目录见 README-env-cn.md。

## 无 Docker 时的安装

若已有或希望自行安装上述组件，可参考：

- [CentOS 9 Stream 构建与安装](../docs/install-centos-9.md)
- [CentOS 7 环境搭建](../docs/install-centos-7.md)
- [Fedora 40 构建与安装](../docs/install-fedora.md)

安装完成后，确保 MySQL、Redis、etcd、Kafka、MinIO 可访问，且服务端机器已安装 ffmpeg，再执行 `make` 与 `teamgramd/bin/runall2.sh` 启动各服务。
