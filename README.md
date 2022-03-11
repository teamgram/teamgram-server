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

#### Docker run demo

### Manual Build and Install
#### Install third_party
##### mysql
- install **mysql5.7**
- init database
    - create database teamgram

    - init teamgram database
  
    ```
    mysql -uroot teamgram < teamgram2.sql
    ```

##### redis
- install [redis](https://redis.io/)

##### etcd
- install [etcd](https://etcd.io/)

##### kafka
- install [kafka](https://kafka.apache.org/quickstart)

##### minio
- install [minio](https://docs.min.io/docs/minio-quickstart-guide.html#GNU/Linux)
- init minio buckets, bucket names:
    - `documents`
    - `encryptedfiles`
    - `photos`
    - `videos`

##### ffmpeg
- install [ffmpeg](https://www.johnvansickle.com/ffmpeg/)

#### Install Teamgram

- Get source code　
```
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

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

### TODO

## Feedback
### TODO

## Feedback
Please report bugs, concerns, suggestions by issues, or join telegram group [Teamgram中文社区](https://t.me/cnteamgram) Or [Teamgram](https://t.me/enteamgram) to discuss problems around source code.

## Notes
If need enterprise edition, please PM the **[author](https://t.me/benqi)**
