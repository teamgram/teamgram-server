# Code Service Handler Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Implement `internal/core/*_handler.go` for the `app/service/biz/code` service — replacing stubs with real phone code cache logic via KV storage.

**Architecture:** Layered: generated wrapper (`code_service_impl.go`) logs at RPC boundary → core handlers own business decisions and error mapping → repository provides semantic cache methods (`GetCachePhoneCode`/`PutCachePhoneCode`/`DeleteCachePhoneCode`) → xkv adapter wraps `kv.ExtStore` with JSON serialization.

**Tech Stack:** Go 1.22+, `github.com/teamgram/marmota/pkg/stores/kv` (kv.ExtStore), `github.com/zeromicro/go-zero/core/stores/kv` (kv.KvConf), `github.com/teamgram/teamgram-server/v2/pkg/proto/crypto` (crypto.GenerateStringNonce), `crypto/rand` (randomNumeric)

**Conventions:** Follow `error-handling-guidelines.md`, `logging-guidelines.md`, `repository-design-guidelines.md`. Service semantic errors in `code/errors.go`, no `tg.Err*` in service layer, no duplicate logging in handlers.

---

## Prep: Read Conventions

Read the three architecture docs before starting:
- `/Users/wubenqi/go/src/teamgram.io/tgdocs-dev/server/teamgram-server-v2/architecture/error-handling-guidelines.md`
- `/Users/wubenqi/go/src/teamgram.io/tgdocs-dev/server/teamgram-server-v2/architecture/logging-guidelines.md`
- `/Users/wubenqi/go/src/teamgram.io/tgdocs-dev/server/teamgram-server-v2/architecture/repository-design-guidelines.md`

Also keep open for reference:
- `app/service/authsession/internal/repository/xkv/future_salts_model.go` — xkv pattern reference
- `app/service/idgen/idgen/errors.go` — service semantic errors pattern

---

### Task 1: Add KV config field

**Files:**
- Modify: `app/service/biz/code/internal/config/config.go`

**Step 1: Add KV and go-zero kv imports, add KV field to Config struct**

```go
package config

import (
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex"
	"github.com/zeromicro/go-zero/core/stores/kv"
)

type Config struct {
	kitex.RpcServerConf
	KV kv.KvConf
}
```

**Step 2: Update etc/code.yaml to include KV config**

Add a KV section to `app/service/biz/code/etc/code.yaml`:

```yaml
Name: service.biz_service.code
ListenOn: 0.0.0.0:29500
KV:
  - Host: 127.0.0.1:6379
    Pass: ""
    Type: node
```

**Step 3: Verify compilation**

Run: `go build ./app/service/biz/code/code/...`
Expected: Compiled successfully

**Step 4: Commit**

```bash
git add app/service/biz/code/internal/config/config.go app/service/biz/code/etc/code.yaml
git commit -m "feat(code): add KV config field"
```

---

### Task 2: Create service semantic errors

**Files:**
- Create: `app/service/biz/code/code/errors.go`

**Step 1: Write errors.go**

```go
package code

import "errors"

var (
	ErrPhoneCodeExpired = errors.New("code: phone code expired")
	ErrPhoneCodeInvalid = errors.New("code: phone code invalid")
	ErrCodeStorage      = errors.New("code: storage failure")
)
```

**Step 2: Verify compilation**

Run: `go build ./app/service/biz/code/code/...`
Expected: Compiled successfully

**Step 3: Commit**

```bash
git add app/service/biz/code/code/errors.go
git commit -m "feat(code): add service semantic errors"
```

---

### Task 3: Create xkv phone code model

**Files:**
- Create: `app/service/biz/code/internal/repository/xkv/phone_code_model.go`

**Step 1: Create xkv directory**

Run: `mkdir -p app/service/biz/code/internal/repository/xkv`

**Step 2: Write phone_code_model.go**

