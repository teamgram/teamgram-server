--
-- Tao bang `message_react_data`
--
CREATE TABLE `message_react_data` (
  `id` int(11) NOT NULL,
  `react_data_id` bigint(20) NOT NULL,
  `react_id` bigint(20) NOT NULL,
  `message_data_id` bigint(20) NOT NULL,
  `sender_user_id` int(11) NOT NULL,
  `date3` int(11) NOT NULL DEFAULT '0',
  `edit_date` int(11) NOT NULL DEFAULT '0',
  `deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
--
-- Them primary key  `message_react_data`
--
ALTER TABLE `message_react_data`
  ADD PRIMARY KEY (`id`);
ALTER TABLE `message_react_data`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- Tao bang  `message_react`
--
CREATE TABLE `message_react`(
`id` int(11) NOT NULL,
`react_id` bigint(20) NOT NULL,
`text` varchar(256) NOT NULL,
  `file_id` bigint(20) NOT NULL,
  `file_hash` bigint(20) NOT NULL,
  `file_size` bigint(20) NOT NULL,
  `width` int(11) DEFAULT '256',
  `height` int(11) DEFAULT '256'
  )ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
--
-- Them primary key  `message_react`
--
ALTER TABLE `message_react`
  ADD PRIMARY KEY (`id`);
ALTER TABLE `message_react`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;
--
-- Them data sample message_react
--
INSERT INTO message_react (react_id,text,file_id, file_hash, file_size,width,height)
VALUES (1,"Like",1, 1, 1,512,512);
INSERT INTO message_react (react_id,text,file_id, file_hash, file_size,width,height)
VALUES (2,"Love",2, 2, 2,512,512);
INSERT INTO message_react (react_id,text,file_id, file_hash, file_size,width,height)
VALUES (3,"Angry",3, 3, 3,512,512);
INSERT INTO message_react (react_id,text,file_id, file_hash, file_size,width,height)
VALUES (4,"Wow",4, 4, 4,512,512);
INSERT INTO message_react (react_id,text,file_id, file_hash, file_size,width,height)
VALUES (5,"Sad",5, 5, 5,512,512);
INSERT INTO message_react (react_id,text,file_id, file_hash, file_size,width,height)
VALUES (6,"Haha",6, 6, 6,512,512);
