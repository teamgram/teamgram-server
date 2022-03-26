ALTER TABLE `photo_sizes` ADD `cached_type` INT NOT NULL DEFAULT '0' AFTER `file_path`;
ALTER TABLE `photo_sizes` ADD `cached_bytes` VARCHAR(4096) NOT NULL DEFAULT '' AFTER `cached_type`;
ALTER TABLE `photo_sizes` DROP INDEX `volume_id`;
ALTER TABLE `photo_sizes` DROP `volume_id`, DROP `local_id`, DROP `secret`, DROP `has_stripped`, DROP `stripped_bytes`;
ALTER TABLE `video_sizes` DROP INDEX `volume_id`;
ALTER TABLE `video_sizes` DROP `volume_id`, DROP `local_id`, DROP `secret`;