```go
package xkv

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	// phoneCodeDefaultTTL is the default TTL for phone codes (3 minutes).
	phoneCodeDefaultTTL = 180
)

// PhoneCodeModel abstracts KV operations for phone verification codes.
type PhoneCodeModel interface {
	GetPhoneCode(ctx context.Context, authKeyId int64, phone string) (*code.PhoneCodeTransaction, error)
	PutPhoneCode(ctx context.Context, authKeyId int64, phone string, data *code.PhoneCodeTransaction) error
	DeletePhoneCode(ctx context.Context, authKeyId int64, phone string) error
}

type phoneCodeModel struct {
	kv     kv.ExtStore
	prefix string
}

// NewPhoneCodeModel creates a kv-backed phone code model.
func NewPhoneCodeModel(kv kv.ExtStore, prefix string) PhoneCodeModel {
	return &phoneCodeModel{
		kv:     kv,
		prefix: prefix,
	}
}

func (m *phoneCodeModel) cacheKey(authKeyId int64, phone string) string {
	if m.prefix == "" {
		return fmt.Sprintf("phone_code#%d:%s", authKeyId, phone)
	}
	return fmt.Sprintf("%s:phone_code#%d:%s", m.prefix, authKeyId, phone)
}

func (m *phoneCodeModel) GetPhoneCode(ctx context.Context, authKeyId int64, phone string) (*code.PhoneCodeTransaction, error) {
	val, err := m.kv.GetCtx(ctx, m.cacheKey(authKeyId, phone))
	if err != nil {
		return nil, fmt.Errorf("phone_code.GetPhoneCode kv get: %w", err)
	}
	if val == "" {
		return nil, nil
	}

	var txn code.PhoneCodeTransaction
	if err := json.Unmarshal([]byte(val), &txn); err != nil {
		logx.WithContext(ctx).Errorf("phone_code.GetPhoneCode json.Unmarshal(%s) error(%v)", val, err)
		return nil, nil
	}

	return &txn, nil
}

func (m *phoneCodeModel) PutPhoneCode(ctx context.Context, authKeyId int64, phone string, data *code.PhoneCodeTransaction) error {
	if data == nil {
		return nil
	}

	b, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("phone_code.PutPhoneCode json marshal: %w", err)
	}

	return m.kv.SetexCtx(ctx, m.cacheKey(authKeyId, phone), string(b), phoneCodeDefaultTTL)
}

func (m *phoneCodeModel) DeletePhoneCode(ctx context.Context, authKeyId int64, phone string) error {
	_, err := m.kv.DelCtx(ctx, m.cacheKey(authKeyId, phone))
	return err
}
```

**Step 3: Verify compilation**

Run: `go build ./app/service/biz/code/...`
Expected: Compiled successfully

**Step 4: Commit**

```bash
git add app/service/biz/code/internal/repository/xkv/phone_code_model.go
git commit -m "feat(code): add xkv phone code model"
```

---

### Task 4: Wire repository with KV store and semantic methods

**Files:**
- Modify: `app/service/biz/code/internal/repository/repository.go`
- Modify: `app/service/biz/code/internal/repository/repository_type.go`

**Step 1: Update repository.go with kv store, phone code model, and semantic methods**

Replace the entire file content:

```go
package repository

import (
	"github.com/teamgram/marmota/pkg/stores/kv"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/internal/config"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/internal/repository/xkv"
)

// Repository is the dependency container for repository instances.
type Repository struct {
	kv             kv.ExtStore
	phoneCodeModel PhoneCodeModelType
}

// NewRepository creates a new Repository.
func NewRepository(c config.Config) *Repository {
	kv2 := kv.NewStore(c.KV)

	return &Repository{
		kv:             kv2,
		phoneCodeModel: xkv.NewPhoneCodeModel(kv2, "code"),
	}
}

// Close releases repository-owned clients.
func (r *Repository) Close() error {
	if r == nil {
		return nil
	}
	return nil
}
```

**Step 2: Create repo_phone_code.go with semantic methods**

Create `app/service/biz/code/internal/repository/repo_phone_code.go`:

