# privacysettings Module Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Implement the privacysettings BFF module on v2 (Kitex) by porting master's business logic.

**Architecture:** Single BFF module with Repository holding RPC clients (UserClient, ChatClient, AuthsessionClient). 8 core handlers handle privacy rules, global privacy settings, default TTL, and enterprise-blocked methods. SyncClient calls in AccountSetPrivacy are replaced with TODO comments.

**Tech Stack:** Go, Kitex RPC, go-zero logx, Teamgram proto/tg types

---

### Task 1: Add RPC Client Configs

**Files:**
- Modify: `app/bff/privacysettings/internal/config/config.go`
- Modify: `app/bff/privacysettings/etc/privacysettings.yaml`

**Step 1: Update config.go**

Add three `kitex.RpcClientConf` fields:

```go
package config

import (
    "github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
)

type Config struct {
    kitex.RpcServerConf
    UserClient        kitex.RpcClientConf
    ChatClient        kitex.RpcClientConf
    AuthsessionClient kitex.RpcClientConf
}
```

**Step 2: Update privacysettings.yaml**

Add client configurations with etcd service keys:

```yaml
Name: bff.privacysettings
ListenOn: 0.0.0.0:27716
Etcd:
  Hosts:
    - 127.0.0.1:2379
  Key: bff.privacysettings

UserClient:
  DestService: service.biz.user
  Codec: zrpc
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: service.biz.user

ChatClient:
  DestService: service.biz.chat
  Codec: zrpc
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: service.biz.chat

AuthsessionClient:
  DestService: service.authsession
  Codec: zrpc
  Etcd:
    Hosts:
      - 127.0.0.1:2379
    Key: service.authsession
```

**Step 3: Build check**

Run: `go build ./app/bff/privacysettings/cmd/privacysettings/`
Expected: success

**Step 4: Commit**

```bash
git add app/bff/privacysettings/internal/config/config.go app/bff/privacysettings/etc/privacysettings.yaml
git commit -m "feat(privacysettings): add RPC client configs for user, chat, authsession"
```

---

### Task 2: Wire RPC Clients in Repository

**Files:**
- Modify: `app/bff/privacysettings/internal/repository/repository.go`
- Modify: `app/bff/privacysettings/internal/repository/repository_type.go`

**Step 1: Update repository.go**

Wire UserClient, ChatClient, and AuthsessionClient:

```go
package repository

import (
    authsessionclient "github.com/teamgram/teamgram-server/v2/app/service/authsession/client"
    chatclient "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/client"
    userclient "github.com/teamgram/teamgram-server/v2/app/service/biz/user/client"

    "github.com/teamgram/teamgram-server/v2/app/bff/privacysettings/internal/config"
)

type Repository struct {
    UserClient        userclient.UserClient
    ChatClient        chatclient.ChatClient
    AuthsessionClient authsessionclient.AuthsessionClient
}

func NewRepository(c config.Config) *Repository {
    return &Repository{
        UserClient:        userclient.NewUserClient(userclient.MustNewKitexClient(c.UserClient)),
        ChatClient:        chatclient.NewChatClient(chatclient.MustNewKitexClient(c.ChatClient)),
        AuthsessionClient: authsessionclient.NewAuthsessionClient(authsessionclient.MustNewKitexClient(c.AuthsessionClient)),
    }
}

func (r *Repository) Close() error {
    if r == nil {
        return nil
    }
    return nil
}
```

**Step 2: Build check**

Run: `go build ./app/bff/privacysettings/cmd/privacysettings/`
Expected: success

**Step 3: Commit**

```bash
git add app/bff/privacysettings/internal/repository/
git commit -m "feat(privacysettings): wire UserClient, ChatClient, AuthsessionClient in Repository"
```

---

### Task 3: Implement account.getPrivacy Handler

**Files:**
- Modify: `app/bff/privacysettings/internal/core/account.getPrivacy_handler.go`

**Step 1: Replace stub with implementation**

