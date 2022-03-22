ALTER TABLE `chat_participants` ADD `groupcall_default_join_as_peer_type` INT NOT NULL DEFAULT '0' AFTER `left_at`, ADD `groupcall_default_join_as_peer_id` BIGINT NOT NULL DEFAULT '0' AFTER `groupcall_default_join_as_peer_type`;
ALTER TABLE `dialogs` ADD `theme_emoticon` VARCHAR(64) NOT NULL DEFAULT '' AFTER `has_scheduled`;
