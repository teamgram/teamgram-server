-- Schema for the idgen SeqStore (alloc.NewMySQLStore).
--
-- One row per logical key (e.g. an inbox or conversation id). max_seq holds
-- the next id to be allocated; the allocator advances it under SELECT ... FOR
-- UPDATE so concurrent callers serialize on the row lock.
--
-- min_seq is reserved for future use (range trimming / archival watermark)
-- and currently always defaults to 0.
--
-- Apply this schema in every database referenced by the idgen service's
-- `Mysql.DSN` config before enabling the service. The table name must match
-- alloc.DefaultSeqTable (or the explicit name passed to
-- alloc.NewMySQLStoreWithTable).

CREATE TABLE IF NOT EXISTS `seq_conversations` (
  `conversation_id` VARCHAR(191) NOT NULL,
  `min_seq`         BIGINT       NOT NULL DEFAULT 0,
  `max_seq`         BIGINT       NOT NULL DEFAULT 0,
  PRIMARY KEY (`conversation_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
