-- migrate-20260209.sql
-- Add index for auth_users table to optimize user_id based queries
-- Resolves full table scan on: SelectAuthKeyIds, SelectListByUserId, DeleteUser

-- ALTER TABLE `auth_users`
--  ADD INDEX `idx_user_deleted` (`user_id`, `deleted`);

-- [可选] 移除冗余索引（需先验证线上查询计划）
-- ALTER TABLE `auth_users` DROP KEY `auth_key_id_2`;  
