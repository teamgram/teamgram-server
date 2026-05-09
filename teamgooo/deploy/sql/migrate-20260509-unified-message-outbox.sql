-- Unified outbox attrs and forward-ref payloads for canonical messages.

DROP PROCEDURE IF EXISTS add_unified_outbox_column_if_missing;
DELIMITER $$
CREATE PROCEDURE add_unified_outbox_column_if_missing(
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

CALL add_unified_outbox_column_if_missing('canonical_messages', 'message_attrs_schema_version', 'int NOT NULL DEFAULT ''0'' AFTER `media_ref_payload`');
CALL add_unified_outbox_column_if_missing('canonical_messages', 'message_attrs_payload', 'blob AFTER `message_attrs_schema_version`');
CALL add_unified_outbox_column_if_missing('canonical_messages', 'forward_ref_schema_version', 'int NOT NULL DEFAULT ''0'' AFTER `message_attrs_payload`');
CALL add_unified_outbox_column_if_missing('canonical_messages', 'forward_ref_payload', 'blob AFTER `forward_ref_schema_version`');

DROP PROCEDURE IF EXISTS add_unified_outbox_column_if_missing;
