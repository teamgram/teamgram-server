-- phpMyAdmin SQL Dump
-- version 5.1.3
-- https://www.phpmyadmin.net/
--
-- 主机： 127.0.0.1
-- 生成日期： 2024-01-03 22:36:21
-- 服务器版本： 8.0.28
-- PHP 版本： 8.2.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";

--
-- 数据库： `teamgram`
--

-- --------------------------------------------------------

--
-- 表的结构 `close_friends`
--

CREATE TABLE IF NOT EXISTS `close_friends` (
                                               `id` bigint NOT NULL AUTO_INCREMENT,
                                               `user_id` bigint NOT NULL,
                                               `close_friend_id` bigint NOT NULL,
                                               `date` bigint NOT NULL,
                                               `deleted` tinyint(1) NOT NULL DEFAULT '0',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `user_id` (`user_id`,`close_friend_id`)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
COMMIT;
