-- migrate-20260209.sql
-- Add index for auth_users table to optimize user_id based queries
-- Resolves full table scan on: SelectAuthKeyIds, SelectListByUserId, DeleteUser

ALTER TABLE `auth_users`
  ADD INDEX `idx_user_deleted` (`user_id`, `deleted`);
