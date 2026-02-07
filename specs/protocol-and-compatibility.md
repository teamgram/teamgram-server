# 协议与兼容性

本文档说明 Teamgram Server 所支持的 MTProto 子协议、API Layer 及与官方客户端的兼容范围。

## MTProto 2.0

Teamgram 实现 [MTProto 2.0](https://core.telegram.org/mtproto)，支持以下传输格式：

- **Abridged**：压缩格式，常见于移动端
- **Intermediate**：中间格式
- **Padded intermediate**：带填充的中间格式
- **Full**：完整格式

客户端与 gnetway 建立 TCP 或 WebSocket 连接后，通过上述任一格式进行握手与后续 RPC 通信。

## API Layer

- 当前支持的 **API Layer** 为 **220**，与 Telegram 官方 API 的对应版本一致。
- 具体 TL 方法与类型以项目依赖的 [teamgram/proto](https://github.com/teamgram/proto) 为准；若官方 API 升级，需同步更新 proto 与服务端实现。

## 兼容客户端

以下客户端可对接 Teamgram Server，需将服务器地址与端口改为自建实例：

| 客户端 | 说明 |
|--------|------|
| [Android (teamgram-android)](../clients/teamgram-android.md) | 基于 Telegram 官方 Android 客户端修改，需修改 `ConnectionsManager.cpp` 中的地址与端口 |
| [iOS (teamgram-ios)](../clients/teamgram-ios.md) | iOS 客户端，需配置连接至自建服务器 |
| [TDesktop (teamgram-tdesktop)](../clients/teamgram-tdesktop.md) | 桌面端，需配置连接至自建服务器 |

**重要**：默认登录验证码为 **12345**，仅用于开发与测试；生产环境应配置自己的验证方式（见企业版或自改逻辑）。

## 部署相关说明

- 客户端需能够访问 gnetway 对外端口（默认 10443 / 11443 / 5222），并关闭 TLS 或使用自签名证书时按客户端要求配置。
- 若通过 Nginx 等反向代理暴露 WebSocket，需保证协议与路径与 gnetway 的 WebSocket 配置一致。
