# CentOS7 teamgram-server环境搭建
> 原文: https://www.jianshu.com/p/60a64203ca13
>
> 作者 @saeipi

### 1、安装dnf
  ```
    yum install epel-release -y
    yum install dnf -y
    dnf --version
  ```

### 2、Install Mysql Server
详见：https://www.jianshu.com/p/a47de83610c7

### 3、Install Redis Server
```
    dnf install redis -y
    systemctl enable redis
    systemctl start redis
```

### 4、Install Etcd.io
```
    4.1
    mkdir -p /tmp/etcd-download-test

    ETCD_VER=v3.5.0

    # choose either URL
    GOOGLE_URL=https://storage.googleapis.com/etcd
    GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
    DOWNLOAD_URL=${GOOGLE_URL}

    rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
    rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test

    4.2
    curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
    tar xzvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
    rm -f /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz

    /tmp/etcd-download-test/etcd --version
    /tmp/etcd-download-test/etcdctl version
    /tmp/etcd-download-test/etcdutl version

    4.3
    # start a local etcd server
    /tmp/etcd-download-test/etcd &
```

###  5、Install APACHE KAFKA
```
    5.1 安装kafka
    dnf install java
    cd /tmp
    wget https://dlcdn.apache.org/kafka/3.2.0/kafka_2.12-3.2.0.tgz
    tar -xzf kafka_2.12-3.2.0.tgz
    cd kafka_2.12-3.2.0

    5.2 zookeeper
    bin/zookeeper-server-start.sh -daemon config/zookeeper.properties
    sleep 2

    export KAFKA_HEAP_OPTS="-Xmx1G -Xms1G"
    bin/kafka-server-start.sh -daemon config/server.properties

    5.3 检查是否安装成功
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

    6.1 安装minio
    cd /tmp
    wget https://dl.min.io/server/minio/release/linux-amd64/minio
    chmod +x minio
    ./minio server /data &

    6.2 Create Bucket 并设置访问权限
    documents
    encryptedfiles
    photos
    videos
```

### 7、Install PIKA
```
  default port : 9221

    cd /tmp
    wget https://github.com/OpenAtomFoundation/pika/releases/download/v3.3.6/pika-linux-x86_64-v3.3.6.tar.bz2
    tar -xf pika-linux-x86_64-v3.3.6.tar.bz2
    /tmp/pika/bin/pika -c /tmp/pika/conf/pika.conf &

    解压后文件可能不是pika，需手动重命名
```

### 8、Install FFMpeg
```
    
    cd /tmp
    wget https://johnvansickle.com/ffmpeg/releases/ffmpeg-release-amd64-static.tar.xz
    xz -d ffmpeg-release-amd64-static.tar.xz
    tar xf ffmpeg-release-amd64-static.tar
    cd ffmpeg-5.0.1-amd64-static
    cp ff* /usr/local/bin/
```

### 9、Get source code
```
    git clone https://github.com/teamgram/teamgram-server.git
    cd teamgram-server
```
