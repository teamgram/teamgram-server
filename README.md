# NebulaChat - Open source [mtproto](https://core.telegram.org/mtproto) server written in golang
> 打造高性能、稳定并且功能完善的开源mtproto服务端，建设开源telegram客户端生态系统非官方首选服务！

## Chinese

### 简介
Open source [mtproto](https://core.telegram.org/mtproto) server written in golang

### 架构图
![架构图](doc/image/architecture-001.jpeg)

### 文档
[Diffie–Hellman key exchange](doc/dh-key-exchange.md)

[Creating an Authorization Key](doc/Creating_an_Authorization_Key.md)

[Mobile Protocol: Detailed Description (v.1.0, DEPRECATED)](doc/Mobile_Protocol-Detailed_Description_v.1.0_DEPRECATED.md)

[Encrypted CDNs for Speed and Security](doc/cdn.md) [@steedfly](https://github.com/steedfly)翻译

### 编译和安装
#### 简单安装
- 准备
    ```
    mkdir $GOPATH/src/github.com/nebula-chat/
    cd $GOPATH/src/github.com/nebula-chat/
    git clone https://github.com/nebula-chat/chatengine.git
    ```

- 编译代码
    ```
    编译frontend
        cd $GOPATH/src/github.com/nebula-chat/chatengine/access/frontend
        go build
    
    编译auth_key
        cd $GOPATH/src/github.com/nebula-chat/chatengine/access/auth_key
        go build

    编译auth_session
        cd $GOPATH/src/github.com/nebula-chat/chatengine/service/auth_session
        go build
        
    编译sync
        cd $GOPATH/src/github.com/nebula-chat/chatengine/messenger/sync
        go build
    
    编译upload
        cd $GOPATH/src/github.com/nebula-chat/chatengine/messenger/upload
        go build
    
    编译document
        cd $GOPATH/src/github.com/nebula-chat/chatengine/service/document
        go build

    编译biz_server
        cd $GOPATH/src/github.com/nebula-chat/chatengine/messenger/biz_server
        go build
        
    编译session
        cd $GOPATH/src/github.com/nebula-chat/chatengine/access/session
        go build
    ```

- 运行
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

#### 更多文档
[Build document](doc/build.md)

[Build script](scripts/build.sh)

[Prerequisite script](scripts/prerequisite.sh)


### TODO
- Secret Chats
- bots
- payments

### 技术交流群
Bug反馈，意见和建议欢迎加入[Telegramd中文技术交流群](https://t.me/joinchat/D8b0DQ-CrvZXZv2dWn310g)讨论。

## English

### Introduce
open source mtproto server implement by golang, which compatible telegram client.

### Install
[Build and install](doc/build.md)

[build](scripts/build.sh)

[prerequisite](scripts/prerequisite.sh)

## Feedback
Please report bugs, concerns, suggestions by issues, or join telegram group [Telegramd](https://t.me/joinchat/D8b0DRJiuH8EcIHNZQmCxQ) to discuss problems around source code.

