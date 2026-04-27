-- Schemas for the idgen SeqStore (alloc.NewMySQLStore).
--
-- One row per table-local id. max_seq holds the next sequence value to be
-- allocated; the allocator advances it under SELECT ... FOR UPDATE so
-- concurrent callers serialize on the row lock.
--
-- min_seq is reserved for future use (range trimming / archival watermark)
-- and currently always defaults to 0.
--
-- Apply this schema in every database referenced by the idgen service's
-- `Mysql.DSN` config before enabling the service.

CREATE TABLE IF NOT EXISTS `message_data_ngen` (
  `id`      BIGINT NOT NULL,
  `min_seq` BIGINT NOT NULL DEFAULT 0,
  `max_seq` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `message_box_ngen` (
  `id`      BIGINT NOT NULL,
  `min_seq` BIGINT NOT NULL DEFAULT 0,
  `max_seq` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `channel_message_box_ngen` (
  `id`      BIGINT NOT NULL,
  `min_seq` BIGINT NOT NULL DEFAULT 0,
  `max_seq` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `seq_updates_ngen` (
  `id`      BIGINT NOT NULL,
  `min_seq` BIGINT NOT NULL DEFAULT 0,
  `max_seq` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `pts_updates_ngen` (
  `id`      BIGINT NOT NULL,
  `min_seq` BIGINT NOT NULL DEFAULT 0,
  `max_seq` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `qts_updates_ngen` (
  `id`      BIGINT NOT NULL,
  `min_seq` BIGINT NOT NULL DEFAULT 0,
  `max_seq` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `channel_pts_updates_ngen` (
  `id`      BIGINT NOT NULL,
  `min_seq` BIGINT NOT NULL DEFAULT 0,
  `max_seq` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `scheduled_ngen` (
  `id`      BIGINT NOT NULL,
  `min_seq` BIGINT NOT NULL DEFAULT 0,
  `max_seq` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `bot_updates_ngen` (
  `id`      BIGINT NOT NULL,
  `min_seq` BIGINT NOT NULL DEFAULT 0,
  `max_seq` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `story_ngen` (
  `id`      BIGINT NOT NULL,
  `min_seq` BIGINT NOT NULL DEFAULT 0,
  `max_seq` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE IF NOT EXISTS `channel_story_ngen` (
  `id`      BIGINT NOT NULL,
  `min_seq` BIGINT NOT NULL DEFAULT 0,
  `max_seq` BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
