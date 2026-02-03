# Gnetway Build Error Fixes

## 🐛 Fixed Compilation Errors

### 1. Unused Import: `github.com/lxzan/gws` in message_handler.go
**Error:**
```
message_handler.go:25:2: "github.com/lxzan/gws" imported and not used
```

**Fix:**
- Removed unused import from `message_handler.go`
- gws is only used in `websocket_handler.go`

**File:** `internal/server/netserver/message_handler.go`

---

### 2. Integer Overflow: constant -404 overflows uint32
**Error:**
```
message_handler.go:243:50: constant -404 overflows uint32
```

**Fix:**
```go
// Before:
binary.LittleEndian.PutUint32(out2, uint32(-404))

// After:
var code int32 = -404
binary.LittleEndian.PutUint32(out2, uint32(code))
```

**Explanation:**
- `-404` is a negative int constant
- Direct cast `uint32(-404)` causes overflow
- First convert to `int32`, then to `uint32` (preserves bit pattern)

**File:** `internal/server/netserver/message_handler.go:242`

---

### 3. Undefined Method: WriteBinary on gws.Conn
**Error:**
```
message_handler.go:284:21: c.gwsConn.WriteBinary undefined
```

**Fix:**
```go
// Before:
return c.gwsConn.WriteBinary(data)

// After:
return c.gwsConn.WriteMessage(1, data) // 1 = OpcodeBinary
```

**Explanation:**
- lxzan/gws uses `WriteMessage(opcode, data)` instead of `WriteBinary(data)`
- Opcode 1 = Binary frame (per WebSocket RFC 6455)

**File:** `internal/server/netserver/message_handler.go:284`

---

### 4. Unknown Field: CompressEnabled in gws.ServerOption
**Error:**
```
websocket_handler.go:121:3: unknown field CompressEnabled in struct literal
```

**Fix:**
```go
// Before:
&gws.ServerOption{
    ReadBufferSize:   65536,
    WriteBufferSize:  65536,
    CompressEnabled:  false,
    CheckUtf8Enabled: false,
}

// After:
&gws.ServerOption{
    ReadBufferSize:  65536,
    WriteBufferSize: 65536,
}
```

**Explanation:**
- gws v1.8.9 doesn't have `CompressEnabled` or `CheckUtf8Enabled` fields
- Removed unsupported fields from ServerOption struct

**File:** `internal/server/netserver/websocket_handler.go:118-123`

---

### 5. Incorrect Upgrade Arguments
**Error:**
```
websocket_handler.go:125:32: not enough arguments in call to upgrader.Upgrade
    have (net.Conn)
    want (http.ResponseWriter, *http.Request)
```

**Fix:**
Created custom WebSocket upgrade handler:

```go
// New method: handleWebSocketUpgrade
func (s *Server) handleWebSocketUpgrade(c *connection, upgrader *gws.Upgrader) error {
    // 1. Read HTTP upgrade request
    req, err := http.ReadRequest(c.reader)
    if err != nil {
        return fmt.Errorf("failed to read HTTP request: %w", err)
    }

    // 2. Create custom ResponseWriter
    respWriter := &responseWriter{
        conn:   c,
        header: make(http.Header),
    }

    // 3. Perform upgrade
    socket, err := upgrader.Upgrade(respWriter, req)
    if err != nil {
        return fmt.Errorf("failed to upgrade: %w", err)
    }

    c.gwsConn = socket
    socket.ReadLoop()
    return nil
}

// Custom http.ResponseWriter implementation
type responseWriter struct {
    conn       *connection
    header     http.Header
    statusCode int
    written    bool
}
```

**Explanation:**
- gws requires `http.ResponseWriter` and `*http.Request` for upgrade
- Implemented custom `responseWriter` that writes to `net.Conn`
- Reads HTTP upgrade request from connection using `http.ReadRequest`
- Writes HTTP response headers manually to connection

**File:** `internal/server/netserver/websocket_handler.go:125-270`

---

### 6. Unused Import: `github.com/panjf2000/gnet/v2` in handshake.go
**Error:**
```
handshake.go:40:2: "github.com/panjf2000/gnet/v2" imported and not used
```

**Fix:**
- Removed `"github.com/panjf2000/gnet/v2"` import
- This file was migrated from gnet package but no longer uses gnet

**File:** `internal/server/netserver/handshake.go:40`

---

### 7. Unused Import: `context` in server.go
**Error:**
```
server.go:18:2: "context" imported and not used
```

**Fix:**
- Removed unused `"context"` import
- Not needed in the netserver implementation

**File:** `internal/server/netserver/server.go:18`

---

## 📦 Build Result

```bash
$ go build ./app/interface/gnetway/cmd/gnetway
Build successful!
-rwxr-xr-x  1 user  staff  98M Feb  3 15:19 gnetway
```

**Binary Size:** 98 MB
**Status:** ✅ All errors fixed, builds successfully

---

## 🔍 Technical Details

### WebSocket Upgrade Flow

The new implementation handles WebSocket upgrades as follows:

```
1. Accept TCP connection
2. Read HTTP GET request with Upgrade header
   ┌─────────────────────────────────────┐
   │ GET /chat HTTP/1.1                  │
   │ Host: server.example.com            │
   │ Upgrade: websocket                  │
   │ Connection: Upgrade                 │
   │ Sec-WebSocket-Key: ...              │
   │ Sec-WebSocket-Version: 13           │
   └─────────────────────────────────────┘
3. Validate upgrade request (gws.Upgrader)
4. Write HTTP 101 Switching Protocols response
   ┌─────────────────────────────────────┐
   │ HTTP/1.1 101 Switching Protocols    │
   │ Upgrade: websocket                  │
   │ Connection: Upgrade                 │
   │ Sec-WebSocket-Accept: ...           │
   └─────────────────────────────────────┘
5. Switch to WebSocket protocol
6. Handle WebSocket frames (via gws)
```

### Custom ResponseWriter Implementation

The `responseWriter` type implements `http.ResponseWriter` interface:

| Method | Implementation |
|--------|----------------|
| `Header()` | Returns `http.Header` map |
| `Write([]byte)` | Writes data to `net.Conn` |
| `WriteHeader(int)` | Writes HTTP status line and headers |

This allows gws to perform the upgrade without requiring a full `http.Server`.

---

## ✅ Testing

### Verify Build
```bash
go build ./app/interface/gnetway/cmd/gnetway
```

### Run Server
```bash
./gnetway -f app/interface/gnetway/etc/gnetway.yaml
```

### Test TCP Connection
```bash
telnet localhost 10443
```

### Test WebSocket Connection
```bash
wscat -c ws://localhost:8080
```

---

## 📚 Dependencies

No changes to dependencies. Still using:
- `github.com/lxzan/gws v1.8.9`
- `github.com/dolthub/maphash v0.1.0`

---

## 🎯 Next Steps

1. ✅ Build verification
2. ⏭️ Integration testing with Telegram clients
3. ⏭️ Performance benchmarking
4. ⏭️ Production deployment

---

## 📝 Related Files

- `MIGRATION.md` - Complete migration documentation
- `internal/server/netserver/` - New netserver package
- `internal/server/gnet/` - Old gnet package (deprecated)

---

## 👥 Fixed By

Claude Code - February 3, 2026
