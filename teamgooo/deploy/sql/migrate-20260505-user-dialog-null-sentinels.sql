-- Normalize legacy nullable user_dialogs timestamps for generated DAL string fields.

UPDATE user_dialogs
SET deleted_at = '1970-01-01 00:00:00.000000'
WHERE deleted_at IS NULL;

UPDATE user_dialogs
SET last_pts_at = '1970-01-01 00:00:00.000000'
WHERE last_pts_at IS NULL;

ALTER TABLE user_dialogs
  MODIFY COLUMN deleted_at DATETIME(6) NULL DEFAULT '1970-01-01 00:00:00.000000',
  MODIFY COLUMN last_pts_at DATETIME(6) NULL DEFAULT '1970-01-01 00:00:00.000000';