```go
package repository

import (
	"context"
	"fmt"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
)

// GetCachePhoneCode retrieves a phone code from cache.
func (r *Repository) GetCachePhoneCode(ctx context.Context, authKeyId int64, phone string) (*code.PhoneCodeTransaction, error) {
	txn, err := r.phoneCodeModel.GetPhoneCode(ctx, authKeyId, phone)
	if err != nil {
		return nil, fmt.Errorf("%w: get phone code cache: %w", code.ErrCodeStorage, err)
	}
	return txn, nil
}

// PutCachePhoneCode stores a phone code into cache.
func (r *Repository) PutCachePhoneCode(ctx context.Context, authKeyId int64, phone string, data *code.PhoneCodeTransaction) error {
	if err := r.phoneCodeModel.PutPhoneCode(ctx, authKeyId, phone, data); err != nil {
		return fmt.Errorf("%w: put phone code cache: %w", code.ErrCodeStorage, err)
	}
	return nil
}

// DeleteCachePhoneCode removes a phone code from cache.
func (r *Repository) DeleteCachePhoneCode(ctx context.Context, authKeyId int64, phone string) error {
	if err := r.phoneCodeModel.DeletePhoneCode(ctx, authKeyId, phone); err != nil {
		return fmt.Errorf("%w: delete phone code cache: %w", code.ErrCodeStorage, err)
	}
	return nil
}
```

**Step 3: Update repository_type.go with type aliases**

Replace:

```go
package repository

// Type aliases for convenience in the Logic layer.
type (
	// PhoneCodeModelType is the type alias for phone code model.
	PhoneCodeModelType = xkv.PhoneCodeModel
)
```

This needs the import: add `"github.com/teamgram/teamgram-server/v2/app/service/biz/code/internal/repository/xkv"`

**Step 4: Verify compilation**

Run: `go build ./app/service/biz/code/...`
Expected: Compiled successfully

**Step 5: Commit**

```bash
git add app/service/biz/code/internal/repository/repository.go app/service/biz/code/internal/repository/repo_phone_code.go app/service/biz/code/internal/repository/repository_type.go
git commit -m "feat(code): wire repository with KV store and semantic phone code methods"
```

---

### Task 5: Write failing handler tests

**Files:**
- Create: `app/service/biz/code/internal/core/code_handlers_test.go`

**Step 1: Write test file**

Follow the idgen handler test pattern (`app/service/idgen/internal/core/idgen_handlers_test.go`) — hand-rolled mock, no framework.

