ALTER TABLE `chat_participants` ADD `is_bot` BOOLEAN NOT NULL DEFAULT FALSE AFTER `groupcall_default_join_as_peer_id`;
