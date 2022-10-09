ALTER TABLE `dialogs` ADD `ttl_period` INT NOT NULL DEFAULT '0' AFTER `has_scheduled`;
ALTER TABLE `messages` ADD `ttl_period` INT NOT NULL DEFAULT '0' AFTER `date2`;
