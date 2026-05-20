CREATE TABLE IF NOT EXISTS `auth_update_payloads` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `payload_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `update_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `codec` int NOT NULL DEFAULT '0',
  `layer` int NOT NULL DEFAULT '0',
  `tl_bytes` blob NOT NULL,
  `payload_hash` binary(32) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `expire_at` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `payload_id` (`payload_id`),
  KEY `payload_hash` (`payload_hash`),
  KEY `idx_expire_at` (`expire_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `auth_seq_deliveries` (
  `user_id` bigint NOT NULL,
  `perm_auth_key_id` bigint NOT NULL,
  `seq` bigint NOT NULL DEFAULT '0',
  `date` bigint NOT NULL DEFAULT '0',
  `payload_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `replay_policy` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `source_perm_auth_key_id` bigint NOT NULL DEFAULT '0',
  `visibility_policy` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `operation_id` varchar(160) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `expire_at` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_id`,`perm_auth_key_id`,`seq`),
  UNIQUE KEY `uk_user_auth_operation` (`user_id`,`perm_auth_key_id`,`operation_id`),
  KEY `idx_user_auth_date` (`user_id`,`perm_auth_key_id`,`date`,`seq`),
  KEY `idx_payload_id` (`payload_id`),
  KEY `idx_expire_at` (`expire_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `auth_seq_state` (
  `user_id` bigint NOT NULL,
  `perm_auth_key_id` bigint NOT NULL,
  `seq` bigint NOT NULL DEFAULT '0',
  `date` bigint NOT NULL DEFAULT '0',
  `row_version` bigint NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`perm_auth_key_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
