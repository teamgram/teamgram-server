# 协议与兼容性

本文档说明 Teamgram Server 所支持的 MTProto 子协议、API Layer 及与官方客户端的兼容范围。

## MTProto 2.0

Teamgram 实现 [MTProto 2.0](https://core.telegram.org/mtproto)，支持以下传输格式：

- **Abridged**
- **Intermediate**
- **Padded intermediate**
- **Full**

客户端通过 TCP 或 WebSocket 连接 gnetway，使用上述任一格式进行握手与 RPC。

## API Layer

- 当前支持的 **API Layer** 为 **222**（见主 README；随 proto 更新可能变动）。
- TL 方法与类型以项目依赖的 [teamgram/proto](https://github.com/teamgram/proto) 为准。

## 兼容客户端

| 客户端 | 说明 |
|--------|------|
| [Android (teamgram-android)](../clients/teamgram-android.md) | 修改 ConnectionsManager.cpp 中的服务器地址与端口 |
| [iOS (teamgram-ios)](../clients/teamgram-ios.md) | 配置连接自建服务器 |
| [桌面端 (teamgram-tdesktop)](../clients/teamgram-tdesktop.md) | 配置连接自建服务器 |

**重要**：默认登录验证码为 **12345**（仅开发用；生产环境请修改）。

## 部署说明

- 客户端需能访问 gnetway 端口（默认 10443、11443、5222）。TLS 或反向代理（如 Nginx WebSocket）需按需配置。
