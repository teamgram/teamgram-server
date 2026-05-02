-- user_global_privacy_settings
ALTER TABLE `user_global_privacy_settings` ADD `keep_archived_unmuted` BOOLEAN NOT NULL DEFAULT FALSE AFTER `archive_and_mute_new_noncontact_peers`, ADD `keep_archived_folders` BOOLEAN NOT NULL DEFAULT FALSE AFTER `keep_archived_unmuted`, ADD `hide_read_marks` BOOLEAN NOT NULL DEFAULT FALSE AFTER `keep_archived_folders`, ADD `new_noncontact_peers_require_premium` BOOLEAN NOT NULL DEFAULT FALSE AFTER `hide_read_marks`;

-- users: birthday
ALTER TABLE `users` ADD `birthday` CHAR(10) NOT NULL DEFAULT '' AFTER `profile_color_background_emoji_id`;
