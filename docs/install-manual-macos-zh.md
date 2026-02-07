# Teamgram 服务端 macOS 手动安装指南

本文档介绍在 macOS（Intel 与 Apple Silicon）上从零手动安装 Teamgram 服务端及其依赖。推荐使用 [Homebrew](https://brew.sh/) 管理依赖。

> [English (primary)](./install-manual-macos.md) | 中文  
> 若使用 Docker 部署，请参阅 [install-docker.md](./install-docker.md)。

---

## 一、环境要求

- **macOS**：10.15 或更高（建议 11+，Apple Silicon 已支持）
- **Go**：1.21 及以上
- **MySQL**：8.x
- **Redis**：6.x
- **Etcd**：3.5.x
- **Kafka**：2.x / 3.x（需 Zookeeper，Homebrew 版 Kafka 通常自带）
- **MinIO**：对象存储
- **Pika**（可选）：需从源码或 release 安装，默认端口 9221
- **FFmpeg**：音视频处理

请先安装 Homebrew（如未安装）：

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

---

## 二、安装依赖

### 2.1 安装 MySQL

```bash
brew install mysql
brew services start mysql
```

**配置并创建数据库：**

```bash
# 首次安装可能需设置 root 密码，按提示操作
mysql_secure_installation

mysql -uroot -p

mysql> CREATE DATABASE teamgram;
mysql> UPDATE mysql.user SET authentication_string='' WHERE user='root';
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY '';
mysql> FLUSH PRIVILEGES;
mysql> exit
```

### 2.2 安装 Redis

```bash
brew install redis
brew services start redis
```

默认监听 `127.0.0.1:6379`。

### 2.3 安装 Etcd

```bash
brew install etcd
brew services start etcd
```

默认监听 `127.0.0.1:2379`。

### 2.4 安装 Kafka

Kafka 依赖 Java，Homebrew 会拉取 OpenJDK：

```bash
brew install kafka
```

**启动 Zookeeper 与 Kafka：**

```bash
# 先启动 Zookeeper（后台）
zookeeper-server-start -daemon $(brew --prefix kafka)/libexec/config/zookeeper.properties
sleep 2

# 再启动 Kafka
kafka-server-start -daemon $(brew --prefix kafka)/libexec/config/server.properties
```

或使用 brew services（若 formula 支持）：

```bash
brew services start zookeeper
brew services start kafka
```

**检查进程：**

```bash
ps aux | grep kafka | wc -l
# 约 3 个进程表示正常
```

### 2.5 安装 MinIO

```bash
brew install minio
```

**启动 MinIO（指定数据目录）：**

```bash
mkdir -p /tmp/minio-data
minio server /tmp/minio-data &
```

或使用 brew services：

```bash
brew services start minio
```

- 控制台：http://localhost:9001  
- 默认账号：minioadmin / minioadmin（或 miniostorage，视版本而定）

**创建存储桶：** 在 MinIO 控制台创建以下 bucket：

- documents  
- encryptedfiles  
- photos  
- videos  

### 2.6 安装 Pika（可选，端口 9221）

Homebrew 可能不提供 Pika，需从 GitHub 下载 macOS 版（若有）或从源码编译：

```bash
cd /tmp
# 若有 macOS 预编译包，例如：
# curl -L -o pika.tar.bz2 https://github.com/OpenAtomFoundation/pika/releases/download/v3.3.6/pika-darwin-x86_64-v3.3.6.tar.bz2
# 解压后运行：
# /tmp/pika/bin/pika -c /tmp/pika/conf/pika.conf &
```

若无可用的 macOS 预编译包，可暂时不使用 Pika，在配置中使用 Redis 的 SSDB 兼容配置（见下文配置说明）。

### 2.7 安装 FFmpeg

```bash
brew install ffmpeg
```

### 2.8 安装 Go

```bash
brew install go
go version   # 需 1.21+
```

或从 [Go 官网](https://go.dev/dl/) 安装。

---

## 三、获取源码并初始化数据库

### 3.1 克隆代码

```bash
cd ~  # 或你希望放置的目录
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

### 3.2 初始化数据库

在项目根目录执行（SQL 位于 `teamgramd/deploy/sql/`）：

```bash
mysql -uroot -e "CREATE DATABASE IF NOT EXISTS teamgram;"

mysql -uroot teamgram < teamgramd/deploy/sql/1_teamgram.sql
for f in teamgramd/deploy/sql/migrate-*.sql; do mysql -uroot teamgram < "$f"; done
mysql -uroot teamgram < teamgramd/deploy/sql/z_init.sql
```

---

## 四、编译 Teamgram 服务端

在项目根目录执行：

```bash
make
```

编译产物在 `teamgramd/bin/` 下。

---

## 五、修改配置文件

配置文件在 `teamgramd/etc/` 下。macOS 本机部署时，各服务通常为 `127.0.0.1`，需确认以下项：

1. **dfs.yaml**  
   - `Minio.Endpoint`：`localhost:9000`  
   - `Minio.AccessKeyID` / `SecretAccessKey`：与 MinIO 一致（如 minioadmin / minioadmin）  
   - `SSDB`：使用 Pika 则为 `127.0.0.1:9221`，否则用 Redis `127.0.0.1:6379`

2. **gnetway.yaml**  
   - `Gnetway.Server.Addresses`：如 `0.0.0.0:10443`、`0.0.0.0:5222`

3. **所有 YAML**  
   - `Etcd.Hosts`：`127.0.0.1:2379`  
   - MySQL：`127.0.0.1:3306`  
   - Redis：`127.0.0.1:6379`  
   - Kafka：`127.0.0.1:9092`

示例（dfs.yaml 片段）：

```yaml
Minio:
  Endpoint: localhost:9000
  AccessKeyID: minioadmin
  SecretAccessKey: minioadmin
  UseSSL: false
SSDB:
  - Host: 127.0.0.1:6379   # 未装 Pika 时用 Redis
```

---

## 六、启动服务

确认 MySQL、Redis、Etcd、Kafka（及 Zookeeper）、MinIO 均已启动后：

```bash
cd teamgramd/bin
./runall2.sh
```

如需从局域网访问，可在 `gnetway.yaml` 中确认监听 `0.0.0.0`，并确保本机防火墙允许 10443、5222 等端口。

---

## 七、停止服务

在 `teamgramd/bin` 下执行：

```bash
./killall.sh
```

停止依赖服务（按需）：

```bash
brew services stop mysql
brew services stop redis
brew services stop etcd
# Kafka / Zookeeper 需手动 kill 或通过 brew services 停止
```

---

## 八、常见问题

- **Apple Silicon (M1/M2/M3)**：Go、MySQL、Redis、Etcd、Kafka、MinIO 的 Homebrew 版均支持 ARM，一般无需额外配置。  
- **Kafka 路径**：使用 `brew --prefix kafka` 查看安装路径，配置文件在其 `libexec/config/` 下。  
- **MinIO 数据目录**：生产环境请将 `/tmp/minio-data` 改为持久化目录。  
- **Pika**：若无 macOS 预编译包，可不装 Pika，配置中 SSDB 使用 Redis 即可。  
- **端口占用**：若 3306、6379、2379、9092、9000 等被占用，需修改对应服务与 teamgram 配置中的端口一致。

完成以上步骤后，Teamgram 服务端即可在 macOS 上以手动方式运行。
