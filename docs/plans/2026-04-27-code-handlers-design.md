# Code Service Handler Implementation Design

Date: 2026-04-27
Status: draft

## Overview

Implement `internal/core/*_handler.go` for the `app/service/biz/code` service, replacing stubs that return `tg.ErrMethodNotImpl` with real business logic. The code service manages phone verification codes via KV cache (Redis).

## Architecture

```
code_service_impl.go (generated wrapper — logs request/error/reply)
  └─ core/*_handler.go (business logic, error mapping)
       └─ repository/*.go (semantic methods: GetCachePhoneCode/PutCachePhoneCode/DeleteCachePhoneCode)
            └─ xkv/phone_code_model.go (KV adapter: kv.ExtStore + JSON serialization)
```

## Layer Rules

| Layer | Allowed | Not Allowed |
|-------|---------|-------------|
| `core/handler` | Business decisions, service semantic errors (`code.Err*`), log only swallowed/rewritten errors | `tg.Err*`, duplicate logging |
| `repository` | Aggregate persistence, translate xkv errors → service errors | `tg.Err*`, duplicate logging |
| `xkv` | KV get/put/delete, JSON encode/decode, TTL management | Business logic, `tg.Err*`, log returned errors |

## Files

### New
- `code/errors.go` — `ErrPhoneCodeExpired`, `ErrPhoneCodeInvalid`, `ErrCodeStorage`
- `internal/repository/xkv/phone_code_model.go` — `PhoneCodeModel` interface + `kv.ExtStore` impl
- `internal/core/code_handlers_test.go` — Table-driven tests with in-memory mock

### Modified
- `internal/config/config.go` — Add `KV kv.KvConf`
- `internal/repository/repository.go` — Init `kv.ExtStore` + `PhoneCodeModel`, add semantic methods
- `internal/core/code.createPhoneCode_handler.go`
- `internal/core/code.getPhoneCode_handler.go`
- `internal/core/code.deletePhoneCode_handler.go`
- `internal/core/code.updatePhoneCodeData_handler.go`

## Handler Logic

| Handler | Behavior |
|---------|----------|
| `createPhoneCode` | GetCache → if missing or session changed, generate new code (5-digit + 16-byte hex nonce, 180s TTL) → return |
| `getPhoneCode` | GetCache → validate hash → return data / `ErrPhoneCodeExpired` / `ErrPhoneCodeInvalid` |
| `deletePhoneCode` | DeleteCache → return `BoolTrue` |
| `updatePhoneCodeData` | PutCache → return `BoolTrue` / `ErrCodeStorage` |

## Logging Strategy

- **No log** for: pass-through errors (wrapper handles), success paths
- **Errorf log** for: swallowed errors (cache fail → generate new code), rewritten errors (NotFound → `ErrPhoneCodeExpired`)

## Error Contract

Internal RPC service returns service semantic errors only. BFF layer maps them to `tg.ErrPhoneCodeExpired` / `tg.ErrPhoneCodeInvalid`.

## Testing

Hand-written mock store (in-memory `map[string][]byte`), table-driven tests covering: cache miss, cache hit, session change, hash mismatch, storage failure.
