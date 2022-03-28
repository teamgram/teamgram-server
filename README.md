# Teamgram - Open source [mtproto](https://core.telegram.org/mtproto) server written in golang
> open source mtproto server implemented in golang with compatible telegram client.

English | [简体中文](readme-cn.md)

### Introduce
Open source [mtproto](https://core.telegram.org/mtproto) server written in golang

### Architecture
![Architecture](docs/image/architecture-001.png)

### Documents
[Diffie–Hellman key exchange](docs/dh-key-exchange.md)

[Creating an Authorization Key](docs/Creating_an_Authorization_Key.md)

[Mobile Protocol: Detailed Description (v.1.0, DEPRECATED)](docs/Mobile_Protocol-Detailed_Description_v.1.0_DEPRECATED.md)

[Encrypted CDNs for Speed and Security](docs/cdn.md) Translate By [@steedfly](https://github.com/steedfly)

### Quick start with Docker
> TODO...

### [Centos 9 Stream Build and Install](docs/install-centos-9.md) [@A Feel]

### Manual Build and Install
#### Depends
- **mysql5.7**
- [redis](https://redis.io/)
- [etcd](https://etcd.io/)
- [kafka](https://kafka.apache.org/quickstart)
- [minio](https://docs.min.io/docs/minio-quickstart-guide.html#GNU/Linux)
- [ffmpeg](https://www.johnvansickle.com/ffmpeg/)

##### Run depends with docker-compose
```bash
# pull docker images
docker-compose pull

# run docker-compose
docker-compose up -d
```

#### Install Teamgram
- Get source code　
```
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

- init database
```
1. create database teamgram
2. init teamgram database
   mysql -uroot teamgram < teamgramd/sql/teamgram2.sql
   mysql -uroot teamgram < teamgramd/sql/migrate_20220321.sql
   mysql -uroot teamgram < teamgramd/sql/migrate_20220326.sql
   mysql -uroot teamgram < teamgramd/sql/migrate_20220328.sql
```

- init minio buckets, bucket names:
  - `documents`
  - `encryptedfiles`
  - `photos`
  - `videos`

- Build
```
cd scripts
./build.sh
```

- Run
```
cd ../teamgramd/bin
./runall2.sh
```


### Compatible clients
**Important**: default signIn and signOut verify code is **12345**

[Android client for Teamgram](https://github.com/teamgram/teamgram-android)

[iOS client for Teamgram](https://github.com/teamgram/teamgram-ios)

[tdesktop for Teamgram](https://github.com/teamgram/teamgram-tdesktop)

## Feedback
Please report bugs, concerns, suggestions by issues, or join telegram group [Teamgram中文社区](https://t.me/cnteamgram) Or [Teamgram](https://t.me/enteamgram) to discuss problems around source code.

## Notes
If need enterprise edition, please PM the **[author](https://t.me/benqi)**

## Give a Star! ⭐

If you like or are using this project to learn or start your solution, please give it a star. Thanks!
