# Teamgram - Open source [mtproto](https://core.telegram.org/mtproto) server written in golang
> open source mtproto server implemented in golang with compatible telegram client.

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
#### Depends
- mysql
- redis
- etcd
- kafka
- minio

#### Build

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
cd teamgramd/bin
./runall2.sh
```
#### More


#### **Note**

### Compatible clients
**Important**: default signIn and signOut verify code is **12345**

[Android client for Teamgram](https://github.com/teamgram/teamgram-android)

[iOS client for Teamgram](https://github.com/teamgram/teamgram-ios)

[tdesktop for Teamgram](https://github.com/teamgram/teamgram-tdesktop)

### TODO

## Feedback
Please report bugs, concerns, suggestions by issues.

## Notes
If need enterprise edition, please PM the [author](https://t.me/benqi)
