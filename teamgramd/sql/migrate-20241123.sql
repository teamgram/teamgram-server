ALTER TABLE `auth_users` DROP `layer`, DROP `device_model`, DROP `platform`, DROP `system_version`, DROP `api_id`, DROP `app_name`, DROP `app_version`, DROP `ip`, DROP `country`, DROP `region`;
ALTER TABLE `auth_users` CHANGE `date_actived` `date_active` BIGINT NOT NULL DEFAULT '0';
ALTER TABLE `auth_users` ADD `authorization_ttl_days` INT NOT NULL DEFAULT '180' AFTER `android_push_session_id`;
