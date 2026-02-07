# Teamgram Server — Manual Installation (macOS)

This guide describes how to install the Teamgram server and its dependencies from scratch on macOS (Intel and Apple Silicon). [Homebrew](https://brew.sh/) is recommended for managing dependencies.

> English (primary) | [中文](./install-manual-macos-zh.md)  
> For Docker deployment, see [install-docker.md](./install-docker.md).

---

## 1. Requirements

- **macOS**: 10.15 or later (11+ recommended; Apple Silicon supported)
- **Go**: 1.21 or later
- **MySQL**: 8.x
- **Redis**: 6.x
- **Etcd**: 3.5.x
- **Kafka**: 2.x / 3.x (with Zookeeper; Homebrew Kafka usually includes it)
- **MinIO**: object storage
- **Pika** (optional): install from source or release; default port 9221
- **FFmpeg**: media processing

Install Homebrew if needed:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

---

## 2. Install Dependencies

### 2.1 Install MySQL

```bash
brew install mysql
brew services start mysql
```

**Configure and create database:**

```bash
# Set root password on first install if prompted
mysql_secure_installation

mysql -uroot -p

mysql> CREATE DATABASE teamgram;
mysql> UPDATE mysql.user SET authentication_string='' WHERE user='root';
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY '';
mysql> FLUSH PRIVILEGES;
mysql> exit
```

### 2.2 Install Redis

```bash
brew install redis
brew services start redis
```

Defaults to `127.0.0.1:6379`.

### 2.3 Install Etcd

```bash
brew install etcd
brew services start etcd
```

Defaults to `127.0.0.1:2379`.

### 2.4 Install Kafka

Kafka depends on Java; Homebrew will install OpenJDK:

```bash
brew install kafka
```

**Start Zookeeper and Kafka:**

```bash
# Start Zookeeper in background
zookeeper-server-start -daemon $(brew --prefix kafka)/libexec/config/zookeeper.properties
sleep 2

# Start Kafka
kafka-server-start -daemon $(brew --prefix kafka)/libexec/config/server.properties
```

Or use brew services if the formula supports it:

```bash
brew services start zookeeper
brew services start kafka
```

**Check processes:**

```bash
ps aux | grep kafka | wc -l
# About 3 processes means it's running
```

### 2.5 Install MinIO

```bash
brew install minio
```

**Start MinIO (with a data directory):**

```bash
mkdir -p /tmp/minio-data
minio server /tmp/minio-data &
```

Or with brew services:

```bash
brew services start minio
```

- Console: http://localhost:9001  
- Default credentials: minioadmin / minioadmin (or miniostorage, depending on version)

**Create buckets** in the MinIO console:

- documents  
- encryptedfiles  
- photos  
- videos  

### 2.6 Install Pika (optional, port 9221)

Pika may not be in Homebrew. Install a macOS build from GitHub if available, or build from source:

```bash
cd /tmp
# If a macOS binary exists, e.g.:
# curl -L -o pika.tar.bz2 https://github.com/OpenAtomFoundation/pika/releases/download/v3.3.6/pika-darwin-x86_64-v3.3.6.tar.bz2
# Then extract and run:
# /tmp/pika/bin/pika -c /tmp/pika/conf/pika.conf &
```

If no macOS build is available, you can skip Pika and use Redis for SSDB in the config (see configuration section below).

### 2.7 Install FFmpeg

```bash
brew install ffmpeg
```

### 2.8 Install Go

```bash
brew install go
go version   # must be 1.21+
```

Or install from [go.dev/dl](https://go.dev/dl/).

---

## 3. Get Source Code and Initialize Database

### 3.1 Clone repository

```bash
cd ~   # or your preferred directory
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

### 3.2 Initialize database

From the project root (SQL files are under `teamgramd/deploy/sql/`):

```bash
mysql -uroot -e "CREATE DATABASE IF NOT EXISTS teamgram;"

mysql -uroot teamgram < teamgramd/deploy/sql/1_teamgram.sql
for f in teamgramd/deploy/sql/migrate-*.sql; do mysql -uroot teamgram < "$f"; done
mysql -uroot teamgram < teamgramd/deploy/sql/z_init.sql
```

---

## 4. Build Teamgram Server

From the project root:

```bash
make
```

Binaries are produced in `teamgramd/bin/`.

---

## 5. Edit Configuration

Configuration files are under `teamgramd/etc/`. For a local macOS setup, services are usually on `127.0.0.1`. Confirm:

1. **dfs.yaml**  
   - `Minio.Endpoint`: `localhost:9000`  
   - `Minio.AccessKeyID` / `SecretAccessKey`: match your MinIO (e.g. minioadmin / minioadmin)  
   - `SSDB`: `127.0.0.1:9221` if using Pika, otherwise Redis `127.0.0.1:6379`

2. **gnetway.yaml**  
   - `Gnetway.Server.Addresses`: e.g. `0.0.0.0:10443`, `0.0.0.0:5222`

3. **All YAMLs**  
   - `Etcd.Hosts`: `127.0.0.1:2379`  
   - MySQL: `127.0.0.1:3306`  
   - Redis: `127.0.0.1:6379`  
   - Kafka: `127.0.0.1:9092`

Example (dfs.yaml snippet):

```yaml
Minio:
  Endpoint: localhost:9000
  AccessKeyID: minioadmin
  SecretAccessKey: minioadmin
  UseSSL: false
SSDB:
  - Host: 127.0.0.1:6379   # Redis when Pika is not installed
```

---

## 6. Start Services

Ensure MySQL, Redis, Etcd, Kafka (and Zookeeper), and MinIO are running, then:

```bash
cd teamgramd/bin
./runall2.sh
```

For LAN access, ensure `gnetway.yaml` listens on `0.0.0.0` and that the firewall allows ports 10443, 5222, etc.

---

## 7. Stop Services

From `teamgramd/bin`:

```bash
./killall.sh
```

To stop dependency services (as needed):

```bash
brew services stop mysql
brew services stop redis
brew services stop etcd
# Kafka / Zookeeper may need to be stopped manually or via brew services
```

---

## 8. Troubleshooting

- **Apple Silicon (M1/M2/M3)**: Homebrew builds for Go, MySQL, Redis, Etcd, Kafka, and MinIO support ARM; no extra setup usually needed.  
- **Kafka path**: Use `brew --prefix kafka` to see the install path; config files are under `libexec/config/`.  
- **MinIO data dir**: For production, use a persistent directory instead of `/tmp/minio-data`.  
- **Pika**: If there is no macOS build, omit Pika and set SSDB to Redis in the config.  
- **Port conflicts**: If 3306, 6379, 2379, 9092, 9000, etc. are in use, change the corresponding service and Teamgram config to match.

After completing these steps, the Teamgram server runs manually on macOS.
