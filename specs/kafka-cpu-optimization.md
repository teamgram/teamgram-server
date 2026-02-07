# Kafka CPU 偏高原因与优化建议

## 现象简述

- 日常 CPU 约 60–80%，偶发尖峰接近 180%
- 某时段掉到接近 0% 后恢复（可能重启或无流量）
- 当前为单节点 Bitnami Kafka 3.5.1（KRaft，broker + controller 同机）

## 可能原因

### 1. 单节点同时跑 Broker + Controller（KRaft）

- 一个进程既处理客户端读写（9092），又处理元数据/选主（9093），CPU 叠加。
- 元数据变更、选主、副本管理都会增加 CPU，在消息量大或 topic 多时更明显。

### 2. JVM 未调优

- 未设置 `KAFKA_HEAP_OPTS` 时，堆可能偏大或偏小，导致：
  - Full GC 频繁 → CPU 尖峰（图中 180% 可能与 GC 有关）
  - 堆过小导致频繁 Young GC，同样抬升 CPU
- 未使用 G1GC 或未调参时，GC 开销会放大。

### 3. 业务消息量

- **Inbox-T**、**Sync-T** 被 msg/sync/inbox 等服务生产和消费，用户活跃时消息量大，broker 序列化/反序列化、网络、日志写入都会推高 CPU。
- 若消息体大或 QPS 高，单节点容易成为瓶颈。

### 4. 磁盘与日志段

- 日志落盘、段滚动、保留/压缩策略（若开启）会占用 IO 和 CPU。
- 数据目录在 `./data/kafka`，若磁盘慢或与其它服务共享，会间接拉高 CPU 使用时间。

### 5. 容器无资源限制

- 未设置 `deploy.resources.limits` 时，Kafka 可用满整机 CPU，监控中的百分比是相对整机；若主机核数少，单核打满会显示 100%+（如 180%≈2 核）。

---

## 优化建议

### 1. 为 Kafka 设置 JVM 堆与 GC（优先）

在 `docker-compose-env.yaml` 的 kafka 服务中增加环境变量，控制堆大小并启用 G1GC，减轻 GC 导致的 CPU 尖峰：

```yaml
environment:
  # 现有 KAFKA_CFG_* ...
  KAFKA_HEAP_OPTS: "-Xmx1G -Xms1G"
  KAFKA_OPTS: "-XX:+UseG1GC -XX:MaxGCPauseMillis=20 -XX:InitiatingHeapOccupancyPercent=35"
```

- 内存紧张可改为 `-Xmx512m -Xms512m`；若宿主机内存充足且消息量大，可适当增大（如 2G），但不要超过宿主机可用内存的 50%。
- 观察一段时间，看 180% 尖峰是否减少或消失。

### 2. 限制容器 CPU（可选）

避免单容器占满整机，便于观察是否仍偏高：

```yaml
kafka:
  deploy:
    resources:
      limits:
        cpus: '2.0'
      reservations:
        cpus: '0.5'
```

### 3. 确认业务与 Topic 负载

- 在 Grafana/Prometheus 中查看 Kafka 的 `kafka_server_BrokerTopicMetrics_MessagesInPerSec`、`BytesInPerSec` 等，确认是否在尖峰时段有流量突增。
- 若 Inbox-T / Sync-T 的 producer/consumer 有批量过大或发送频率过高，可在业务侧做限流或批处理，降低 broker 压力。

### 4. 检查 go-stash 与其它消费者

- 若 go-stash 或其它服务消费 lag 大，会触发大量 catch-up 读取，增加 broker CPU。
- 当前 `teamgramd/deploy/go-stash/config.yaml` 中 topic 为 `your-log-topic`，若未改为真实 topic，可确认是否实际在消费；若在消费大 topic，可考虑错峰或增加消费能力。

### 5. 磁盘与数据目录

- 保证 `./data/kafka` 所在磁盘 IO 充足，避免与 Elasticsearch、MySQL 等 IO 密集型服务强争同一块盘。
- 若有日志保留或压缩需求，可适当放宽保留时间、减少压缩频率，以减轻 CPU（需权衡存储）。

---

## 快速可做的改动

1. **在 compose 中为 Kafka 加上 JVM 参数**（见上文 `KAFKA_HEAP_OPTS`、`KAFKA_OPTS`），重启 Kafka 观察数小时。
2. **为 kafka 容器加上 CPU limits**，再看监控曲线是否仍经常顶满。
3. **结合 Prometheus/Grafana** 看 Kafka 的 GC、网络、消息速率指标，确认尖峰是否与 GC 或流量一致。

若按上述调整后 CPU 仍偏高，再考虑将 broker 与 controller 拆到不同节点（多节点 KRaft）或增加 broker 节点做分区与负载分散。
