# 架构与数据流

本文档描述 Teamgram Server 的服务拓扑、请求路径及与 MTProto/API 的对应关系。更详细的端口表与 Mermaid 图见 [服务拓扑与配置](../docs/service-topology-zh.md)。

## 架构图

项目 README 中的整体架构图：

![Architecture](../docs/image/architecture-001.png)

## 服务列表与职责

| 服务 | 默认 RPC 端口 | 对客户端暴露（示例） | 职责 |
|------|----------------|----------------------|------|
| **gnetway** | 20110 | 10443 (TCP MTProto)、11443 (WebSocket)、5222 (TCP) | 协议网关：接收客户端 MTProto，转发到 session |
| **session** | 20120 | — | 会话层：连接状态、鉴权、将 RPC 路由到 BFF |
| **bff** | 20010 | — | 聚合层：MTProto RPC 转后端 gRPC |
| **authsession** | 20450 | — | 鉴权会话：登录态、会话校验 |
| **biz** | 20020 | — | 业务核心：用户、聊天、对话、消息、更新等 gRPC |
| **msg** | 20030 | — | 消息服务：存储、投递、inbox |
| **sync** | 20420 | — | 同步：多端更新 |
| **dfs** | 20640 | 11701 (HTTP 可选) | 分布式文件：上传/下载路由与存储 |
| **media** | 20650 | — | 媒体：图片/文档/视频元数据与处理 |
| **idgen** | 20660 | — | 分布式 ID 生成 |
| **status** | 20670 | — | 用户在线状态 |
| **httpserver** | 8801 | 8801 (HTTP) | 可选：HTTP API 或 Web 回调 |

发现与配置通过 **etcd**；配置文件在 `teamgramd/etc/`（Docker 可用 `teamgramd/etc2/`）。启动顺序见 `teamgramd/bin/runall2.sh` 或 `teamgramd/bin/runall-docker.sh`。

## 请求路径（概要）

客户端 → gnetway (MTProto) → session → BFF → biz / msg / dfs / media / sync；消息与事件经 Kafka。

## 数据与存储

| 组件 | 用途 |
|------|------|
| **MySQL** | 业务数据；执行 `teamgramd/deploy/sql/` 下初始化与 migrate 脚本 |
| **Redis** | 缓存、会话、去重 |
| **Kafka** | 消息与事件管道（msg、sync、inbox） |
| **MinIO** | 对象存储：桶 `documents`、`encryptedfiles`、`photos`、`videos` |
| **etcd** | 服务发现与配置 |

## 端口与配置

- 对客户端暴露的端口以 **gnetway** 为准（默认 10443、11443、5222）。其余为内部 RPC/HTTP。
- YAML 配置：`teamgramd/etc/` 或 `teamgramd/etc2/`；二进制：`make` → `teamgramd/bin/`。