```go
package core

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/internal/svc"
)

// mockRepo implements the subset of repository methods used by core handlers.
type mockRepo struct {
	store map[string][]byte // key -> JSON of code.PhoneCodeTransaction
	getErr error
	putErr error
	delErr error
}

func (m *mockRepo) cacheKey(authKeyId int64, phone string) string {
	return fmt.Sprintf("phone_code#%d:%s", authKeyId, phone)
}

func (m *mockRepo) GetCachePhoneCode(ctx context.Context, authKeyId int64, phone string) (*code.PhoneCodeTransaction, error) {
	if m.getErr != nil {
		return nil, m.getErr
	}
	val, ok := m.store[m.cacheKey(authKeyId, phone)]
	if !ok || val == nil {
		return nil, nil
	}
	var txn code.PhoneCodeTransaction
	if err := json.Unmarshal(val, &txn); err != nil {
		return nil, err
	}
	return &txn, nil
}

func (m *mockRepo) PutCachePhoneCode(ctx context.Context, authKeyId int64, phone string, data *code.PhoneCodeTransaction) error {
	if m.putErr != nil {
		return m.putErr
	}
	b, _ := json.Marshal(data)
	m.store[m.cacheKey(authKeyId, phone)] = b
	return nil
}

func (m *mockRepo) DeleteCachePhoneCode(ctx context.Context, authKeyId int64, phone string) error {
	delete(m.store, m.cacheKey(authKeyId, phone))
	return m.delErr
}

// repoInterface is the interface core expects from the real Repository.
// Using an interface lets the test inject a mock without depending on the
// concrete repository type.
type repoInterface interface {
	GetCachePhoneCode(ctx context.Context, authKeyId int64, phone string) (*code.PhoneCodeTransaction, error)
	PutCachePhoneCode(ctx context.Context, authKeyId int64, phone string, data *code.PhoneCodeTransaction) error
	DeleteCachePhoneCode(ctx context.Context, authKeyId int64, phone string) error
}

func newTestCore(t *testing.T, repo repoInterface) *CodeCore {
	t.Helper()
	return New(context.Background(), &svc.ServiceContext{
		Repo: repo,
	})
}

func TestCreatePhoneCode_NewCode(t *testing.T) {
	c := newTestCore(t, &mockRepo{store: make(map[string][]byte)})

	in := &code.TLCodeCreatePhoneCode{
		AuthKeyId:    1,
		SessionId:    100,
		Phone:        "1234567890",
		SentCodeType: 1,
	}

	result, err := c.CodeCreatePhoneCode(in)
	if err != nil {
		t.Fatalf("CodeCreatePhoneCode() err = %v", err)
	}
	if result == nil {
		t.Fatal("CodeCreatePhoneCode() result is nil")
	}
	if result.PhoneCode == "" {
		t.Error("PhoneCode is empty")
	}
	if result.PhoneCodeHash == "" {
		t.Error("PhoneCodeHash is empty")
	}
	if result.PhoneCodeExpired <= int32(time.Now().Unix()) {
		t.Error("PhoneCodeExpired is in the past")
	}
	if result.State != codeStateSend {
		t.Errorf("State = %d, want %d", result.State, codeStateSend)
	}
}

func TestCreatePhoneCode_SessionChanged(t *testing.T) {
	store := make(map[string][]byte)
	existing, _ := json.Marshal(&code.TLPhoneCodeTransaction{
		AuthKeyId:     1,
		SessionId:     200,
		Phone:         "1234567890",
		PhoneCode:     "12345",
		PhoneCodeHash: "abc",
	})
	store["phone_code#1:1234567890"] = existing

	c := newTestCore(t, &mockRepo{store: store})

	in := &code.TLCodeCreatePhoneCode{
		AuthKeyId:    1,
		SessionId:    100,
		Phone:        "1234567890",
		SentCodeType: 1,
	}

	result, err := c.CodeCreatePhoneCode(in)
	if err != nil {
		t.Fatalf("CodeCreatePhoneCode() err = %v", err)
	}
	if result.SessionId != 100 {
		t.Errorf("SessionId = %d, want 100", result.SessionId)
	}
}

func TestGetPhoneCode_Success(t *testing.T) {
	store := make(map[string][]byte)
	existing, _ := json.Marshal(&code.TLPhoneCodeTransaction{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "abc",
	})
	store["phone_code#1:1234567890"] = existing

	c := newTestCore(t, &mockRepo{store: store})

	result, err := c.CodeGetPhoneCode(&code.TLCodeGetPhoneCode{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "abc",
	})
	if err != nil {
		t.Fatalf("CodeGetPhoneCode() err = %v", err)
	}
	if result.PhoneCodeHash != "abc" {
		t.Errorf("PhoneCodeHash = %s, want abc", result.PhoneCodeHash)
	}
}

func TestGetPhoneCode_Expired(t *testing.T) {
	c := newTestCore(t, &mockRepo{store: make(map[string][]byte)})

	_, err := c.CodeGetPhoneCode(&code.TLCodeGetPhoneCode{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "abc",
	})
	if !errors.Is(err, code.ErrPhoneCodeExpired) {
		t.Errorf("err = %v, want ErrPhoneCodeExpired", err)
	}
}

func TestGetPhoneCode_InvalidHash(t *testing.T) {
	store := make(map[string][]byte)
	existing, _ := json.Marshal(&code.TLPhoneCodeTransaction{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "abc",
	})
	store["phone_code#1:1234567890"] = existing

	c := newTestCore(t, &mockRepo{store: store})

	_, err := c.CodeGetPhoneCode(&code.TLCodeGetPhoneCode{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "wrong",
	})
	if !errors.Is(err, code.ErrPhoneCodeInvalid) {
		t.Errorf("err = %v, want ErrPhoneCodeInvalid", err)
	}
}

func TestGetPhoneCode_StorageError(t *testing.T) {
	storageErr := errors.New("redis connection refused")
	c := newTestCore(t, &mockRepo{store: make(map[string][]byte), getErr: storageErr})

	_, err := c.CodeGetPhoneCode(&code.TLCodeGetPhoneCode{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "abc",
	})
	if !errors.Is(err, code.ErrPhoneCodeExpired) {
		t.Errorf("err = %v, want ErrPhoneCodeExpired", err)
	}
}

func TestDeletePhoneCode(t *testing.T) {
	c := newTestCore(t, &mockRepo{store: make(map[string][]byte)})

	result, err := c.CodeDeletePhoneCode(&code.TLCodeDeletePhoneCode{
		AuthKeyId:     1,
		Phone:         "1234567890",
		PhoneCodeHash: "abc",
	})
	if err != nil {
		t.Fatalf("CodeDeletePhoneCode() err = %v", err)
	}
	if result == nil {
		t.Fatal("CodeDeletePhoneCode() result is nil")
	}
}

func TestUpdatePhoneCodeData_StorageError(t *testing.T) {
	storageErr := errors.New("redis connection refused")
	c := newTestCore(t, &mockRepo{store: make(map[string][]byte), putErr: storageErr})

	_, err := c.CodeUpdatePhoneCodeData(&code.TLCodeUpdatePhoneCodeData{
		AuthKeyId: 1,
		Phone:     "1234567890",
		CodeData: code.MakeTLPhoneCodeTransaction(&code.TLPhoneCodeTransaction{
			AuthKeyId: 1,
			Phone:     "1234567890",
		}),
	})
	if !errors.Is(err, code.ErrCodeStorage) {
		t.Errorf("err = %v, want ErrCodeStorage", err)
	}
}
```

