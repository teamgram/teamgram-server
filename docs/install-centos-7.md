# CentOS7 teamgram-server环境搭建
> 原文: https://www.jianshu.com/p/60a64203ca13
>
> 作者 @saeipi
>

## 1、安装dnf
```
yum install epel-release -y
yum install dnf -y
dnf --version
```
## 2、Install Mysql Server
```
### 旧的:
dnf install mysql-server -y
systemctl enable --now mysqld

新的：
下载mysql
wget https://repo.mysql.com//mysql80-community-release-el7-3.noarch.rpm
sudo rpm -ivh mysql80-community-release-el7-3.noarch.rpm

安装mysql
yum install mysql mysql-server

启动mysql
sudo systemctl start mysqld

检查mysql状态
sudo systemctl status mysqld


遇到错误及处理
错误1:
总计                                                                                                                                 2.2 MB/s |  74 MB  00:00:32     
从 file:///etc/pki/rpm-gpg/RPM-GPG-KEY-mysql 检索密钥
源 "MySQL 8.0 Community Server" 的 GPG 密钥已安装，但是不适用于此软件包。请检查源的公钥 URL 是否配置正确。
失败的软件包是：mysql-community-icu-data-files-8.0.29-1.el7.x86_64
GPG  密钥配置为：file:///etc/pki/rpm-gpg/RPM-GPG-KEY-mysql


rpm --checksig mysql-community-server-8.0.27-1.el7.x86_64.rpm
gpg --export -a 3a79bd29 > 3a79bd29.asc
rpm --import 3a79bd29.asc
rpm --import https://repo.mysql.com/RPM-GPG-KEY-mysql-2022

错误2:
● mysqld.service - MySQL Server
Loaded: loaded (/usr/lib/systemd/system/mysqld.service; enabled; vendor preset: disabled)
Active: failed (Result: exit-code) since 四 2022-05-26 09:00:28 CST; 27s ago
Docs: man:mysqld(8)
http://dev.mysql.com/doc/refman/en/using-systemd.html
Process: 56583 ExecStart=/usr/sbin/mysqld $MYSQLD_OPTS (code=exited, status=1/FAILURE)
Process: 56554 ExecStartPre=/usr/bin/mysqld_pre_systemd (code=exited, status=0/SUCCESS)
Main PID: 56583 (code=exited, status=1/FAILURE)
Status: "Data Dictionary upgrade from MySQL 5.7 in progress"
Error: 2 (没有那个文件或目录)

cat /var/log/mysqld.log
rm -rf /var/lib/mysql/*

错误3
ERROR 1045 (28000): Access denied for user 'root'@'localhost' (using password: NO)

vim /etc/my.cnf
在最后加上
skip-grant-tables
再次启动
systemctl start mysqld.service
```

## 3、Install Redis Server
```
dnf install redis -y
systemctl enable redis
systemctl start redis
```

## 4、Install Etcd.io
### 4.1
```
mkdir -p /tmp/etcd-download-test

ETCD_VER=v3.5.0

# choose either URL
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GOOGLE_URL}

rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test
```

### 4.2
```
curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

/tmp/etcd-download-test/etcd --version
/tmp/etcd-download-test/etcdctl version
/tmp/etcd-download-test/etcdutl version
```

### 4.3
```
# start a local etcd server
/tmp/etcd-download-test/etcd &
```

## 5、Install APACHE KAFKA
### 5.1 安装kafka
```
dnf install java
cd /tmp
wget https://dlcdn.apache.org/kafka/3.2.0/kafka_2.12-3.2.0.tgz
tar -xzf kafka_2.12-3.2.0.tgz
cd kafka_2.12-3.2.0
```

### 5.2 zookeeper
```
bin/zookeeper-server-start.sh -daemon config/zookeeper.properties
sleep 2

export KAFKA_HEAP_OPTS="-Xmx1G -Xms1G"
bin/kafka-server-start.sh -daemon config/server.properties
```

### 5.3 检查是否安装成功
```
ps aux | grep kafka | wc -l
#show number 3 is ok !
#else run
bin/kafka-server-start.sh -daemon config/server.properties
```

### 6、Install Minio
```
# 关闭防火墙
systemctl stop firewalld.service

# 查看防火墙状态
firewall-cmd --state
```

### 6.1 安装minio
```
cd /tmp
wget https://dl.min.io/server/minio/release/linux-amd64/minio
chmod +x minio
./minio server /data &
```

### 6.2 Create Bucket 并设置访问权限
```
documents
encryptedfiles
photos
videos
```

### 7、nstall PIKA
```
default port : 9221

cd /tmp
wget https://github.com/OpenAtomFoundation/pika/releases/download/v3.3.6/pika-linux-x86_64-v3.3.6.tar.bz2
tar -xf pika-linux-x86_64-v3.3.6.tar.bz2
/tmp/pika/bin/pika -c /tmp/pika/conf/pika.conf &

解压后文件可能不是pika，需手动重命名
```

## 8、Install FFMpeg
```
cd /tmp
wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
xz -d ffmpeg-release-amd64-static.tar.xz
tar xf ffmpeg-release-amd64-static.tar
cd ffmpeg-5.0.1-amd64-static
cp ff* /usr/local/bin/
```

## 9、Get source code
```
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```
