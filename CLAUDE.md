# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

TeamGram Server (`teamgram-server`) is a complete, production-ready Telegram-compatible server implementation written in Go. It implements MTProto (Telegram's binary protocol) and provides a full microservices architecture for running a Telegram-like messaging platform.

**Key characteristics:**
- Go 1.22+ module-based project
- Microservices architecture with 12+ independent services
- Protocol layer (~99% generated from TL schema files via `mtprotoc`)
- Uses CloudWeGo Kitex framework for RPC
- Uses etcd for service discovery and configuration
- Licensed under Apache License 2.0
- Primary module: `github.com/teamgooo/teamgooo-server`
- Version: v0.211.0-teamgooo-server

## Development Commands

### Building All Services

**Build all services at once:**
```bash
make all
# Builds: idgen, status, dfs, media, authsession, biz, msg, sync, bff, session, gnetway
# Output directory: /opt/data/teamgooo/bin/
```

**Clean all built binaries:**
```bash
make clean
```

### Building Individual Services

**Infrastructure Services:**
```bash
make idgen        # ID generator service
make status       # Status service
make dfs          # Distributed file system
make media        # Media processing service
```

**Authentication & Session:**
```bash
make authsession  # Authentication session service
make session      # Session management
```

**Business Logic:**
```bash
make biz          # Core business logic (users, chats, etc.)
```

**Messaging:**
```bash
make msg          # Message service
make inbox        # Inbox service (not in 'all' target)
make sync         # Sync service
```

**Frontend Services:**
```bash
make bff          # Backend-for-Frontend (API gateway)
make gnetway      # Gateway (MTProto protocol handler)
```

### Running Services

**Run a service directly:**
```bash
go run ./app/service/biz/biz/cmd/biz
go run ./app/interface/gnetway/cmd/gnetway
# etc.
```

**Run built binaries:**
```bash
/opt/data/teamgooo/bin/biz
/opt/data/teamgooo/bin/gnetway
# etc.
```

**Note:** Services require proper configuration files in their `etc/` directories and etcd to be running.

### Testing

**Run all tests:**
```bash
go test ./...
```

**Run tests in a specific package:**
```bash
go test ./pkg/proto/crypto/
go test ./pkg/proto/bin/
```

**Run a specific test:**
```bash
go test -v -run TestAES256IGECryptor ./pkg/proto/crypto/
go test -v -run TestInt128 ./pkg/proto/bin/
```

**Run tests with coverage:**
```bash
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

### Dependency Management

**Download dependencies:**
```bash
go mod download
```

**Update dependencies:**
```bash
go get -u ./...
go mod tidy
```

**Verify dependencies:**
```bash
go mod verify
```

## Architecture

### High-Level Microservices Architecture

TeamGram Server follows a layered microservices architecture:

```
┌─────────────────────────────────────────────────────────────┐
│  Interface Layer (MTProto Gateway)                          │
│  ├── gnetway     - MTProto protocol handler & gateway       │
│  └── session     - Session management                       │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  BFF Layer (Backend-for-Frontend)                           │
│  ├── bff         - Core BFF service                         │
│  ├── account, authorization, contacts, dialogs              │
│  ├── chats, messages, files, drafts                         │
│  ├── users, userprofile, usernames                          │
│  ├── notification, updates, qrcode                          │
│  ├── premium, privacysettings, chatinvites                  │
│  └── ... (26 BFF modules total)                             │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  Messenger Layer                                            │
│  ├── msg         - Message service (includes inbox)         │
│  └── sync        - Client synchronization service           │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  Service Layer (Business Logic & Infrastructure)            │
│  ├── biz         - Core business logic (users, chats, etc.) │
│  ├── authsession - Authentication & session management      │
│  ├── dfs         - Distributed file system                  │
│  ├── media       - Media processing & storage               │
│  ├── status      - User status tracking                     │
│  └── idgen       - Distributed ID generator                 │
└─────────────────────────────────────────────────────────────┘
                            ↓
┌─────────────────────────────────────────────────────────────┐
│  Protocol Layer (MTProto Implementation)                    │
│  pkg/proto/ - See "Protocol Package Structure" below        │
└─────────────────────────────────────────────────────────────┘
```

### Application Structure

```
app/
├── interface/        - Interface layer (client-facing services)
│   ├── gnetway/      - MTProto gateway (TCP/HTTP connections)
│   └── session/      - Session management service
│
├── bff/              - Backend-for-Frontend layer (26 modules)
│   ├── bff/          - Core BFF service
│   ├── account/      - Account management
│   ├── authorization/- Auth & login flows
│   ├── contacts/     - Contact list management
│   ├── dialogs/      - Dialog/conversation list
│   ├── messages/     - Message operations
│   ├── chats/        - Chat/group management
│   ├── users/        - User operations
│   ├── files/        - File upload/download
│   └── ...           - (17 more modules)
│
├── messenger/        - Messaging layer
│   ├── msg/          - Message routing & delivery
│   │   ├── msg/      - Main message service
│   │   └── inbox/    - Inbox service
│   └── sync/         - Client sync & push notifications
│
└── service/          - Core services
    ├── biz/          - Business logic service
    │   └── user/     - User business logic
    ├── authsession/  - Auth session management
    ├── dfs/          - Distributed file system
    ├── media/        - Media processing
    ├── status/       - Status service
    └── idgen/        - ID generation
```

### Protocol Package Structure

```
pkg/proto/
├── bin/          - Binary protocol layer (TL serialization/deserialization)
├── crypto/       - MTProto cryptography (AES-256-IGE, RSA, SRP, hashing)
├── iface/        - Core TLObject interface and type registry
├── mt/           - MTProto low-level protocol types (key exchange, sessions)
├── rpc/          - Kitex RPC framework integration and custom codec
│   └── examples/ - Working examples (echo server/client)
├── tg/           - High-level Telegram API types (322k lines, 96% of codebase)
└── tgapi/        - API aggregation wrapper layer
```

### Utility Packages

```
pkg/
├── code/         - Error codes & status codes
├── conf/         - Configuration utilities
├── env2/         - Environment utilities
├── fasttime/     - Fast time utilities
├── goffmpeg/     - FFmpeg wrapper for media processing
├── hashx/        - Hashing utilities
├── mention/      - Mention parsing (@username, #hashtag)
├── net/          - Network utilities
├── phonenumber/  - Phone number validation
└── proto/        - Protocol implementation (see above)
```

### Service Communication Flow

```
Client (MTProto)
    ↓
gnetway (Protocol Handler)
    ↓
session (Session Validation)
    ↓
bff/* (API Endpoints)
    ↓
msg/sync (Messaging)
    ↓
biz (Business Logic)
    ↓
dfs/media/status/idgen (Infrastructure Services)
```

### Key Concepts

**TLObject Interface** (`pkg/proto/iface/iface.go`):
- Core abstraction for all protocol objects
- Methods: `Encode(x *bin.Encoder, layer int32) error` and `Decode(d *bin.Decoder) error`
- All generated types implement this interface

**Code Generation:**
- TL schema files → `mtprotoc` → Generated `.tl.go` files
- Generated files have warning header: `Created from 'scheme.tl' by 'mtprotoc'`
- Never manually edit generated files (`.tl.go`, `*_registers.tl.go`, `schema.tl.*.pb.go`)

**Binary Protocol** (`pkg/proto/bin/`):
- TL protocol uses 4-byte word alignment (`WordLen = 4`)
- `Encoder` serializes TLObject → binary
- `Decoder` deserializes binary → TLObject
- Uses `bytebufferpool` for efficient buffer management

**Cryptography** (`pkg/proto/crypto/`):
- AES-256-IGE: Primary cipher for MTProto (`aes256_ige_cryptor.go`)
- AES-CBC and AES-CTR: Alternative modes
- RSA: Initial key exchange (`rsa_cryptor.go`)
- SRP: Secure Remote Password protocol (`srp_util.go`)
- Utilities: SHA256, SHA512, PBKDF2, random number generation

**MTProto Types** (`pkg/proto/mt/`):
- Low-level protocol messages: `req_pq`, `req_DH_params`, `set_client_DH_params`
- Session management: `new_session_created`, `destroy_session`, `msgs_ack`
- Error handling: `rpc_error`, `bad_msg_notification`, `bad_server_salt`

**Telegram API** (`pkg/proto/tg/`):
- High-level Telegram types: users, chats, channels, messages, files
- Auth states: `AuthStateNew`, `AuthStateNormal`, `AuthStateNeedPassword`, etc.
- Peer types: `PEER_USER`, `PEER_CHAT`, `PEER_CHANNEL`
- Auth key types: `AuthKeyTypePerm`, `AuthKeyTypeTemp`, `AuthKeyTypeMediaTemp`
- Helper utilities: `peer_util.go`, `auth_helper.go`, `message_build_help.go`

**RPC Integration** (`pkg/proto/rpc/`):
- Custom `ZRpcCodec` for Kitex framework
- Wraps TL binary protocol in JSON metadata
- See `pkg/proto/rpc/examples/echo/` for complete server/client implementation

### Service Structure Pattern

All services in `app/` follow a consistent structure:

```
app/{layer}/{service}/
├── cmd/{service}/
│   └── main.go              # Entry point
├── internal/
│   ├── config/              # Configuration structs
│   ├── core/                # Business logic implementation
│   ├── dao/                 # Data access objects
│   ├── dal/                 # Data access layer
│   ├── server/              # Server initialization
│   └── svc/                 # Service context
├── etc/
│   └── {service}.yaml       # Configuration file
└── {service}.go             # Package marker
```

**Standard service main.go pattern:**
```go
package main

import (
    "github.com/teamgram/marmota/pkg/commands"
    "github.com/teamgooo/teamgooo-server/app/{layer}/{service}/internal/server"
)

func main() {
    commands.Run(server.New())
}
```

**Configuration file format (YAML):**
```yaml
Name: service.name
ListenOn: 127.0.0.1:20110
Etcd:
  Hosts:
    - 127.0.0.1:2379
  Key: service.name
# Service-specific config...
```

### Example: Echo Service (RPC Example)

```
pkg/proto/rpc/examples/echo/
├── cmd/echo/main.go         # Server entry point
├── client/t/main.go         # Client entry point
├── echo/                    # Generated API types
├── internal/
│   ├── config/              # Configuration
│   ├── core/                # Business logic
│   ├── server/              # Server initialization
│   └── svc/                 # Service context
└── etc/etc.yaml             # Config file
```

**Echo server pattern:**
```go
svr := server.New()
_ = svr.Initialize()
svr.RunLoop()
svr.Destroy()
```

**Echo client pattern:**
```go
zCodec := codec.NewZRpcCodec(true)
cli, err := echo.NewClient("echo",
    client.WithHostPorts("0.0.0.0:8888"),
    client.WithCodec(zCodec))
req := &api.TLEchoEcho{...}
resp, err := cli.EchosEcho(context.Background(), req)
```

## Important Implementation Notes

### Generated vs Manual Code

**Never modify:**
- Files with `Created from 'scheme.tl' by 'mtprotoc'` header
- `*.tl.go` files (generated type definitions)
- `*_registers.tl.go` files (type registries)
- `schema.tl.*.pb.go` files (protobuf marshaling)

**Safe to modify:**
- Helper files: `*_helper.go`, `*_util.go`
- Implementation files in `internal/` directories
- Test files: `*_test.go`
- Configuration files

### Type Registration

All TL types must be registered with a class ID:
```go
const ClazzID_echo = 0x05162463

func init() {
    iface.RegisterClazzID(ClazzID_echo, func() iface.TLObject {
        return &TLEcho{}
    })
}
```

### Encoding/Decoding Pattern

```go
// Encoding
x := bin.NewEncoder()
_ = tlObject.Encode(x, layer)
payload := x.Bytes()

// Decoding
d := bin.NewDecoder(payload)
tlObject := &TLSomeType{}
_ = tlObject.Decode(d)
```

### Working with Peer Types

Use constants from `pkg/proto/tg/peer_util.go`:
- `PEER_EMPTY`, `PEER_SELF`, `PEER_USER`, `PEER_CHAT`, `PEER_CHANNEL`
- Helper functions available for peer manipulation

### Authentication States

Auth flow states in `pkg/proto/tg/auth_helper.go`:
- `AuthStateNew` → `AuthStateUnauthorized` → `AuthStateNormal`
- May require `AuthStateNeedPassword` for 2FA

### Cryptographic Operations

**AES-256-IGE encryption:**
```go
key := []byte("32-byte-key...")
iv := []byte("32-byte-iv...")
cryptor := crypto.NewAES256IGECryptor(key, iv)
encrypted, _ := cryptor.Encrypt(plaintext)
decrypted, _ := cryptor.Decrypt(encrypted)
```

**Access hash generation:**
```go
hash := crypto.GenerateAccessHash(userId, salt)
```

## External Dependencies

**Core frameworks:**
- `github.com/cloudwego/kitex v0.13.1` - RPC framework
- `github.com/teamgram/marmota` - Service framework & utilities
- `golang.org/x/crypto v0.28.0` - Cryptographic primitives

**Infrastructure:**
- **etcd** - Service discovery and configuration (required for all services)
- Uses Kitex for inter-service RPC communication

**Utilities:**
- `github.com/valyala/bytebufferpool` - Buffer pooling for encoders
- `github.com/apache/thrift v0.13.0` - IDL support
- `github.com/stretchr/testify v1.9.0` - Testing utilities

## Project Context

This is a **complete, production-ready Telegram-compatible server implementation**, not just a protocol library. It provides:

**Protocol Layer (`pkg/proto/`):**
- MTProto binary protocol encoding/decoding
- Complete cryptographic suite (AES-256-IGE, RSA, SRP)
- Generated Telegram API type definitions
- RPC framework integration (Kitex)

**Application Layer (`app/`):**
- 12+ independent microservices
- Complete implementation of Telegram features:
  - User authentication & authorization
  - Message sending/receiving/routing
  - Group chats & channels
  - File/media upload/download/processing
  - Contact management
  - Push notifications & sync
  - User profiles & status
  - And much more...

**Infrastructure:**
- Distributed ID generation
- Distributed file system
- Media processing (FFmpeg integration)
- Service discovery (etcd)
- Session management

This server can run independently and handle Telegram-compatible clients.

## BFF (Backend-for-Frontend) Modules

The BFF layer contains 26 modules that implement Telegram API endpoints:

**Core & Authentication:**
- `bff` - Main BFF service coordinator
- `account` - Account settings & management
- `authorization` - Login, logout, signup flows

**User & Social:**
- `users` - User operations & info
- `userprofile` - Profile management
- `usernames` - Username handling
- `contacts` - Contact list management

**Messaging:**
- `messages` - Send, edit, delete messages
- `dialogs` - Conversation list
- `drafts` - Draft messages
- `savedmessagedialogs` - Saved messages

**Group Communication:**
- `chats` - Group chat operations
- `chatinvites` - Invite link management

**Content & Media:**
- `files` - File upload/download
- `autodownload` - Auto-download settings

**Settings & Privacy:**
- `privacysettings` - Privacy controls
- `notification` - Notification settings
- `configuration` - App configuration

**Features:**
- `updates` - Update delivery
- `qrcode` - QR code login
- `premium` - Premium features
- `nsfw` - Content filtering
- `passport` - Telegram Passport
- `tos` - Terms of Service
- `sponsoredmessages` - Sponsored content
- `miscellaneous` - Misc utilities

Each BFF module follows the standard service structure with `cmd/`, `internal/`, and `etc/` directories.

## Development Notes

### Prerequisites

1. **etcd must be running** - All services use etcd for service discovery
2. **Dependencies installed** - Run `go mod download`
3. **Build directory** - Default output: `/opt/data/teamgooo/bin/`

### Service Startup Order

For a complete deployment, services should start in this order:

1. **Infrastructure:**
   - `idgen` - ID generation (required by most services)
   - `status` - Status tracking
   - `dfs` - File system
   - `media` - Media processing

2. **Core Services:**
   - `authsession` - Authentication
   - `biz` - Business logic

3. **Messaging:**
   - `msg` (and `inbox`)
   - `sync`

4. **Frontend:**
   - `bff` - API gateway
   - `session` - Session management
   - `gnetway` - MTProto gateway (client entry point)

### Configuration

Each service has a YAML config file in its `etc/` directory:
- etcd connection settings
- Service-specific ports and addresses
- Logging configuration
- Database connections (if applicable)

**Important:** Update etcd hosts, service addresses, and RSA key paths before running in production.

### Build Flags

The Makefile uses these build-time variables:
- `version` - Version string (v0.211.0-teamgooo-server)
- `gitTag` - Git tag
- `gitBranch` - Git branch
- `gitCommit` - Git commit hash
- `buildDate` - Build timestamp
- `gitTreeState` - Clean or dirty working tree

Built with `-ldflags` to embed version info and `-tags=jsoniter` for faster JSON.

### Development Workflow

1. Make changes to service code in `app/{layer}/{service}/internal/`
2. Build the service: `make {service}`
3. Update config if needed: `app/{layer}/{service}/etc/{service}.yaml`
4. Run the service: `/opt/data/teamgooo/bin/{service}`
5. Check logs and etcd registration

### Testing

Run protocol layer tests:
```bash
# Test specific crypto operations
go test -v ./pkg/proto/crypto/ -run TestAES256IGE

# Test binary encoding
go test -v ./pkg/proto/bin/

# All protocol tests
go test ./pkg/proto/...
```

**Note:** Application service tests may require database and etcd to be running.
