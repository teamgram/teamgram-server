/*
 Navicat Premium Data Transfer

 Source Server         : localhost-chatengine
 Source Server Type    : MySQL
 Source Server Version : 50717
 Source Host           : localhost:3307
 Source Schema         : chatengine

 Target Server Type    : MySQL
 Target Server Version : 50717
 File Encoding         : 65001

 Date: 04/03/2020 00:03:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_configs
-- ----------------------------
DROP TABLE IF EXISTS `app_configs`;
CREATE TABLE `app_configs` (
  `app_id` int(11) NOT NULL AUTO_INCREMENT,
  `config_key` int(11) NOT NULL,
  `config_value` int(11) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` int(11) NOT NULL,
  `updated_at` int(11) NOT NULL,
  PRIMARY KEY (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for app_ios_push_certs
-- ----------------------------
DROP TABLE IF EXISTS `app_ios_push_certs`;
CREATE TABLE `app_ios_push_certs` (
  `cert_id` int(11) NOT NULL AUTO_INCREMENT,
  `app_id` int(11) NOT NULL,
  `bundle_id` int(11) NOT NULL,
  `cert_type` int(11) NOT NULL,
  `cert_memo` int(11) NOT NULL,
  `uploaded` int(11) NOT NULL,
  `expired` int(11) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` int(11) NOT NULL,
  `updated_at` int(11) NOT NULL,
  PRIMARY KEY (`cert_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for app_keys
-- ----------------------------
DROP TABLE IF EXISTS `app_keys`;
CREATE TABLE `app_keys` (
  `app_id` int(11) NOT NULL,
  `app_key` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `app_secret` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` int(11) NOT NULL,
  `refresher` int(11) NOT NULL,
  `refreshed_at` int(11) NOT NULL,
  PRIMARY KEY (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for apps
-- ----------------------------
DROP TABLE IF EXISTS `apps`;
CREATE TABLE `apps` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `api_id` int(11) NOT NULL,
  `api_hash` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `title` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `short_name` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `api_id` (`api_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='apps';

-- ----------------------------
-- Table structure for auth_keys
-- ----------------------------
DROP TABLE IF EXISTS `auth_keys`;
CREATE TABLE `auth_keys` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL COMMENT 'auth_id',
  `body` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'auth_key，原始数据为256的二进制数据，存储时转换成base64格式',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for auth_op_logs
-- ----------------------------
DROP TABLE IF EXISTS `auth_op_logs`;
CREATE TABLE `auth_op_logs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(11) NOT NULL,
  `ip` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `op_type` int(11) NOT NULL DEFAULT '1',
  `log_text` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for auth_phone_transactions
-- ----------------------------
DROP TABLE IF EXISTS `auth_phone_transactions`;
CREATE TABLE `auth_phone_transactions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL,
  `phone_number` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL,
  `code_expired` int(11) NOT NULL DEFAULT '0',
  `code_msg_id` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `transaction_hash` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `sent_code_type` tinyint(4) NOT NULL DEFAULT '0',
  `flash_call_pattern` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `next_code_type` tinyint(4) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `api_id` int(11) NOT NULL,
  `api_hash` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `attempts` int(11) NOT NULL DEFAULT '0',
  `created_time` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `transaction_hash` (`transaction_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for auth_salts
-- ----------------------------
DROP TABLE IF EXISTS `auth_salts`;
CREATE TABLE `auth_salts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL,
  `salt` bigint(20) NOT NULL,
  `valid_since` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `auth` (`auth_key_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for auth_seq_updates
-- ----------------------------
DROP TABLE IF EXISTS `auth_seq_updates`;
CREATE TABLE `auth_seq_updates` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `auth_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `seq` int(11) NOT NULL,
  `update_type` int(11) NOT NULL,
  `update_data` blob NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_id` (`auth_id`,`user_id`,`seq`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for auth_updates_state
-- ----------------------------
DROP TABLE IF EXISTS `auth_updates_state`;
CREATE TABLE `auth_updates_state` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `pts` int(11) NOT NULL DEFAULT '0',
  `pts2` int(11) NOT NULL DEFAULT '0',
  `qts` int(11) NOT NULL DEFAULT '0',
  `qts2` int(11) NOT NULL DEFAULT '0',
  `seq` int(11) NOT NULL DEFAULT '-1',
  `seq2` int(11) NOT NULL DEFAULT '-1',
  `date` int(11) NOT NULL,
  `date2` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for auth_users
-- ----------------------------
DROP TABLE IF EXISTS `auth_users`;
CREATE TABLE `auth_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `hash` bigint(20) NOT NULL DEFAULT '0',
  `layer` int(11) NOT NULL DEFAULT '0',
  `device_model` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `platform` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `system_version` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `api_id` int(11) NOT NULL DEFAULT '0',
  `app_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `app_version` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `date_created` int(11) NOT NULL DEFAULT '0',
  `date_actived` int(11) NOT NULL DEFAULT '0',
  `ip` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `country` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `region` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`,`user_id`),
  KEY `auth_key_id_2` (`auth_key_id`,`user_id`,`deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for auths
-- ----------------------------
DROP TABLE IF EXISTS `auths`;
CREATE TABLE `auths` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL,
  `layer` int(11) NOT NULL DEFAULT '0',
  `api_id` int(11) NOT NULL,
  `device_model` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `system_version` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `app_version` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `system_lang_code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `lang_pack` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `lang_code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `client_ip` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for banned
-- ----------------------------
DROP TABLE IF EXISTS `banned`;
CREATE TABLE `banned` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `banned_time` bigint(20) NOT NULL,
  `expires` bigint(20) NOT NULL DEFAULT '0',
  `banned_reason` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `log` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for blocks
-- ----------------------------
DROP TABLE IF EXISTS `blocks`;
CREATE TABLE `blocks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `type` int(11) NOT NULL DEFAULT '0',
  `block_id` int(11) NOT NULL DEFAULT '0',
  `blocked_id` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for bot_commands
-- ----------------------------
DROP TABLE IF EXISTS `bot_commands`;
CREATE TABLE `bot_commands` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bot_id` int(11) NOT NULL,
  `command` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `description` varchar(10240) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `bot_id` (`bot_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for bots
-- ----------------------------
DROP TABLE IF EXISTS `bots`;
CREATE TABLE `bots` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `bot_id` int(11) NOT NULL,
  `bot_type` tinyint(4) NOT NULL DEFAULT '0',
  `description` varchar(10240) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `bot_chat_history` tinyint(4) NOT NULL DEFAULT '0',
  `bot_nochats` tinyint(4) NOT NULL DEFAULT '1',
  `verified` tinyint(4) NOT NULL DEFAULT '0',
  `bot_inline_geo` tinyint(4) NOT NULL DEFAULT '0',
  `bot_info_version` int(11) NOT NULL DEFAULT '1',
  `bot_inline_placeholder` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `bot_id` (`bot_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for channel_messages2
-- ----------------------------
DROP TABLE IF EXISTS `channel_messages2`;
CREATE TABLE `channel_messages2` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `channel_id` int(11) NOT NULL,
  `sender_user_id` int(11) NOT NULL,
  `channel_message_box_id` int(11) NOT NULL,
  `dialog_message_id` bigint(20) NOT NULL,
  `random_id` bigint(20) NOT NULL,
  `message_type` tinyint(4) NOT NULL DEFAULT '0',
  `message_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `date2` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for chat_participants
-- ----------------------------
DROP TABLE IF EXISTS `chat_participants`;
CREATE TABLE `chat_participants` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `chat_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `participant_type` tinyint(4) DEFAULT '0',
  `inviter_user_id` int(11) NOT NULL DEFAULT '0',
  `invited_at` int(11) NOT NULL DEFAULT '0',
  `kicked_at` int(11) NOT NULL DEFAULT '0',
  `left_at` int(11) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `chat_id_2` (`chat_id`,`user_id`),
  KEY `chat_id` (`chat_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for chats
-- ----------------------------
DROP TABLE IF EXISTS `chats`;
CREATE TABLE `chats` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `creator_user_id` int(11) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `random_id` bigint(20) NOT NULL,
  `participant_count` int(11) NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `link` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `photo_id` bigint(20) NOT NULL DEFAULT '0',
  `admins_enabled` tinyint(4) NOT NULL DEFAULT '0',
  `migrated_to` int(11) NOT NULL DEFAULT '0',
  `deactivated` tinyint(4) NOT NULL DEFAULT '0',
  `version` int(11) NOT NULL DEFAULT '1',
  `date` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for devices
-- ----------------------------
DROP TABLE IF EXISTS `devices`;
CREATE TABLE `devices` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `token_type` tinyint(4) NOT NULL,
  `token` varchar(190) COLLATE utf8mb4_unicode_ci NOT NULL,
  `app_sandbox` tinyint(4) NOT NULL DEFAULT '0',
  `secret` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `other_uids` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`,`user_id`,`token_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for documents
-- ----------------------------
DROP TABLE IF EXISTS `documents`;
CREATE TABLE `documents` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `document_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `dc_id` int(11) NOT NULL,
  `file_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_size` int(11) NOT NULL,
  `uploaded_file_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `ext` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `mime_type` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `thumb_id` bigint(20) NOT NULL DEFAULT '0',
  `version` int(11) NOT NULL DEFAULT '0',
  `attributes` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for file_parts
-- ----------------------------
DROP TABLE IF EXISTS `file_parts`;
CREATE TABLE `file_parts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `creator_id` bigint(20) NOT NULL DEFAULT '0',
  `creator_user_id` int(11) NOT NULL DEFAULT '0',
  `file_id` bigint(20) NOT NULL DEFAULT '0',
  `file_part_id` bigint(20) NOT NULL,
  `file_part` int(11) NOT NULL DEFAULT '0',
  `is_big_file` tinyint(4) NOT NULL DEFAULT '0',
  `file_total_parts` int(11) NOT NULL DEFAULT '0',
  `file_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `file_size` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `file_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `creator_id` bigint(20) NOT NULL DEFAULT '0',
  `creator_user_id` int(11) NOT NULL DEFAULT '0',
  `file_part_id` bigint(20) NOT NULL DEFAULT '0',
  `file_parts` int(11) NOT NULL DEFAULT '0',
  `file_size` bigint(20) NOT NULL,
  `file_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ext` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `is_big_file` tinyint(4) NOT NULL DEFAULT '0',
  `md5_checksum` char(33) COLLATE utf8mb4_unicode_ci NOT NULL,
  `upload_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for imported_contacts
-- ----------------------------
DROP TABLE IF EXISTS `imported_contacts`;
CREATE TABLE `imported_contacts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `imported_user_id` int(11) NOT NULL,
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`),
  UNIQUE KEY `user_id_2` (`user_id`,`imported_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for message_boxes
-- ----------------------------
DROP TABLE IF EXISTS `message_boxes`;
CREATE TABLE `message_boxes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `user_message_box_id` int(11) NOT NULL,
  `dialog_id` bigint(20) NOT NULL DEFAULT '0',
  `dialog_message_id` int(11) NOT NULL,
  `message_data_id` bigint(20) NOT NULL,
  `pts` int(11) NOT NULL DEFAULT '0',
  `message_box_type` tinyint(4) NOT NULL,
  `reply_to_msg_id` int(11) NOT NULL DEFAULT '0',
  `mentioned` tinyint(4) NOT NULL DEFAULT '0',
  `media_unread` tinyint(4) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`message_data_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for message_datas
-- ----------------------------
DROP TABLE IF EXISTS `message_datas`;
CREATE TABLE `message_datas` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message_data_id` bigint(20) NOT NULL,
  `dialog_id` bigint(20) NOT NULL,
  `dialog_message_id` int(11) NOT NULL DEFAULT '0',
  `sender_user_id` int(11) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `random_id` bigint(20) NOT NULL,
  `message_type` tinyint(4) NOT NULL DEFAULT '0',
  `message_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `media_unread` tinyint(4) NOT NULL DEFAULT '0',
  `has_media_unread` tinyint(4) NOT NULL DEFAULT '0',
  `date` int(11) NOT NULL DEFAULT '0',
  `edit_message` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `edit_date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `dialog_id` (`dialog_id`,`dialog_message_id`),
  UNIQUE KEY `sender_user_id` (`sender_user_id`,`random_id`),
  UNIQUE KEY `message_data_id` (`message_data_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `user_message_box_id` int(11) NOT NULL,
  `dialog_message_id` bigint(20) NOT NULL,
  `sender_user_id` int(11) NOT NULL,
  `message_box_type` tinyint(4) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `random_id` bigint(20) NOT NULL,
  `message_type` tinyint(4) NOT NULL DEFAULT '0',
  `message_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `date2` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for orgs
-- ----------------------------
DROP TABLE IF EXISTS `orgs`;
CREATE TABLE `orgs` (
  `org_id` int(11) NOT NULL AUTO_INCREMENT,
  `account_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `passwd` char(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `org_name` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mail` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mobile` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`org_id`),
  UNIQUE KEY `account_name` (`account_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for phone_books
-- ----------------------------
DROP TABLE IF EXISTS `phone_books`;
CREATE TABLE `phone_books` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL,
  `client_id` bigint(20) NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `first_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `auth_key_id` (`auth_key_id`,`client_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for phone_call_sessions
-- ----------------------------
DROP TABLE IF EXISTS `phone_call_sessions`;
CREATE TABLE `phone_call_sessions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `call_session_id` bigint(20) NOT NULL,
  `admin_id` int(11) NOT NULL,
  `admin_access_hash` bigint(20) NOT NULL,
  `participant_id` int(11) NOT NULL,
  `participant_access_hash` bigint(20) NOT NULL,
  `udp_p2p` tinyint(4) NOT NULL DEFAULT '0',
  `udp_reflector` tinyint(4) NOT NULL DEFAULT '0',
  `min_layer` int(11) NOT NULL,
  `max_layer` int(11) NOT NULL,
  `g_a` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `g_b` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` int(11) NOT NULL DEFAULT '0',
  `admin_debug_data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `participant_debug_data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `call_session_id` (`call_session_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for phones
-- ----------------------------
DROP TABLE IF EXISTS `phones`;
CREATE TABLE `phones` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `region` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'CN',
  `region_code` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '86',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for photo_datas
-- ----------------------------
DROP TABLE IF EXISTS `photo_datas`;
CREATE TABLE `photo_datas` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `photo_id` bigint(20) NOT NULL,
  `photo_type` tinyint(4) NOT NULL,
  `dc_id` int(11) NOT NULL,
  `volume_id` bigint(20) NOT NULL,
  `local_id` int(11) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `width` int(11) NOT NULL,
  `height` int(11) NOT NULL,
  `file_size` int(11) NOT NULL DEFAULT '0',
  `file_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `ext` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for photos
-- ----------------------------
DROP TABLE IF EXISTS `photos`;
CREATE TABLE `photos` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `photo_id` int(11) NOT NULL,
  `has_stickers` int(11) NOT NULL DEFAULT '0',
  `access_hash` int(11) NOT NULL,
  `date` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for popular_contacts
-- ----------------------------
DROP TABLE IF EXISTS `popular_contacts`;
CREATE TABLE `popular_contacts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `importers` int(11) NOT NULL DEFAULT '1',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for push_credentials
-- ----------------------------
DROP TABLE IF EXISTS `push_credentials`;
CREATE TABLE `push_credentials` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `token_type` tinyint(4) NOT NULL,
  `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `auth_id` (`auth_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for reports
-- ----------------------------
DROP TABLE IF EXISTS `reports`;
CREATE TABLE `reports` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `peer_type` int(11) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `reason` tinyint(4) NOT NULL DEFAULT '0',
  `content` varchar(10000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for sticker_data
-- ----------------------------
DROP TABLE IF EXISTS `sticker_data`;
CREATE TABLE `sticker_data` (
  `id` int(11) NOT NULL,
  `pack_id` int(11) NOT NULL,
  `emoji` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `image_128_file_id` bigint(20) NOT NULL,
  `image_128_file_hash` bigint(20) NOT NULL,
  `image_128_file_size` bigint(20) NOT NULL,
  `image_256_file_id` bigint(20) DEFAULT NULL,
  `image_256_file_hash` bigint(20) DEFAULT NULL,
  `image_256_file_size` bigint(20) DEFAULT NULL,
  `image_512_file_id` bigint(20) DEFAULT NULL,
  `image_512_file_hash` bigint(20) DEFAULT NULL,
  `image_512_file_size` bigint(20) DEFAULT NULL,
  `image_128_width` int(11) NOT NULL DEFAULT '128',
  `image_128_height` int(11) NOT NULL DEFAULT '128',
  `image_256_width` int(11) DEFAULT '256',
  `image_256_height` int(11) DEFAULT '256',
  `image_512_width` int(11) DEFAULT '512',
  `image_512_height` int(11) DEFAULT '512',
  PRIMARY KEY (`id`,`pack_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for sticker_packs
-- ----------------------------
DROP TABLE IF EXISTS `sticker_packs`;
CREATE TABLE `sticker_packs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sticker_set_id` bigint(20) NOT NULL,
  `emoticon` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `document_id` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for sticker_sets
-- ----------------------------
DROP TABLE IF EXISTS `sticker_sets`;
CREATE TABLE `sticker_sets` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sticker_set_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `title` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `short_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `count` int(11) NOT NULL DEFAULT '0',
  `hash` int(11) NOT NULL DEFAULT '0',
  `official` tinyint(4) NOT NULL DEFAULT '0',
  `mask` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `sticker_set_id` (`sticker_set_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for tmp_passwords
-- ----------------------------
DROP TABLE IF EXISTS `tmp_passwords`;
CREATE TABLE `tmp_passwords` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `auth_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `password_hash` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `period` int(11) NOT NULL,
  `tmp_password` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `valid_until` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for unread_mentions
-- ----------------------------
DROP TABLE IF EXISTS `unread_mentions`;
CREATE TABLE `unread_mentions` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `mentioned_message_id` int(11) NOT NULL,
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`,`mentioned_message_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for unregistered_contacts
-- ----------------------------
DROP TABLE IF EXISTS `unregistered_contacts`;
CREATE TABLE `unregistered_contacts` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `importer_user_id` int(11) NOT NULL,
  `import_first_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `import_last_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `imported` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `phone` (`phone`,`importer_user_id`),
  KEY `phone_2` (`phone`,`importer_user_id`,`imported`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_blocks
-- ----------------------------
DROP TABLE IF EXISTS `user_blocks`;
CREATE TABLE `user_blocks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `block_id` int(11) NOT NULL DEFAULT '0',
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`block_id`),
  KEY `user_id_2` (`user_id`,`block_id`,`deleted`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_contacts
-- ----------------------------
DROP TABLE IF EXISTS `user_contacts`;
CREATE TABLE `user_contacts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `owner_user_id` int(11) NOT NULL,
  `contact_user_id` int(11) NOT NULL,
  `contact_phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `contact_first_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `contact_last_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `mutual` tinyint(4) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `owner_user_id_2` (`owner_user_id`,`contact_phone`),
  KEY `owner_user_id` (`owner_user_id`,`contact_user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_dialogs
-- ----------------------------
DROP TABLE IF EXISTS `user_dialogs`;
CREATE TABLE `user_dialogs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `is_pinned` tinyint(4) NOT NULL DEFAULT '0',
  `top_message` int(11) NOT NULL DEFAULT '0',
  `read_inbox_max_id` int(11) NOT NULL DEFAULT '0',
  `read_outbox_max_id` int(11) NOT NULL DEFAULT '0',
  `unread_count` int(11) NOT NULL DEFAULT '0',
  `unread_mentions_count` int(11) NOT NULL DEFAULT '0',
  `show_previews` tinyint(4) NOT NULL DEFAULT '1',
  `silent` tinyint(4) NOT NULL DEFAULT '0',
  `mute_until` int(11) NOT NULL DEFAULT '0',
  `sound` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'default',
  `pts` int(11) NOT NULL DEFAULT '0',
  `draft_id` int(11) NOT NULL DEFAULT '0',
  `draft_type` tinyint(4) NOT NULL DEFAULT '0',
  `draft_message_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `date2` int(11) NOT NULL,
  `version` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_notify_settings
-- ----------------------------
DROP TABLE IF EXISTS `user_notify_settings`;
CREATE TABLE `user_notify_settings` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `show_previews` tinyint(1) NOT NULL DEFAULT '0',
  `silent` tinyint(1) NOT NULL DEFAULT '0',
  `mute_until` int(11) NOT NULL DEFAULT '0',
  `sound` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'default',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_passwords
-- ----------------------------
DROP TABLE IF EXISTS `user_passwords`;
CREATE TABLE `user_passwords` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `server_salt` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `hash` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `salt` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `hint` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `email` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `has_recovery` tinyint(4) NOT NULL DEFAULT '0',
  `code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `code_expired` int(11) NOT NULL DEFAULT '0',
  `attempts` int(11) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_presences
-- ----------------------------
DROP TABLE IF EXISTS `user_presences`;
CREATE TABLE `user_presences` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `last_seen_at` bigint(20) NOT NULL,
  `last_seen_auth_key_id` bigint(20) NOT NULL,
  `last_seen_ip` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`),
  KEY `user_id_2` (`user_id`,`last_seen_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_privacys
-- ----------------------------
DROP TABLE IF EXISTS `user_privacys`;
CREATE TABLE `user_privacys` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `key_type` tinyint(4) NOT NULL DEFAULT '0',
  `rules` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id` (`user_id`,`key_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_profile_photos
-- ----------------------------
DROP TABLE IF EXISTS `user_profile_photos`;
CREATE TABLE `user_profile_photos` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `photo_id` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_pts_updates
-- ----------------------------
DROP TABLE IF EXISTS `user_pts_updates`;
CREATE TABLE `user_pts_updates` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `pts` int(11) NOT NULL,
  `pts_count` int(11) NOT NULL,
  `update_type` tinyint(4) NOT NULL DEFAULT '0',
  `update_data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_qts_updates
-- ----------------------------
DROP TABLE IF EXISTS `user_qts_updates`;
CREATE TABLE `user_qts_updates` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `qts` int(11) NOT NULL,
  `update_type` int(11) NOT NULL,
  `update_data` blob NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for user_sticker_sets
-- ----------------------------
DROP TABLE IF EXISTS `user_sticker_sets`;
CREATE TABLE `user_sticker_sets` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `sticker_set_id` bigint(20) NOT NULL DEFAULT '0',
  `archived` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq` (`user_id`,`sticker_set_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for username
-- ----------------------------
DROP TABLE IF EXISTS `username`;
CREATE TABLE `username` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `username` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `peer_type` (`peer_type`,`peer_id`),
  KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_type` tinyint(4) NOT NULL DEFAULT '0',
  `access_hash` bigint(20) NOT NULL,
  `first_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `last_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `country_code` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL,
  `verified` tinyint(4) NOT NULL DEFAULT '0',
  `about` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` int(11) NOT NULL DEFAULT '0',
  `is_bot` tinyint(1) NOT NULL DEFAULT '0',
  `account_days_ttl` int(11) NOT NULL DEFAULT '180',
  `photos` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `min` tinyint(4) NOT NULL DEFAULT '0',
  `restricted` tinyint(4) NOT NULL DEFAULT '0',
  `restriction_reason` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `delete_reason` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for wall_papers
-- ----------------------------
DROP TABLE IF EXISTS `wall_papers`;
CREATE TABLE `wall_papers` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) NOT NULL DEFAULT '0',
  `title` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `color` int(11) NOT NULL DEFAULT '0',
  `bg_color` int(11) NOT NULL DEFAULT '0',
  `photo_id` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

SET FOREIGN_KEY_CHECKS = 1;
