-- Persist temp auth key expiry so reconnect survives Redis lifecycle loss.

DROP PROCEDURE IF EXISTS add_auth_key_expires_at_if_missing;
DELIMITER $$
CREATE PROCEDURE add_auth_key_expires_at_if_missing()
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM information_schema.COLUMNS
    WHERE TABLE_SCHEMA = DATABASE()
      AND TABLE_NAME = 'auth_keys'
      AND COLUMN_NAME = 'expires_at'
  ) THEN
    ALTER TABLE `auth_keys`
      ADD COLUMN `expires_at` bigint NOT NULL DEFAULT 0 AFTER `media_temp_auth_key_id`;
  END IF;
END$$
DELIMITER ;

CALL add_auth_key_expires_at_if_missing();

DROP PROCEDURE IF EXISTS add_auth_key_expires_at_if_missing;

UPDATE `auth_keys`
SET `expires_at` = UNIX_TIMESTAMP(`created_at`) + 604800
WHERE `expires_at` = 0
  AND `auth_key_type` IN (1, 2);
