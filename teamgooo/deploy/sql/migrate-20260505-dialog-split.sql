CREATE TABLE IF NOT EXISTS `dialog_preferences` (
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `peer_dialog_id` bigint NOT NULL,
  `folder_id` int NOT NULL DEFAULT 0,
  `main_pinned_order` bigint NOT NULL DEFAULT 0,
  `folder_pinned_order` bigint NOT NULL DEFAULT 0,
  `preferences_version` bigint NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`peer_type`,`peer_id`),
  UNIQUE KEY `uk_user_peer_dialog` (`user_id`,`peer_dialog_id`),
  KEY `idx_main_pin` (`user_id`,`folder_id`,`main_pinned_order`),
  KEY `idx_folder_pin` (`user_id`,`folder_id`,`folder_pinned_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `dialog_preference_versions` (
  `user_id` bigint NOT NULL,
  `scope_kind` varchar(64) NOT NULL,
  `folder_id` int NOT NULL DEFAULT 0,
  `aggregate_version` bigint NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`scope_kind`,`folder_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `dialog_drafts` (
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `peer_dialog_id` bigint NOT NULL,
  `draft_kind` int NOT NULL DEFAULT 0,
  `message` text NOT NULL,
  `entities_payload` blob NOT NULL,
  `reply_to_peer_seq` bigint NOT NULL DEFAULT 0,
  `draft_payload_schema_version` int NOT NULL DEFAULT 1,
  `draft_payload` blob NOT NULL,
  `date` datetime(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`peer_dialog_id`),
  UNIQUE KEY `uk_user_peer` (`user_id`,`peer_type`,`peer_id`),
  KEY `idx_user_date` (`user_id`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

ALTER TABLE `dialog_filters`
  MODIFY `dialog_filter` json NULL,
  ADD COLUMN `title` varchar(256) NOT NULL DEFAULT '' AFTER `slug`,
  ADD COLUMN `enabled` tinyint(1) NOT NULL DEFAULT 1 AFTER `order_value`,
  ADD COLUMN `filter_schema_version` int NOT NULL DEFAULT 1 AFTER `deleted`,
  ADD COLUMN `filter_payload` blob NULL AFTER `filter_schema_version`,
  ADD KEY `idx_user_deleted_order` (`user_id`,`deleted`,`order_value`);

ALTER TABLE `saved_dialogs`
  MODIFY `top_message` int NULL DEFAULT 0,
  ADD COLUMN `top_peer_seq` bigint NOT NULL DEFAULT 0 AFTER `peer_id`,
  ADD COLUMN `top_canonical_message_id` bigint NOT NULL DEFAULT 0 AFTER `top_peer_seq`,
  ADD COLUMN `top_message_date` datetime(6) NOT NULL DEFAULT '1970-01-01 00:00:00.000000' AFTER `top_canonical_message_id`,
  ADD COLUMN `pin_order` bigint NOT NULL DEFAULT 0 AFTER `pinned`,
  ADD COLUMN `saved_schema_version` int NOT NULL DEFAULT 1 AFTER `deleted`,
  ADD COLUMN `saved_payload` blob NULL AFTER `saved_schema_version`,
  ADD KEY `idx_user_deleted_top` (`user_id`,`deleted`,`top_message_date`),
  ADD KEY `idx_user_pinned_order` (`user_id`,`pinned`,`pin_order`);

CREATE TABLE IF NOT EXISTS `dialog_peer_policy` (
  `scope_type` varchar(64) NOT NULL,
  `scope_id` varchar(160) NOT NULL,
  `peer_type` int NOT NULL DEFAULT 0,
  `peer_id` bigint NOT NULL DEFAULT 0,
  `ttl_period` int NOT NULL DEFAULT 0,
  `theme_emoticon` varchar(64) NOT NULL DEFAULT '',
  `policy_version` bigint NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`scope_type`,`scope_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `dialog_visual_settings` (
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `wallpaper_id` bigint NOT NULL DEFAULT 0,
  `wallpaper_overridden` tinyint(1) NOT NULL DEFAULT 0,
  `visual_version` bigint NOT NULL DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `dialog_auth_seq_outbox` (
  `outbox_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `source_perm_auth_key_id` bigint NOT NULL DEFAULT 0,
  `target_auth_policy` varchar(64) NOT NULL,
  `operation_id` varchar(160) NOT NULL,
  `event_type` varchar(128) NOT NULL,
  `peer_type` int NOT NULL DEFAULT 0,
  `peer_id` bigint NOT NULL DEFAULT 0,
  `payload_schema_version` int NOT NULL DEFAULT 1,
  `payload` blob NOT NULL,
  `payload_hash` binary(32) NOT NULL,
  `status` int NOT NULL,
  `attempt_count` int NOT NULL DEFAULT 0,
  `next_retry_at` datetime(6) NOT NULL,
  `lease_owner` varchar(128) NOT NULL DEFAULT '',
  `lease_until` datetime(6) NOT NULL DEFAULT '1970-01-01 00:00:00.000000',
  `published_seq` bigint NOT NULL DEFAULT 0,
  `published_date` int NOT NULL DEFAULT 0,
  `last_error_kind` varchar(128) NOT NULL DEFAULT '',
  `last_error_message` varchar(1024) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`outbox_id`),
  UNIQUE KEY `uk_user_operation` (`user_id`,`operation_id`),
  KEY `idx_status_retry` (`status`,`next_retry_at`),
  KEY `idx_user_event_peer` (`user_id`,`event_type`,`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `dialog_public_update_outbox` (
  `outbox_id` bigint NOT NULL,
  `source_user_id` bigint NOT NULL,
  `source_perm_auth_key_id` bigint NOT NULL DEFAULT 0,
  `target_user_id` bigint NOT NULL,
  `target_auth_policy` varchar(64) NOT NULL,
  `operation_id` varchar(160) NOT NULL,
  `delivery_path` varchar(64) NOT NULL,
  `public_update_type` varchar(128) NOT NULL,
  `peer_type` int NOT NULL DEFAULT 0,
  `peer_id` bigint NOT NULL DEFAULT 0,
  `payload_schema_version` int NOT NULL DEFAULT 1,
  `payload` blob NOT NULL,
  `payload_hash` binary(32) NOT NULL,
  `status` int NOT NULL,
  `attempt_count` int NOT NULL DEFAULT 0,
  `next_retry_at` datetime(6) NOT NULL,
  `lease_owner` varchar(128) NOT NULL DEFAULT '',
  `lease_until` datetime(6) NOT NULL DEFAULT '1970-01-01 00:00:00.000000',
  `published_pts` bigint NOT NULL DEFAULT 0,
  `published_pts_count` int NOT NULL DEFAULT 0,
  `published_seq` bigint NOT NULL DEFAULT 0,
  `published_date` int NOT NULL DEFAULT 0,
  `last_error_kind` varchar(128) NOT NULL DEFAULT '',
  `last_error_message` varchar(1024) NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`outbox_id`),
  UNIQUE KEY `uk_target_operation_path_type` (`target_user_id`,`operation_id`,`delivery_path`,`public_update_type`),
  KEY `idx_status_retry` (`status`,`next_retry_at`),
  KEY `idx_target_type_peer` (`target_user_id`,`public_update_type`,`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
