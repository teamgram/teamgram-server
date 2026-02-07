# Teamgram Server — Manual Installation (Linux)

This guide describes how to install the Teamgram server and its dependencies from scratch on Linux. It applies to CentOS 7/9, Fedora, Ubuntu/Debian, and similar distributions. Use the package manager commands that match your distribution.

> English (primary) | [中文](./install-manual-linux-zh.md)  
> For Docker deployment, see [install-docker.md](./install-docker.md).

---

## 1. Requirements

- **Go**: 1.21 or later (for building teamgram-server)
- **MySQL**: 8.x (8.0.29 recommended)
- **Redis**: 6.x
- **Etcd**: 3.5.x
- **Kafka**: 2.x / 3.x (with Zookeeper)
- **MinIO**: object storage
- **Pika** (optional): default port 9221; can replace some Redis usage
- **FFmpeg**: media processing

Run the following steps as root or with sudo.

---

## 2. Install Dependencies

### 2.1 Install DNF (CentOS 7 only)

CentOS 7 uses yum by default. You can install dnf for consistency with the rest of this guide:

```bash
yum install epel-release -y
yum install dnf -y
dnf --version
```

### 2.2 Install MySQL

**CentOS 9 / Fedora:**

```bash
dnf install mysql-server -y
systemctl enable --now mysqld
```

**Fedora (community MySQL):**

```bash
dnf install community-mysql-server -y
systemctl enable --now mysqld
```

**Ubuntu/Debian:**

```bash
apt update
apt install mysql-server -y
systemctl enable --now mysql
```

**Configure MySQL:**

```bash
# Optional: run security setup
mysql_secure_installation

# Log in, create database, set empty password (adjust as needed)
mysql -uroot -p

mysql> CREATE DATABASE teamgram;
mysql> UPDATE mysql.user SET authentication_string='' WHERE user='root';
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY '';
mysql> FLUSH PRIVILEGES;
mysql> exit
```

### 2.3 Install Redis

**DNF (CentOS/Fedora):**

```bash
dnf install redis -y
systemctl enable --now redis
```

**APT (Ubuntu/Debian):**

```bash
apt install redis-server -y
systemctl enable --now redis-server
```

### 2.4 Install Etcd

Install Etcd from the official release (distribution-agnostic):

```bash
ETCD_VER=v3.5.0
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GOOGLE_URL}

mkdir -p /tmp/etcd-download-test
curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

# Verify
/tmp/etcd-download-test/etcd --version
/tmp/etcd-download-test/etcdctl version
```

**Start Etcd (foreground or background):**

```bash
# Run in background
/tmp/etcd-download-test/etcd &
```

For long-term use, copy the binaries to `/usr/local/bin` and configure a systemd service.

### 2.5 Install Apache Kafka

Install Java first, then Kafka (Kafka bundles Zookeeper).

**Install Java:**

```bash
# CentOS/Fedora
dnf install java-21-openjdk -y
# or
dnf install java -y

# Ubuntu/Debian
apt install openjdk-17-jdk -y
```

**Install and start Kafka:**

```bash
cd /tmp
# Example for 3.2.0; get the latest from https://kafka.apache.org/downloads
wget https://dlcdn.apache.org/kafka/3.2.0/kafka_2.12-3.2.0.tgz
tar -xzf kafka_2.12-3.2.0.tgz
cd kafka_2.12-3.2.0

# Start Zookeeper
bin/zookeeper-server-start.sh -daemon config/zookeeper.properties
sleep 2

# Start Kafka
export KAFKA_HEAP_OPTS="-Xmx1G -Xms1G"
bin/kafka-server-start.sh -daemon config/server.properties
```

**Check processes:**

```bash
ps aux | grep kafka | wc -l
# A count of 3 is normal; otherwise run again:
# bin/kafka-server-start.sh -daemon config/server.properties
```

### 2.6 Install MinIO

```bash
cd /tmp
wget https://dl.min.io/server/minio/release/linux-amd64/minio
chmod +x minio
./minio server /data &
```

- Console: `http://<IP>:9001` (or the port shown in the terminal)
- Default credentials: RootUser: minioadmin, RootPass: minioadmin (or miniostorage, depending on version)

**Create buckets (via MinIO console or mc CLI):**

- documents  
- encryptedfiles  
- photos  
- videos  

If using a firewall, open the MinIO ports, for example:

```bash
firewall-cmd --zone=public --permanent --add-port=9000/tcp
firewall-cmd --zone=public --permanent --add-port=9001/tcp
firewall-cmd --reload
```

