# privacysettings 模块实现设计

## 背景

参考 `master` 分支的 `app/bff/privacysettings`，在 `v2` 分支上实现完整的 Kitex 版本。

## 对比 master 的主要变更

| 项 | master | v2 |
|---|--------|----|
| RPC 框架 | gRPC (go-zero zrpc) | Kitex |
| 依赖注入 | `internal/dao/` | `internal/repository/` |
| 协议类型 | `mtproto.*` | `tg.*` |
| 同步推送 | SyncClient | TODO（不接入） |

## 修改文件清单

### 1. Config (`internal/config/config.go`)
添加三个 RPC 客户端配置字段：
```go
type Config struct {
    kitex.RpcServerConf
    UserClient        kitex.RpcClientConf
    ChatClient        kitex.RpcClientConf
    AuthsessionClient kitex.RpcClientConf
}
```

### 2. YAML (`etc/privacysettings.yaml`)
对应三个 client 的 etcd 配置。

### 3. Repository (`internal/repository/repository.go`)
```go
type Repository struct {
    UserClient        userclient.UserClient
    ChatClient        chatclient.ChatClient
    AuthsessionClient authsessionclient.AuthsessionClient
}
```

### 4. Core Handlers (8 个文件)
全部从 master 移植业务逻辑，适配 tg 类型：
- `account.getPrivacy_handler.go`
- `account.setPrivacy_handler.go` — SyncClient 部分替换为 TODO
- `account.getGlobalPrivacySettings_handler.go`
- `account.setGlobalPrivacySettings_handler.go`
- `messages.getDefaultHistoryTTL_handler.go`
- `messages.setDefaultHistoryTTL_handler.go`
- `users.getRequirementsToContact_handler.go`
- `users.getIsPremiumRequiredToContact_handler.go` — **新增**

### 5. Service Impl (`internal/server/tg/service/privacysettings_service_impl.go`)
新增 `UsersGetIsPremiumRequiredToContact` forwarder 方法（enterprise-blocked 模式）。

## 不做的变更
- 不接入 SyncClient（AccountSetPrivacy 中 sync 部分写 TODO）
- AuthsessionClient 虽加入 Repository 但尚未被 handler 使用（同 master）
- 不修改已生成的 Kitex 代码（`privacysettings/privacysettingsservice/`）
