# 服务监控

Teamgram 使用 [Prometheus](https://prometheus.io/) 采集指标，[Grafana](https://grafana.com/) 做可视化。在各服务配置中启用 `Prometheus` 块后，go-zero 会自动暴露指标接口。

## 架构

- **各服务**：通过配置的地址暴露 `/metrics` 接口。
- **Prometheus**：按间隔拉取这些接口并存储时序数据。
- **Grafana**：将 Prometheus 添加为数据源，通过仪表盘展示指标。

## 服务端开启指标

在 `teamgramd/etc2/` 下各服务配置（如 `bff.yaml`）中保留或添加 `Prometheus` 块：

```yaml
Prometheus:
  Host: 0.0.0.0
  Port: 20011        # 各服务端口不同，避免冲突
  Path: /metrics
```

- **Host** / **Port**：指标 HTTP 服务监听地址（通常与本机及端口一致）。
- **Path**：指标路径，一般为 `/metrics`。

仓库中的 Prometheus 配置已按服务配置了抓取任务，需保证各服务的 `Host`/`Port`（或 Docker 内对外暴露的地址）与 Prometheus 的 target 一致。

## Prometheus 配置

- **配置路径**：`teamgramd/deploy/prometheus/server/prometheus.yml`
- **Docker**：在 `docker-compose-env.yaml` 中挂载（如 `./teamgramd/deploy/prometheus/server/prometheus.yml:/etc/prometheus/prometheus.yml`）。

`prometheus.yml` 中包括：

- **global**：如 `scrape_interval: 15s`。
- **scrape_configs**：按服务配置 job（如 `bff.bff`、`service.authsession`），`metrics_path: /metrics`，`static_configs` 中填写对应服务的 host:port（如 bff 为 `teamgram:20011`）。

新增服务时，在此增加对应的 scrape job，并保证该服务的 `Prometheus.Port` 与网络可达性正确。

## Grafana

### 1. 启动 Grafana

Grafana 已包含在 `docker-compose-env.yaml` 中，启动环境栈即可：

```bash
docker compose -f docker-compose-env.yaml up -d
```

访问地址一般为 `http://localhost:3000`（或所暴露端口）。默认账号常见为 `admin`/`admin`（以环境变量或文档为准，首次登录可能要求修改密码）。

### 2. 添加 Prometheus 数据源

- **Configuration** → **Data sources** → **Add data source** → 选择 **Prometheus**。
- **URL**：同一 Docker 网络下填 `http://prometheus:9090`，否则填 Prometheus 实际地址。
- 保存并测试。

### 3. 仪表盘

- 可手动创建 Panel（使用 PromQL），或从 Grafana.com 导入现成仪表盘（如 go-zero、Go 运行时相关）。
- 常用指标包括请求总数、延迟直方图等（如 `http_server_requests_total`、`http_server_request_duration_ms` 等，视实际暴露而定）。

## 相关文档

- [日志收集](log-collection-zh.md)
- [链路追踪](link-tracking-zh.md)
