# Roadmap

This document outlines short- and medium-term goals and the boundary between community and enterprise editions.

## Short-term

- **Docs and specs**: Keep specs (architecture, protocol, dependencies, contributing, security, release, roadmap) in sync with the codebase.
- **CI and quality**: Add CI (e.g. GitHub Actions) for build, test, and lint (e.g. golangci-lint); add `test`, `lint`, `fmt` targets to Makefile where useful.
- **Doc consistency**: Align README and other docs (e.g. docker-compose-env.yaml naming, install guides).

## Medium-term

- **Test coverage**: Add unit and integration tests for critical paths; define coverage expectations over time.
- **Observability and ops**: Document runbooks and usage of the monitoring stack (Prometheus, Grafana, Jaeger, ELK).
- **Install and deploy**: Support more environments (e.g. other Linux distros, Kubernetes examples) and keep them consistent with docker-compose-env.

## Community vs enterprise

The following are available in the **enterprise edition** (contact the [author](https://t.me/benqi)); the community edition does not include or only partially supports them:

- sticker / theme / chat_theme / wallpaper / reactions / secret chat / 2FA / SMS / push (APNS / Web / FCM) / web / scheduled / autodelete / …
- channels / megagroups
- audio / video / group / conferenceCall
- bots
- miniapp

See the main README “Enterprise edition” section. Community roadmap focuses on docs, CI, tests, and deployment experience, not the above enterprise features.
