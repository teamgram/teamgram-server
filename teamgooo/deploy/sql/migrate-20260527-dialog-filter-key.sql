-- Dialog filters are keyed by user_id and dialog_filter_id.

DROP PROCEDURE IF EXISTS drop_dialog_filter_slug_unique_key_if_exists;
DELIMITER $$
CREATE PROCEDURE drop_dialog_filter_slug_unique_key_if_exists()
BEGIN
  IF EXISTS (
    SELECT 1
    FROM information_schema.STATISTICS
    WHERE TABLE_SCHEMA = DATABASE()
      AND TABLE_NAME = 'dialog_filters'
      AND INDEX_NAME = 'uk_user_slug'
  ) THEN
    ALTER TABLE `dialog_filters`
      DROP INDEX `uk_user_slug`;
  END IF;
END$$
DELIMITER ;

CALL drop_dialog_filter_slug_unique_key_if_exists();

DROP PROCEDURE IF EXISTS drop_dialog_filter_slug_unique_key_if_exists;