**Note:** The test injects mock via an interface. Since the real `CodeCore` holds `*svc.ServiceContext` which has `*repository.Repository`, we need to either:
- Option A: Define a `RepoInterface` in core and store as interface — cleanest but changes core.go
- Option B: Use the concrete mock by converting the test to be aware of the real type structure

Option A is preferred. This means `core.go` will change — `svcCtx.Repo` becomes an interface.

**Step 2: Run tests to verify they fail**

Run: `go test ./app/service/biz/code/internal/core/... -v -count=1`
Expected: Compilation errors about undefined `codeStateSend` and `Repo` type mismatch

**Step 3: Commit**

```bash
git add app/service/biz/code/internal/core/code_handlers_test.go
git commit -m "test(code): add failing handler tests"
```

---

### Task 6: Add code state constant and helper

**Files:**
- Create: `app/service/biz/code/internal/core/code_state.go` (or inline in handler helper)
- Modify: `app/service/biz/code/internal/core/core.go` — add `Repo` interface type

**Step 1: Define Repo interface in core.go**

Update `core.go` to use an interface for the repository, enabling test injection:

```go
package core

import (
	"context"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/internal/svc"
	"github.com/teamgram/teamgram-server/v2/pkg/net/kitex/metadata"

	"github.com/zeromicro/go-zero/core/logx"
)

// Repo is the interface CodeCore expects from the repository layer.
type Repo interface {
	GetCachePhoneCode(ctx context.Context, authKeyId int64, phone string) (*code.PhoneCodeTransaction, error)
	PutCachePhoneCode(ctx context.Context, authKeyId int64, phone string, data *code.PhoneCodeTransaction) error
	DeleteCachePhoneCode(ctx context.Context, authKeyId int64, phone string) error
}

type CodeCore struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   Repo
	logx.Logger
	MD *metadata.RpcMetadata
}

func New(ctx context.Context, svcCtx *svc.ServiceContext) *CodeCore {
	return &CodeCore{
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   svcCtx.Repo,
		Logger: logx.WithContext(ctx),
		MD:     metadata.RpcMetadataFromIncoming(ctx),
	}
}
```

**Step 2: Run tests — still fail with codeStateSend undefined**

Run: `go test ./app/service/biz/code/internal/core/... -v -count=1 -run TestCreate`
Expected: FAIL — `undefined: codeStateSend`

**Step 3: Add code state constant**

Add a helper file or inline constant. Since it's a small constant, add in the test file first helper (or add to the create handler):

In the create handler, add before handler function:

```go
const codeStateSend = 1 // CodeStateSend — code has been generated/sent
```

**Step 4: Commit**

```bash
git add app/service/biz/code/internal/core/core.go
git commit -m "refactor(code): extract Repo interface in core for testability"
```

---

### Task 7: Implement createPhoneCode handler

**Files:**
- Modify: `app/service/biz/code/internal/core/code.createPhoneCode_handler.go`

**Step 1: Write implementation**

