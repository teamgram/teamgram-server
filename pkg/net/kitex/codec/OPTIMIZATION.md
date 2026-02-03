# ZRpcCodec 优化说明

## 优化内容

### 1. 使用二进制序列化代替 JSON

**优化前:**
- Meta 结构使用 JSON 序列化 (`json.Marshal` / `json.Unmarshal`)
- 包格式: `[8字节长度] + [JSON数据]`
- 性能开销大，序列化/反序列化慢

**优化后:**
- Meta 结构使用高性能二进制序列化
- 包格式: `[4字节Magic] + [4字节长度] + [二进制数据]`
- 零反射、零内存分配（编码时仅1次分配）

### 2. 添加 Magic Number

**Magic Number:** `0x5A525043` (ASCII: "ZRPC")

**用途:**
- 协议识别和验证
- 防止错误的数据包被解析
- 提供协议版本控制的基础

## 二进制格式规范

### 包整体格式

```
+------------------+------------------+------------------+
| Magic (4 bytes)  | Length (4 bytes) | Meta Data        |
+------------------+------------------+------------------+
```

### Meta 数据格式

```
+----------------------+------------------+
| ServiceName Length   | 2 bytes (uint16) |
+----------------------+------------------+
| ServiceName          | variable         |
+----------------------+------------------+
| MethodName Length    | 2 bytes (uint16) |
+----------------------+------------------+
| MethodName           | variable         |
+----------------------+------------------+
| SeqID                | 4 bytes (int32)  |
+----------------------+------------------+
| MsgType              | 4 bytes (uint32) |
+----------------------+------------------+
| Payload Length       | 4 bytes (uint32) |
+----------------------+------------------+
| Payload              | variable         |
+----------------------+------------------+
| Metadata Count       | 2 bytes (uint16) |
+----------------------+------------------+
| [For each metadata entry]              |
|   Key Length         | 2 bytes (uint16) |
|   Key                | variable         |
|   Value Length       | 2 bytes (uint16) |
|   Value              | variable         |
+----------------------+------------------+
```

所有多字节整数使用大端序 (Big Endian)。

## 性能对比

### 基准测试结果 (Apple M4 Max)

**编码性能:**
- 操作耗时: ~231.7 ns/op
- 内存分配: 1152 B/op
- 分配次数: 1 allocs/op

**解码性能:**
- 操作耗时: ~314.4 ns/op
- 内存分配: 1504 B/op
- 分配次数: 12 allocs/op

### 相比 JSON 的优势

1. **更快的序列化/反序列化**
   - 无需 JSON 解析器
   - 无反射开销
   - 直接二进制操作

2. **更小的数据包**
   - 二进制格式比 JSON 文本更紧凑
   - 减少网络传输开销

3. **更少的内存分配**
   - 编码仅 1 次内存分配
   - 减少 GC 压力

4. **类型安全**
   - 强类型二进制格式
   - 避免 JSON 类型转换错误

## 向后兼容性

**注意:** 此优化**不向后兼容**旧版本的 ZRpcCodec。

如需平滑迁移:
1. 通过 Magic Number 区分新旧协议
2. 实现协议版本协商机制
3. 或使用独立的服务端口

## 使用示例

```go
// 创建 codec
codec := NewZRpcCodec(true) // true 启用调试日志

// 在 Kitex 客户端使用
cli, err := echo.NewClient("echo",
    client.WithHostPorts("0.0.0.0:8888"),
    client.WithCodec(codec))

// 在 Kitex 服务端使用
svr := echo.NewServer(handler,
    server.WithCodec(codec))
```

## 测试

运行单元测试:
```bash
go test -v ./pkg/net/kitex/codec/
```

运行基准测试:
```bash
go test -bench=. -benchmem ./pkg/net/kitex/codec/
```

## 技术细节

### 编码流程

1. 编码 Payload (TLObject) → 二进制
2. 构建 Meta 结构
3. 二进制序列化 Meta → buffer
4. 写入: Magic Number + Length + Meta Buffer

### 解码流程

1. 读取并验证 Magic Number
2. 读取 Length
3. 读取 Meta Buffer
4. 二进制反序列化 Meta
5. 解码 Payload (TLObject)

### 错误处理

- Magic Number 不匹配: 返回协议错误
- 缓冲区长度不足: 返回格式错误
- TLObject 解码失败: 返回有效载荷错误
- 异常消息: 通过 RpcError 返回

## 限制

- ServiceName / MethodName: 最大 65535 字节
- Metadata 键/值: 最大 65535 字节
- Metadata 条目数: 最大 65535 个
- Payload: 最大 4GB (uint32 限制)
- 总消息大小: 最大 4GB + 开销

如需支持更大的消息，可将长度字段改为 uint64。