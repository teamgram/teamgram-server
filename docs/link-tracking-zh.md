# 链路追踪（分布式追踪）

Teamgram 使用 go-zero 内置的分布式追踪能力，可对接 [Jaeger](https://www.jaegertracing.io/) 或 Zipkin，收集并分析请求在各服务间的调用链。

## 原理

各服务将 span 上报到追踪后端（Jaeger 或 Zipkin），后端按 trace ID 串联成完整调用链，便于查看一次请求的完整路径（如 gateway → bff → biz_service → 数据库）。

## 启用方式

### 1. 配置 Telemetry

在 `teamgramd/etc2/` 下各服务配置（如 `bff.yaml`、`gateway.yaml`）中增加或修改 `Telemetry` 块：

```yaml
Telemetry:
  Name: bff.bff                    # 在追踪界面中显示的服务名
  Endpoint: http://jaeger:14268/api/traces
  Sampler: 1.0                     # 采样率，1.0 表示 100% 采样
  Batcher: jaeger                  # 或 zipkin
```

- **Name**：在追踪界面中显示的服务名称。
- **Endpoint**：Jaeger 的 collector 地址（如 `http://jaeger:14268/api/traces`）或 Zipkin 的 endpoint。
- **Sampler**：采样率，`1.0` 表示全量采样；生产环境可按需调小（如 `0.1`）。
- **Batcher**：`jaeger` 或 `zipkin`。

### 2. 运行 Jaeger（或 Zipkin）

`docker-compose-env.yaml` 中已包含 Jaeger 服务，启动环境栈即可：

```bash
docker compose -f docker-compose-env.yaml up -d
```

Jaeger all-in-one 会接收上报（如 14268），并提供 Web UI（默认 16686）。在浏览器打开 `http://localhost:16686`，可按服务、操作、时间范围查询链路。

### 3. 可选：Jaeger 使用 Elasticsearch 存储

生产环境可将 span 存到 Elasticsearch。使用配置了 ES 的 Jaeger 部署替代 all-in-one 内存存储即可，各服务配置中的 `Telemetry.Endpoint` 指向该 Jaeger collector。

## Jaeger 界面

- **搜索**：选择服务、操作、时间范围后点击 “Find Traces”。
- **链路详情**：点击某条 trace 可查看时间线和 span 树（谁调用了谁、耗时多少）。

## 相关文档

- [日志收集](log-collection-zh.md)
- [服务监控](service-monitoring-zh.md)
