--
-- 表的结构 `banned`
--

DROP TABLE IF EXISTS `banned`;
CREATE TABLE `banned` (
  `id` int(11) NOT NULL,
  `phone` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `banned_time` bigint(20) NOT NULL,
  `expires` bigint(20) NOT NULL DEFAULT '0',
  `banned_reason` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `log` varchar(1024) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `state` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `banned`
--
ALTER TABLE `banned`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `user_id` (`phone`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `banned`
--
ALTER TABLE `banned`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;
