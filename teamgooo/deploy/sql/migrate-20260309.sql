-- SQL performance / index adjustments collected from specs
-- NOTE: Review and apply selectively in your environments.

-- 1) auth_seq_updates: support queries by (auth_id, user_id, date2) ordered by seq
ALTER TABLE `auth_seq_updates`
  ADD KEY `idx_auth_user_date` (`auth_id`, `user_id`, `date2`, `seq`);

-- 2) auth_users: support queries by user_id + deleted (session lookups & deletes)
ALTER TABLE `auth_users`
  ADD KEY `idx_user_deleted` (`user_id`, `deleted`);

-- 3) auth_op_logs: align columns with IPv6 and default/value expectations
--
-- 表的结构 `auth_op_logs`
--

CREATE TABLE IF NOT EXISTS `auth_op_logs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `auth_key_id` bigint(20) NOT NULL,
  `ip` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `op_type` int(11) NOT NULL DEFAULT '1',
  `log_text` varchar(512) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


ALTER TABLE `auth_op_logs`
  MODIFY COLUMN `ip` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'client ip, IPv6 max 45',
  MODIFY COLUMN `log_text` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'log content',
  COMMENT = 'Authorization operation log';

-- 4) chats: speed up \"last created chat\" lookup by creator
ALTER TABLE `chats`
  ADD KEY `idx_creator_date` (`creator_user_id`, `date`);

-- 5) chat_invite_participants: speed up operations by chat_id / (chat_id, user_id)
ALTER TABLE `chat_invite_participants`
  ADD KEY `idx_chat_user` (`chat_id`, `user_id`),
  ADD KEY `idx_chat_requested` (`chat_id`, `requested`);

-- 6) chat_participants: speed up per-user chat scans
ALTER TABLE `chat_participants`
  ADD KEY `idx_user_state` (`user_id`, `state`);

