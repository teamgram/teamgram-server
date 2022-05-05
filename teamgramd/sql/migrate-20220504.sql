ALTER TABLE `teamgram`.`documents` DROP INDEX `document_id`, ADD UNIQUE `document_id` (`document_id`) USING BTREE;
