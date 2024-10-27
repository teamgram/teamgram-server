# Teamgram Deployment Tutorial on Fedora 40

## by @lingyicute

## Dependency Installation and Running (All operations should be executed as root user, please pay attention to permissions and security)

### Install MySQL Server

```
dnf install community-mysql-server -y
systemctl enable --now mysqld
```

### Configure MySQL

```
mysql_secure_installation
# It will ask some questions. You can set the password to anything; we will delete it later. Answer "yes" to the remaining questions.

mysql -uroot -p
# Enter the password you just set

mysql> CREATE DATABASE teamgram;
# Create the database
mysql> UPDATE mysql.user SET authentication_string='' WHERE user='root';
mysql> ALTER USER 'root'@'localhost' IDENTIFIED BY '';
mysql> FLUSH PRIVILEGES;
# Configure empty password and refresh privileges

mysql> exit

mysql -uroot teamgram < teamgramd/sql/1_teamgram.sql
mysql -uroot teamgram < teamgramd/sql/migrate-*.sql
mysql -uroot teamgram < teamgramd
# Import configuration scripts
```

### Install Redis Server

```
sudo dnf install redis -y
sudo systemctl enable --now redis
```

### Install Etcd

[Etcd](https://etcd.io/docs/v3.5/install/)

### Install Kafka

```
dnf install java-21-openjdk
cd /tmp
wget < Please find the latest version on the official website~ >
tar -xzf kafka-<ver>.tgz
cd kafka-<ver>

# Start Zookeeper
bin/zookeeper-server-start.sh -daemon config/zookeeper.properties
sleep 2

# Start Kafka Server
export KAFKA_HEAP_OPTS="-Xmx1G -Xms1G"
bin/kafka-server-start.sh -daemon config/server.properties

# Check Kafka process
ps aux | grep kafka | wc -l
# If the process count is 3, it started successfully.
# If the process count is not 3, restart the Kafka server:
bin/kafka-server-start.sh -daemon config/server.properties
```

### Install Minio
```
cd /tmp
wget https://dl.min.io/server/minio/release/linux-amd64/minio
chmod +x minio
./minio server /data &

# Console access address
# http://ip:xxxxx ...
# Default username and password
# RootUser: minioadmin
# RootPass: minioadmin

# Open firewall port
sudo firewall-cmd --zone=public --permanent --add-port=xxxxx/tcp
sudo firewall-cmd --reload

# Access http://ip:xxxxx to create buckets
# - documents
# - encryptedfiles
# - photos
# - videos

# Close firewall port
sudo firewall-cmd --zone=public --permanent --remove-port=xxxxx/tcp
sudo firewall-cmd --reload
```

### Install Pika
```
/tmp/pika/bin/pika -c /tmp/pika/conf/pika.conf &
```

### Install FFmpeg
```
sudo cp ff* /usr/local/bin/
```

### Deployment
```
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
make
```
