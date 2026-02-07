# Teamgram 服务端 Linux 手动安装指南

本文档介绍在 Linux 上从零手动安装 Teamgram 服务端及其依赖，适用于 CentOS 7/9、Fedora、Ubuntu/Debian 等发行版。不同发行版包管理器不同，请按说明替换为对应命令。

> [English (primary)](./install-manual-linux.md) | 中文  
> 若使用 Docker 部署，请参阅 [install-docker.md](./install-docker.md)。

---

## 一、环境要求

- **Go**：1.21 及以上（用于编译 teamgram-server）
- **MySQL**：8.x（推荐 8.0.29）
- **Redis**：6.x
- **Etcd**：3.5.x
- **Kafka**：2.x / 3.x（需配合 Zookeeper）
- **MinIO**：对象存储
- **Pika**（可选）：默认端口 9221，可替代部分 Redis 能力
- **FFmpeg**：音视频处理

以下操作建议使用 root 或具备 sudo 权限的账号执行。

---

## 二、安装依赖

### 2.1 安装 DNF（仅 CentOS 7 需要）

CentOS 7 默认仅有 yum，可先安装 dnf 便于与后续文档一致：

```bash
yum install epel-release -y
yum install dnf -y
dnf --version
```

### 2.2 安装 MySQL

**CentOS 9 / Fedora：**

```bash
dnf install mysql-server -y
systemctl enable --now mysqld
```

**Fedora（社区版 MySQL）：**

```bash
dnf install community-mysql-server -y
systemctl enable --now mysqld
```

**Ubuntu/Debian：**

```bash
apt update
apt install mysql-server -y
systemctl enable --now mysql
```

**配置 MySQL：**

```bash
# 安全初始化（可选）
mysql_secure_installation

# 登录并创建数据库、配置空密码（按需修改）
mysql -uroot -p

mysql> CREATE DATABASE teamgram;
mysql> UPDATE mysql.user SET authentication_string='' WHERE user='root';
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY '';
mysql> FLUSH PRIVILEGES;
mysql> exit
```

### 2.3 安装 Redis

**DNF（CentOS/Fedora）：**

```bash
dnf install redis -y
systemctl enable --now redis
```

**APT（Ubuntu/Debian）：**

```bash
apt install redis-server -y
systemctl enable --now redis-server
```

### 2.4 安装 Etcd

Etcd 需从官方发布包安装，与发行版无关：

```bash
ETCD_VER=v3.5.0
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GOOGLE_URL}

mkdir -p /tmp/etcd-download-test
curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

# 验证
/tmp/etcd-download-test/etcd --version
/tmp/etcd-download-test/etcdctl version
```

**启动 Etcd（前台或后台）：**

```bash
# 后台运行
/tmp/etcd-download-test/etcd &
```

若需长期使用，可将二进制拷到 `/usr/local/bin` 并配置 systemd 服务。

### 2.5 安装 Apache Kafka

需先安装 Java，再安装 Kafka（自带 Zookeeper）。

**安装 Java：**

```bash
# CentOS/Fedora
dnf install java-21-openjdk -y
# 或
dnf install java -y

# Ubuntu/Debian
apt install openjdk-17-jdk -y
```

**安装并启动 Kafka：**

```bash
cd /tmp
# 以 3.2.0 为例，可从 https://kafka.apache.org/downloads 获取最新版
wget https://dlcdn.apache.org/kafka/3.2.0/kafka_2.12-3.2.0.tgz
tar -xzf kafka_2.12-3.2.0.tgz
cd kafka_2.12-3.2.0

# 启动 Zookeeper
bin/zookeeper-server-start.sh -daemon config/zookeeper.properties
sleep 2

# 启动 Kafka
export KAFKA_HEAP_OPTS="-Xmx1G -Xms1G"
bin/kafka-server-start.sh -daemon config/server.properties
```

**检查进程：**

```bash
ps aux | grep kafka | wc -l
# 显示为 3 表示正常；否则可再次执行：
# bin/kafka-server-start.sh -daemon config/server.properties
```

### 2.6 安装 MinIO

```bash
cd /tmp
wget https://dl.min.io/server/minio/release/linux-amd64/minio
chmod +x minio
./minio server /data &
```

- 控制台：`http://<IP>:9001`（或终端提示的端口）
- 默认账号：RootUser: minioadmin，RootPass: minioadmin（或 miniostorage，视版本而定）

**创建存储桶（在 MinIO 控制台或 mc 命令行）：**

- documents  
- encryptedfiles  
- photos  
- videos  

若使用防火墙，需放行 MinIO 端口，例如：

```bash
firewall-cmd --zone=public --permanent --add-port=9000/tcp
firewall-cmd --zone=public --permanent --add-port=9001/tcp
firewall-cmd --reload
```

### 2.7 安装 Pika（可选，默认端口 9221）