```go
package core

import (
    "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
    "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
    "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *PrivacySettingsCore) AccountGetPrivacy(in *tg.TLAccountGetPrivacy) (*tg.AccountPrivacyRules, error) {
    key := tg.FromInputPrivacyKeyType(in.Key)

    if key == tg.KEY_TYPE_INVALID {
        c.Logger.Errorf("account.getPrivacy - error: invalid privacy key")
        return nil, tg.ErrPrivacyKeyInvalid
    }

    ruleList, _ := c.svcCtx.Repo.UserClient.UserGetPrivacy(c.ctx, &user.TLUserGetPrivacy{
        UserId:  c.MD.UserId,
        KeyType: int32(key),
    })

    var rVal *tg.AccountPrivacyRules

    if len(ruleList.GetDatas()) == 0 {
        if key == tg.PHONE_NUMBER {
            rVal = tg.MakeTLAccountPrivacyRules(&tg.TLAccountPrivacyRules{
                Rules: []tg.PrivacyRuleClazz{tg.MakeTLPrivacyValueDisallowAll(nil).ToPrivacyRule()},
                Users: []tg.UserClazz{},
                Chats: []tg.ChatClazz{},
            }).ToAccountPrivacyRules()
        } else if key == tg.BIRTHDAY {
            rVal = tg.MakeTLAccountPrivacyRules(&tg.TLAccountPrivacyRules{
                Rules: []tg.PrivacyRuleClazz{tg.MakeTLPrivacyValueAllowContacts(nil).ToPrivacyRule()},
                Users: []tg.UserClazz{},
                Chats: []tg.ChatClazz{},
            }).ToAccountPrivacyRules()
        } else {
            rVal = tg.MakeTLAccountPrivacyRules(&tg.TLAccountPrivacyRules{
                Rules: []tg.PrivacyRuleClazz{tg.MakeTLPrivacyValueAllowAll(nil).ToPrivacyRule()},
                Users: []tg.UserClazz{},
                Chats: []tg.ChatClazz{},
            }).ToAccountPrivacyRules()
        }
    } else {
        rVal = tg.MakeTLAccountPrivacyRules(&tg.TLAccountPrivacyRules{
            Rules: ruleList.GetDatas(),
            Users: []tg.UserClazz{},
            Chats: []tg.ChatClazz{},
        }).ToAccountPrivacyRules()

        idHelper := tg.NewIDListHelper(c.MD.UserId)
        idHelper.PickByRules(ruleList.GetDatas())
        idHelper.Visit(
            func(userIdList []int64) {
                users, _ := c.svcCtx.Repo.UserClient.UserGetMutableUsers(c.ctx,
                    &user.TLUserGetMutableUsers{
                        Id: userIdList,
                    })
                rVal.Users = users.GetUserListByIdList(c.MD.UserId, userIdList...)
            },
            func(chatIdList []int64) {
                chats, _ := c.svcCtx.Repo.ChatClient.ChatGetChatListByIdList(c.ctx,
                    &chat.TLChatGetChatListByIdList{
                        IdList: chatIdList,
                    })
                rVal.Chats = chats.GetChatListByIdList(c.MD.UserId, chatIdList...)
            },
            func(channelIdList []int64) {
                // TODO
            })
    }

    return rVal, nil
}
```

**Step 2: Build check**

Run: `go build ./app/bff/privacysettings/cmd/privacysettings/`
Expected: success

**Step 3: Commit**

```bash
git add app/bff/privacysettings/internal/core/account.getPrivacy_handler.go
git commit -m "feat(privacysettings): implement account.getPrivacy handler"
```

---

### Task 4: Implement account.setPrivacy Handler

**Files:**
- Modify: `app/bff/privacysettings/internal/core/account.setPrivacy_handler.go`

**Step 1: Replace stub with implementation (SyncClient → TODO)**

