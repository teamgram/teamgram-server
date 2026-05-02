-- phpMyAdmin SQL Dump
-- version 5.2.0-dev+20211229.cb650d2a2d
-- https://www.phpmyadmin.net/
--
-- 主机： 127.0.0.1
-- 生成日期： 2022-03-01 12:33:39
-- 服务器版本： 5.7.37
-- PHP 版本： 8.1.2
use teamgram;
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `teamgram2`
--

-- --------------------------------------------------------

--
-- 表的结构 `auths`
--

CREATE TABLE `auths` (
  `id` bigint(20) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL,
  `layer` int(11) NOT NULL DEFAULT '0',
  `api_id` int(11) NOT NULL,
  `device_model` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `system_version` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `app_version` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `system_lang_code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `lang_pack` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `lang_code` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `system_code` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `proxy` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `params` json NOT NULL,
  `client_ip` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `date_active` bigint(20) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `auth_keys`
--

CREATE TABLE `auth_keys` (
  `id` bigint(20) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL COMMENT 'auth_id',
  `body` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'auth_key，原始数据为256的二进制数据，存储时转换成base64格式',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `auth_key_infos`
--

CREATE TABLE `auth_key_infos` (
  `id` bigint(20) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL,
  `auth_key_type` int(11) NOT NULL,
  `perm_auth_key_id` bigint(20) NOT NULL DEFAULT '0',
  `temp_auth_key_id` bigint(20) NOT NULL DEFAULT '0',
  `media_temp_auth_key_id` bigint(20) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `auth_seq_updates`
--

CREATE TABLE `auth_seq_updates` (
  `id` bigint(20) NOT NULL,
  `auth_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `seq` int(11) NOT NULL DEFAULT '0',
  `update_type` int(11) NOT NULL DEFAULT '0',
  `update_data` json NOT NULL,
  `date2` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `auth_users`
--

CREATE TABLE `auth_users` (
  `id` bigint(20) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL DEFAULT '0',
  `hash` bigint(20) NOT NULL DEFAULT '0',
  `layer` int(11) NOT NULL DEFAULT '0',
  `device_model` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `platform` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `system_version` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `api_id` int(11) NOT NULL DEFAULT '0',
  `app_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `app_version` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `date_created` bigint(20) NOT NULL DEFAULT '0',
  `date_actived` bigint(20) NOT NULL DEFAULT '0',
  `ip` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `country` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `region` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `bots`
--

CREATE TABLE `bots` (
  `id` bigint(20) NOT NULL,
  `bot_id` bigint(20) NOT NULL,
  `bot_type` int(11) NOT NULL DEFAULT '0',
  `creator_user_id` bigint(20) NOT NULL DEFAULT '0',
  `token` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `description` varchar(10240) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `bot_chat_history` tinyint(1) NOT NULL DEFAULT '0',
  `bot_nochats` tinyint(1) NOT NULL DEFAULT '1',
  `verified` tinyint(1) NOT NULL DEFAULT '0',
  `bot_inline_geo` tinyint(1) NOT NULL DEFAULT '0',
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
  `id` bigint(20) NOT NULL,
  `bot_id` bigint(20) NOT NULL,
  `command` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `description` varchar(10240) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `chats`
--

CREATE TABLE `chats` (
  `id` bigint(20) NOT NULL,
  `creator_user_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `random_id` bigint(20) NOT NULL,
  `participant_count` int(11) NOT NULL,
  `title` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `about` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `photo_id` bigint(20) NOT NULL DEFAULT '0',
  `default_banned_rights` bigint(20) NOT NULL DEFAULT '0',
  `migrated_to_id` bigint(20) NOT NULL DEFAULT '0',
  `migrated_to_access_hash` bigint(20) NOT NULL DEFAULT '0',
  `deactivated` tinyint(1) NOT NULL DEFAULT '0',
  `version` int(11) NOT NULL DEFAULT '1',
  `date` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `chat_invites`
--

CREATE TABLE `chat_invites` (
  `id` bigint(20) NOT NULL,
  `chat_id` bigint(20) NOT NULL,
  `admin_id` bigint(20) NOT NULL,
  `migrated_to_id` bigint(20) NOT NULL DEFAULT '0',
  `link` varchar(64) COLLATE utf8mb4_bin NOT NULL,
  `permanent` tinyint(1) NOT NULL DEFAULT '0',
  `revoked` tinyint(1) NOT NULL DEFAULT '0',
  `request_needed` tinyint(1) NOT NULL DEFAULT '0',
  `start_date` bigint(20) NOT NULL DEFAULT '0',
  `expire_date` bigint(20) NOT NULL DEFAULT '0',
  `usage_limit` int(11) NOT NULL DEFAULT '0',
  `usage2` int(11) NOT NULL DEFAULT '0',
  `requested` int(11) NOT NULL DEFAULT '0',
  `title` varchar(64) COLLATE utf8mb4_bin NOT NULL DEFAULT '',
  `date2` bigint(20) NOT NULL,
  `state` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- 表的结构 `chat_invite_participants`
--

CREATE TABLE `chat_invite_participants` (
  `id` bigint(20) NOT NULL,
  `link` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `date2` bigint(20) NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- --------------------------------------------------------

--
-- 表的结构 `chat_participants`
--

CREATE TABLE `chat_participants` (
  `id` bigint(20) NOT NULL,
  `chat_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `participant_type` int(11) DEFAULT '0',
  `link` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `usage2` int(11) NOT NULL DEFAULT '0',
  `admin_rights` int(11) NOT NULL DEFAULT '0',
  `inviter_user_id` bigint(20) NOT NULL DEFAULT '0',
  `invited_at` bigint(20) NOT NULL DEFAULT '0',
  `kicked_at` bigint(20) NOT NULL DEFAULT '0',
  `left_at` bigint(20) NOT NULL DEFAULT '0',
  `state` int(11) NOT NULL DEFAULT '0',
  `date2` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `devices`
--

CREATE TABLE `devices` (
  `id` bigint(20) NOT NULL,
  `auth_key_id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `token_type` int(11) NOT NULL,
  `token` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `no_muted` tinyint(1) NOT NULL DEFAULT '0',
  `locked_period` int(11) NOT NULL DEFAULT '0',
  `app_sandbox` tinyint(1) NOT NULL DEFAULT '0',
  `secret` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `other_uids` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `dialogs`
--

CREATE TABLE `dialogs` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `peer_type` int(11) NOT NULL DEFAULT '2',
  `peer_id` bigint(20) NOT NULL,
  `peer_dialog_id` bigint(20) NOT NULL,
  `pinned` bigint(20) NOT NULL DEFAULT '0',
  `top_message` int(11) NOT NULL DEFAULT '0',
  `pinned_msg_id` int(11) NOT NULL DEFAULT '0',
  `read_inbox_max_id` int(11) NOT NULL DEFAULT '0',
  `read_outbox_max_id` int(11) NOT NULL DEFAULT '0',
  `unread_count` int(11) NOT NULL DEFAULT '0',
  `unread_mentions_count` int(11) NOT NULL DEFAULT '0',
  `unread_mark` tinyint(1) NOT NULL DEFAULT '0',
  `draft_type` int(11) NOT NULL DEFAULT '0',
  `draft_message_data` json NOT NULL,
  `folder_id` int(11) NOT NULL DEFAULT '0',
  `folder_pinned` bigint(20) NOT NULL DEFAULT '0',
  `has_scheduled` tinyint(1) NOT NULL DEFAULT '0',
  `date2` bigint(20) NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `dialog_filters`
--

CREATE TABLE `dialog_filters` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `dialog_filter_id` int(11) NOT NULL,
  `dialog_filter` json NOT NULL,
  `order_value` bigint(20) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
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
  `video_thumb_id` bigint(20) NOT NULL DEFAULT '0',
  `version` int(11) NOT NULL DEFAULT '0',
  `attributes` json NOT NULL,
  `date2` bigint(20) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `encrypted_files`
--

CREATE TABLE `encrypted_files` (
  `id` bigint(20) NOT NULL,
  `encrypted_file_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `dc_id` int(11) NOT NULL,
  `file_size` int(11) NOT NULL,
  `key_fingerprint` int(11) NOT NULL,
  `md5_checksum` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `file_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `hash_tags`
--

CREATE TABLE `hash_tags` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `peer_type` int(11) NOT NULL,
  `peer_id` bigint(20) NOT NULL,
  `hash_tag` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
  `hash_tag_message_id` int(11) NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `imported_contacts`
--

CREATE TABLE `imported_contacts` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `imported_user_id` bigint(20) NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `messages`
--

CREATE TABLE `messages` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `user_message_box_id` int(11) NOT NULL,
  `dialog_id1` bigint(20) NOT NULL,
  `dialog_id2` bigint(20) NOT NULL,
  `dialog_message_id` bigint(20) NOT NULL,
  `sender_user_id` bigint(20) NOT NULL,
  `peer_type` int(11) NOT NULL,
  `peer_id` bigint(20) NOT NULL,
  `random_id` bigint(20) NOT NULL DEFAULT '0',
  `message_filter_type` int(11) NOT NULL DEFAULT '0',
  `message_data` json DEFAULT NULL,
  `message` varchar(6000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `mentioned` tinyint(1) NOT NULL DEFAULT '0',
  `media_unread` tinyint(1) NOT NULL DEFAULT '0',
  `pinned` tinyint(1) NOT NULL DEFAULT '0',
  `date2` bigint(20) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `phone_books`
--

CREATE TABLE `phone_books` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL DEFAULT '0',
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
-- 表的结构 `photos`
--

CREATE TABLE `photos` (
  `id` bigint(20) NOT NULL,
  `photo_id` bigint(20) NOT NULL,
  `access_hash` bigint(20) NOT NULL,
  `has_stickers` tinyint(1) NOT NULL DEFAULT '0',
  `dc_id` int(11) NOT NULL DEFAULT '2',
  `date2` bigint(20) NOT NULL DEFAULT '0',
  `has_video` tinyint(1) NOT NULL DEFAULT '0',
  `size_id` bigint(20) NOT NULL DEFAULT '0',
  `video_size_id` bigint(20) NOT NULL DEFAULT '0',
  `input_file_name` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `ext` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `photo_sizes`
--

CREATE TABLE `photo_sizes` (
  `id` bigint(20) NOT NULL,
  `photo_size_id` bigint(20) NOT NULL,
  `size_type` char(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `volume_id` bigint(20) NOT NULL,
  `local_id` int(11) NOT NULL,
  `secret` bigint(20) NOT NULL,
  `width` int(11) NOT NULL,
  `height` int(11) NOT NULL,
  `file_size` int(11) NOT NULL,
  `file_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `has_stripped` tinyint(1) NOT NULL DEFAULT '0',
  `stripped_bytes` varchar(4096) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
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
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `predefined_users`
--

CREATE TABLE `predefined_users` (
  `id` bigint(20) NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `first_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `last_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `username` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `code` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `verified` tinyint(1) NOT NULL DEFAULT '0',
  `registered_user_id` bigint(20) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
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
  `importer_user_id` bigint(20) NOT NULL,
  `import_first_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `import_last_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `imported` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `username`
--

CREATE TABLE `username` (
  `id` bigint(20) NOT NULL,
  `username` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `peer_type` int(11) NOT NULL DEFAULT '0',
  `peer_id` bigint(20) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `user_type` int(11) NOT NULL DEFAULT '2',
  `access_hash` bigint(20) NOT NULL,
  `secret_key_id` bigint(20) NOT NULL DEFAULT '0',
  `first_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `last_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `username` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `country_code` varchar(3) COLLATE utf8mb4_unicode_ci NOT NULL,
  `verified` tinyint(1) NOT NULL DEFAULT '0',
  `support` tinyint(1) NOT NULL DEFAULT '0',
  `scam` tinyint(1) NOT NULL DEFAULT '0',
  `fake` tinyint(1) NOT NULL DEFAULT '0',
  `about` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` int(11) NOT NULL DEFAULT '0',
  `is_bot` tinyint(1) NOT NULL DEFAULT '0',
  `account_days_ttl` int(11) NOT NULL DEFAULT '180',
  `photo_id` bigint(20) NOT NULL DEFAULT '0',
  `restricted` tinyint(1) NOT NULL DEFAULT '0',
  `restriction_reason` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `archive_and_mute_new_noncontact_peers` tinyint(1) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `delete_reason` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_contacts`
--

CREATE TABLE `user_contacts` (
  `id` bigint(20) NOT NULL,
  `owner_user_id` bigint(20) NOT NULL,
  `contact_user_id` bigint(20) NOT NULL,
  `contact_phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `contact_first_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `contact_last_name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `mutual` tinyint(1) NOT NULL DEFAULT '0',
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `date2` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_global_privacy_settings`
--

CREATE TABLE `user_global_privacy_settings` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `archive_and_mute_new_noncontact_peers` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_notify_settings`
--

CREATE TABLE `user_notify_settings` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `peer_type` int(11) NOT NULL,
  `peer_id` bigint(20) NOT NULL,
  `show_previews` int(11) NOT NULL DEFAULT '-1',
  `silent` int(11) NOT NULL DEFAULT '-1',
  `mute_until` int(11) NOT NULL DEFAULT '-1',
  `sound` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'default',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_peer_blocks`
--

CREATE TABLE `user_peer_blocks` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `peer_type` int(11) NOT NULL,
  `peer_id` bigint(20) NOT NULL,
  `date` bigint(20) NOT NULL DEFAULT '0',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_peer_settings`
--

CREATE TABLE `user_peer_settings` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `peer_type` int(11) NOT NULL,
  `peer_id` bigint(20) NOT NULL,
  `hide` tinyint(1) NOT NULL DEFAULT '0',
  `report_spam` tinyint(1) NOT NULL DEFAULT '0',
  `add_contact` tinyint(1) NOT NULL DEFAULT '0',
  `block_contact` tinyint(1) NOT NULL DEFAULT '0',
  `share_contact` tinyint(1) NOT NULL DEFAULT '0',
  `need_contacts_exception` tinyint(1) NOT NULL DEFAULT '0',
  `report_geo` tinyint(1) NOT NULL DEFAULT '0',
  `autoarchived` tinyint(1) NOT NULL DEFAULT '0',
  `invite_members` tinyint(1) NOT NULL DEFAULT '0',
  `geo_distance` int(11) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_presences`
--

CREATE TABLE `user_presences` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `last_seen_at` bigint(20) NOT NULL,
  `expires` int(10) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_privacies`
--

CREATE TABLE `user_privacies` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `key_type` int(11) NOT NULL DEFAULT '0',
  `rules` json NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_profile_photos`
--

CREATE TABLE `user_profile_photos` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `photo_id` bigint(20) NOT NULL,
  `date2` bigint(20) NOT NULL COMMENT '排序',
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- --------------------------------------------------------

--
-- 表的结构 `user_pts_updates`
--

CREATE TABLE `user_pts_updates` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `pts` int(11) NOT NULL,
  `pts_count` int(11) NOT NULL,
  `update_type` tinyint(4) NOT NULL DEFAULT '0',
  `update_data` json NOT NULL,
  `date2` bigint(20) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `user_settings`
--

CREATE TABLE `user_settings` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `key2` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `deleted` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- --------------------------------------------------------

--
-- 表的结构 `video_sizes`
--

CREATE TABLE `video_sizes` (
  `id` bigint(20) NOT NULL,
  `video_size_id` bigint(20) NOT NULL,
  `size_type` char(1) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `volume_id` bigint(20) NOT NULL,
  `local_id` int(11) NOT NULL,
  `secret` bigint(20) NOT NULL,
  `width` int(11) NOT NULL,
  `height` int(11) NOT NULL,
  `file_size` int(11) NOT NULL DEFAULT '0',
  `video_start_ts` double NOT NULL DEFAULT '0',
  `file_path` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- 转储表的索引
--

--
-- 表的索引 `auths`
--
ALTER TABLE `auths`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`);

--
-- 表的索引 `auth_keys`
--
ALTER TABLE `auth_keys`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`);

--
-- 表的索引 `auth_key_infos`
--
ALTER TABLE `auth_key_infos`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`);

--
-- 表的索引 `auth_seq_updates`
--
ALTER TABLE `auth_seq_updates`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_id` (`auth_id`,`user_id`,`seq`);

--
-- 表的索引 `auth_users`
--
ALTER TABLE `auth_users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`,`user_id`),
  ADD KEY `auth_key_id_2` (`auth_key_id`,`user_id`,`deleted`);

--
-- 表的索引 `bots`
--
ALTER TABLE `bots`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `bot_id` (`bot_id`),
  ADD UNIQUE KEY `token` (`token`);

--
-- 表的索引 `bot_commands`
--
ALTER TABLE `bot_commands`
  ADD PRIMARY KEY (`id`),
  ADD KEY `bot_id` (`bot_id`);

--
-- 表的索引 `chats`
--
ALTER TABLE `chats`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `chat_invites`
--
ALTER TABLE `chat_invites`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `link` (`link`),
  ADD KEY `chat_id` (`chat_id`,`permanent`,`admin_id`),
  ADD KEY `chat_id_2` (`chat_id`,`admin_id`);

--
-- 表的索引 `chat_invite_participants`
--
ALTER TABLE `chat_invite_participants`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `link` (`link`,`user_id`),
  ADD KEY `link_2` (`link`);

--
-- 表的索引 `chat_participants`
--
ALTER TABLE `chat_participants`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `chat_id_2` (`chat_id`,`user_id`),
  ADD KEY `chat_id` (`chat_id`);

--
-- 表的索引 `devices`
--
ALTER TABLE `devices`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`,`user_id`,`token_type`);

--
-- 表的索引 `dialogs`
--
ALTER TABLE `dialogs`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`),
  ADD UNIQUE KEY `user_id_2` (`user_id`,`peer_dialog_id`);

--
-- 表的索引 `dialog_filters`
--
ALTER TABLE `dialog_filters`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `documents`
--
ALTER TABLE `documents`
  ADD PRIMARY KEY (`id`),
  ADD KEY `document_id` (`document_id`);

--
-- 表的索引 `encrypted_files`
--
ALTER TABLE `encrypted_files`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `hash_tags`
--
ALTER TABLE `hash_tags`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id_4` (`user_id`,`hash_tag`,`hash_tag_message_id`),
  ADD KEY `user_id` (`user_id`,`hash_tag`),
  ADD KEY `user_id_2` (`user_id`,`peer_type`,`peer_id`,`hash_tag`),
  ADD KEY `user_id_3` (`user_id`,`hash_tag_message_id`);

--
-- 表的索引 `imported_contacts`
--
ALTER TABLE `imported_contacts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`),
  ADD UNIQUE KEY `user_id_2` (`user_id`,`imported_user_id`);

--
-- 表的索引 `messages`
--
ALTER TABLE `messages`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id_2` (`user_id`,`user_message_box_id`),
  ADD KEY `user_id` (`user_id`,`dialog_id1`,`dialog_id2`);

--
-- 表的索引 `phone_books`
--
ALTER TABLE `phone_books`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `auth_key_id` (`auth_key_id`,`client_id`);

--
-- 表的索引 `photos`
--
ALTER TABLE `photos`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `photo_id` (`photo_id`);

--
-- 表的索引 `photo_sizes`
--
ALTER TABLE `photo_sizes`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `photo_size_id` (`photo_size_id`,`size_type`),
  ADD KEY `photo_id` (`photo_size_id`),
  ADD KEY `volume_id` (`volume_id`,`local_id`);

--
-- 表的索引 `popular_contacts`
--
ALTER TABLE `popular_contacts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `phone` (`phone`);

--
-- 表的索引 `predefined_users`
--
ALTER TABLE `predefined_users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `phone` (`phone`);

--
-- 表的索引 `unregistered_contacts`
--
ALTER TABLE `unregistered_contacts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `phone` (`phone`,`importer_user_id`),
  ADD KEY `phone_2` (`phone`,`importer_user_id`,`imported`),
  ADD KEY `phone_3` (`phone`);

--
-- 表的索引 `username`
--
ALTER TABLE `username`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`);

--
-- 表的索引 `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `phone` (`phone`);

--
-- 表的索引 `user_contacts`
--
ALTER TABLE `user_contacts`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `owner_user_id` (`owner_user_id`,`contact_user_id`);

--
-- 表的索引 `user_global_privacy_settings`
--
ALTER TABLE `user_global_privacy_settings`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`);

--
-- 表的索引 `user_notify_settings`
--
ALTER TABLE `user_notify_settings`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`);

--
-- 表的索引 `user_peer_blocks`
--
ALTER TABLE `user_peer_blocks`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id_2` (`user_id`,`peer_type`,`peer_id`),
  ADD KEY `user_id` (`user_id`);

--
-- 表的索引 `user_peer_settings`
--
ALTER TABLE `user_peer_settings`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`peer_type`,`peer_id`);

--
-- 表的索引 `user_presences`
--
ALTER TABLE `user_presences`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`);

--
-- 表的索引 `user_privacies`
--
ALTER TABLE `user_privacies`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`key_type`);

--
-- 表的索引 `user_profile_photos`
--
ALTER TABLE `user_profile_photos`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`photo_id`),
  ADD KEY `user_id_2` (`user_id`);

--
-- 表的索引 `user_pts_updates`
--
ALTER TABLE `user_pts_updates`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`,`pts`);

--
-- 表的索引 `user_settings`
--
ALTER TABLE `user_settings`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`key2`);

--
-- 表的索引 `video_sizes`
--
ALTER TABLE `video_sizes`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `video_size_id` (`video_size_id`,`size_type`),
  ADD KEY `volume_id` (`volume_id`,`local_id`),
  ADD KEY `video_size_id_2` (`video_size_id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `auths`
--
ALTER TABLE `auths`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=202910;

--
-- 使用表AUTO_INCREMENT `auth_keys`
--
ALTER TABLE `auth_keys`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=213484;

--
-- 使用表AUTO_INCREMENT `auth_seq_updates`
--
ALTER TABLE `auth_seq_updates`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `auth_users`
--
ALTER TABLE `auth_users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=120061;

--
-- 使用表AUTO_INCREMENT `bots`
--
ALTER TABLE `bots`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- 使用表AUTO_INCREMENT `bot_commands`
--
ALTER TABLE `bot_commands`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=44;

--
-- 使用表AUTO_INCREMENT `chats`
--
ALTER TABLE `chats`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=120031;

--
-- 使用表AUTO_INCREMENT `chat_invites`
--
ALTER TABLE `chat_invites`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=35;

--
-- 使用表AUTO_INCREMENT `chat_invite_participants`
--
ALTER TABLE `chat_invite_participants`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `chat_participants`
--
ALTER TABLE `chat_participants`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=120068;

--
-- 使用表AUTO_INCREMENT `devices`
--
ALTER TABLE `devices`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- 使用表AUTO_INCREMENT `dialogs`
--
ALTER TABLE `dialogs`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=152283;

--
-- 使用表AUTO_INCREMENT `dialog_filters`
--
ALTER TABLE `dialog_filters`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- 使用表AUTO_INCREMENT `documents`
--
ALTER TABLE `documents`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=120762;

--
-- 使用表AUTO_INCREMENT `encrypted_files`
--
ALTER TABLE `encrypted_files`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `hash_tags`
--
ALTER TABLE `hash_tags`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `imported_contacts`
--
ALTER TABLE `imported_contacts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `messages`
--
ALTER TABLE `messages`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=155566;

--
-- 使用表AUTO_INCREMENT `phone_books`
--
ALTER TABLE `phone_books`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=91195;

--
-- 使用表AUTO_INCREMENT `photos`
--
ALTER TABLE `photos`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=90105;

--
-- 使用表AUTO_INCREMENT `photo_sizes`
--
ALTER TABLE `photo_sizes`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=121104;

--
-- 使用表AUTO_INCREMENT `popular_contacts`
--
ALTER TABLE `popular_contacts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `predefined_users`
--
ALTER TABLE `predefined_users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;

--
-- 使用表AUTO_INCREMENT `unregistered_contacts`
--
ALTER TABLE `unregistered_contacts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=91062;

--
-- 使用表AUTO_INCREMENT `username`
--
ALTER TABLE `username`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=90183;

--
-- 使用表AUTO_INCREMENT `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=136907713;

--
-- 使用表AUTO_INCREMENT `user_contacts`
--
ALTER TABLE `user_contacts`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=60039;

--
-- 使用表AUTO_INCREMENT `user_global_privacy_settings`
--
ALTER TABLE `user_global_privacy_settings`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=90025;

--
-- 使用表AUTO_INCREMENT `user_notify_settings`
--
ALTER TABLE `user_notify_settings`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=67;

--
-- 使用表AUTO_INCREMENT `user_peer_blocks`
--
ALTER TABLE `user_peer_blocks`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- 使用表AUTO_INCREMENT `user_peer_settings`
--
ALTER TABLE `user_peer_settings`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=90018;

--
-- 使用表AUTO_INCREMENT `user_presences`
--
ALTER TABLE `user_presences`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=90075;

--
-- 使用表AUTO_INCREMENT `user_privacies`
--
ALTER TABLE `user_privacies`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=90193;

--
-- 使用表AUTO_INCREMENT `user_profile_photos`
--
ALTER TABLE `user_profile_photos`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- 使用表AUTO_INCREMENT `user_pts_updates`
--
ALTER TABLE `user_pts_updates`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=156759;

--
-- 使用表AUTO_INCREMENT `user_settings`
--
ALTER TABLE `user_settings`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `video_sizes`
--
ALTER TABLE `video_sizes`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
