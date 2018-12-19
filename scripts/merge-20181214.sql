ALTER TABLE `chats` ADD `link` VARCHAR(128) NOT NULL DEFAULT '' AFTER `title`;
ALTER TABLE `chats` ADD `migrated_to` int(11) NOT NULL DEFAULT 0 AFTER `admins_enabled`;