```go
package core

import (
    "github.com/teamgram/teamgram-server/v2/app/service/biz/chat/chat"
    "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
    "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *PrivacySettingsCore) AccountSetPrivacy(in *tg.TLAccountSetPrivacy) (*tg.AccountPrivacyRules, error) {
    key := tg.FromInputPrivacyKeyType(in.Key)

    if key == tg.KEY_TYPE_INVALID {
        c.Logger.Errorf("account.setPrivacy - error: invalid privacy key")
        return nil, tg.ErrPrivacyKeyInvalid
    }

    ruleList := tg.ToPrivacyRuleListByInput(c.MD.UserId, in.Rules)

    if _, err := c.svcCtx.Repo.UserClient.UserSetPrivacy(c.ctx, &user.TLUserSetPrivacy{
        UserId:  c.MD.UserId,
        KeyType: int32(key),
        Rules:   ruleList,
    }); err != nil {
        c.Logger.Errorf("account.setPrivacy - error: %v", err)
        return nil, err
    }

    rVal := tg.MakeTLAccountPrivacyRules(&tg.TLAccountPrivacyRules{
        Rules: ruleList,
        Users: []tg.UserClazz{},
        Chats: []tg.ChatClazz{},
    }).ToAccountPrivacyRules()

    // TODO: syncUpdatesNotMe
    // syncUpdates := tg.MakeUpdatesByUpdates(tg.MakeTLUpdatePrivacy(&tg.TLUpdate{
    //     Key:   tg.ToPrivacyKey(key),
    //     Rules: ruleList,
    // }).ToUpdate())
    //
    // idHelper := tg.NewIDListHelper(c.MD.UserId)
    // idHelper.PickByRules(ruleList)
    // idHelper.Visit(
    //     func(userIdList []int64) {
    //         users, _ := c.svcCtx.Repo.UserClient.UserGetMutableUsers(c.ctx,
    //             &user.TLUserGetMutableUsers{
    //                 Id: userIdList,
    //             })
    //         rVal.Users = users.GetUserListByIdList(c.MD.UserId, userIdList...)
    //         syncUpdates.PushUser(rVal.Users...)
    //     },
    //     func(chatIdList []int64) {
    //         chats, _ := c.svcCtx.Repo.ChatClient.ChatGetChatListByIdList(c.ctx,
    //             &chat.TLChatGetChatListByIdList{
    //                 IdList: chatIdList,
    //             })
    //         rVal.Chats = chats.GetChatListByIdList(c.MD.UserId, chatIdList...)
    //         syncUpdates.PushChat(rVal.Chats...)
    //     },
    //     func(channelIdList []int64) {
    //         // TODO
    //     })
    //
    // c.svcCtx.Repo.SyncClient.SyncUpdatesNotMe(c.ctx, &sync.TLSyncUpdatesNotMe{
    //     UserId:        c.MD.UserId,
    //     PermAuthKeyId: c.MD.PermAuthKeyId,
    //     Updates:       syncUpdates,
    // })

    return rVal, nil
}
```

**Step 2: Build check**

Run: `go build ./app/bff/privacysettings/cmd/privacysettings/`
Expected: success

**Step 3: Commit**

```bash
git add app/bff/privacysettings/internal/core/account.setPrivacy_handler.go
git commit -m "feat(privacysettings): implement account.setPrivacy handler (sync TODO)"
```

---

### Task 5: Implement account.getGlobalPrivacySettings Handler

**Files:**
- Modify: `app/bff/privacysettings/internal/core/account.getGlobalPrivacySettings_handler.go`

**Step 1: Replace stub**

```go
package core

import (
    "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
    "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *PrivacySettingsCore) AccountGetGlobalPrivacySettings(in *tg.TLAccountGetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error) {
    globalPrivacySettings, err := c.svcCtx.Repo.UserClient.UserGetGlobalPrivacySettings(c.ctx, &user.TLUserGetGlobalPrivacySettings{
        UserId: c.MD.UserId,
    })
    if err != nil {
        c.Logger.Errorf("account.getGlobalPrivacySettings - error: %v", err)
        return nil, err
    }
    return globalPrivacySettings, nil
}
```

**Step 2: Build check**

