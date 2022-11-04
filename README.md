# Teamgram - Unofficial open source [mtproto](https://core.telegram.org/mtproto) server written in golang
> open source mtproto server implemented in golang with compatible telegram client.

## Introduce
Open source [mtproto](https://core.telegram.org/mtproto) server implementation written in golang, support private deployment.

## Features
- MTProto 2.0
  - Abridged
  - Intermediate
  - Padded intermediate
  - Full
- API Layer: 147
- private chat
- basic group
- contacts

## Architecture
![Architecture](docs/image/architecture-001.png)

### Documents
[Diffie–Hellman key exchange](docs/dh-key-exchange.md)

[Creating an Authorization Key](docs/Creating_an_Authorization_Key.md)

[Mobile Protocol: Detailed Description (v.1.0, DEPRECATED)](docs/Mobile_Protocol-Detailed_Description_v.1.0_DEPRECATED.md)

[Encrypted CDNs for Speed and Security](docs/cdn.md) Translate By [@steedfly](https://github.com/steedfly)

## Installing Teamgram 
`Teamgram` relies on open source high-performance components: 

- **mysql5.7**
- [redis](https://redis.io/)
- [etcd](https://etcd.io/)
- [kafka](https://kafka.apache.org/quickstart)
- [minio](https://docs.min.io/docs/minio-quickstart-guide.html#GNU/Linux)
- [ffmpeg](https://www.johnvansickle.com/ffmpeg/)

Privatization deployment Before `Teamgram`, please make sure that the above five components have been installed. If your server does not have the above components, you must first install Missing components. 

- [Centos9 Stream Build and Install](docs/install-centos-9.md) [@A Feel]
- [CentOS7 teamgram-server环境搭建](docs/install-centos-7.md) [@saeipi]

If you have the above components, it is recommended to use them directly. If not, it is recommended to use `docker-compose-env.yaml`.


### Source code deployment
#### Install [Go environment](https://go.dev/doc/install). Make sure Go version is at least 1.17.


#### Get source code　

```
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

#### Init data
- init database

	```
	1. create database teamgram
	2. init teamgram database
	   mysql -uroot teamgram < teamgramd/sql/teamgram2.sql
	   mysql -uroot teamgram < teamgramd/sql/migrate-*.sql
	```

- init minio buckets
	- bucket names
	  - `documents`
	  - `encryptedfiles`
	  - `photos`
	  - `videos`
	- Access `http://ip:xxxxx` and create


#### Build
	
```
make
```

#### Run

```
cd teamgramd/bin
./runall2.sh
```

### Docker deployment
#### Install [Docker](https://docs.docker.com/get-docker/)

#### Install [Docker Compose](https://docs.docker.com/compose/install/)

#### Get source code

```
git clone https://github.com/teamgram/teamgram-server.git
cd teamgram-server
```

#### Install depends
- **change `192.168.1.150` to your ip in `docker-compose-env.yaml`**
- install depends

  ```
  # pull docker images
  docker-compose -f docker-compose-env.yaml pull
  
  # run docker-compose
  docker-compose -f docker-compose-env.yaml up -d
  ```
  
#### Init data
- init database
	
	```
	mysql -uteamgram -h127.0.0.1 -pteamgram teamgram < teamgramd/sql/teamgram2.sql
	mysql -uteamgram -h127.0.0.1 -pteamgram teamgram < teamgramd/sql/migrate-*.sql
	mysql -uteamgram -h127.0.0.1 -pteamgram teamgram < teamgramd/sql/init.sql
	```

- init minio buckets
	- bucket names:
	  - `documents`
	  - `encryptedfiles`
	  - `photos`
	  - `videos`
	- create buckets
		
		```
		# get mc
		docker run -it --entrypoint=/bin/bash minio/mc
		   
		# change 192.168.1.150 to your ip    
		mc alias set minio http://192.168.1.150:9000 minio miniostorage
		
		# create buckets
		mc mb minio/documents
		mc mb minio/encryptedfiles
		mc mb minio/photos
		mc mb minio/videos
  
		# quit docker minio/mc
		exit
		```

#### Run

```  
# run docker-compose
docker-compose up -d
```
	
## Compatible clients
**Important**: default signIn verify code is **12345**

[Android client for Teamgram](clients/teamgram-android.md)

[iOS client for Teamgram](clients/teamgram-ios.md)

[tdesktop for Teamgram](clients/teamgram-tdesktop.md)

## Feedback
Please report bugs, concerns, suggestions by issues, or join telegram group [Teamgram中文社区](https://t.me/+S1_22-6EM1BaffXS) Or [Teamgram](https://t.me/+TjD5LZJ5XLRlCYLF) to discuss problems around source code.

## Notes
If need enterprise edition:

- sticker/theme/wallpaper/reactions/2fa/secretchat/sms/push(apns/web/fcm)/web...
- channel/megagroup
- audiocall/videocall/groupcall/`rtmp live stream`
- bots

please PM the **[author](https://t.me/benqi)**

## Give a Star! ⭐

If you like or are using this project to learn or start your solution, please give it a star. Thanks!
