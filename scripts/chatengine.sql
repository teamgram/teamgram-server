-- phpMyAdmin SQL Dump
-- version 4.8.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 2018-11-25 12:14:06
-- 服务器版本： 5.7.23
-- PHP Version: 7.1.16

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `chatengine`
--

-- --------------------------------------------------------

--
-- 表的结构 `apps`
--

CREATE TABLE `apps` (
  `id` int(11) NOT NULL,
  `api_id` int(11) NOT NULL,
  `api_hash` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `title` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `short_name` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='apps';

-- --------------------------------------------------------

--
-- 表的结构 `app_configs`
--

CREATE TABLE `app_configs` (
  `app_id` int(11) NOT NULL,
  `config_key` int(11) NOT NULL,
  `config_value` int(11) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` int(11) NOT NULL,
  `updated_at` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `app_ios_push_certs`
--

CREATE TABLE `app_ios_push_certs` (
  `cert_id` int(11) NOT NULL,
  `app_id` int(11) NOT NULL,
  `bundle_id` int(11) NOT NULL,
  `cert_type` int(11) NOT NULL,
  `cert_memo` int(11) NOT NULL,
  `uploaded` int(11) NOT NULL,
  `expired` int(11) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` int(11) NOT NULL,
  `updated_at` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `app_keys`
--

CREATE TABLE `app_keys` (
  `app_id` int(11) NOT NULL,
  `app_key` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `app_secret` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` int(11) NOT NULL,
  `refresher` int(11) NOT NULL,
  `refreshed_at` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `auths`
--

CREATE TABLE `auths` (
  `id` int(11) NOT NULL,
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
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `auth_channel_updates_state`
--

CREATE TABLE `auth_channel_updates_state` (
  `id` int(11) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `pts` int(11) NOT NULL DEFAULT '0',
  `pts2` int(11) NOT NULL DEFAULT '0',
  `date` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `auth_keys`
--

CREATE TABLE `auth_keys` (
  `id` int(11) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL COMMENT 'auth_id',
  `body` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'auth_key，原始数据为256的二进制数据，存储时转换成base64格式',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `auth_op_logs`
--

CREATE TABLE `auth_op_logs` (
  `id` bigint(20) NOT NULL,
  `auth_key_id` bigint(11) NOT NULL,
  `ip` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `op_type` int(11) NOT NULL DEFAULT '1',
  `log_text` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `auth_phone_transactions`
--

CREATE TABLE `auth_phone_transactions` (
  `id` bigint(20) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL,
  `phone_number` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL,
  `code_expired` int(11) NOT NULL DEFAULT '0',
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
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `auth_salts`
--

CREATE TABLE `auth_salts` (
  `id` int(11) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL,
  `salt` bigint(20) NOT NULL,
  `valid_since` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `auth_seq_updates`
--

CREATE TABLE `auth_seq_updates` (
  `id` bigint(20) NOT NULL,
  `auth_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `seq` int(11) NOT NULL,
  `update_type` int(11) NOT NULL,
  `update_data` blob NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `auth_updates_state`
--

CREATE TABLE `auth_updates_state` (
  `id` int(11) NOT NULL,
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
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `auth_users`
--

CREATE TABLE `auth_users` (
  `id` int(11) NOT NULL,
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
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `banned`
--

CREATE TABLE `banned` (
  `id` int(11) NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `banned_time` bigint(20) NOT NULL,
  `expires` bigint(20) NOT NULL,
  `banned_reason` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `log` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `state` tinyint(4) NOT NULL,
  `created_at` int(11) NOT NULL,
  `updated_at` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `blocks`
--

CREATE TABLE `blocks` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `type` int(11) NOT NULL DEFAULT '0',
  `block_id` int(11) NOT NULL DEFAULT '0',
  `blocked_id` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `bots`
--

CREATE TABLE `bots` (
  `id` int(11) NOT NULL,
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
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `bot_commands`
--

CREATE TABLE `bot_commands` (
  `id` int(11) NOT NULL,
  `bot_id` int(11) NOT NULL,
  `command` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `description` varchar(10240) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `channels`
--

CREATE TABLE `channels` (
  `id` int(11) NOT NULL,
  `creator_user_id` int(11) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `random_id` bigint(20) NOT NULL,
  `top_message` int(11) NOT NULL DEFAULT '0',
  `participant_count` int(11) NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `about` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `photo_id` bigint(20) NOT NULL DEFAULT '0',
  `public` tinyint(4) NOT NULL DEFAULT '0',
  `link` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `broadcast` tinyint(4) NOT NULL DEFAULT '0',
  `verified` tinyint(4) NOT NULL DEFAULT '0',
  `megagroup` tinyint(4) NOT NULL DEFAULT '0',
  `democracy` tinyint(4) NOT NULL DEFAULT '0',
  `signatures` tinyint(4) NOT NULL DEFAULT '0',
  `admins_enabled` tinyint(4) NOT NULL DEFAULT '0',
  `deactivated` tinyint(4) NOT NULL DEFAULT '0',
  `version` int(11) NOT NULL DEFAULT '1',
  `date` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `channel_admin_logs`
--

CREATE TABLE `channel_admin_logs` (
  `id` bigint(20) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `admin_user_id` int(11) NOT NULL,
  `event` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `channel_media_unread`
--

CREATE TABLE `channel_media_unread` (
  `id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `channel_message_id` int(11) NOT NULL,
  `media_unread` tinyint(4) NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `channel_messages`
--

CREATE TABLE `channel_messages` (
  `id` int(11) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `channel_message_id` int(11) NOT NULL,
  `sender_user_id` int(11) NOT NULL,
  `random_id` bigint(20) NOT NULL,
  `message_data_id` bigint(20) NOT NULL,
  `message_type` tinyint(4) NOT NULL DEFAULT '0',
  `message_data` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `has_media_unread` tinyint(4) NOT NULL DEFAULT '0',
  `edit_message` text COLLATE utf8mb4_unicode_ci NOT NULL,
  `edit_date` int(11) NOT NULL DEFAULT '0',
  `views` int(11) NOT NULL DEFAULT '1',
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `channel_messages2`
--

CREATE TABLE `channel_messages2` (
  `id` int(11) NOT NULL,
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
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `channel_message_boxes`
--

CREATE TABLE `channel_message_boxes` (
  `id` int(11) NOT NULL,
  `sender_user_id` int(11) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `dialog_id` bigint(20) NOT NULL,
  `dialog_message_id` int(11) NOT NULL,
  `channel_message_box_id` int(11) NOT NULL,
  `message_id` bigint(20) NOT NULL,
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `channel_participants`
--

CREATE TABLE `channel_participants` (
  `id` bigint(11) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `is_creator` int(11) NOT NULL DEFAULT '0',
  `participant_type` tinyint(4) DEFAULT '0',
  `inviter_user_id` int(11) NOT NULL DEFAULT '0',
  `invited_at` int(11) NOT NULL DEFAULT '0',
  `joined_at` int(11) NOT NULL DEFAULT '0',
  `promoted_by` int(11) NOT NULL DEFAULT '0',
  `admin_rights` int(11) NOT NULL DEFAULT '0',
  `promoted_at` int(11) NOT NULL DEFAULT '0',
  `is_left` tinyint(4) NOT NULL DEFAULT '0',
  `hidden_prehistory` tinyint(4) NOT NULL DEFAULT '0',
  `hidden_prehistory_message_id` int(11) NOT NULL DEFAULT '0',
  `left_at` int(11) NOT NULL DEFAULT '0',
  `is_kicked` tinyint(4) NOT NULL DEFAULT '0',
  `kicked_by` int(11) NOT NULL DEFAULT '0',
  `kicked_at` int(11) NOT NULL DEFAULT '0',
  `banned_rights` int(11) NOT NULL DEFAULT '0',
  `banned_until_date` int(11) NOT NULL DEFAULT '0',
  `banned_at` int(11) NOT NULL DEFAULT '0',
  `read_inbox_max_id` int(11) NOT NULL DEFAULT '0',
  `read_outbox_max_id` int(11) DEFAULT '0',
  `date` int(11) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `channel_pts_updates`
--

CREATE TABLE `channel_pts_updates` (
  `id` bigint(20) NOT NULL,
  `channel_id` int(11) NOT NULL,
  `pts` int(11) NOT NULL,
  `pts_count` int(11) NOT NULL,
  `update_type` tinyint(4) NOT NULL DEFAULT '0',
  `update_data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `chats`
--

CREATE TABLE `chats` (
  `id` int(11) NOT NULL,
  `creator_user_id` int(11) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `random_id` bigint(20) NOT NULL,
  `participant_count` int(11) NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `photo_id` bigint(20) NOT NULL DEFAULT '0',
  `admins_enabled` tinyint(4) NOT NULL DEFAULT '0',
  `deactivated` tinyint(4) NOT NULL DEFAULT '0',
  `version` int(11) NOT NULL DEFAULT '1',
  `date` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `chat_participants`
--

CREATE TABLE `chat_participants` (
  `id` int(11) NOT NULL,
  `chat_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `participant_type` tinyint(4) DEFAULT '0',
  `inviter_user_id` int(11) NOT NULL DEFAULT '0',
  `invited_at` int(11) NOT NULL DEFAULT '0',
  `kicked_at` int(11) NOT NULL DEFAULT '0',
  `left_at` int(11) NOT NULL DEFAULT '0',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `config`
--

CREATE TABLE `config` (
  `id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `devices`
--

CREATE TABLE `devices` (
  `id` bigint(20) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `token_type` tinyint(4) NOT NULL,
  `token` varchar(190) COLLATE utf8mb4_unicode_ci NOT NULL,
  `app_sandbox` tinyint(4) NOT NULL DEFAULT '0',
  `secret` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `other_uids` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `documents`
--

CREATE TABLE `documents` (
  `id` bigint(20) NOT NULL,
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
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `files`
--

CREATE TABLE `files` (
  `id` bigint(20) NOT NULL,
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
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `file_parts`
--

CREATE TABLE `file_parts` (
  `id` bigint(20) NOT NULL,
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
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `imported_contacts`
--

CREATE TABLE `imported_contacts` (
  `id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `imported_user_id` int(11) NOT NULL,
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `messages`
--

CREATE TABLE `messages` (
  `id` int(11) NOT NULL,
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
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `message_boxes`
--

CREATE TABLE `message_boxes` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `user_message_box_id` int(11) NOT NULL,
  `dialog_id` bigint(20) NOT NULL DEFAULT '0',
  `dialog_message_id` int(11) NOT NULL,
  `message_data_id` bigint(20) NOT NULL,
  `message_box_type` tinyint(4) NOT NULL,
  `reply_to_msg_id` int(11) NOT NULL DEFAULT '0',
  `mentioned` tinyint(4) NOT NULL DEFAULT '0',
  `media_unread` tinyint(4) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `message_datas`
--

CREATE TABLE `message_datas` (
  `id` int(11) NOT NULL,
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
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `orgs`
--

CREATE TABLE `orgs` (
  `org_id` int(11) NOT NULL,
  `account_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `passwd` char(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `org_name` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mail` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mobile` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `phones`
--

CREATE TABLE `phones` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `region` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'CN',
  `region_code` varchar(8) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '86',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `phone_books`
--

CREATE TABLE `phone_books` (
  `id` bigint(20) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL,
  `client_id` bigint(20) NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `first_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `last_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `phone_call_sessions`
--

CREATE TABLE `phone_call_sessions` (
  `id` int(11) NOT NULL,
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
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `photos`
--

CREATE TABLE `photos` (
  `id` int(11) NOT NULL,
  `photo_id` int(11) NOT NULL,
  `has_stickers` int(11) NOT NULL DEFAULT '0',
  `access_hash` int(11) NOT NULL,
  `date` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `photo_datas`
--

CREATE TABLE `photo_datas` (
  `id` int(11) NOT NULL,
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
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `popular_contacts`
--

CREATE TABLE `popular_contacts` (
  `id` bigint(20) NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `importers` int(11) NOT NULL DEFAULT '1',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `push_credentials`
--

CREATE TABLE `push_credentials` (
  `id` int(11) NOT NULL,
  `auth_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `token_type` tinyint(4) NOT NULL,
  `token` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `reports`
--

CREATE TABLE `reports` (
  `id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `peer_type` int(11) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `reason` tinyint(4) NOT NULL DEFAULT '0',
  `content` varchar(10000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `secret_messages`
--

CREATE TABLE `secret_messages` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `peer_type` int(11) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `random_id` bigint(20) NOT NULL,
  `message_content_header` int(11) NOT NULL,
  `message_content_data` blob NOT NULL,
  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `sticker_data`
--

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
  `image_512_height` int(11) DEFAULT '512'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `sticker_packs`
--

CREATE TABLE `sticker_packs` (
  `id` int(11) NOT NULL,
  `sticker_set_id` bigint(20) NOT NULL,
  `emoticon` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `document_id` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `sticker_sets`
--

CREATE TABLE `sticker_sets` (
  `id` int(11) NOT NULL,
  `sticker_set_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `title` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `short_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `count` int(11) NOT NULL DEFAULT '0',
  `hash` int(11) NOT NULL DEFAULT '0',
  `official` tinyint(4) NOT NULL DEFAULT '0',
  `mask` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `tmp_passwords`
--

CREATE TABLE `tmp_passwords` (
  `id` int(11) NOT NULL,
  `auth_id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `password_hash` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `period` int(11) NOT NULL,
  `tmp_password` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `valid_until` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `unread_mentions`
--

CREATE TABLE `unread_mentions` (
  `id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `mentioned_message_id` int(11) NOT NULL,
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `unregistered_contacts`
--

CREATE TABLE `unregistered_contacts` (
  `id` bigint(20) NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `importer_user_id` int(11) NOT NULL,
  `import_first_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `import_last_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `imported` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `username`
--

CREATE TABLE `username` (
  `id` bigint(20) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `username` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
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
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_contacts`
--

CREATE TABLE `user_contacts` (
  `id` int(11) NOT NULL,
  `owner_user_id` int(11) NOT NULL,
  `contact_user_id` int(11) NOT NULL,
  `contact_phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `contact_first_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `contact_last_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `mutual` tinyint(4) NOT NULL DEFAULT '0',
  `is_blocked` tinyint(1) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_dialogs`
--

CREATE TABLE `user_dialogs` (
  `id` int(11) NOT NULL,
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
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_notify_settings`
--

CREATE TABLE `user_notify_settings` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `peer_type` tinyint(4) NOT NULL,
  `peer_id` int(11) NOT NULL,
  `show_previews` tinyint(1) NOT NULL DEFAULT '0',
  `silent` tinyint(1) NOT NULL DEFAULT '0',
  `mute_until` int(11) NOT NULL DEFAULT '0',
  `sound` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'default',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_passwords`
--

CREATE TABLE `user_passwords` (
  `id` bigint(20) NOT NULL,
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
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_presences`
--

CREATE TABLE `user_presences` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `last_seen_at` bigint(20) NOT NULL,
  `last_seen_auth_key_id` bigint(20) NOT NULL,
  `last_seen_ip` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_privacys`
--

CREATE TABLE `user_privacys` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `key_type` tinyint(4) NOT NULL DEFAULT '0',
  `rules` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_profile_photos`
--

CREATE TABLE `user_profile_photos` (
  `id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `photo_id` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_pts_updates`
--

CREATE TABLE `user_pts_updates` (
  `id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `pts` int(11) NOT NULL,
  `pts_count` int(11) NOT NULL,
  `update_type` tinyint(4) NOT NULL DEFAULT '0',
  `update_data` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_qts_updates`
--

CREATE TABLE `user_qts_updates` (
  `id` bigint(20) NOT NULL,
  `user_id` int(11) NOT NULL,
  `qts` int(11) NOT NULL,
  `update_type` int(11) NOT NULL,
  `update_data` blob NOT NULL,
  `date2` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_sticker_sets`
--

CREATE TABLE `user_sticker_sets` (
  `id` int(10) UNSIGNED NOT NULL,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `sticker_set_id` bigint(20) NOT NULL DEFAULT '0',
  `archived` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `wall_papers`
--

CREATE TABLE `wall_papers` (
  `id` int(11) NOT NULL,
  `type` tinyint(4) NOT NULL DEFAULT '0',
  `title` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `color` int(11) NOT NULL DEFAULT '0',
  `bg_color` int(11) NOT NULL DEFAULT '0',
  `photo_id` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` bigint(20) NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `apps`
--
ALTER TABLE `apps`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `api_id` (`api_id`);

--
-- Indexes for table `app_configs`
--
ALTER TABLE `app_configs`
  ADD PRIMARY KEY (`app_id`);

--
-- Indexes for table `app_ios_push_certs`
--
ALTER TABLE `app_ios_push_certs`
  ADD PRIMARY KEY (`cert_id`);

--
-- Indexes for table `app_keys`
--
ALTER TABLE `app_keys`
  ADD PRIMARY KEY (`app_id`);

--
-- Indexes for table `auths`
--
ALTER TABLE `auths`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`);

--
-- Indexes for table `auth_channel_updates_state`
--
ALTER TABLE `auth_channel_updates_state`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`,`channel_id`);

--
-- Indexes for table `auth_keys`
--
ALTER TABLE `auth_keys`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`);

--
-- Indexes for table `auth_op_logs`
--
ALTER TABLE `auth_op_logs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `auth_phone_transactions`
--
ALTER TABLE `auth_phone_transactions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `transaction_hash` (`transaction_hash`);

--
-- Indexes for table `auth_salts`
--
ALTER TABLE `auth_salts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `auth` (`auth_key_id`);

--
-- Indexes for table `auth_seq_updates`
--
ALTER TABLE `auth_seq_updates`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_id` (`auth_id`,`user_id`,`seq`);

--
-- Indexes for table `auth_updates_state`
--
ALTER TABLE `auth_updates_state`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`,`user_id`);

--
-- Indexes for table `auth_users`
--
ALTER TABLE `auth_users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`,`user_id`),
  ADD KEY `auth_key_id_2` (`auth_key_id`,`user_id`,`deleted`);

--
-- Indexes for table `banned`
--
ALTER TABLE `banned`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`phone`);

--
-- Indexes for table `blocks`
--
ALTER TABLE `blocks`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `bots`
--
ALTER TABLE `bots`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `bot_id` (`bot_id`);

--
-- Indexes for table `bot_commands`
--
ALTER TABLE `bot_commands`
  ADD PRIMARY KEY (`id`),
  ADD KEY `bot_id` (`bot_id`);

--
-- Indexes for table `channels`
--
ALTER TABLE `channels`
  ADD PRIMARY KEY (`id`),
  ADD KEY `creator_user_id_3` (`creator_user_id`,`access_hash`);

--
-- Indexes for table `channel_admin_logs`
--
ALTER TABLE `channel_admin_logs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `channel_media_unread`
--
ALTER TABLE `channel_media_unread`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `channel_messages`
--
ALTER TABLE `channel_messages`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `sender_user_id` (`sender_user_id`,`random_id`),
  ADD UNIQUE KEY `channel_id` (`channel_id`,`channel_message_id`),
  ADD UNIQUE KEY `message_data_id` (`message_data_id`);

--
-- Indexes for table `channel_messages2`
--
ALTER TABLE `channel_messages2`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `channel_message_boxes`
--
ALTER TABLE `channel_message_boxes`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `channel_participants`
--
ALTER TABLE `channel_participants`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `channel_id` (`channel_id`,`user_id`),
  ADD KEY `chat_id` (`channel_id`);

--
-- Indexes for table `channel_pts_updates`
--
ALTER TABLE `channel_pts_updates`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `chats`
--
ALTER TABLE `chats`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `chat_participants`
--
ALTER TABLE `chat_participants`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `chat_id_2` (`chat_id`,`user_id`),
  ADD KEY `chat_id` (`chat_id`);

--
-- Indexes for table `config`
--
ALTER TABLE `config`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `devices`
--
ALTER TABLE `devices`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`,`user_id`,`token_type`);

--
-- Indexes for table `documents`
--
ALTER TABLE `documents`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `files`
--
ALTER TABLE `files`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `file_parts`
--
ALTER TABLE `file_parts`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `imported_contacts`
--
ALTER TABLE `imported_contacts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`),
  ADD UNIQUE KEY `user_id_2` (`user_id`,`imported_user_id`);

--
-- Indexes for table `messages`
--
ALTER TABLE `messages`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `message_boxes`
--
ALTER TABLE `message_boxes`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`message_data_id`);

--
-- Indexes for table `message_datas`
--
ALTER TABLE `message_datas`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `dialog_id` (`dialog_id`,`dialog_message_id`),
  ADD UNIQUE KEY `sender_user_id` (`sender_user_id`,`random_id`),
  ADD UNIQUE KEY `message_data_id` (`message_data_id`);

--
-- Indexes for table `orgs`
--
ALTER TABLE `orgs`
  ADD PRIMARY KEY (`org_id`),
  ADD UNIQUE KEY `account_name` (`account_name`);

--
-- Indexes for table `phones`
--
ALTER TABLE `phones`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `phone` (`phone`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `phone_books`
--
ALTER TABLE `phone_books`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`,`client_id`);

--
-- Indexes for table `phone_call_sessions`
--
ALTER TABLE `phone_call_sessions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `call_session_id` (`call_session_id`);

--
-- Indexes for table `photos`
--
ALTER TABLE `photos`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `photo_datas`
--
ALTER TABLE `photo_datas`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `popular_contacts`
--
ALTER TABLE `popular_contacts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `phone` (`phone`);

--
-- Indexes for table `push_credentials`
--
ALTER TABLE `push_credentials`
  ADD PRIMARY KEY (`id`),
  ADD KEY `auth_id` (`auth_id`,`user_id`);

--
-- Indexes for table `reports`
--
ALTER TABLE `reports`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `secret_messages`
--
ALTER TABLE `secret_messages`
  ADD PRIMARY KEY (`id`),
  ADD KEY `message_content_header` (`message_content_header`);

--
-- Indexes for table `sticker_data`
--
ALTER TABLE `sticker_data`
  ADD PRIMARY KEY (`id`,`pack_id`);

--
-- Indexes for table `sticker_packs`
--
ALTER TABLE `sticker_packs`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `sticker_sets`
--
ALTER TABLE `sticker_sets`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `sticker_set_id` (`sticker_set_id`);

--
-- Indexes for table `tmp_passwords`
--
ALTER TABLE `tmp_passwords`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `unread_mentions`
--
ALTER TABLE `unread_mentions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`,`mentioned_message_id`);

--
-- Indexes for table `unregistered_contacts`
--
ALTER TABLE `unregistered_contacts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `phone` (`phone`,`importer_user_id`),
  ADD KEY `phone_2` (`phone`,`importer_user_id`,`imported`);

--
-- Indexes for table `username`
--
ALTER TABLE `username`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `peer_type` (`peer_type`,`peer_id`),
  ADD KEY `username` (`username`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_contacts`
--
ALTER TABLE `user_contacts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `owner_user_id_2` (`owner_user_id`,`contact_phone`),
  ADD KEY `owner_user_id` (`owner_user_id`,`contact_user_id`);

--
-- Indexes for table `user_dialogs`
--
ALTER TABLE `user_dialogs`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`);

--
-- Indexes for table `user_notify_settings`
--
ALTER TABLE `user_notify_settings`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`);

--
-- Indexes for table `user_passwords`
--
ALTER TABLE `user_passwords`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`);

--
-- Indexes for table `user_presences`
--
ALTER TABLE `user_presences`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`),
  ADD KEY `user_id_2` (`user_id`,`last_seen_at`);

--
-- Indexes for table `user_privacys`
--
ALTER TABLE `user_privacys`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`key_type`);

--
-- Indexes for table `user_profile_photos`
--
ALTER TABLE `user_profile_photos`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_pts_updates`
--
ALTER TABLE `user_pts_updates`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_qts_updates`
--
ALTER TABLE `user_qts_updates`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_sticker_sets`
--
ALTER TABLE `user_sticker_sets`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uniq` (`user_id`,`sticker_set_id`) USING BTREE,
  ADD KEY `user_id` (`user_id`) USING BTREE;

--
-- Indexes for table `wall_papers`
--
ALTER TABLE `wall_papers`
  ADD PRIMARY KEY (`id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `apps`
--
ALTER TABLE `apps`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `app_configs`
--
ALTER TABLE `app_configs`
  MODIFY `app_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `app_ios_push_certs`
--
ALTER TABLE `app_ios_push_certs`
  MODIFY `cert_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `auths`
--
ALTER TABLE `auths`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `auth_channel_updates_state`
--
ALTER TABLE `auth_channel_updates_state`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `auth_keys`
--
ALTER TABLE `auth_keys`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `auth_op_logs`
--
ALTER TABLE `auth_op_logs`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `auth_phone_transactions`
--
ALTER TABLE `auth_phone_transactions`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `auth_salts`
--
ALTER TABLE `auth_salts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `auth_seq_updates`
--
ALTER TABLE `auth_seq_updates`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `auth_updates_state`
--
ALTER TABLE `auth_updates_state`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `auth_users`
--
ALTER TABLE `auth_users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `banned`
--
ALTER TABLE `banned`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `blocks`
--
ALTER TABLE `blocks`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `bots`
--
ALTER TABLE `bots`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `bot_commands`
--
ALTER TABLE `bot_commands`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `channels`
--
ALTER TABLE `channels`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `channel_admin_logs`
--
ALTER TABLE `channel_admin_logs`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `channel_media_unread`
--
ALTER TABLE `channel_media_unread`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `channel_messages`
--
ALTER TABLE `channel_messages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `channel_messages2`
--
ALTER TABLE `channel_messages2`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `channel_message_boxes`
--
ALTER TABLE `channel_message_boxes`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `channel_participants`
--
ALTER TABLE `channel_participants`
  MODIFY `id` bigint(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `channel_pts_updates`
--
ALTER TABLE `channel_pts_updates`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `chats`
--
ALTER TABLE `chats`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `chat_participants`
--
ALTER TABLE `chat_participants`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `devices`
--
ALTER TABLE `devices`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `documents`
--
ALTER TABLE `documents`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `files`
--
ALTER TABLE `files`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `file_parts`
--
ALTER TABLE `file_parts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `imported_contacts`
--
ALTER TABLE `imported_contacts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `messages`
--
ALTER TABLE `messages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `message_boxes`
--
ALTER TABLE `message_boxes`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `message_datas`
--
ALTER TABLE `message_datas`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `orgs`
--
ALTER TABLE `orgs`
  MODIFY `org_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `phones`
--
ALTER TABLE `phones`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `phone_books`
--
ALTER TABLE `phone_books`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `phone_call_sessions`
--
ALTER TABLE `phone_call_sessions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `photos`
--
ALTER TABLE `photos`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `photo_datas`
--
ALTER TABLE `photo_datas`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `popular_contacts`
--
ALTER TABLE `popular_contacts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `push_credentials`
--
ALTER TABLE `push_credentials`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `reports`
--
ALTER TABLE `reports`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `secret_messages`
--
ALTER TABLE `secret_messages`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `sticker_packs`
--
ALTER TABLE `sticker_packs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `sticker_sets`
--
ALTER TABLE `sticker_sets`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `tmp_passwords`
--
ALTER TABLE `tmp_passwords`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `unread_mentions`
--
ALTER TABLE `unread_mentions`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `unregistered_contacts`
--
ALTER TABLE `unregistered_contacts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `username`
--
ALTER TABLE `username`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_contacts`
--
ALTER TABLE `user_contacts`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_dialogs`
--
ALTER TABLE `user_dialogs`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_notify_settings`
--
ALTER TABLE `user_notify_settings`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_passwords`
--
ALTER TABLE `user_passwords`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_presences`
--
ALTER TABLE `user_presences`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_privacys`
--
ALTER TABLE `user_privacys`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_profile_photos`
--
ALTER TABLE `user_profile_photos`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_pts_updates`
--
ALTER TABLE `user_pts_updates`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_qts_updates`
--
ALTER TABLE `user_qts_updates`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `user_sticker_sets`
--
ALTER TABLE `user_sticker_sets`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `wall_papers`
--
ALTER TABLE `wall_papers`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