Run: `go build ./app/bff/privacysettings/cmd/privacysettings/`
Expected: success

**Step 3: Commit**

```bash
git add app/bff/privacysettings/internal/core/account.getGlobalPrivacySettings_handler.go
git commit -m "feat(privacysettings): implement account.getGlobalPrivacySettings handler"
```

---

### Task 6: Implement account.setGlobalPrivacySettings Handler

**Files:**
- Modify: `app/bff/privacysettings/internal/core/account.setGlobalPrivacySettings_handler.go`

**Step 1: Replace stub**

```go
package core

import (
    "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
    "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *PrivacySettingsCore) AccountSetGlobalPrivacySettings(in *tg.TLAccountSetGlobalPrivacySettings) (*tg.GlobalPrivacySettings, error) {
    rSettings := in.GetSettings()

    _, err := c.svcCtx.Repo.UserClient.UserSetGlobalPrivacySettings(c.ctx, &user.TLUserSetGlobalPrivacySettings{
        UserId:   c.MD.UserId,
        Settings: rSettings,
    })
    if err != nil {
        c.Logger.Errorf("account.setGlobalPrivacySettings - error: %v", err)
        return nil, err
    }

    return rSettings, nil
}
```

**Step 2: Build check**

Run: `go build ./app/bff/privacysettings/cmd/privacysettings/`
Expected: success

**Step 3: Commit**

```bash
git add app/bff/privacysettings/internal/core/account.setGlobalPrivacySettings_handler.go
git commit -m "feat(privacysettings): implement account.setGlobalPrivacySettings handler"
```

---

### Task 7: Implement messages.getDefaultHistoryTTL Handler

**Files:**
- Modify: `app/bff/privacysettings/internal/core/messages.getDefaultHistoryTTL_handler.go`

**Step 1: Replace stub**

```go
package core

import (
    "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
    "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *PrivacySettingsCore) MessagesGetDefaultHistoryTTL(in *tg.TLMessagesGetDefaultHistoryTTL) (*tg.DefaultHistoryTTL, error) {
    rV, err := c.svcCtx.Repo.UserClient.UserGetDefaultHistoryTTL(c.ctx, &user.TLUserGetDefaultHistoryTTL{
        UserId: c.MD.UserId,
    })
    if err != nil {
        c.Logger.Errorf("user.getDefaultHistoryTTL - error: %v", err)
        return nil, err
    }
    return rV, nil
}
```

**Step 2: Build check**

Run: `go build ./app/bff/privacysettings/cmd/privacysettings/`
Expected: success

**Step 3: Commit**

```bash
git add app/bff/privacysettings/internal/core/messages.getDefaultHistoryTTL_handler.go
git commit -m "feat(privacysettings): implement messages.getDefaultHistoryTTL handler"
```

---

### Task 8: Implement messages.setDefaultHistoryTTL Handler

**Files:**
- Modify: `app/bff/privacysettings/internal/core/messages.setDefaultHistoryTTL_handler.go`

**Step 1: Replace stub**

```go
package core

import (
    "github.com/teamgram/teamgram-server/v2/app/service/biz/user/user"
    "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *PrivacySettingsCore) MessagesSetDefaultHistoryTTL(in *tg.TLMessagesSetDefaultHistoryTTL) (*tg.Bool, error) {
    period := in.GetPeriod()

    if period < 0 || period > 366*86400 {
        c.Logger.Errorf("messages.setDefaultHistoryTTL - error: invalid period %v", period)
        return nil, tg.ErrTtlPeriodInvalid
    }

    _, _ = c.svcCtx.Repo.UserClient.UserSetDefaultHistoryTTL(c.ctx, &user.TLUserSetDefaultHistoryTTL{
        UserId: c.MD.UserId,
        Ttl:    period,
    })

    return tg.BoolTrue, nil
}
```

**Step 2: Build check**

Run: `go build ./app/bff/privacysettings/cmd/privacysettings/`
Expected: success

**Step 3: Commit**

