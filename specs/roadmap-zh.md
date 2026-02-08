# 路线图与目标

本文档描述短期与中期目标，以及社区版与企业版的功能边界。

## 短期目标

- **文档与规范**：保持 specs（架构、协议、依赖、贡献、安全、发布、路线图）与代码一致。
- **CI 与质量**：引入 CI（如 GitHub Actions）做构建、测试、lint（如 golangci-lint）；在 Makefile 中增加 `test`、`lint`、`fmt` 等目标。
- **文档一致性**：统一 README 与各文档（如 docker-compose-env 命名、安装文档）。

## 中期目标

- **测试覆盖**：增加核心路径的单元测试与集成测试，逐步建立覆盖率要求。
- **可观测与运维**：补充监控栈（Prometheus、Grafana、Jaeger、ELK）使用说明与 runbook。
- **安装与部署**：覆盖更多安装场景（如其他 Linux 发行版、Kubernetes 示例），与 docker-compose-env 文档保持一致。

## 社区版与企业版边界

以下能力在**企业版**中提供，社区版不包含或仅部分支持：

- sticker / theme / chat_theme / wallpaper / reactions / secret chat / 2FA / SMS / push（APNS / Web / FCM）/ web / scheduled / autodelete / …
- channels / megagroups
- audio / video / group / conferenceCall
- bots
- miniapp

详见主 README「企业版」小节。若有企业版需求，请通过 README 中的 [作者](https://t.me/benqi) 联系。社区版路线图聚焦于文档、CI、测试与部署体验，与上述企业功能无直接绑定。