```bash
cd /tmp
wget https://github.com/OpenAtomFoundation/pika/releases/download/v3.3.6/pika-linux-x86_64-v3.3.6.tar.bz2
tar -xf pika-linux-x86_64-v3.3.6.tar.bz2
# 解压后目录可能为 output，可重命名为 pika
mv output pika 2>/dev/null || true
/tmp/pika/bin/pika -c /tmp/pika/conf/pika.conf &
```

### 2.8 安装 FFmpeg

```bash
cd /tmp
wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
xz -d ffmpeg-release-amd64-static.tar.xz
tar xf ffmpeg-release-amd64-static.tar
cd ffmpeg-*-amd64-static
cp ff* /usr/local/bin/
```

或使用包管理器（版本可能略旧）：

```bash
# CentOS/Fedora
dnf install ffmpeg -y

# Ubuntu/Debian
apt install ffmpeg -y
```

---

## 三、获取源码并初始化数据库

### 3.1 克隆代码

```bash
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

### 3.2 初始化数据库

在项目根目录下执行（SQL 位于 `teamgramd/deploy/sql/`）：

```bash
# 创建数据库（若已创建可跳过）
mysql -uroot -e "CREATE DATABASE IF NOT EXISTS teamgram;"

# 导入 SQL（按顺序）
mysql -uroot teamgram < teamgramd/deploy/sql/1_teamgram.sql
# ... 按时间顺序导入所有 migrate-*.sql，或使用循环：
for f in teamgramd/deploy/sql/migrate-*.sql; do mysql -uroot teamgram < "$f"; done
mysql -uroot teamgram < teamgramd/deploy/sql/z_init.sql
```

---

## 四、编译 Teamgram 服务端

### 4.1 安装 Go

**DNF：**

```bash
dnf install go -y
```

**APT：**

```bash
apt install golang-go -y
```

或从 [Go 官网](https://go.dev/dl/) 安装 1.21+ 版本。

### 4.2 编译

在项目根目录执行：

```bash
make
```

编译产物在 `teamgramd/bin/` 下（idgen、status、authsession、dfs、media、biz、msg、sync、bff、session、gnetway 等）。

---

## 五、修改配置文件

配置文件在 `teamgramd/etc/` 下，需根据本机环境修改为 **127.0.0.1** 或实际 IP/端口，保证与 MySQL、Redis、Etcd、Kafka、MinIO、Pika 一致。

**重点检查：**

1. **dfs.yaml**  
   - `Minio`：Endpoint、AccessKeyID、SecretAccessKey（与 MinIO 一致）  
   - `SSDB`：若用 Pika，改为 `127.0.0.1:9221`；否则用 Redis `127.0.0.1:6379`

2. **gnetway.yaml**  
   - `Gnetway.Server.Addresses`：对外的 MTProto 端口，如 `0.0.0.0:10443`、`0.0.0.0:5222` 等

3. **各服务 YAML**  
   - `Etcd.Hosts`：`127.0.0.1:2379`  
   - `Mysql.Addr` / `DSN`：`127.0.0.1:3306`  
   - `Cache` / `Redis`：`127.0.0.1:6379`  
   - Kafka 相关配置中的 broker 地址：`127.0.0.1:9092`

示例（dfs.yaml 片段）：

```yaml
Minio:
  Endpoint: localhost:9000
  AccessKeyID: minioadmin
  SecretAccessKey: miniostorage
  UseSSL: false
SSDB:
  - Host: 127.0.0.1:9221   # 使用 Pika 时
  # - Host: 127.0.0.1:6379 # 使用 Redis 时
```

---

## 六、启动服务

确保 MySQL、Redis、Etcd、Kafka（及 Zookeeper）、MinIO（及可选 Pika）均已启动后：

```bash
cd teamgramd/bin
./runall2.sh
```

如需开放外网访问，在防火墙中放行对应端口，例如：

```bash
firewall-cmd --zone=public --permanent --add-port=10443/tcp
firewall-cmd --zone=public --permanent --add-port=5222/tcp
firewall-cmd --reload
```

---

## 七、停止服务

在 `teamgramd/bin` 下执行：

```bash
./killall.sh
```

---

## 八、常见问题

- **CentOS 7 防火墙**：MinIO 等需放行端口，或临时 `systemctl stop firewalld` 排查。
- **Kafka 进程数**：`ps aux | grep kafka | wc -l` 为 3 表示 Zookeeper + Kafka 正常。
- **Pika 解压目录**：若解压后不是 `pika` 目录，需手动重命名或修改启动路径。
- **SQL 路径**：所有初始化 SQL 均在 `teamgramd/deploy/sql/`，不要使用错误的 `teamgramd/sql/`。

完成以上步骤后，Teamgram 服务端即可在 Linux 上以手动方式运行。
