ALTER TABLE `auth_users` ADD `android_push_session_id` BIGINT NOT NULL DEFAULT '0' AFTER `state`;
