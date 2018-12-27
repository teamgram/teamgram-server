# NebulaChat - Open source [mtproto](https://core.telegram.org/mtproto) server written in golang
> open source mtproto server implement by golang, which compatible telegram client.

### Introduce
Open source [mtproto](https://core.telegram.org/mtproto) server written in golang

### Architecture
![Architecture](doc/image/architecture-001.jpeg)

### Documents
[Diffie–Hellman key exchange](doc/dh-key-exchange.md)

[Creating an Authorization Key](doc/Creating_an_Authorization_Key.md)

[Mobile Protocol: Detailed Description (v.1.0, DEPRECATED)](doc/Mobile_Protocol-Detailed_Description_v.1.0_DEPRECATED.md)

[Encrypted CDNs for Speed and Security](doc/cdn.md) [@steedfly](https://github.com/steedfly)翻译

### Build and Install
#### Build
- Depends
    - redis
    - mysql
    - etcd

- Get source code　
```
mkdir $GOPATH/src/github.com/nebula-chat/
cd $GOPATH/src/github.com/nebula-chat/
git clone https://github.com/nebula-chat/chatengine.git

```

- Build
    ```
    build frontend
        cd $GOPATH/src/github.com/nebula-chat/chatengine/access/frontend
        go build
    
    build auth_key
        cd $GOPATH/src/github.com/nebula-chat/chatengine/access/auth_key
        go build

    build auth_session
        cd $GOPATH/src/github.com/nebula-chat/chatengine/service/auth_session
        go build
        
    build sync
        cd $GOPATH/src/github.com/nebula-chat/chatengine/messenger/sync
        go build
    
    build upload
        cd $GOPATH/src/github.com/nebula-chat/chatengine/messenger/upload
        go build
    
    build document
        cd $GOPATH/src/github.com/nebula-chat/chatengine/service/document
        go build

    build biz_server
        cd $GOPATH/src/github.com/nebula-chat/chatengine/messenger/biz_server
        go build
        
    build session
        cd $GOPATH/src/github.com/nebula-chat/chatengine/access/session
        go build
    ```

- Run
    ```
    cd $GOPATH/src/github.comnebula-chat/chatengine/service/auth_session
    ./auth_session
    
    cd $GOPATH/src/github.com/nebula-chat/chatengine/service/document
    ./document

    cd $GOPATH/src/github.com/nebula-chat/chatengine/messenger/sync
    ./sync
    
    cd $GOPATH/src/github.com/nebula-chat/chatengine/messenger/upload
    ./upload

    cd $GOPATH/src/github.com/nebula-chat/chatengine/messenger/biz_server
    ./biz_server

    cd $GOPATH/src/github.com/nebula-chat/chatengine/access/auth_key
    ./auth_key

    cd $GOPATH/src/github.com/nebula-chat/chatengine/access/session
    ./session
    
    cd $GOPATH/src/github.com/nebula-chat/chatengine/access/frontend
    ./frontend
    ```

#### More
[Build document](doc/build.md)

[Build script](scripts/build.sh)

[Prerequisite script](scripts/prerequisite.sh)


### Compatible clients
[Android client for NebulaChat](Telegram-Android)

[FOSS client for NebulaChat](Telegram-FOSS)

[iOS client for NebulaChat](Telegram-iOS)

[tdesktop for NebulaChat](tdesktop)


### TODO
- stickers
- phone call
- channel
- megagroup
- secret chats
- bots

## Feedback
Please report bugs, concerns, suggestions by issues, or join telegram group [Telegramd中文技术交流群](https://t.me/cntelegramd) Or [Telegramd](https://t.me/entelegramd) to discuss problems around source code.

