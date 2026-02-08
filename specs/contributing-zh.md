# 贡献指南

本文档说明分支策略、PR 规范、代码风格、提交信息及本地开发与测试。根目录 [CONTRIBUTING.md](../CONTRIBUTING.md) 为简要入口。

## 分支策略

- **主干**：默认分支为 `master`。主干的变更应通过 Pull Request 合并，不直接 push。
- 在 fork 上从最新主干拉出功能分支，开发后向上游发起 PR。

## Pull Request 规范

- **标题**：简洁说明改动；若有关联 issue 可注明（如 `Fix #123`）。
- **CI**：若已配置 CI，PR 需通过构建与测试。
- **Review**：需至少一名 maintainer 审核（若项目有明确维护者）。
- **范围**：单次 PR 尽量聚焦单一功能或修复。

## 代码风格

- **Go**：使用 `gofmt`；推荐 `golangci-lint`（若项目已启用）。
- 遵守项目内约定（如生成的 pb 代码、dalgen）。

## 提交信息

- 推荐 **Conventional Commits**（如 `feat: 添加 xxx`、`fix: 修复 xxx`、`docs: 更新 xxx`）。
- 保持简洁，便于 CHANGELOG 与回溯。

## 本地开发与测试

### 构建

- **Go**：1.21+。
- 仓库根目录执行 `make`；产物在 `teamgramd/bin/`。

### 依赖环境

- MySQL、Redis、etcd、Kafka、MinIO（及可选 FFmpeg）。可选：`docker compose -f docker-compose-env.yaml up -d`。
- **数据库**：创建库 `teamgram`，按顺序执行 `teamgramd/deploy/sql/` 下所有 SQL（见主 README）。
- **MinIO**：桶 `documents`、`encryptedfiles`、`photos`、`videos`（使用 docker-compose-env 时可由 minio-mc 自动创建）。

### 运行服务

- `cd teamgramd/bin && ./runall2.sh`（Docker 下可用 `./runall-docker.sh`）。

### 测试

- 仓库根目录：`go test ./...`。若部分测试依赖外部服务，可先跑 `go test ./pkg/...`。

### 代码生成

- 修改 proto 或 dal 表结构后，按项目内脚本重新生成代码（如 dalgen）。
