## Centos 9 Stream Build and Install 

## Depends Install and Run

### Install Mysql Server 
```
dnf install mysql-server -y
systemctl enable --now mysqld
```

### Install Redis Server 
```
dnf install redis -y
systemctl enable redis
systemctl start redis
```

### Install Etcd.io 
```
ETCD_VER=v3.5.0

# choose either URL
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GOOGLE_URL}

rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

/tmp/etcd-download-test/etcd --version
/tmp/etcd-download-test/etcdctl version
/tmp/etcd-download-test/etcdutl version
```
```
# start a local etcd server
/tmp/etcd-download-test/etcd &
```

### Install APACHE KAFKA
```
dnf install java
cd /tmp
wget http://archive.apache.org/dist/kafka/2.2.1/kafka_2.11-2.2.1.tgz
tar -xzf kafka_2.11-2.2.1.tgz
cd kafka_2.11-2.2.1

```
```
bin/zookeeper-server-start.sh -daemon config/zookeeper.properties
sleep 2

export KAFKA_HEAP_OPTS="-Xmx1G -Xms1G"
bin/kafka-server-start.sh -daemon config/server.properties
```
```
ps aux | grep kafka | wc -l
#show number 3 is ok !
#else run 
bin/kafka-server-start.sh -daemon config/server.properties
```

### Install Minio 
```
cd /tmp
wget https://dl.min.io/server/minio/release/linux-amd64/minio
chmod +x minio
./minio server /data &
```
Console: http://ip:xxxxx ...
RootUser: minioadmin
RootPass: minioadmin
```
firewall-cmd --zone=public --permanent --add-port=xxxxx/tcp
firewall-cmd --reload
```
Access  http://ip:xxxxx
Create Bucket 
 - documents
 - encryptedfiles
 - photos
 - videos
```
firewall-cmd --zone=public --permanent --remove-port=xxxxx/tcp
firewall-cmd --reload
```

### Install PIKA
default port : 9221
```
cd /tmp
wget https://github.com/OpenAtomFoundation/pika/releases/download/v3.3.6/pika-linux-x86_64-v3.3.6.tar.bz2
tar -xf pika-linux-x86_64-v3.3.6.tar.bz2
mv output pika
/tmp/pika/bin/pika -c /tmp/pika/conf/pika.conf &
```
### Install FFMpeg 
```
cd /tmp
wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
xz -d ffmpeg-release-amd64-static.tar.xz
tar xf ffmpeg-release-amd64-static.tar
cd ffmpeg-5.0-amd64-static/
cp ff* /usr/local/bin/
```

## Install Teamgram Server

### Get source code 
```
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

### Init Database 
- create teamgram database
```
mysql -uroot -p
mysql> create database teamgram;
mysql> exit
```

- import sql scripts
```
mysql -uroot teamgram < teamgramd/sql/1_teamgram.sql
mysql -uroot teamgram < teamgramd/sql/migrate-*.sql
mysql -uroot teamgram < teamgramd/sql/z_init.sql
```

### Build
```
dnf install go -y
make
```

### Modify config file 
```
vim ../teamgramd/etc/dfs.yaml
{
AccessKeyID: minioadmin
SecretAccessKey: minioadmin

SSDB:
  - Host: 127.0.0.1:9221
}

vim ../teamgramd/etc/gateway.yaml
{
Addrs:
    - 0.0.0.0:10443  #modity listen port..
}
```

```
cd ../teamgramd/bin/
./runall2.sh
firewall-cmd --zone=public --permanent --add-port=10443/tcp
firewall-cmd --reload
```