```go
package core

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
)

const (
	codeStateSend      = 1
	codeLen            = 5
	codeHashLen        = 16
	codeExpireDuration = 3 * 60 // seconds
)

func randomNumeric(n int) string {
	const digits = "0123456789"
	b := make([]byte, n)
	for i := range b {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		b[i] = digits[idx.Int64()]
	}
	return string(b)
}

func randomHexString(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}

// CodeCreatePhoneCode
// code.createPhoneCode flags:# auth_key_id:long session_id:long phone:string phone_number_registered:flags.0?true sent_code_type:int next_code_type:int state:int = PhoneCodeTransaction;
func (c *CodeCore) CodeCreatePhoneCode(in *code.TLCodeCreatePhoneCode) (*code.PhoneCodeTransaction, error) {
	codeData, err := c.repo.GetCachePhoneCode(c.ctx, in.AuthKeyId, in.Phone)
	if err != nil {
		c.Logger.Errorf("code.createPhoneCode - get cache failed: auth_key_id: %d, phone: %s, err: %v",
			in.AuthKeyId, in.Phone, err)
	}

	if codeData == nil || in.SessionId != codeData.SessionId {
		codeData = code.MakeTLPhoneCodeTransaction(&code.TLPhoneCodeTransaction{
			AuthKeyId:             in.AuthKeyId,
			SessionId:             in.SessionId,
			Phone:                 in.Phone,
			PhoneNumberRegistered: in.PhoneNumberRegistered,
			PhoneCode:             randomNumeric(codeLen),
			PhoneCodeHash:         randomHexString(codeHashLen),
			PhoneCodeExpired:      int32(time.Now().Unix() + codeExpireDuration),
			SentCodeType:          in.SentCodeType,
			FlashCallPattern:      "*",
			NextCodeType:          in.NextCodeType,
			State:                 codeStateSend,
		})
	}

	return codeData, nil
}
```

**Step 2: Run create tests**

Run: `go test ./app/service/biz/code/internal/core/... -v -count=1 -run TestCreate`
Expected: PASS

**Step 3: Commit**

```bash
git add app/service/biz/code/internal/core/code.createPhoneCode_handler.go
git commit -m "feat(code): implement createPhoneCode handler"
```

---

### Task 8: Implement getPhoneCode handler

**Files:**
- Modify: `app/service/biz/code/internal/core/code.getPhoneCode_handler.go`

**Step 1: Write implementation**

```go
package core

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
)

// CodeGetPhoneCode
// code.getPhoneCode auth_key_id:long phone:string phone_code_hash:string = PhoneCodeTransaction;
func (c *CodeCore) CodeGetPhoneCode(in *code.TLCodeGetPhoneCode) (*code.PhoneCodeTransaction, error) {
	codeData, err := c.repo.GetCachePhoneCode(c.ctx, in.AuthKeyId, in.Phone)
	if err != nil {
		c.Logger.Errorf("code.getPhoneCode - get cache failed: auth_key_id: %d, phone: %s, err: %v",
			in.AuthKeyId, in.Phone, err)
		return nil, code.ErrPhoneCodeExpired
	}
	if codeData == nil {
		c.Logger.Errorf("code.getPhoneCode - not found: auth_key_id: %d, phone: %s",
			in.AuthKeyId, in.Phone)
		return nil, code.ErrPhoneCodeExpired
	}
	if codeData.PhoneCodeHash != in.PhoneCodeHash {
		c.Logger.Errorf("code.getPhoneCode - hash mismatch: auth_key_id: %d, phone: %s",
			in.AuthKeyId, in.Phone)
		return nil, code.ErrPhoneCodeInvalid
	}

	return codeData, nil
}
```

**Step 2: Run getPhoneCode tests**

Run: `go test ./app/service/biz/code/internal/core/... -v -count=1 -run TestGetPhoneCode`
Expected: PASS

**Step 3: Commit**

```bash
git add app/service/biz/code/internal/core/code.getPhoneCode_handler.go
git commit -m "feat(code): implement getPhoneCode handler"
```

---

### Task 9: Implement deletePhoneCode handler

**Files:**
- Modify: `app/service/biz/code/internal/core/code.deletePhoneCode_handler.go`

**Step 1: Write implementation**

```go
package core

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// CodeDeletePhoneCode
// code.deletePhoneCode auth_key_id:long phone:string phone_code_hash:string = Bool;
func (c *CodeCore) CodeDeletePhoneCode(in *code.TLCodeDeletePhoneCode) (*tg.Bool, error) {
	_ = c.repo.DeleteCachePhoneCode(c.ctx, in.AuthKeyId, in.Phone)
	return tg.BoolTrue, nil
}
```