```bash
git add app/bff/privacysettings/internal/core/messages.setDefaultHistoryTTL_handler.go
git commit -m "feat(privacysettings): implement messages.setDefaultHistoryTTL handler"
```

---

### Task 9: Implement users.getRequirementsToContact Handler

**Files:**
- Modify: `app/bff/privacysettings/internal/core/users.getRequirementsToContact_handler.go`

**Step 1: Replace stub**

```go
package core

import (
    "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *PrivacySettingsCore) UsersGetRequirementsToContact(in *tg.TLUsersGetRequirementsToContact) (*tg.VectorRequirementToContact, error) {
    c.Logger.Errorf("users.getRequirementsToContact blocked, License key from https://teamgram.net required to unlock enterprise features.")
    return nil, tg.ErrEnterpriseIsBlocked
}
```

**Step 2: Build check**

Run: `go build ./app/bff/privacysettings/cmd/privacysettings/`
Expected: success

**Step 3: Commit**

```bash
git add app/bff/privacysettings/internal/core/users.getRequirementsToContact_handler.go
git commit -m "feat(privacysettings): implement users.getRequirementsToContact handler (enterprise blocked)"
```

---

### Task 10: Add users.getIsPremiumRequiredToContact to Core + Service Impl

**Files:**
- Create: `app/bff/privacysettings/internal/core/users.getIsPremiumRequiredToContact_handler.go`
- Modify: `app/bff/privacysettings/internal/server/tg/service/privacysettings_service_impl.go`

**Step 1: Create core handler file**

```go
package core

import (
    "github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

func (c *PrivacySettingsCore) UsersGetIsPremiumRequiredToContact(in *tg.TLUsersGetIsPremiumRequiredToContact) (*tg.VectorBool, error) {
    c.Logger.Errorf("users.getIsPremiumRequiredToContact blocked, License key from https://teamgram.net required to unlock enterprise features.")
    return nil, tg.ErrEnterpriseIsBlocked
}
```

**Step 2: Add service impl method**

Add after UsersGetRequirementsToContact in `privacysettings_service_impl.go`:

```go
// UsersGetIsPremiumRequiredToContact
// users.getIsPremiumRequiredToContact#a622aa10 id:Vector<InputUser> = Vector<Bool>;
func (s *Service) UsersGetIsPremiumRequiredToContact(ctx context.Context, request *tg.TLUsersGetIsPremiumRequiredToContact) (*tg.VectorBool, error) {
    c := core.New(ctx, s.svcCtx)
    c.Logger.Debugf("users.getIsPremiumRequiredToContact - metadata: %s, request: %s", c.MD, request)

    r, err := c.UsersGetIsPremiumRequiredToContact(request)
    if err != nil {
        c.Logger.Errorf("users.getIsPremiumRequiredToContact - error: request: %s, err: %v", request, err)
        return nil, err
    }

    c.Logger.Debugf("users.getIsPremiumRequiredToContact - reply: %s", r)
    return r, err
}
```

**Step 3: Build check**

Run: `go build ./app/bff/privacysettings/cmd/privacysettings/`
Expected: success

**Step 4: Commit**

```bash
git add app/bff/privacysettings/internal/core/users.getIsPremiumRequiredToContact_handler.go app/bff/privacysettings/internal/server/tg/service/privacysettings_service_impl.go
git commit -m "feat(privacysettings): add UsersGetIsPremiumRequiredToContact handler and service impl"
```

---

### Task 11: Full Build Verification

**Step 1: Build all**

Run: `go build ./...`
Expected: success (no compile errors anywhere)

**Step 2: Run tests**

Run: `go test ./app/bff/privacysettings/...`
Expected: success (no test failures, or no tests at all)

**Step 3: Final check — verify all 8 handlers exist**

Run: `grep -c "func (c \*PrivacySettingsCore)" app/bff/privacysettings/internal/core/*.go`
Expected: 8 matches (one per handler file)

**Step 4: Final commit**

```bash
git add -A
git commit -m "feat(privacysettings): complete module implementation with all 8 handlers"
```
