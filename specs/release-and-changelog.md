# 版本与发布

本文档约定版本号规则、CHANGELOG 格式及发布检查清单。

## 版本号规则

- 采用**语义化版本**思想，格式示例：`v0.96.0-teamgram-server`。
  - **主版本号**：不兼容的 API 或架构变更。
  - **次版本号**：向后兼容的功能新增。
  - **修订号**：向后兼容的问题修复与小改进。
- 后缀 `-teamgram-server` 用于与依赖库（如 proto）区分，发布时可按需保留或简化。

## CHANGELOG

- **位置**：仓库根目录 **CHANGELOG.md**（若已创建）。
- **格式**：推荐遵循 [Keep a Changelog](https://keepachangelog.com/)（按版本分段，分为 Added / Changed / Deprecated / Removed / Fixed / Security 等）。
- **责任**：每次发布前更新 CHANGELOG，将本次发布的改动归入对应版本与类型下。

## 发布检查清单

发布新版本时建议完成：

1. **版本号**：在 Makefile 或版本注入处更新 `VERSION`（若适用）。
2. **CHANGELOG**：在 CHANGELOG.md 中新增该版本条目并填写变更摘要。
3. **Tag**：在 Git 中打 tag（如 `v0.96.0`），并推送到远程。
4. **二进制/镜像**：若提供预编译二进制或 Docker 镜像，按项目惯例构建并发布。
5. **README**：若主 README 中有固定版本号或“当前版本”说明，按需更新。
