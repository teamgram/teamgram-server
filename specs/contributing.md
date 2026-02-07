# 贡献指南

本文档说明如何为 Teamgram Server 贡献代码：分支策略、PR 规范、代码风格、提交信息及本地开发与测试。根目录 [CONTRIBUTING.md](../CONTRIBUTING.md) 为简要入口。

## 分支策略

- **主干**：默认主干为 `master`（或仓库配置的默认分支）。主干的提交应通过 Pull Request 合并，不直接 push。
- **贡献流程**：在 fork 上从最新主干拉出功能分支，在分支上开发并提交，再向上游仓库发起 PR。

## Pull Request 规范

- **标题**：简洁说明改动内容；若有关联 issue，可在标题或描述中注明（如 `Fix #123`）。
- **通过 CI**：若仓库已配置 CI，PR 需通过构建与测试后再合并。
- **Review**：需至少一名 maintainer 审核通过（若项目有明确维护者）。
- **范围**：单次 PR 尽量聚焦单一功能或修复，便于 review 与回滚。

## 代码风格

- **Go**：使用标准 `gofmt` 格式化；推荐使用 `golangci-lint` 进行静态检查（若 CI 中已启用，请保证本地通过）。
- **项目约定**：若某子目录或模块有额外约定（如生成的 pb 代码不手改），以项目内说明或 README 为准。

## 提交信息

- 推荐使用 **Conventional Commits** 或项目约定格式，例如：
  - `feat: 添加 xxx 接口`
  - `fix: 修复登录超时问题`
  - `docs: 更新安装文档`
- 提交信息应简明扼要，便于生成 CHANGELOG 与回溯。

## 本地开发与测试

### 构建

- 需要 Go 1.21 及以上。
- 在仓库根目录执行：`make`，产物在 `teamgramd/bin/`。

### 依赖环境

- 需已部署 MySQL、Redis、etcd、Kafka、MinIO（及可选 ffmpeg）。可使用仓库提供的环境栈：
  - `docker compose -f docker-compose-env.yaml up -d`
- 数据库：创建库 `teamgram`，按顺序执行 `teamgramd/sql/` 下所有 SQL（见主 README 的 Init data 部分）。
- MinIO：创建 bucket `documents`、`encryptedfiles`、`photos`、`videos`（使用 docker-compose-env 时可由 minio-mc 自动创建）。

### 运行服务

- `cd teamgramd/bin && ./runall2.sh`，按脚本顺序启动各服务。

### 运行测试

- 在仓库根目录执行：`go test ./...`。若后续 Makefile 增加 `make test`，以 Makefile 为准。
- 部分测试可能依赖本地 MySQL/Redis 等，可先运行不依赖外部资源的包（如 `go test ./pkg/...`）验证改动。

### 代码生成

- 若修改了 proto 或 dal 表结构，需按项目内脚本重新生成代码（如 `teamgramd/deploy/sql/README.md` 及各 service 下的 dalgen 脚本）。
