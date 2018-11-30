ALTER TABLE `message_boxes` ADD `pts` INT NOT NULL DEFAULT '0' AFTER `message_data_id`;
RENAME TABLE `chatengine`.`blocks` TO `chatengine`.`user_blocks`;
DROP TABLE `channels`, `channel_admin_logs`, `channel_media_unread`, `channel_messages`, `channel_message_boxes`, `channel_participants`, `channel_pts_updates`;
DROP TABLE `secret_messages`;
DROP TABLE `auth_channel_updates_state`;
DROP TABLE `config`;

