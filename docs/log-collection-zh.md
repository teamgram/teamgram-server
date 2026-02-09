# 日志收集

Teamgram 使用 [Filebeat](https://www.elastic.co/beats/filebeat) 采集日志，[Kafka](https://kafka.apache.org/) 作为缓冲，[go-stash](https://github.com/kevwan/go-stash) 从 Kafka 消费并写入 Elasticsearch（替代资源占用较大的 Logstash），最后通过 [Elasticsearch](https://www.elastic.co/elasticsearch/) 与 [Kibana](https://www.elastic.co/kibana/) 进行存储与检索。

## 架构

![日志采集架构](image/log-pipeline-architecture.png)

```
┌─────────────┐     ┌─────────┐     ┌─────────┐     ┌─────────────────┐     ┌────────┐
│  应用日志   │────▶│ Filebeat│────▶│  Kafka  │────▶│    go-stash     │────▶│   ES   │
│ (容器日志)  │     │         │     │(teamgram│     │ (过滤/转换)     │     │        │
└─────────────┘     └─────────┘     │ -log)   │     └─────────────────┘     └───┬────┘
                                    └─────────┘                                  │
                                                                                 ▼
                                                                           ┌─────────┐
                                                                           │ Kibana  │
                                                                           └─────────┘
```

- **Filebeat**：采集容器日志并发送到 Kafka。
- **Kafka**：主题 `teamgram-log` 缓冲日志，解耦采集与写入。
- **go-stash**：从 Kafka 消费，过滤（如丢弃 debug、清理字段）后写入 Elasticsearch。
- **Elasticsearch**：存储日志，索引模式为 `teamgram-{{yyyy-MM-dd}}`。
- **Kibana**：查询与可视化。

## 配置说明

### Filebeat

- **配置路径**：`teamgramd/deploy/filebeat/conf/filebeat.yml`
- **Docker**：在 `docker-compose-env.yaml` 中挂载；需根据本机 Docker 的容器日志路径（如 `/var/lib/docker/containers`）正确配置。

主要项：

- `filebeat.inputs`：类型为 container，路径指向容器日志，并做 JSON 解析。
- `output.kafka`：Kafka 地址、主题 `teamgram-log`、压缩等。

### go-stash

- **配置路径**：`teamgramd/deploy/go-stash/etc/config.yaml`
- **Docker**：目录 `teamgramd/deploy/go-stash/etc` 挂载到 go-stash 容器。

主要项：

- **Input.Kafka**：Brokers、主题 `teamgram-log`、消费组、并发数。
- **Filters**：如丢弃 debug、删除无用字段、将 `message` 映射到 `data`。
- **Output.ElasticSearch**：ES 地址、索引 `teamgram-{{yyyy-MM-dd}}`。

### Elasticsearch 与 Kibana

由 `docker-compose-env.yaml` 启动。默认暴露端口：Elasticsearch 9200，Kibana 5601。

## Kibana 使用

1. 打开 Kibana（如 `http://localhost:5601`）。
2. **Stack Management** → **Index Patterns**（新版本中可能为 **Index Management**）。
3. 创建索引模式：`teamgram-*`，时间字段选 `@timestamp`。
4. 在 **Discover** 中按时间、级别、服务、内容等检索日志。

## 常见问题

- **Kibana 无数据**：确认 Filebeat、go-stash 容器正常，Kafka 主题 `teamgram-log` 有数据，Elasticsearch 存在 `teamgram-*` 索引。查看 go-stash、Filebeat 日志是否有报错。
- **Docker 日志路径**：Filebeat 需能读到容器日志文件；`filebeat.yml` 中的路径需与 Docker 的 containers 目录一致（如 `/var/lib/docker/containers`）。若 Docker 根目录不同，需在 `docker-compose-env.yaml` 中调整挂载。
- **go-stash 镜像**：使用 `docker-compose-env.yaml` 中指定的镜像（如 `kevinwan/go-stash:1.1.1`）。若配置格式不兼容，请参考 [go-stash](https://github.com/kevwan/go-stash) 文档。

## 相关文档

- [链路追踪](link-tracking-zh.md)
- [服务监控](service-monitoring-zh.md)