### 2.7 Install Pika (optional, default port 9221)

```bash
cd /tmp
wget https://github.com/OpenAtomFoundation/pika/releases/download/v3.3.6/pika-linux-x86_64-v3.3.6.tar.bz2
tar -xf pika-linux-x86_64-v3.3.6.tar.bz2
# Extracted dir may be named output; rename to pika if needed
mv output pika 2>/dev/null || true
/tmp/pika/bin/pika -c /tmp/pika/conf/pika.conf &
```

### 2.8 Install FFmpeg

```bash
cd /tmp
wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
xz -d ffmpeg-release-amd64-static.tar.xz
tar xf ffmpeg-release-amd64-static.tar
cd ffmpeg-*-amd64-static
cp ff* /usr/local/bin/
```

Or use the package manager (may be an older version):

```bash
# CentOS/Fedora
dnf install ffmpeg -y

# Ubuntu/Debian
apt install ffmpeg -y
```

---

## 3. Get Source Code and Initialize Database

### 3.1 Clone repository

```bash
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

### 3.2 Initialize database

From the project root (SQL files are under `teamgramd/deploy/sql/`):

```bash
# Create database (skip if already created)
mysql -uroot -e "CREATE DATABASE IF NOT EXISTS teamgram;"

# Import SQL in order
mysql -uroot teamgram < teamgramd/deploy/sql/1_teamgram.sql
# Import all migrate-*.sql (in order), e.g. with a loop:
for f in teamgramd/deploy/sql/migrate-*.sql; do mysql -uroot teamgram < "$f"; done
mysql -uroot teamgram < teamgramd/deploy/sql/z_init.sql
```

---

## 4. Build Teamgram Server

### 4.1 Install Go

**DNF:**

```bash
dnf install go -y
```

**APT:**

```bash
apt install golang-go -y
```

Or install Go 1.21+ from [go.dev/dl](https://go.dev/dl/).

### 4.2 Build

From the project root:

```bash
make
```

Binaries are produced in `teamgramd/bin/` (idgen, status, authsession, dfs, media, biz, msg, sync, bff, session, gnetway, etc.).

---

## 5. Edit Configuration

Configuration files are under `teamgramd/etc/`. Set addresses to **127.0.0.1** or your actual IP/ports so they match MySQL, Redis, Etcd, Kafka, MinIO, and Pika.

**Check in particular:**

1. **dfs.yaml**  
   - `Minio`: Endpoint, AccessKeyID, SecretAccessKey (must match your MinIO setup)  
   - `SSDB`: use `127.0.0.1:9221` if using Pika, otherwise Redis `127.0.0.1:6379`

2. **gnetway.yaml**  
   - `Gnetway.Server.Addresses`: MTProto ports, e.g. `0.0.0.0:10443`, `0.0.0.0:5222`

3. **All service YAMLs**  
   - `Etcd.Hosts`: `127.0.0.1:2379`  
   - `Mysql.Addr` / `DSN`: `127.0.0.1:3306`  
   - `Cache` / `Redis`: `127.0.0.1:6379`  
   - Kafka broker address: `127.0.0.1:9092`

Example (dfs.yaml snippet):

```yaml
Minio:
  Endpoint: localhost:9000
  AccessKeyID: minioadmin
  SecretAccessKey: miniostorage
  UseSSL: false
SSDB:
  - Host: 127.0.0.1:9221   # when using Pika
  # - Host: 127.0.0.1:6379 # when using Redis
```

---

## 6. Start Services

Ensure MySQL, Redis, Etcd, Kafka (and Zookeeper), MinIO (and optionally Pika) are running, then:

```bash
cd teamgramd/bin
./runall2.sh
```

To allow external access, open the gateway ports in the firewall, for example:

```bash
firewall-cmd --zone=public --permanent --add-port=10443/tcp
firewall-cmd --zone=public --permanent --add-port=5222/tcp
firewall-cmd --reload
```

---

## 7. Stop Services

From `teamgramd/bin`:

```bash
./killall.sh
```

---

## 8. Troubleshooting

- **CentOS 7 firewall**: Open required ports for MinIO etc., or temporarily run `systemctl stop firewalld` to debug.
- **Kafka process count**: `ps aux | grep kafka | wc -l` should be 3 when Zookeeper and Kafka are running.
- **Pika extract dir**: If the archive does not extract to a directory named `pika`, rename it or adjust the start command.
- **SQL path**: All init SQL files live under `teamgramd/deploy/sql/`; do not use `teamgramd/sql/`.

After completing these steps, the Teamgram server runs manually on Linux.
