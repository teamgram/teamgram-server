ALTER TABLE `message_boxes` ADD `pts` INT NOT NULL DEFAULT '0' AFTER `message_data_id`;
-- RENAME TABLE `chatengine`.`blocks` TO `chatengine`.`user_blocks`;
DROP TABLE `channels`, `channel_admin_logs`, `channel_media_unread`, `channel_messages`, `channel_message_boxes`, `channel_participants`, `channel_pts_updates`;
DROP TABLE `secret_messages`;
DROP TABLE `auth_channel_updates_state`;
DROP TABLE `config`;

--
-- 表的结构 `user_blocks`
--

CREATE TABLE `user_blocks` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `block_id` int(11) NOT NULL DEFAULT '0',
  `date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `user_blocks`
--
ALTER TABLE `user_blocks`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`user_id`,`block_id`),
  ADD KEY `user_id_2` (`user_id`,`block_id`,`deleted`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `user_blocks`
--
ALTER TABLE `user_blocks`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

-- drop is_blocked
ALTER TABLE `user_contacts`
  DROP `is_blocked`;