**Step 2: Run delete test**

Run: `go test ./app/service/biz/code/internal/core/... -v -count=1 -run TestDeletePhoneCode`
Expected: PASS

**Step 3: Commit**

```bash
git add app/service/biz/code/internal/core/code.deletePhoneCode_handler.go
git commit -m "feat(code): implement deletePhoneCode handler"
```

---

### Task 10: Implement updatePhoneCodeData handler

**Files:**
- Modify: `app/service/biz/code/internal/core/code.updatePhoneCodeData_handler.go`

**Step 1: Write implementation**

```go
package core

import (
	"github.com/teamgram/teamgram-server/v2/app/service/biz/code/code"
	"github.com/teamgram/teamgram-server/v2/pkg/proto/tg"
)

// CodeUpdatePhoneCodeData
// code.updatePhoneCodeData auth_key_id:long phone:string phone_code_hash:string code_data:PhoneCodeTransaction = Bool;
func (c *CodeCore) CodeUpdatePhoneCodeData(in *code.TLCodeUpdatePhoneCodeData) (*tg.Bool, error) {
	var phoneCodeData *code.TLPhoneCodeTransaction
	if cd, ok := in.CodeData.(*code.TLPhoneCodeTransaction); ok {
		phoneCodeData = cd
	} else {
		phoneCodeData = code.MakeTLPhoneCodeTransaction(&code.TLPhoneCodeTransaction{
			AuthKeyId: in.AuthKeyId,
			Phone:     in.Phone,
		})
	}

	if err := c.repo.PutCachePhoneCode(c.ctx, in.AuthKeyId, in.Phone, phoneCodeData); err != nil {
		return nil, err
	}

	return tg.BoolTrue, nil
}
```

Note: The `PhoneCodeTransactionClazz` is an interface type (`*TLPhoneCodeTransaction`). The `in.CodeData` field is of type `PhoneCodeTransactionClazz` which is `*TLPhoneCodeTransaction`. We can type-assert or use directly.

**Step 2: Run updatePhoneCodeData test**

Run: `go test ./app/service/biz/code/internal/core/... -v -count=1 -run TestUpdatePhoneCodeData`
Expected: PASS

**Step 3: Commit**

```bash
git add app/service/biz/code/internal/core/code.updatePhoneCodeData_handler.go
git commit -m "feat(code): implement updatePhoneCodeData handler"
```

---

### Task 11: Run full test suite and verify

**Step 1: Run all tests with race detector**

Run: `go test -race ./app/service/biz/code/... -v -count=1`
Expected: All tests PASS

**Step 2: Build all services**

Run: `go build ./app/service/biz/code/...`
Expected: Successful

**Step 3: Check no tg.Err* leaks in service layer**

Run: `grep -r "tg.Err" app/service/biz/code/internal/`
Expected: No matches (only `tg.Bool`/`tg.BoolTrue` for return types, NO `tg.Err*`)

**Step 4: Commit**

```bash
git add -A
git commit -m "chore(code): finalize handler implementations and tests"
```

---

### Task 12: Review against conventions

**Step 1: Review logging compliance**

For each handler, verify:
- `createPhoneCode`: Only logs swallowed error (cache get failure), per logging guidelines
- `getPhoneCode`: Only logs rewritten errors (not-found → ErrPhoneCodeExpired, hash mismatch → ErrPhoneCodeInvalid), per logging guidelines
- `deletePhoneCode`: No handler logs (errors swallowed), per logging guidelines
- `updatePhoneCodeData`: No handler logs (errors passed through), per logging guidelines

**Step 2: Review error contract**

Verify:
- Handlers return `code.Err*` (service semantic errors), never `tg.Err*`
- Repository returns `code.ErrCodeStorage` wrapped with `%w`
- BFF layer will map `code.Err*` → `tg.Err*` at protocol boundary

**Step 3: Review repository design**

Verify:
- xkv owns only KV get/put/delete + serialization
- Repository owns aggregate cache semantics
- Core accesses only `c.repo.GetCachePhoneCode` etc., not xkv internals

**Step 4: Final commit if any fixes needed**
