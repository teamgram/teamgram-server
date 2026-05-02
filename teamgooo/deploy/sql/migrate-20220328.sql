ALTER TABLE `chats` ADD `available_reactions` VARCHAR(128) NOT NULL DEFAULT '' AFTER `migrated_to_access_hash`;
ALTER TABLE `messages` ADD `reaction` VARCHAR(16) NOT NULL DEFAULT '' AFTER `pinned`, ADD `reaction_date` BIGINT NOT NULL DEFAULT '0' AFTER `reaction`;
