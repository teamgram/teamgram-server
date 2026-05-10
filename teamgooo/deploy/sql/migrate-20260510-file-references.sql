-- Media-owned opaque file_reference handles.

ALTER TABLE `photo_sizes`
  MODIFY COLUMN `stripped_bytes` varbinary(4096) NOT NULL DEFAULT '';

CREATE TABLE IF NOT EXISTS `file_references` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `ref_hash` varbinary(25) NOT NULL,
  `domain` varchar(32) NOT NULL DEFAULT '',
  `media_id` bigint NOT NULL DEFAULT '0',
  `access_hash` bigint NOT NULL DEFAULT '0',
  `object_id` varchar(255) NOT NULL DEFAULT '',
  `origin_domain` varchar(64) NOT NULL DEFAULT '',
  `origin_id` bigint NOT NULL DEFAULT '0',
  `expire_at` bigint NOT NULL DEFAULT '0',
  `revoked_at` bigint NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `ref_hash` (`ref_hash`),
  KEY `media_lookup` (`domain`,`media_id`,`access_hash`),
  KEY `expire_at` (`expire_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
