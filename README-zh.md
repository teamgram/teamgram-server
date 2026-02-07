# Teamgram Server

**中文** | [English](README.md)

---

基于 Go 实现的开源 [MTProto](https://core.telegram.org/mtproto) 服务端，非官方实现。兼容 Telegram 客户端，支持私有化部署。

## 功能特性

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

## 架构

![架构图](docs/image/architecture-001.png)

服务拓扑、数据流与端口说明见 [架构说明（specs）](specs/architecture.md)。

## 前置依赖

| 组件 | 用途 |
|------|------|
| [MySQL](https://www.mysql.com/) 5.7+ / 8.0 | 主数据存储 |
| [Redis](https://redis.io/) | 缓存、会话、去重 |
| [etcd](https://etcd.io/) | 服务发现与配置 |
| [Kafka](https://kafka.apache.org/) | 消息与事件管道 |
| [MinIO](https://minio.io/) | 对象存储 |
| [FFmpeg](https://ffmpeg.org/) | 媒体转码（需在服务端安装） |

版本建议与可选监控栈详见 [依赖与运行环境（specs）](specs/dependencies-and-runtime.md)。

**无 Docker 时的安装文档：**

- [CentOS 9 Stream](docs/install-centos-9.md)
- [CentOS 7](docs/install-centos-7.md)
- [Fedora 40](docs/install-fedora.md)

**一键环境（Docker）：** 使用 [docker-compose-env.yaml](docker-compose-env.yaml)，详见 [README-env-cn.md](README-env-cn.md) / [README-env-en.md](README-env-en.md)。

---

## 手动安装

从源码构建并运行服务。需自行安装并配置依赖，并**手动初始化数据库与 MinIO**。

**需要 Go 1.21+。**

### 1. 克隆仓库

```bash
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

### 2. 依赖环境

安装 MySQL、Redis、etcd、Kafka、MinIO、FFmpeg（见上文安装文档），或仅用 Docker 启动依赖栈：

```bash
docker compose -f docker-compose-env.yaml up -d
```

### 3. 初始化数据

**数据库**

1. 创建数据库：`teamgram`
2. 按顺序执行 SQL（在仓库根目录）：

```bash
for f in teamgramd/deploy/sql/1_teamgram.sql teamgramd/deploy/sql/migrate-*.sql teamgramd/deploy/sql/z_init.sql; do
  mysql -uroot teamgram < "$f"
done
```

也可逐条执行；参见 [teamgramd/deploy/sql/README.md](teamgramd/deploy/sql/README.md)。

**MinIO**

创建桶：`documents`、`encryptedfiles`、`photos`、`videos`（可通过 MinIO 控制台 `http://<主机>:9001`）。若已用 Docker 启动环境栈，`minio-mc` 服务会自动创建。

### 4. 构建与运行

```bash
make
cd teamgramd/bin
./runall2.sh
```

---

## Docker 安装

使用 Docker 一键运行完整栈。**无需手动初始化数据**：依赖栈首次启动时会自动初始化数据库（挂载 SQL）和 MinIO 桶（通过 `minio-mc`）。

### 1. 克隆仓库

```bash
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

### 2. 启动依赖栈

将启动 MySQL、Redis、etcd、Kafka、MinIO 及可选监控组件。数据库与 MinIO 桶会自动完成初始化。

```bash
docker compose -f docker-compose-env.yaml up -d
```

### 3. 启动应用

```bash
docker compose up -d
```

---

## 兼容客户端

**默认登录验证码：** `12345`（生产环境请修改。）

| Platform | Repository | Patch Link |
|----------|------------|------------|
| Android | [https://github.com/teamgram/teamgram-android](https://github.com/teamgram/teamgram-android) | [teamgram-android](clients/teamgram-android.md) |
| iOS | [https://github.com/teamgram/teamgram-ios](https://github.com/teamgram/teamgram-ios) | [teamgram-ios](clients/teamgram-ios.md) |
| Desktop (TDesktop) | [https://github.com/teamgram/teamgram-tdesktop](https://github.com/teamgram/teamgram-tdesktop) | [teamgram-tdesktop](clients/teamgram-tdesktop.md) |

---

## 文档

- [项目规范与设计文档（specs）](specs/README.md) — 架构、协议、依赖、贡献、安全、路线图
- [CONTRIBUTING](CONTRIBUTING.md) · [SECURITY](SECURITY.md) · [CHANGELOG](CHANGELOG.md)

---

## 社区与反馈

- **Issues：** 缺陷与功能建议
- **Telegram：** [Teamgram 群组](https://t.me/+TjD5LZJ5XLRlCYLF)

---

## 企业版

以下能力在企业版中提供，请联系[作者](https://t.me/benqi)：

- sticker/theme/chat_theme/wallpaper/reactions/secretchat/2fa/sms/push(apns/web/fcm)/web/scheduled/autodelete/... 
- channels/megagroups
- audio/video/group/conferenceCall
- bots
- miniapp

社区版与企业版边界见 [specs/roadmap.md](specs/roadmap.md)。

---

## 许可证

[Apache License 2.0](LICENSE)。

---

## Star 与访问统计

若本项目对你有帮助，欢迎 Star。

[![Hits](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fteamgram%2Fteamgram-server&count_bg=%2379C83D&title_bg=%23555555&icon=&icon_color=%23E7E7E7&title=visitors&edge_flat=false)](https://hits.seeyoufarm.com)
