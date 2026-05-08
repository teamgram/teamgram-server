-- Public per-user message id storage for single-chat user projections.
-- This migration preserves peer_seq as the internal dialog ordering key.

CREATE TABLE IF NOT EXISTS `user_message_sequences` (
  `user_id` bigint NOT NULL,
  `next_user_message_id` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP PROCEDURE IF EXISTS add_user_message_id_column_if_missing;
DELIMITER $$
CREATE PROCEDURE add_user_message_id_column_if_missing(
  IN p_table_name varchar(64),
  IN p_column_name varchar(64),
  IN p_column_definition text
)
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM information_schema.COLUMNS
    WHERE TABLE_SCHEMA = DATABASE()
      AND TABLE_NAME = p_table_name
      AND COLUMN_NAME = p_column_name
  ) THEN
    SET @ddl = CONCAT('ALTER TABLE `', p_table_name, '` ADD COLUMN `', p_column_name, '` ', p_column_definition);
    PREPARE stmt FROM @ddl;
    EXECUTE stmt;
    DEALLOCATE PREPARE stmt;
  END IF;
END$$
DELIMITER ;

CALL add_user_message_id_column_if_missing('user_message_views', 'user_message_id', 'bigint NOT NULL DEFAULT 0 AFTER peer_seq');

CALL add_user_message_id_column_if_missing('user_dialogs', 'top_user_message_id', 'bigint NOT NULL DEFAULT 0 AFTER top_peer_seq');
CALL add_user_message_id_column_if_missing('user_dialogs', 'read_inbox_max_user_message_id', 'bigint NOT NULL DEFAULT 0 AFTER read_inbox_max_peer_seq');
CALL add_user_message_id_column_if_missing('user_dialogs', 'read_outbox_max_user_message_id', 'bigint NOT NULL DEFAULT 0 AFTER read_outbox_max_peer_seq');
CALL add_user_message_id_column_if_missing('user_dialogs', 'pinned_user_message_id', 'bigint NOT NULL DEFAULT 0 AFTER pinned_peer_seq');
CALL add_user_message_id_column_if_missing('user_dialogs', 'available_min_user_message_id', 'bigint NOT NULL DEFAULT 0 AFTER available_min_peer_seq');

CALL add_user_message_id_column_if_missing('hash_tags', 'hash_tag_user_message_id', 'bigint NOT NULL DEFAULT 0 AFTER hash_tag_message_id');

DROP PROCEDURE IF EXISTS add_user_message_id_column_if_missing;

UPDATE user_message_views v
JOIN (
  SELECT
    user_id,
    peer_type,
    peer_id,
    peer_seq,
    canonical_message_id,
    ROW_NUMBER() OVER (
      PARTITION BY user_id
      ORDER BY date ASC, peer_seq ASC, canonical_message_id ASC
    ) AS assigned_user_message_id
  FROM user_message_views
) numbered
  ON numbered.user_id = v.user_id
  AND numbered.peer_type = v.peer_type
  AND numbered.peer_id = v.peer_id
  AND numbered.peer_seq = v.peer_seq
  AND numbered.canonical_message_id = v.canonical_message_id
SET v.user_message_id = numbered.assigned_user_message_id
WHERE v.user_message_id = 0;

INSERT INTO user_message_sequences (user_id, next_user_message_id)
SELECT
  user_id,
  COALESCE(MAX(user_message_id), 0) + 1 AS next_user_message_id
FROM user_message_views
GROUP BY user_id
ON DUPLICATE KEY UPDATE
  next_user_message_id = GREATEST(user_message_sequences.next_user_message_id, VALUES(next_user_message_id));

UPDATE user_dialogs d
LEFT JOIN user_message_views top_v
  ON top_v.user_id = d.user_id
  AND top_v.peer_type = d.peer_type
  AND top_v.peer_id = d.peer_id
  AND top_v.peer_seq = d.top_peer_seq
LEFT JOIN user_message_views read_in_v
  ON read_in_v.user_id = d.user_id
  AND read_in_v.peer_type = d.peer_type
  AND read_in_v.peer_id = d.peer_id
  AND read_in_v.peer_seq = d.read_inbox_max_peer_seq
LEFT JOIN user_message_views read_out_v
  ON read_out_v.user_id = d.user_id
  AND read_out_v.peer_type = d.peer_type
  AND read_out_v.peer_id = d.peer_id
  AND read_out_v.peer_seq = d.read_outbox_max_peer_seq
LEFT JOIN user_message_views pinned_v
  ON pinned_v.user_id = d.user_id
  AND pinned_v.peer_type = d.peer_type
  AND pinned_v.peer_id = d.peer_id
  AND pinned_v.peer_seq = d.pinned_peer_seq
LEFT JOIN user_message_views available_v
  ON available_v.user_id = d.user_id
  AND available_v.peer_type = d.peer_type
  AND available_v.peer_id = d.peer_id
  AND available_v.peer_seq = d.available_min_peer_seq
SET
  d.top_user_message_id = COALESCE(top_v.user_message_id, 0),
  d.read_inbox_max_user_message_id = COALESCE(read_in_v.user_message_id, 0),
  d.read_outbox_max_user_message_id = COALESCE(read_out_v.user_message_id, 0),
  d.pinned_user_message_id = COALESCE(pinned_v.user_message_id, 0),
  d.available_min_user_message_id = COALESCE(available_v.user_message_id, 0);

UPDATE hash_tags h
JOIN user_message_views v
  ON v.user_id = h.user_id
  AND v.peer_type = h.peer_type
  AND v.peer_id = h.peer_id
  AND v.peer_seq = h.hash_tag_message_id
SET h.hash_tag_user_message_id = v.user_message_id
WHERE h.hash_tag_user_message_id = 0;

DROP PROCEDURE IF EXISTS add_user_message_id_index_if_missing;
DELIMITER $$
CREATE PROCEDURE add_user_message_id_index_if_missing(
  IN p_table_name varchar(64),
  IN p_index_name varchar(64),
  IN p_index_definition text
)
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM information_schema.STATISTICS
    WHERE TABLE_SCHEMA = DATABASE()
      AND TABLE_NAME = p_table_name
      AND INDEX_NAME = p_index_name
  ) THEN
    SET @ddl = CONCAT('ALTER TABLE `', p_table_name, '` ADD ', p_index_definition);
    PREPARE stmt FROM @ddl;
    EXECUTE stmt;
    DEALLOCATE PREPARE stmt;
  END IF;
END$$
DELIMITER ;

CALL add_user_message_id_index_if_missing('user_message_views', 'uk_user_message_id', 'UNIQUE KEY `uk_user_message_id` (`user_id`, `user_message_id`)');
CALL add_user_message_id_index_if_missing('user_message_views', 'uk_user_canonical', 'UNIQUE KEY `uk_user_canonical` (`user_id`, `canonical_message_id`)');
CALL add_user_message_id_index_if_missing('user_message_views', 'idx_user_peer_message_id', 'KEY `idx_user_peer_message_id` (`user_id`, `peer_type`, `peer_id`, `user_message_id`)');
CALL add_user_message_id_index_if_missing('user_message_views', 'idx_user_peer_seq', 'KEY `idx_user_peer_seq` (`user_id`, `peer_type`, `peer_id`, `peer_seq`)');
CALL add_user_message_id_index_if_missing('hash_tags', 'idx_hash_tag_user_message_id', 'KEY `idx_hash_tag_user_message_id` (`user_id`, `hash_tag_user_message_id`)');

DROP PROCEDURE IF EXISTS add_user_message_id_index_if_missing;
