#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

if ! command -v rg >/dev/null 2>&1; then
  echo "lint-tg-primitives: ripgrep (rg) is required" >&2
  exit 2
fi

file_list="$(mktemp)"
trap 'rm -f "$file_list"' EXIT

rg --files app | rg '/internal/core/.*_handler\.go$' >"$file_list" || true
if [[ ! -s "$file_list" ]]; then
  exit 0
fi

pattern='tg\.MakeTL(BoolTrue|BoolFalse|Int32|Int64|String)\('

if xargs rg -n --color=never "$pattern" <"$file_list"; then
  cat >&2 <<'EOF'

lint-tg-primitives: hand-written handlers must use TL primitive helpers/constants.

Use:
  tg.MakeInt32(...)
  tg.MakeInt64(...)
  tg.MakeString(...)
  tg.ToBool(...)
  tg.BoolTrue / tg.BoolFalse

Do not use verbose constructors such as:
  tg.MakeTLInt64(&tg.TLInt64{...}).ToInt64()
  tg.MakeTLString(&tg.TLString{...}).ToString()
  tg.MakeTLBoolTrue(&tg.TLBoolTrue{}).ToBool()
EOF
  exit 1
fi
