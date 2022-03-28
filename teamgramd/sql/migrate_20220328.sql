ALTER TABLE `chats` ADD `available_reactions` VARCHAR(128) NOT NULL DEFAULT '' AFTER `migrated_to_access_hash`;
