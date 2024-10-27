# Fedora 40 上的 teamgram 部署教程

## by @lingyicute

## 简体中文 | [English](./install-fedora-en.md) 

## 依赖安装与运行（以下全部操作使用root用户执行，请注意权限和安全性）

### 安装 MySQL 服务器

```
dnf install community-mysql-server -y
systemctl enable --now mysqld
```

### 配置 MySQL

```
mysql_secure_installation
# 它将会询问一些问题。密码可以随便设置，我们接下来会删除它。剩下的问题，全部回答“是”即可。

mysql -uroot -p
# 输入你刚才设置的密码

mysql> CREATE DATABASE teamgram;
# 创建数据库
mysql> UPDATE mysql.user SET authentication_string='' WHERE user='root';
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY '';
mysql> FLUSH PRIVILEGES;
# 配置空密码并刷新权限

mysql> exit

mysql -uroot teamgram < teamgramd/sql/1_teamgram.sql
mysql -uroot teamgram < teamgramd/sql/migrate-*.sql
mysql -uroot teamgram < teamgramd
# 导入配置脚本
```

### 安装 Redis 服务器

```
sudo dnf install redis -y
sudo systemctl enable --now redis
```

### 安装 Etcd

[Etcd](https://etcd.io/docs/v3.5/install/)

### 安装 Kafka

```
dnf install java-21-openjdk
cd /tmp
wget < 请在官网上寻找最新版本吧~ >
tar -xzf kafka-<ver>.tgz
cd kafka-<ver>

# 启动 Zookeeper
bin/zookeeper-server-start.sh -daemon config/zookeeper.properties
sleep 2

# 启动 Kafka 服务器
export KAFKA_HEAP_OPTS="-Xmx1G -Xms1G"
bin/kafka-server-start.sh -daemon config/server.properties

# 检查 Kafka 进程
ps aux | grep kafka | wc -l
# 进程数为 3 就启动成功啦。
# 如果进程数不为 3，则重新启动 Kafka 服务器：
bin/kafka-server-start.sh -daemon config/server.properties
```

### 安装 Minio
```
cd /tmp
wget https://dl.min.io/server/minio/release/linux-amd64/minio
chmod +x minio
./minio server /data &

# 控制台访问地址
# http://ip:xxxxx ...
# 默认用户名和密码
# RootUser: minioadmin
# RootPass: minioadmin

# 开放防火墙端口
sudo firewall-cmd --zone=public --permanent --add-port=xxxxx/tcp
sudo firewall-cmd --reload

# 访问 http://ip:xxxxx 创建存储桶
# - documents
# - encryptedfiles
# - photos
# - videos

# 关闭防火墙端口
sudo firewall-cmd --zone=public --permanent --remove-port=xxxxx/tcp
sudo firewall-cmd --reload
```

### 安装 Pika
```
cd /tmp
wget https://github.com/OpenAtomFoundation/pika/releases/download/v3.3.6/pika-linux-x86_64-v3.3.6.tar.bz2
tar -xf pika-linux-x86_64-v3.3.6.tar.bz2
mv output pika
/tmp/pika/bin/pika -c /tmp/pika/conf/pika.conf &
```

### 安装 FFmpeg
```
cd /tmp
wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
xz -d ffmpeg-release-amd64-static.tar.xz
tar xf ffmpeg-release-amd64-static.tar
cd ffmpeg-5.0-amd64-static/
sudo cp ff* /usr/local/bin/
```
### 部署
```
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
make
```
