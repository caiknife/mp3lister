# ************************************************************
# Sequel Ace SQL dump
# 版本号： 20073
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# 主机: 127.0.0.1 (MySQL 8.4.2)
# 数据库: music_data
# 生成时间: 2024-09-30 15:31:57 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# 转储表 books
# ------------------------------------------------------------

DROP TABLE IF EXISTS `books`;

CREATE TABLE `books` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '书籍名称',
  `author` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '作者',
  `genre` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类',
  PRIMARY KEY (`id`),
  KEY `title` (`title`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# 转储表 cars
# ------------------------------------------------------------

DROP TABLE IF EXISTS `cars`;

CREATE TABLE `cars` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `type` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '车辆类型',
  `fuel` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '燃料',
  `transmission` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '传动结构',
  `brand` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '品牌',
  `model` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '型号',
  `year` int NOT NULL DEFAULT '0' COMMENT '生产年份',
  PRIMARY KEY (`id`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# 转储表 movies
# ------------------------------------------------------------

DROP TABLE IF EXISTS `movies`;

CREATE TABLE `movies` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '电影名称',
  `genre` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分类',
  `version` bigint unsigned NOT NULL DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  KEY `name` (`name`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# 转储表 players
# ------------------------------------------------------------

DROP TABLE IF EXISTS `players`;

CREATE TABLE `players` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '姓名',
  `phone` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '电话',
  `email` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮件地址',
  `gold` bigint NOT NULL DEFAULT '0' COMMENT '金币数量',
  `extra` json DEFAULT NULL COMMENT '扩展信息',
  PRIMARY KEY (`id`),
  KEY `deleted_at` (`deleted_at`),
  KEY `phone` (`phone`),
  KEY `email` (`email`),
  CONSTRAINT `players_chk_1` CHECK (json_valid(`extra`))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# 转储表 songs
# ------------------------------------------------------------

DROP TABLE IF EXISTS `songs`;

CREATE TABLE `songs` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '歌曲名称',
  `artist` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '艺术家',
  `album` varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '专辑',
  `bpm` int NOT NULL DEFAULT '0' COMMENT 'BPM',
  `origin_file` varchar(500) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '源文件路径',
  `length` double DEFAULT NULL COMMENT '歌曲长度',
  `version` bigint unsigned NOT NULL DEFAULT '0' COMMENT '乐观锁版本号',
  PRIMARY KEY (`id`),
  KEY `title` (`title`),
  KEY `origin_file` (`origin_file`),
  KEY `deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
