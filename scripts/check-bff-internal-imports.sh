#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

forbidden='github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection'

if rg -n --glob '*.go' "$forbidden" app/messenger app/service; then
  echo "forbidden import found: app/messenger and app/service must use app/service/biz/user/userprojection" >&2
  exit 1
fi
