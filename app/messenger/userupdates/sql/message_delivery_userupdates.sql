-- Message delivery / userupdates first-slice storage contract.
--
-- Source spec:
--   /Users/wubenqi/go/src/teamgram.io/tgdocs-dev/server/teamgram-server-v2/specs/2026-04-29-message-delivery-storage-repository-contract.zh.md
--
-- This script contains the production MySQL tables needed by the first
-- message delivery vertical slice. It intentionally includes both msg-owned
-- and userupdates-owned tables because tgctl model generation needs the real
-- table definitions before repository code is generated.
--
-- MySQL is the first production backend. MongoDB remains a future/fake adapter
-- boundary and is not represented by this script.
--
-- user_peer_read_states is intentionally not created here. First version keeps
-- read/unread state embedded in user_dialogs.

SET NAMES utf8mb4;
SET time_zone = '+00:00';

-- ---------------------------------------------------------------------------
-- msg-owned storage
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS message_peer_sequences (
  peer_type             INT NOT NULL,
  peer_id               BIGINT NOT NULL,
  next_peer_seq         BIGINT NOT NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (peer_type, peer_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS canonical_messages (
  canonical_message_id        BIGINT NOT NULL,
  peer_type                   INT NOT NULL,
  peer_id                     BIGINT NOT NULL,
  peer_seq                    BIGINT NOT NULL,
  from_user_id                BIGINT NOT NULL,
  message_kind                INT NOT NULL,
  message_text                TEXT NULL,
  entities_payload_schema_version INT NOT NULL DEFAULT 1,
  entities_payload            BLOB NULL,
  media_ref_schema_version    INT NOT NULL DEFAULT 1,
  media_ref_payload           BLOB NULL,
  service_action_schema_version INT NOT NULL DEFAULT 1,
  service_action_payload      BLOB NULL,
  message_status              INT NOT NULL,
  edit_version                INT NOT NULL DEFAULT 0,
  date                        DATETIME(6) NOT NULL,
  edit_date                   DATETIME(6) NULL,
  deleted_at                  DATETIME(6) NULL,
  storage_schema_version      INT NOT NULL DEFAULT 1,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (canonical_message_id),
  UNIQUE KEY uk_peer_seq (peer_type, peer_id, peer_seq),
  KEY idx_peer_date (peer_type, peer_id, date),
  KEY idx_from_date (from_user_id, date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS message_client_randoms (
  sender_user_id        BIGINT NOT NULL,
  peer_type             INT NOT NULL,
  peer_id               BIGINT NOT NULL,
  client_random_id      BIGINT NOT NULL,
  canonical_message_id  BIGINT NOT NULL,
  send_state_id         BIGINT NOT NULL,
  request_payload_hash  BINARY(32) NOT NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (sender_user_id, peer_type, peer_id, client_random_id),
  UNIQUE KEY uk_canonical_message (canonical_message_id),
  KEY idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS message_send_states (
  send_state_id         BIGINT NOT NULL,
  sender_user_id        BIGINT NOT NULL,
  peer_type             INT NOT NULL,
  peer_id               BIGINT NOT NULL,
  client_random_id      BIGINT NOT NULL,
  canonical_message_id  BIGINT NULL,
  peer_seq              BIGINT NULL,
  status                INT NOT NULL,
  request_payload_schema_version INT NOT NULL,
  request_payload_hash  BINARY(32) NOT NULL,
  sender_operation_id   VARCHAR(160) NULL,
  sender_pts            BIGINT NULL,
  sender_pts_count      INT NULL,
  sender_update_schema_version INT NULL,
  sender_update_payload BLOB NULL,
  sender_update_payload_hash BINARY(32) NULL,
  receiver_manifest_id  BIGINT NULL,
  last_error_category   INT NULL,
  last_error_code       VARCHAR(128) NULL,
  last_error_message    VARCHAR(512) NULL,
  retry_count           INT NOT NULL DEFAULT 0,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  completed_at          DATETIME(6) NULL,
  PRIMARY KEY (send_state_id),
  UNIQUE KEY uk_random (sender_user_id, peer_type, peer_id, client_random_id),
  UNIQUE KEY uk_sender_operation (sender_operation_id),
  KEY idx_status_updated (status, updated_at),
  KEY idx_status_completed (status, completed_at),
  KEY idx_canonical_message (canonical_message_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS message_fanout_manifests (
  manifest_id           BIGINT NOT NULL,
  canonical_message_id  BIGINT NOT NULL,
  peer_type             INT NOT NULL,
  peer_id               BIGINT NOT NULL,
  peer_seq              BIGINT NOT NULL,
  actor_user_id         BIGINT NOT NULL,
  affected_user_count   INT NOT NULL,
  status                INT NOT NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  completed_at          DATETIME(6) NULL,
  PRIMARY KEY (manifest_id),
  UNIQUE KEY uk_canonical_message (canonical_message_id),
  KEY idx_peer_seq (peer_type, peer_id, peer_seq),
  KEY idx_status_updated (status, updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS message_fanout_receivers (
  manifest_id           BIGINT NOT NULL,
  receiver_user_id      BIGINT NOT NULL,
  operation_id          VARCHAR(160) NOT NULL,
  operation_payload_schema_version INT NOT NULL,
  operation_payload_codec INT NOT NULL,
  operation_payload     BLOB NOT NULL,
  operation_payload_hash BINARY(32) NOT NULL,
  kafka_topic           VARCHAR(128) NULL,
  kafka_partition       INT NULL,
  kafka_offset          BIGINT NULL,
  status                INT NOT NULL,
  retry_count           INT NOT NULL DEFAULT 0,
  next_retry_at         DATETIME(6) NULL,
  last_attempt_at       DATETIME(6) NULL,
  last_error_code       VARCHAR(128) NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (manifest_id, receiver_user_id),
  UNIQUE KEY uk_operation (receiver_user_id, operation_id),
  KEY idx_retry (status, next_retry_at),
  KEY idx_status_updated (status, updated_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ---------------------------------------------------------------------------
-- userupdates-owned storage
-- ---------------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS userupdates_partition_fences (
  partition_id          INT NOT NULL,
  owner_epoch           BIGINT NOT NULL,
  owner_instance_id     VARCHAR(128) NOT NULL,
  lease_id              VARCHAR(128) NULL,
  lease_expires_at      DATETIME(6) NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (partition_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS user_pts_state (
  user_id               BIGINT NOT NULL,
  pts                   BIGINT NOT NULL,
  pts_updated_at        DATETIME(6) NOT NULL,
  partition_id          INT NOT NULL,
  owner_epoch           BIGINT NOT NULL,
  row_version           BIGINT NOT NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id),
  KEY idx_partition (partition_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS user_pts_events (
  user_id               BIGINT NOT NULL,
  pts                   BIGINT NOT NULL,
  pts_count             INT NOT NULL,
  operation_id          VARCHAR(160) NOT NULL,
  op_type               INT NOT NULL,
  event_type            INT NOT NULL,
  peer_type             INT NOT NULL,
  peer_id               BIGINT NOT NULL,
  canonical_message_id  BIGINT NULL,
  peer_seq              BIGINT NULL,
  actor_user_id         BIGINT NULL,
  event_schema_version  INT NOT NULL,
  event_codec           INT NOT NULL,
  event_payload         BLOB NOT NULL,
  event_payload_hash    BINARY(32) NOT NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, pts),
  UNIQUE KEY uk_user_operation (user_id, operation_id),
  KEY idx_peer (peer_type, peer_id, peer_seq)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS user_auth_seq_state (
  user_id               BIGINT NOT NULL,
  seq                   BIGINT NOT NULL,
  date                  INT NOT NULL,
  row_version           BIGINT NOT NULL DEFAULT 0,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS user_auth_seq_events (
  user_id               BIGINT NOT NULL,
  seq                   BIGINT NOT NULL,
  date                  INT NOT NULL,
  operation_id          VARCHAR(160) NOT NULL,
  source_perm_auth_key_id BIGINT NOT NULL DEFAULT 0,
  target_auth_policy    VARCHAR(64) NOT NULL,
  public_update_type    VARCHAR(128) NOT NULL,
  peer_type             INT NOT NULL DEFAULT 0,
  peer_id               BIGINT NOT NULL DEFAULT 0,
  event_schema_version  INT NOT NULL,
  event_codec           INT NOT NULL,
  event_payload         BLOB NOT NULL,
  event_payload_hash    BINARY(32) NOT NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, seq),
  UNIQUE KEY uk_user_operation (user_id, operation_id),
  KEY idx_user_date (user_id, date, seq)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS user_operation_results (
  user_id               BIGINT NOT NULL,
  operation_id          VARCHAR(160) NOT NULL,
  op_type               INT NOT NULL,
  status                INT NOT NULL,
  pts                   BIGINT NULL,
  pts_count             INT NULL,
  payload_hash          BINARY(32) NOT NULL,
  response_schema_version INT NOT NULL,
  response_codec        INT NOT NULL,
  response_payload      BLOB NOT NULL,
  response_payload_hash BINARY(32) NOT NULL,
  terminal_error_code   VARCHAR(128) NULL,
  completed_at          DATETIME(6) NOT NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, operation_id),
  KEY idx_status_completed (status, completed_at, user_id, operation_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS user_message_views (
  user_id               BIGINT NOT NULL,
  peer_type             INT NOT NULL,
  peer_id               BIGINT NOT NULL,
  peer_seq              BIGINT NOT NULL,
  canonical_message_id  BIGINT NOT NULL,
  from_user_id          BIGINT NOT NULL,
  outgoing              BOOLEAN NOT NULL,
  message_kind          INT NOT NULL,
  message_status        INT NOT NULL,
  edit_version          INT NOT NULL DEFAULT 0,
  date                  DATETIME(6) NOT NULL,
  edit_date             DATETIME(6) NULL,
  deleted_at            DATETIME(6) NULL,
  view_schema_version   INT NOT NULL,
  view_payload          BLOB NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, peer_type, peer_id, peer_seq),
  UNIQUE KEY uk_user_canonical (user_id, canonical_message_id),
  KEY idx_user_peer_date (user_id, peer_type, peer_id, date),
  KEY idx_user_date (user_id, date)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS user_dialogs (
  user_id               BIGINT NOT NULL,
  peer_type             INT NOT NULL,
  peer_id               BIGINT NOT NULL,
  top_peer_seq          BIGINT NULL,
  top_canonical_message_id BIGINT NULL,
  top_message_date      DATETIME(6) NULL,
  top_message_status    INT NOT NULL DEFAULT 1,
  read_inbox_max_peer_seq BIGINT NOT NULL DEFAULT 0,
  read_outbox_max_peer_seq BIGINT NOT NULL DEFAULT 0,
  unread_count          INT NOT NULL DEFAULT 0,
  unread_mentions_count INT NOT NULL DEFAULT 0,
  unread_reactions_count INT NOT NULL DEFAULT 0,
  unread_mark           BOOLEAN NOT NULL DEFAULT FALSE,
  pinned_peer_seq       BIGINT NOT NULL DEFAULT 0,
  pinned_canonical_message_id BIGINT NOT NULL DEFAULT 0,
  has_scheduled         BOOLEAN NOT NULL DEFAULT FALSE,
  available_min_peer_seq BIGINT NOT NULL DEFAULT 0,
  hidden                BOOLEAN NOT NULL DEFAULT FALSE,
  deleted_at            DATETIME(6) NULL DEFAULT '1970-01-01 00:00:00.000000',
  last_pts              BIGINT NOT NULL DEFAULT 0,
  last_pts_at           DATETIME(6) NULL DEFAULT '1970-01-01 00:00:00.000000',
  dialog_schema_version INT NOT NULL DEFAULT 1,
  dialog_payload        BLOB NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (user_id, peer_type, peer_id),
  KEY idx_user_hidden_top (user_id, hidden, top_message_date, top_peer_seq),
  KEY idx_peer_read_inbox (peer_type, peer_id, read_inbox_max_peer_seq)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS push_task_outbox (
  task_id               BIGINT NOT NULL,
  user_id               BIGINT NOT NULL,
  pts                   BIGINT NOT NULL,
  push_type             INT NOT NULL,
  peer_type             INT NOT NULL,
  peer_id               BIGINT NOT NULL,
  operation_id          VARCHAR(160) NOT NULL,
  push_partition_id     INT NOT NULL,
  task_schema_version   INT NOT NULL,
  task_codec            INT NOT NULL,
  task_payload          BLOB NOT NULL,
  status                INT NOT NULL,
  publish_attempts      INT NOT NULL DEFAULT 0,
  available_at          DATETIME(6) NOT NULL,
  next_retry_at         DATETIME(6) NULL,
  published_topic       VARCHAR(128) NULL,
  published_partition   INT NULL,
  published_offset      BIGINT NULL,
  last_error_code       VARCHAR(128) NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  published_at          DATETIME(6) NULL,
  PRIMARY KEY (task_id),
  UNIQUE KEY uk_user_pts_push (user_id, pts, push_type),
  KEY idx_status_retry (status, next_retry_at),
  KEY idx_status_available (status, available_at, task_id),
  KEY idx_user_created (user_id, created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS dialog_side_effect_outbox (
  side_effect_id                 BIGINT NOT NULL,
  kind                           VARCHAR(64) NOT NULL,
  user_id                        BIGINT NOT NULL,
  peer_type                      INT NOT NULL,
  peer_id                        BIGINT NOT NULL,
  source_perm_auth_key_id        BIGINT NOT NULL DEFAULT 0,
  source_operation_id            VARCHAR(160) NOT NULL,
  source_message_date            DATETIME(6) NOT NULL,
  source_peer_seq                BIGINT NOT NULL DEFAULT 0,
  source_canonical_message_id    BIGINT NOT NULL DEFAULT 0,
  clear_before_date              DATETIME(6) NOT NULL DEFAULT '1970-01-01 00:00:00.000000',
  payload_schema_version         INT NOT NULL DEFAULT 1,
  payload                        BLOB NOT NULL,
  payload_hash                   BINARY(32) NOT NULL,
  status                         INT NOT NULL,
  attempt_count                  INT NOT NULL DEFAULT 0,
  next_retry_at                  DATETIME(6) NOT NULL,
  lease_owner                    VARCHAR(128) NOT NULL DEFAULT '',
  lease_until                    DATETIME(6) NOT NULL DEFAULT '1970-01-01 00:00:00.000000',
  last_error_code                VARCHAR(128) NOT NULL DEFAULT '',
  created_at                     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at                     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (side_effect_id),
  UNIQUE KEY uk_kind_operation (kind, source_operation_id),
  KEY idx_status_retry (status, next_retry_at, side_effect_id),
  KEY idx_user_kind (user_id, kind, created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS delivery_failed_operations (
  failed_id             BIGINT NOT NULL,
  user_id               BIGINT NOT NULL,
  operation_id          VARCHAR(160) NOT NULL,
  op_type               INT NOT NULL,
  bucket_id             INT NOT NULL,
  kafka_topic           VARCHAR(128) NOT NULL,
  kafka_partition       INT NOT NULL,
  kafka_offset          BIGINT NOT NULL,
  payload_schema_version INT NOT NULL,
  payload_hash          BINARY(32) NOT NULL,
  failure_category      INT NOT NULL,
  failure_code          VARCHAR(128) NOT NULL,
  failure_message       VARCHAR(1024) NULL,
  retry_count           INT NOT NULL DEFAULT 0,
  status                INT NOT NULL,
  failed_at             DATETIME(6) NOT NULL,
  replayed_at           DATETIME(6) NULL,
  replayed_by           VARCHAR(128) NULL,
  created_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at            TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (failed_id),
  UNIQUE KEY uk_operation_offset (kafka_topic, kafka_partition, kafka_offset),
  KEY idx_user_operation (user_id, operation_id),
  KEY idx_bucket_status_failed (bucket_id, status, failed_at, failed_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ---------------------------------------------------------------------------
-- Initial fence rows
-- ---------------------------------------------------------------------------
--
-- receiver_partition_count = 256. Each partition starts unowned at epoch 0.
-- Ownership acquisition must use CAS:
--
--   UPDATE userupdates_partition_fences
--   SET owner_epoch = owner_epoch + 1,
--       owner_instance_id = :new_instance,
--       lease_id = :lease_id,
--       lease_expires_at = :lease_expires_at
--   WHERE partition_id = :partition_id
--     AND owner_epoch = :prev_epoch;

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
      NULL
    );

    SET i = i + 1;
  END WHILE;
END//

DELIMITER ;

CALL init_userupdates_partition_fences();
DROP PROCEDURE init_userupdates_partition_fences;
