CREATE DATABASE IF NOT EXISTS `teamgooo`
  DEFAULT CHARACTER SET utf8mb4
  COLLATE utf8mb4_general_ci;

USE `teamgooo`;

/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `auth_keys` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint NOT NULL COMMENT 'auth_id',
  `body` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'auth_key，原始数据为256的二进制数据，存储时转换成base64格式',
  `auth_key_type` int NOT NULL DEFAULT '-1',
  `perm_auth_key_id` bigint NOT NULL DEFAULT '0',
  `temp_auth_key_id` bigint NOT NULL DEFAULT '0',
  `media_temp_auth_key_id` bigint NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `auth_op_logs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint NOT NULL,
  `ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `op_type` int NOT NULL,
  `log_text` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `auth_users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint NOT NULL,
  `user_id` bigint NOT NULL DEFAULT '0',
  `hash` bigint NOT NULL DEFAULT '0',
  `date_created` bigint NOT NULL DEFAULT '0',
  `date_active` bigint NOT NULL DEFAULT '0',
  `state` int NOT NULL DEFAULT '0',
  `android_push_session_id` bigint NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`,`user_id`),
  KEY `auth_key_id_2` (`auth_key_id`,`user_id`,`deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `auths` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint NOT NULL,
  `layer` int NOT NULL DEFAULT '0',
  `api_id` int NOT NULL DEFAULT '0',
  `device_model` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `system_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `app_version` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `system_lang_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `lang_pack` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `lang_code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `system_code` varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `proxy` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `params` json NOT NULL,
  `client_ip` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `date_active` bigint NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `bot_commands` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `bot_id` bigint NOT NULL,
  `command` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `description` varchar(10240) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_bot_id_command` (`bot_id`,`command`),
  KEY `bot_id` (`bot_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `bots` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `bot_id` bigint NOT NULL,
  `bot_type` int DEFAULT NULL,
  `creator_user_id` bigint DEFAULT NULL,
  `token` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `description` varchar(10240) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `bot_chat_history` tinyint(1) DEFAULT NULL,
  `bot_nochats` tinyint(1) DEFAULT NULL,
  `verified` tinyint(1) DEFAULT NULL,
  `bot_inline_geo` tinyint(1) DEFAULT NULL,
  `bot_info_version` int DEFAULT NULL,
  `bot_inline_placeholder` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `bot_attach_menu` tinyint(1) DEFAULT NULL,
  `attach_menu_enabled` tinyint(1) DEFAULT NULL,
  `bot_business` tinyint(1) DEFAULT NULL,
  `bot_has_main_app` tinyint(1) DEFAULT NULL,
  `bot_active_users` int DEFAULT NULL,
  `has_menu_button` tinyint(1) DEFAULT NULL,
  `menu_button_text` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `menu_button_url` varchar(256) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `bot_can_edit` tinyint(1) DEFAULT NULL,
  `has_preview_medias` tinyint(1) DEFAULT NULL,
  `description_photo_id` bigint DEFAULT NULL,
  `description_document_id` bigint DEFAULT NULL,
  `main_app_url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `has_app_settings` tinyint(1) DEFAULT NULL,
  `placeholder_path` varchar(4096) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `background_color` int DEFAULT NULL,
  `background_dark_color` int DEFAULT NULL,
  `header_color` int DEFAULT NULL,
  `header_dark_color` int DEFAULT NULL,
  `privacy_policy_url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `mode` int DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `bot_id` (`bot_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `canonical_messages` (
  `canonical_message_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `peer_seq` bigint NOT NULL,
  `from_user_id` bigint NOT NULL,
  `message_kind` int NOT NULL,
  `message_text` text COLLATE utf8mb4_general_ci,
  `entities_payload_schema_version` int NOT NULL DEFAULT '1',
  `entities_payload` blob,
  `media_ref_schema_version` int NOT NULL DEFAULT '1',
  `media_ref_payload` blob,
  `service_action_schema_version` int NOT NULL DEFAULT '1',
  `service_action_payload` blob,
  `message_status` int NOT NULL,
  `edit_version` int NOT NULL DEFAULT '0',
  `date` datetime(6) NOT NULL,
  `edit_date` datetime(6) DEFAULT NULL,
  `deleted_at` datetime(6) DEFAULT NULL,
  `storage_schema_version` int NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`canonical_message_id`),
  UNIQUE KEY `uk_peer_seq` (`peer_type`,`peer_id`,`peer_seq`),
  KEY `idx_peer_date` (`peer_type`,`peer_id`,`date`),
  KEY `idx_from_date` (`from_user_id`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `chat_invite_participants` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `chat_id` bigint NOT NULL,
  `link` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` bigint NOT NULL,
  `requested` tinyint(1) NOT NULL DEFAULT '0',
  `approved_by` bigint NOT NULL DEFAULT '0',
  `date2` bigint NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `link` (`link`,`user_id`),
  KEY `link_2` (`link`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `chat_invites` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `chat_id` bigint NOT NULL,
  `admin_id` bigint NOT NULL,
  `migrated_to_id` bigint NOT NULL DEFAULT '0',
  `link` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `permanent` tinyint(1) NOT NULL DEFAULT '0',
  `revoked` tinyint(1) NOT NULL DEFAULT '0',
  `request_needed` tinyint(1) NOT NULL DEFAULT '0',
  `start_date` bigint NOT NULL DEFAULT '0',
  `expire_date` bigint NOT NULL DEFAULT '0',
  `usage_limit` int NOT NULL DEFAULT '0',
  `usage2` int NOT NULL DEFAULT '0',
  `requested` int NOT NULL DEFAULT '0',
  `title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `date2` bigint NOT NULL,
  `state` int NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `link` (`link`),
  KEY `chat_id` (`chat_id`,`permanent`,`admin_id`),
  KEY `chat_id_2` (`chat_id`,`admin_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `chat_participants` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `chat_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `participant_type` int DEFAULT '0',
  `link` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `usage2` int NOT NULL DEFAULT '0',
  `admin_rights` int NOT NULL DEFAULT '0',
  `inviter_user_id` bigint NOT NULL DEFAULT '0',
  `invited_at` bigint NOT NULL DEFAULT '0',
  `kicked_at` bigint NOT NULL DEFAULT '0',
  `left_at` bigint NOT NULL DEFAULT '0',
  `groupcall_default_join_as_peer_type` int NOT NULL DEFAULT '0',
  `groupcall_default_join_as_peer_id` bigint NOT NULL DEFAULT '0',
  `is_bot` tinyint(1) NOT NULL DEFAULT '0',
  `state` int NOT NULL DEFAULT '0',
  `date2` bigint NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `chat_id_2` (`chat_id`,`user_id`),
  KEY `chat_id` (`chat_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `chats` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `creator_user_id` bigint NOT NULL,
  `access_hash` bigint NOT NULL,
  `random_id` bigint NOT NULL,
  `participant_count` int NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `about` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `photo_id` bigint NOT NULL DEFAULT '0',
  `default_banned_rights` bigint NOT NULL DEFAULT '0',
  `migrated_to_id` bigint NOT NULL DEFAULT '0',
  `migrated_to_access_hash` bigint NOT NULL DEFAULT '0',
  `available_reactions_type` int NOT NULL DEFAULT '0',
  `available_reactions` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `deactivated` tinyint(1) NOT NULL DEFAULT '0',
  `noforwards` tinyint(1) NOT NULL DEFAULT '0',
  `ttl_period` int NOT NULL DEFAULT '0',
  `version` int NOT NULL DEFAULT '1',
  `date` bigint NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `default_history_ttl` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `period` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `delivery_failed_operations` (
  `failed_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `operation_id` varchar(160) COLLATE utf8mb4_general_ci NOT NULL,
  `op_type` int NOT NULL,
  `bucket_id` int NOT NULL,
  `kafka_topic` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `kafka_partition` int NOT NULL,
  `kafka_offset` bigint NOT NULL,
  `payload_schema_version` int NOT NULL,
  `payload_hash` binary(32) NOT NULL,
  `failure_category` int NOT NULL,
  `failure_code` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `failure_message` varchar(1024) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `retry_count` int NOT NULL DEFAULT '0',
  `status` int NOT NULL,
  `failed_at` datetime(6) NOT NULL,
  `replayed_at` datetime(6) DEFAULT NULL,
  `replayed_by` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`failed_id`),
  UNIQUE KEY `uk_operation_offset` (`kafka_topic`,`kafka_partition`,`kafka_offset`),
  KEY `idx_user_operation` (`user_id`,`operation_id`),
  KEY `idx_bucket_status_failed` (`bucket_id`,`status`,`failed_at`,`failed_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `dialog_auth_seq_outbox` (
  `outbox_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `source_perm_auth_key_id` bigint NOT NULL DEFAULT '0',
  `target_auth_policy` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `operation_id` varchar(160) COLLATE utf8mb4_general_ci NOT NULL,
  `event_type` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `peer_type` int NOT NULL DEFAULT '0',
  `peer_id` bigint NOT NULL DEFAULT '0',
  `payload_schema_version` int NOT NULL DEFAULT '1',
  `payload` blob NOT NULL,
  `payload_hash` binary(32) NOT NULL,
  `status` int NOT NULL,
  `attempt_count` int NOT NULL DEFAULT '0',
  `next_retry_at` datetime(6) NOT NULL,
  `lease_owner` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `lease_until` datetime(6) NOT NULL DEFAULT '1970-01-01 00:00:00.000000',
  `published_seq` bigint NOT NULL DEFAULT '0',
  `published_date` int NOT NULL DEFAULT '0',
  `last_error_kind` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `last_error_message` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`outbox_id`),
  UNIQUE KEY `uk_user_operation` (`user_id`,`operation_id`),
  KEY `idx_status_retry` (`status`,`next_retry_at`),
  KEY `idx_user_event_peer` (`user_id`,`event_type`,`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `dialog_drafts` (
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `peer_dialog_id` bigint NOT NULL,
  `draft_kind` int NOT NULL DEFAULT '0',
  `message` text COLLATE utf8mb4_general_ci NOT NULL,
  `entities_payload` blob NOT NULL,
  `reply_to_peer_seq` bigint NOT NULL DEFAULT '0',
  `draft_payload_schema_version` int NOT NULL DEFAULT '1',
  `draft_payload` blob NOT NULL,
  `date` datetime(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`peer_dialog_id`),
  UNIQUE KEY `uk_user_peer` (`user_id`,`peer_type`,`peer_id`),
  KEY `idx_user_date` (`user_id`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `dialog_filters` (
  `user_id` bigint NOT NULL,
  `dialog_filter_id` int NOT NULL,
  `slug` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `title` varchar(256) COLLATE utf8mb4_general_ci NOT NULL,
  `order_value` bigint NOT NULL DEFAULT '0',
  `enabled` tinyint(1) NOT NULL DEFAULT '1',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `filter_schema_version` int NOT NULL DEFAULT '1',
  `filter_payload` blob NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`dialog_filter_id`),
  UNIQUE KEY `uk_user_slug` (`user_id`,`slug`),
  KEY `idx_user_deleted_order` (`user_id`,`deleted`,`order_value`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `dialog_peer_policy` (
  `scope_type` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `scope_id` varchar(160) COLLATE utf8mb4_general_ci NOT NULL,
  `peer_type` int NOT NULL DEFAULT '0',
  `peer_id` bigint NOT NULL DEFAULT '0',
  `ttl_period` int NOT NULL DEFAULT '0',
  `theme_emoticon` varchar(64) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `policy_version` bigint NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`scope_type`,`scope_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `dialog_preference_versions` (
  `user_id` bigint NOT NULL,
  `scope_kind` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `folder_id` int NOT NULL DEFAULT '0',
  `aggregate_version` bigint NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`scope_kind`,`folder_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `dialog_preferences` (
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `peer_dialog_id` bigint NOT NULL,
  `folder_id` int NOT NULL DEFAULT '0',
  `main_pinned_order` bigint NOT NULL DEFAULT '0',
  `folder_pinned_order` bigint NOT NULL DEFAULT '0',
  `preferences_version` bigint NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`peer_type`,`peer_id`),
  UNIQUE KEY `uk_user_peer_dialog` (`user_id`,`peer_dialog_id`),
  KEY `idx_main_pin` (`user_id`,`folder_id`,`main_pinned_order`),
  KEY `idx_folder_pin` (`user_id`,`folder_id`,`folder_pinned_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `dialog_public_update_outbox` (
  `outbox_id` bigint NOT NULL,
  `source_user_id` bigint NOT NULL,
  `source_perm_auth_key_id` bigint NOT NULL DEFAULT '0',
  `target_user_id` bigint NOT NULL,
  `target_auth_policy` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `operation_id` varchar(160) COLLATE utf8mb4_general_ci NOT NULL,
  `delivery_path` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `public_update_type` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `peer_type` int NOT NULL DEFAULT '0',
  `peer_id` bigint NOT NULL DEFAULT '0',
  `payload_schema_version` int NOT NULL DEFAULT '1',
  `payload` blob NOT NULL,
  `payload_hash` binary(32) NOT NULL,
  `status` int NOT NULL,
  `attempt_count` int NOT NULL DEFAULT '0',
  `next_retry_at` datetime(6) NOT NULL,
  `lease_owner` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `lease_until` datetime(6) NOT NULL DEFAULT '1970-01-01 00:00:00.000000',
  `published_pts` bigint NOT NULL DEFAULT '0',
  `published_pts_count` int NOT NULL DEFAULT '0',
  `published_seq` bigint NOT NULL DEFAULT '0',
  `published_date` int NOT NULL DEFAULT '0',
  `last_error_kind` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `last_error_message` varchar(1024) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`outbox_id`),
  UNIQUE KEY `uk_target_operation_path_type` (`target_user_id`,`operation_id`,`delivery_path`,`public_update_type`),
  KEY `idx_status_retry` (`status`,`next_retry_at`),
  KEY `idx_target_type_peer` (`target_user_id`,`public_update_type`,`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `dialog_side_effect_outbox` (
  `side_effect_id` bigint NOT NULL,
  `kind` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `source_perm_auth_key_id` bigint NOT NULL DEFAULT '0',
  `source_operation_id` varchar(160) COLLATE utf8mb4_general_ci NOT NULL,
  `source_message_date` datetime(6) NOT NULL,
  `source_peer_seq` bigint NOT NULL DEFAULT '0',
  `source_canonical_message_id` bigint NOT NULL DEFAULT '0',
  `clear_before_date` datetime(6) NOT NULL DEFAULT '1970-01-01 00:00:00.000000',
  `payload_schema_version` int NOT NULL DEFAULT '1',
  `payload` blob NOT NULL,
  `payload_hash` binary(32) NOT NULL,
  `status` int NOT NULL,
  `attempt_count` int NOT NULL DEFAULT '0',
  `next_retry_at` datetime(6) NOT NULL,
  `lease_owner` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `lease_until` datetime(6) NOT NULL DEFAULT '1970-01-01 00:00:00.000000',
  `last_error_code` varchar(128) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`side_effect_id`),
  UNIQUE KEY `uk_kind_operation` (`kind`,`source_operation_id`),
  KEY `idx_status_retry` (`status`,`next_retry_at`,`side_effect_id`),
  KEY `idx_user_kind` (`user_id`,`kind`,`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `dialog_visual_settings` (
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `wallpaper_id` bigint NOT NULL DEFAULT '0',
  `wallpaper_overridden` tinyint(1) NOT NULL DEFAULT '0',
  `visual_version` bigint NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `dialogs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL DEFAULT '2',
  `peer_id` bigint NOT NULL,
  `peer_dialog_id` bigint NOT NULL,
  `pinned` bigint NOT NULL DEFAULT '0',
  `top_message` int NOT NULL DEFAULT '0',
  `pinned_msg_id` int NOT NULL DEFAULT '0',
  `read_inbox_max_id` int NOT NULL DEFAULT '0',
  `read_outbox_max_id` int NOT NULL DEFAULT '0',
  `unread_count` int NOT NULL DEFAULT '0',
  `unread_mentions_count` int NOT NULL DEFAULT '0',
  `unread_reactions_count` int NOT NULL DEFAULT '0',
  `unread_mark` tinyint(1) NOT NULL DEFAULT '0',
  `draft_type` int NOT NULL DEFAULT '0',
  `draft_message_data` json NOT NULL,
  `folder_id` int NOT NULL DEFAULT '0',
  `folder_pinned` bigint NOT NULL DEFAULT '0',
  `has_scheduled` tinyint(1) NOT NULL DEFAULT '0',
  `ttl_period` int NOT NULL DEFAULT '0',
  `theme_emoticon` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `wallpaper_id` bigint NOT NULL DEFAULT '0',
  `wallpaper_overridden` tinyint(1) NOT NULL DEFAULT '0',
  `date2` bigint NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`),
  UNIQUE KEY `user_id_2` (`user_id`,`peer_dialog_id`),
  KEY `user_id_3` (`user_id`,`peer_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `documents` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `document_id` bigint NOT NULL,
  `access_hash` bigint NOT NULL,
  `dc_id` int NOT NULL,
  `file_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `file_size` bigint NOT NULL,
  `uploaded_file_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `ext` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `mime_type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `thumb_id` bigint NOT NULL DEFAULT '0',
  `video_thumb_id` bigint NOT NULL DEFAULT '0',
  `version` int NOT NULL DEFAULT '0',
  `attributes` json NOT NULL,
  `date2` bigint NOT NULL DEFAULT '0',
  `import_document_id` bigint NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `document_id` (`document_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `drafts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `peer_dialog_id` bigint NOT NULL,
  `draft_type` int NOT NULL,
  `draft_message_data` json NOT NULL,
  `date2` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`peer_dialog_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `hash_tags` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `hash_tag` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `hash_tag_message_id` int NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_4` (`user_id`,`hash_tag`,`hash_tag_message_id`),
  KEY `user_id` (`user_id`,`hash_tag`),
  KEY `user_id_2` (`user_id`,`peer_type`,`peer_id`,`hash_tag`),
  KEY `user_id_3` (`user_id`,`hash_tag_message_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `imported_contacts` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `imported_user_id` bigint NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`),
  UNIQUE KEY `user_id_2` (`user_id`,`imported_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `message_client_randoms` (
  `sender_user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `client_random_id` bigint NOT NULL,
  `canonical_message_id` bigint NOT NULL,
  `send_state_id` bigint NOT NULL,
  `request_payload_hash` binary(32) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`sender_user_id`,`peer_type`,`peer_id`,`client_random_id`),
  UNIQUE KEY `uk_canonical_message` (`canonical_message_id`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `message_fanout_manifests` (
  `manifest_id` bigint NOT NULL,
  `canonical_message_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `peer_seq` bigint NOT NULL,
  `actor_user_id` bigint NOT NULL,
  `affected_user_count` int NOT NULL,
  `status` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `completed_at` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`manifest_id`),
  UNIQUE KEY `uk_canonical_message` (`canonical_message_id`),
  KEY `idx_peer_seq` (`peer_type`,`peer_id`,`peer_seq`),
  KEY `idx_status_updated` (`status`,`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `message_fanout_receivers` (
  `manifest_id` bigint NOT NULL,
  `receiver_user_id` bigint NOT NULL,
  `operation_id` varchar(160) COLLATE utf8mb4_general_ci NOT NULL,
  `operation_payload_schema_version` int NOT NULL,
  `operation_payload_codec` int NOT NULL,
  `operation_payload` blob NOT NULL,
  `operation_payload_hash` binary(32) NOT NULL,
  `kafka_topic` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `kafka_partition` int DEFAULT NULL,
  `kafka_offset` bigint DEFAULT NULL,
  `status` int NOT NULL,
  `retry_count` int NOT NULL DEFAULT '0',
  `next_retry_at` datetime(6) DEFAULT NULL,
  `last_attempt_at` datetime(6) DEFAULT NULL,
  `last_error_code` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`manifest_id`,`receiver_user_id`),
  UNIQUE KEY `uk_operation` (`receiver_user_id`,`operation_id`),
  KEY `idx_retry` (`status`,`next_retry_at`),
  KEY `idx_status_updated` (`status`,`updated_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `message_peer_sequences` (
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `next_peer_seq` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `message_read_outbox` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `read_user_id` bigint NOT NULL,
  `read_outbox_max_id` int NOT NULL,
  `read_outbox_max_date` datetime(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_outbox_read` (`user_id`,`peer_type`,`peer_id`,`read_user_id`,`read_outbox_max_id`),
  KEY `idx_outbox_read_lookup` (`user_id`,`peer_type`,`peer_id`,`read_user_id`,`read_outbox_max_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `message_send_states` (
  `send_state_id` bigint NOT NULL,
  `sender_user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `client_random_id` bigint NOT NULL,
  `canonical_message_id` bigint DEFAULT NULL,
  `peer_seq` bigint DEFAULT NULL,
  `status` int NOT NULL,
  `request_payload_schema_version` int NOT NULL,
  `request_payload_hash` binary(32) NOT NULL,
  `sender_operation_id` varchar(160) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sender_pts` bigint DEFAULT NULL,
  `sender_pts_count` int DEFAULT NULL,
  `sender_update_schema_version` int DEFAULT NULL,
  `sender_update_payload` blob,
  `sender_update_payload_hash` binary(32) DEFAULT NULL,
  `receiver_manifest_id` bigint DEFAULT NULL,
  `last_error_category` int DEFAULT NULL,
  `last_error_code` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `last_error_message` varchar(512) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `retry_count` int NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `completed_at` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`send_state_id`),
  UNIQUE KEY `uk_random` (`sender_user_id`,`peer_type`,`peer_id`,`client_random_id`),
  UNIQUE KEY `uk_sender_operation` (`sender_operation_id`),
  KEY `idx_status_updated` (`status`,`updated_at`),
  KEY `idx_status_completed` (`status`,`completed_at`),
  KEY `idx_canonical_message` (`canonical_message_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `phone_books` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL DEFAULT '0',
  `auth_key_id` bigint NOT NULL,
  `client_id` bigint NOT NULL,
  `phone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `first_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `last_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`,`client_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `photo_sizes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `photo_size_id` bigint NOT NULL,
  `size_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `volume_id` bigint NOT NULL DEFAULT '0',
  `local_id` int NOT NULL DEFAULT '0',
  `secret` bigint NOT NULL DEFAULT '0',
  `width` int NOT NULL,
  `height` int NOT NULL,
  `file_size` int NOT NULL,
  `file_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `has_stripped` tinyint(1) NOT NULL DEFAULT '0',
  `stripped_bytes` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `cached_type` int NOT NULL DEFAULT '0',
  `cached_bytes` varchar(4096) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `photo_size_id` (`photo_size_id`,`size_type`),
  KEY `photo_id` (`photo_size_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `photos` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `photo_id` bigint NOT NULL,
  `access_hash` bigint NOT NULL,
  `has_stickers` tinyint(1) NOT NULL DEFAULT '0',
  `dc_id` int NOT NULL DEFAULT '2',
  `date2` bigint NOT NULL DEFAULT '0',
  `has_video` tinyint(1) NOT NULL DEFAULT '0',
  `size_id` bigint NOT NULL DEFAULT '0',
  `video_size_id` bigint NOT NULL DEFAULT '0',
  `input_file_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `ext` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `photo_id` (`photo_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `popular_contacts` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `phone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `importers` int NOT NULL DEFAULT '1',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `predefined_users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `phone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `first_name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `last_name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `verified` tinyint(1) DEFAULT NULL,
  `registered_user_id` bigint DEFAULT NULL,
  `deleted` tinyint(1) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `pts_updates_ngen` (
  `id` bigint NOT NULL,
  `min_seq` bigint NOT NULL DEFAULT '0',
  `max_seq` bigint NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `push_task_outbox` (
  `task_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `pts` bigint NOT NULL,
  `push_type` int NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `operation_id` varchar(160) COLLATE utf8mb4_general_ci NOT NULL,
  `push_partition_id` int NOT NULL,
  `task_schema_version` int NOT NULL,
  `task_codec` int NOT NULL,
  `task_payload` blob NOT NULL,
  `status` int NOT NULL,
  `publish_attempts` int NOT NULL DEFAULT '0',
  `available_at` datetime(6) NOT NULL,
  `next_retry_at` datetime(6) DEFAULT NULL,
  `published_topic` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `published_partition` int DEFAULT NULL,
  `published_offset` bigint DEFAULT NULL,
  `last_error_code` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `published_at` datetime(6) DEFAULT NULL,
  PRIMARY KEY (`task_id`),
  UNIQUE KEY `uk_user_pts_push` (`user_id`,`pts`,`push_type`),
  KEY `idx_status_retry` (`status`,`next_retry_at`),
  KEY `idx_user_created` (`user_id`,`created_at`),
  KEY `idx_status_available` (`status`,`available_at`,`task_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `saved_dialogs` (
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `top_peer_seq` bigint NOT NULL DEFAULT '0',
  `top_canonical_message_id` bigint NOT NULL DEFAULT '0',
  `top_message_date` datetime(6) NOT NULL,
  `pinned` tinyint(1) NOT NULL DEFAULT '0',
  `pin_order` bigint NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `saved_schema_version` int NOT NULL DEFAULT '1',
  `saved_payload` blob NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`peer_type`,`peer_id`),
  KEY `idx_user_deleted_top` (`user_id`,`deleted`,`top_message_date`),
  KEY `idx_user_pinned_order` (`user_id`,`pinned`,`pin_order`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `unregistered_contacts` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `phone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `importer_user_id` bigint NOT NULL,
  `import_first_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `import_last_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `imported` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`,`importer_user_id`),
  KEY `phone_2` (`phone`,`importer_user_id`,`imported`),
  KEY `phone_3` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_auth_seq_events` (
  `user_id` bigint NOT NULL,
  `seq` bigint NOT NULL,
  `date` int NOT NULL,
  `operation_id` varchar(160) COLLATE utf8mb4_general_ci NOT NULL,
  `source_perm_auth_key_id` bigint NOT NULL DEFAULT '0',
  `target_auth_policy` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
  `public_update_type` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `peer_type` int NOT NULL DEFAULT '0',
  `peer_id` bigint NOT NULL DEFAULT '0',
  `event_schema_version` int NOT NULL,
  `event_codec` int NOT NULL,
  `event_payload` blob NOT NULL,
  `event_payload_hash` binary(32) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`seq`),
  UNIQUE KEY `uk_user_operation` (`user_id`,`operation_id`),
  KEY `idx_user_date` (`user_id`,`date`,`seq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_auth_seq_state` (
  `user_id` bigint NOT NULL,
  `seq` bigint NOT NULL,
  `date` int NOT NULL,
  `row_version` bigint NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_contacts` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `owner_user_id` bigint NOT NULL,
  `contact_user_id` bigint NOT NULL,
  `contact_phone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `contact_first_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `contact_last_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `mutual` tinyint(1) NOT NULL DEFAULT '0',
  `close_friend` tinyint(1) NOT NULL DEFAULT '0',
  `stories_hidden` tinyint(1) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `date2` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `owner_user_id` (`owner_user_id`,`contact_user_id`),
  KEY `owner_user_id_2` (`owner_user_id`),
  KEY `contact_user_id` (`contact_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_dialogs` (
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `top_peer_seq` bigint DEFAULT NULL,
  `top_canonical_message_id` bigint DEFAULT NULL,
  `top_message_date` datetime(6) DEFAULT NULL,
  `top_message_status` int NOT NULL DEFAULT '1',
  `unread_count` int NOT NULL DEFAULT '0',
  `unread_mentions_count` int NOT NULL DEFAULT '0',
  `unread_reactions_count` int NOT NULL DEFAULT '0',
  `unread_mark` tinyint(1) NOT NULL DEFAULT '0',
  `pinned_peer_seq` bigint NOT NULL DEFAULT '0',
  `pinned_canonical_message_id` bigint NOT NULL DEFAULT '0',
  `has_scheduled` tinyint(1) NOT NULL DEFAULT '0',
  `available_min_peer_seq` bigint NOT NULL DEFAULT '0',
  `hidden` tinyint(1) NOT NULL DEFAULT '0',
  `deleted_at` datetime(6) DEFAULT '1970-01-01 00:00:00.000000',
  `last_pts` bigint NOT NULL DEFAULT '0',
  `last_pts_at` datetime(6) DEFAULT '1970-01-01 00:00:00.000000',
  `read_inbox_max_peer_seq` bigint NOT NULL DEFAULT '0',
  `read_outbox_max_peer_seq` bigint NOT NULL DEFAULT '0',
  `dialog_schema_version` int NOT NULL DEFAULT '1',
  `dialog_payload` blob,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`peer_type`,`peer_id`),
  KEY `idx_user_hidden_top` (`user_id`,`hidden`,`top_message_date`,`top_peer_seq`),
  KEY `idx_peer_read_inbox` (`peer_type`,`peer_id`,`read_inbox_max_peer_seq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_global_privacy_settings` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `archive_and_mute_new_noncontact_peers` tinyint(1) NOT NULL DEFAULT '0',
  `keep_archived_unmuted` tinyint(1) NOT NULL DEFAULT '0',
  `keep_archived_folders` tinyint(1) NOT NULL DEFAULT '0',
  `hide_read_marks` tinyint(1) NOT NULL DEFAULT '0',
  `new_noncontact_peers_require_premium` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_message_views` (
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `peer_seq` bigint NOT NULL,
  `canonical_message_id` bigint NOT NULL,
  `from_user_id` bigint NOT NULL,
  `outgoing` tinyint(1) NOT NULL,
  `message_kind` int NOT NULL,
  `message_status` int NOT NULL,
  `edit_version` int NOT NULL DEFAULT '0',
  `date` datetime(6) NOT NULL,
  `edit_date` datetime(6) DEFAULT NULL,
  `deleted_at` datetime(6) DEFAULT NULL,
  `view_schema_version` int NOT NULL,
  `view_payload` blob,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`peer_type`,`peer_id`,`peer_seq`),
  UNIQUE KEY `uk_user_canonical` (`user_id`,`canonical_message_id`),
  KEY `idx_user_peer_date` (`user_id`,`peer_type`,`peer_id`,`date`),
  KEY `idx_user_date` (`user_id`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_notify_settings` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `show_previews` int NOT NULL DEFAULT '-1',
  `silent` int NOT NULL DEFAULT '-1',
  `mute_until` int NOT NULL DEFAULT '-1',
  `sound` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'default',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_operation_results` (
  `user_id` bigint NOT NULL,
  `operation_id` varchar(160) COLLATE utf8mb4_general_ci NOT NULL,
  `op_type` int NOT NULL,
  `status` int NOT NULL,
  `pts` bigint DEFAULT NULL,
  `pts_count` int DEFAULT NULL,
  `payload_hash` binary(32) NOT NULL,
  `response_schema_version` int NOT NULL,
  `response_codec` int NOT NULL,
  `response_payload` blob NOT NULL,
  `response_payload_hash` binary(32) NOT NULL,
  `terminal_error_code` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `completed_at` datetime(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`operation_id`),
  KEY `idx_status_completed` (`status`,`completed_at`,`user_id`,`operation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_peer_blocks` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `date` bigint NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_2` (`user_id`,`peer_type`,`peer_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_peer_settings` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `hide` tinyint(1) NOT NULL DEFAULT '0',
  `report_spam` tinyint(1) NOT NULL DEFAULT '0',
  `add_contact` tinyint(1) NOT NULL DEFAULT '0',
  `block_contact` tinyint(1) NOT NULL DEFAULT '0',
  `share_contact` tinyint(1) NOT NULL DEFAULT '0',
  `need_contacts_exception` tinyint(1) NOT NULL DEFAULT '0',
  `report_geo` tinyint(1) NOT NULL DEFAULT '0',
  `autoarchived` tinyint(1) NOT NULL DEFAULT '0',
  `invite_members` tinyint(1) NOT NULL DEFAULT '0',
  `geo_distance` int NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_presences` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `last_seen_at` bigint NOT NULL,
  `expires` int NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_privacies` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `key_type` int NOT NULL DEFAULT '0',
  `rules` json NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`key_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_profile_photos` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `photo_id` bigint NOT NULL,
  `date2` bigint NOT NULL COMMENT '排序',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`photo_id`),
  KEY `user_id_2` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_pts_events` (
  `user_id` bigint NOT NULL,
  `pts` bigint NOT NULL,
  `pts_count` int NOT NULL,
  `operation_id` varchar(160) COLLATE utf8mb4_general_ci NOT NULL,
  `op_type` int NOT NULL,
  `event_type` int NOT NULL,
  `peer_type` int NOT NULL,
  `peer_id` bigint NOT NULL,
  `canonical_message_id` bigint DEFAULT NULL,
  `peer_seq` bigint DEFAULT NULL,
  `actor_user_id` bigint DEFAULT NULL,
  `event_schema_version` int NOT NULL,
  `event_codec` int NOT NULL,
  `event_payload` blob NOT NULL,
  `event_payload_hash` binary(32) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`,`pts`),
  UNIQUE KEY `uk_user_operation` (`user_id`,`operation_id`),
  KEY `idx_peer` (`peer_type`,`peer_id`,`peer_seq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_pts_state` (
  `user_id` bigint NOT NULL,
  `pts` bigint NOT NULL,
  `pts_updated_at` datetime(6) NOT NULL,
  `partition_id` int NOT NULL,
  `owner_epoch` bigint NOT NULL,
  `row_version` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`),
  KEY `idx_partition` (`partition_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_pts_updates` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `pts` int NOT NULL,
  `pts_count` int NOT NULL,
  `update_type` tinyint NOT NULL DEFAULT '0',
  `update_data` json NOT NULL,
  `date2` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `user_id` (`user_id`,`pts`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_saved_music` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `saved_music_id` bigint NOT NULL,
  `order2` bigint NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`saved_music_id`),
  KEY `user_id_2` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `user_settings` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `key2` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `value` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`key2`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `username` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `peer_type` int NOT NULL DEFAULT '0',
  `peer_id` bigint NOT NULL DEFAULT '0',
  `editable` tinyint(1) NOT NULL DEFAULT '1',
  `active` tinyint(1) NOT NULL DEFAULT '1',
  `order2` bigint NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_type` int NOT NULL DEFAULT '2',
  `access_hash` bigint NOT NULL,
  `secret_key_id` bigint NOT NULL DEFAULT '0',
  `first_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `last_name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `username` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `phone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `country_code` varchar(3) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `verified` tinyint(1) NOT NULL DEFAULT '0',
  `support` tinyint(1) NOT NULL DEFAULT '0',
  `scam` tinyint(1) NOT NULL DEFAULT '0',
  `fake` tinyint(1) NOT NULL DEFAULT '0',
  `premium` tinyint(1) NOT NULL DEFAULT '0',
  `premium_expire_date` bigint NOT NULL DEFAULT '0',
  `about` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `state` int NOT NULL DEFAULT '0',
  `is_bot` tinyint(1) NOT NULL DEFAULT '0',
  `account_days_ttl` int NOT NULL DEFAULT '180',
  `photo_id` bigint NOT NULL DEFAULT '0',
  `restricted` tinyint(1) NOT NULL DEFAULT '0',
  `restriction_reason` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `archive_and_mute_new_noncontact_peers` tinyint(1) NOT NULL DEFAULT '0',
  `emoji_status_document_id` bigint NOT NULL DEFAULT '0',
  `emoji_status_until` int NOT NULL DEFAULT '0',
  `stories_max_id` int NOT NULL DEFAULT '0',
  `color` int NOT NULL DEFAULT '0',
  `color_background_emoji_id` bigint NOT NULL DEFAULT '0',
  `profile_color` int NOT NULL DEFAULT '0',
  `profile_color_background_emoji_id` bigint NOT NULL DEFAULT '0',
  `birthday` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `personal_channel_id` bigint NOT NULL DEFAULT '0',
  `authorization_ttl_days` int NOT NULL DEFAULT '180',
  `saved_music_id` bigint NOT NULL DEFAULT '0',
  `main_tab` int NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `delete_reason` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `userupdates_partition_fences` (
  `partition_id` int NOT NULL,
  `owner_epoch` bigint NOT NULL,
  `owner_instance_id` varchar(128) COLLATE utf8mb4_general_ci NOT NULL,
  `lease_id` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `lease_expires_at` datetime(6) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`partition_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

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

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE IF NOT EXISTS `video_sizes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `video_size_id` bigint NOT NULL,
  `size_type` char(1) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `width` int NOT NULL,
  `height` int NOT NULL,
  `file_size` int NOT NULL DEFAULT '0',
  `video_start_ts` double NOT NULL DEFAULT '0',
  `file_path` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `video_size_id` (`video_size_id`,`size_type`),
  KEY `video_size_id_2` (`video_size_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
/*!40101 SET character_set_client = @saved_cs_client */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
