# Contributing

This document covers branching, PR rules, code style, commit messages, and local dev/test. The root [CONTRIBUTING.md](../CONTRIBUTING.md) is a short entry point.

## Branching

- **Default branch**: `master`. Changes should land via Pull Request, not direct push.
- Work in a fork: create a feature branch from the latest default branch, then open a PR.

## Pull requests

- **Title**: Clear and concise; reference issues (e.g. `Fix #123`) when relevant.
- **CI**: PRs should pass build and tests when CI is enabled.
- **Review**: At least one maintainer review when applicable.
- **Scope**: One logical change per PR when possible.

## Code style

- **Go**: Use `gofmt`; prefer `golangci-lint` if the project uses it.
- Follow any project-specific rules (e.g. generated code, dalgen).

## Commit messages

- Prefer **Conventional Commits** (e.g. `feat: add X`, `fix: Y`, `docs: Z`).
- Keep messages concise for CHANGELOG and history.

## Local dev and test

### Build

- **Go**: 1.21+.
- From repo root: `make`; binaries go to `teamgramd/bin/`.

### Dependencies

- MySQL, Redis, etcd, Kafka, MinIO (and FFmpeg if needed). Optional: `docker compose -f docker-compose-env.yaml up -d`.
- **Database**: Create DB `teamgram`, run all scripts under `teamgramd/deploy/sql/` in order (see main README).
- **MinIO**: Buckets `documents`, `encryptedfiles`, `photos`, `videos` (auto-created when using docker-compose-env).

### Run services

- `cd teamgramd/bin && ./runall2.sh` (or `./runall-docker.sh` in Docker).

### Tests

- From repo root: `go test ./...`. Use `go test ./pkg/...` if some tests need no external deps.

### Code generation

- After changing proto or dal schema, re-run the project’s codegen (e.g. dalgen scripts).
