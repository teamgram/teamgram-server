ALTER TABLE `teamgram`.`documents` DROP INDEX `document_id`, ADD UNIQUE `document_id` (`document_id`) USING BTREE;
ALTER TABLE `user_contacts` ADD INDEX(`owner_user_id`);
