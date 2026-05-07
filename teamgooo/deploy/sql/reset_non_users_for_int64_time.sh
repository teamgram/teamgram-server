#!/usr/bin/env bash
set -euo pipefail

usage() {
  cat >&2 <<'EOF'
Usage: reset_non_users_for_int64_time.sh [--dry-run]

Required environment:
  TEAMGOOO_RESET_ENV             local | rebuild | disposable-integration
  TEAMGOOO_RESET_DSN             Go MySQL DSN, for example root:@tcp(127.0.0.1:3306)/teamgooo

Required only for destructive mode:
  TEAMGOOO_RESET_CONFIRM_DB      must exactly match the parsed database name

Required outside local destructive mode:
  TEAMGOOO_RESET_BACKUP_ID       backup/snapshot identifier
    or
  TEAMGOOO_RESET_SKIP_BACKUP=1   explicit operator override
EOF
}

die() {
  echo "refusing: $*" >&2
  exit 1
}

dry_run=0
print_destructive_sql=0
if [[ $# -gt 1 ]]; then
  usage
  exit 2
fi
if [[ $# -eq 1 ]]; then
  case "$1" in
    --dry-run)
      dry_run=1
      ;;
    --print-destructive-sql)
      print_destructive_sql=1
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      usage
      exit 2
      ;;
  esac
fi

case "${TEAMGOOO_RESET_ENV:-}" in
  local|rebuild|disposable-integration)
    ;;
  "")
    die "TEAMGOOO_RESET_ENV must be one of: local, rebuild, disposable-integration"
    ;;
  *)
    die "TEAMGOOO_RESET_ENV must be one of: local, rebuild, disposable-integration"
    ;;
esac

dsn="${TEAMGOOO_RESET_DSN:-}"
if [[ -z "$dsn" ]]; then
  die "TEAMGOOO_RESET_DSN is required"
fi

dsn_regex='^([^:@/]+):([^@]*)@tcp\(([^:)]+):([0-9]+)\)/([^?]+)(\?.*)?$'
if [[ ! "$dsn" =~ $dsn_regex ]]; then
  die "unsupported TEAMGOOO_RESET_DSN shape; expected user:password@tcp(host:port)/database"
fi

mysql_user="${BASH_REMATCH[1]}"
mysql_pass="${BASH_REMATCH[2]}"
mysql_host="${BASH_REMATCH[3]}"
mysql_port="${BASH_REMATCH[4]}"
database="${BASH_REMATCH[5]}"

if [[ -z "$database" || ! "$database" =~ ^[A-Za-z0-9_$]+$ ]]; then
  die "could not parse a safe database name from TEAMGOOO_RESET_DSN"
fi

mysql_args=(
  --protocol=tcp
  -h "$mysql_host"
  -P "$mysql_port"
  -u "$mysql_user"
  --batch
  --skip-column-names
)
if [[ -n "$mysql_pass" ]]; then
  mysql_args+=(--password="$mysql_pass")
fi

list_tables_sql=$(
  cat <<SQL
SELECT TABLE_NAME
FROM information_schema.TABLES
WHERE TABLE_SCHEMA = '${database}'
  AND TABLE_TYPE = 'BASE TABLE'
  AND TABLE_NAME <> 'users'
ORDER BY TABLE_NAME;
SQL
)

seed_sql=$(
  cat <<'SQL'
DROP PROCEDURE IF EXISTS init_userupdates_partition_fences;

DELIMITER //

CREATE PROCEDURE init_userupdates_partition_fences()
BEGIN
  DECLARE i INT DEFAULT 0;

  WHILE i < 256 DO
    INSERT IGNORE INTO userupdates_partition_fences (
      partition_id,
      owner_epoch,
      owner_instance_id,
      lease_id,
      lease_expires_at
    ) VALUES (
      i,
      0,
      'unassigned',
      NULL,
      0
    );

    SET i = i + 1;
  END WHILE;
END//

DELIMITER ;

CALL init_userupdates_partition_fences();
DROP PROCEDURE init_userupdates_partition_fences;
SQL
)

if [[ "$print_destructive_sql" -eq 1 ]]; then
  cat <<SQL
SET FOREIGN_KEY_CHECKS = 0;
-- TRUNCATE statements are generated from information_schema base tables
-- where TABLE_SCHEMA = '${database}' and TABLE_NAME <> 'users'.
SET FOREIGN_KEY_CHECKS = 1;
${seed_sql}
SQL
  exit 0
fi

tables=()
while IFS= read -r table_name; do
  tables+=("$table_name")
done < <(mysql "${mysql_args[@]}" -e "$list_tables_sql")

if [[ "$dry_run" -eq 1 ]]; then
  echo "mode: dry-run"
  echo "database: ${database}"
  echo "preserve: users"
  echo "tables:"
  if [[ "${#tables[@]}" -eq 0 ]]; then
    echo "  (none)"
  else
    printf '  %s\n' "${tables[@]}"
  fi
  exit 0
fi

if [[ "${TEAMGOOO_RESET_CONFIRM_DB:-}" != "$database" ]]; then
  die "TEAMGOOO_RESET_CONFIRM_DB must exactly equal parsed database name: ${database}"
fi

if [[ "${TEAMGOOO_RESET_ENV}" != "local" ]]; then
  if [[ -z "${TEAMGOOO_RESET_BACKUP_ID:-}" && "${TEAMGOOO_RESET_SKIP_BACKUP:-}" != "1" ]]; then
    die "outside local, set TEAMGOOO_RESET_BACKUP_ID or TEAMGOOO_RESET_SKIP_BACKUP=1 before destructive reset"
  fi
fi

truncate_select_sql=$(
  cat <<SQL
SELECT CONCAT('TRUNCATE TABLE \`', REPLACE(TABLE_NAME, '\`', '\`\`'), '\`;')
FROM information_schema.TABLES
WHERE TABLE_SCHEMA = '${database}'
  AND TABLE_TYPE = 'BASE TABLE'
  AND TABLE_NAME <> 'users'
ORDER BY TABLE_NAME;
SQL
)

echo "mode: destructive"
echo "database: ${database}"
echo "preserve: users"
echo "backup: ${TEAMGOOO_RESET_BACKUP_ID:-${TEAMGOOO_RESET_SKIP_BACKUP:+skipped-by-operator}}"
echo "tables:"
if [[ "${#tables[@]}" -eq 0 ]]; then
  echo "  (none)"
else
  printf '  %s\n' "${tables[@]}"
fi

mysql "${mysql_args[@]}" "$database" < <(
  echo "SET FOREIGN_KEY_CHECKS = 0;"
  mysql "${mysql_args[@]}" -e "$truncate_select_sql"
  echo "SET FOREIGN_KEY_CHECKS = 1;"
)
mysql "${mysql_args[@]}" "$database" <<<"$seed_sql"
