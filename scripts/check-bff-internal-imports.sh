#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/.."

for forbidden in \
  'github.com/teamgram/teamgram-server/v2/app/bff/internal/userprojection' \
  'github.com/teamgram/teamgram-server/v2/app/bff/internal/chatprojection'
do
  if rg -n --glob '*.go' "$forbidden" app/messenger app/service; then
    echo "forbidden import found: app/messenger and app/service must use owner-owned projection helpers" >&2
    exit 1
  fi
done
