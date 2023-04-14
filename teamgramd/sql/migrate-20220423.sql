ALTER TABLE `messages` ADD `reaction_unread` BOOLEAN NOT NULL DEFAULT FALSE AFTER `reaction_date`;
ALTER TABLE `messages` ADD `has_reaction` BOOLEAN NOT NULL DEFAULT FALSE AFTER `pinned`;
