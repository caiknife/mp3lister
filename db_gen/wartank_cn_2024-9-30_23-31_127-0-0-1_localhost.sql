# ************************************************************
# Sequel Ace SQL dump
# 版本号： 20073
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# 主机: 127.0.0.1 (MySQL 8.4.2)
# 数据库: wartank_cn
# 生成时间: 2024-09-30 15:31:39 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# 转储表 charge_refund
# ------------------------------------------------------------

DROP TABLE IF EXISTS `charge_refund`;

CREATE TABLE `charge_refund` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `game_center_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `player_id` bigint NOT NULL DEFAULT '0',
  `total_charge` double NOT NULL DEFAULT '0',
  `diamonds` bigint NOT NULL DEFAULT '0',
  `acquired` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `game_center_id` (`game_center_id`),
  KEY `player_id` (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



# 转储表 wt_device
# ------------------------------------------------------------

DROP TABLE IF EXISTS `wt_device`;

CREATE TABLE `wt_device` (
  `id` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '设备ID',
  `player_id` bigint NOT NULL COMMENT '玩家ID',
  `device_region` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '设备地区',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='玩家设备关系表';



# 转储表 wt_gamecenter
# ------------------------------------------------------------

DROP TABLE IF EXISTS `wt_gamecenter`;

CREATE TABLE `wt_gamecenter` (
  `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `bundle_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'BUNDLE ID',
  `player_id` bigint NOT NULL COMMENT '玩家ID',
  `gc_display_name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT 'GC displayName',
  `gc_game_player_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `gc_team_player_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `gc_player_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `bundle_id` (`bundle_id`),
  KEY `player_id` (`player_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='玩家和GameCenter关系表';



# 转储表 wt_gift
# ------------------------------------------------------------

DROP TABLE IF EXISTS `wt_gift`;

CREATE TABLE `wt_gift` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键',
  `gift_code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '礼品码',
  `gift_product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '奖励商品ID',
  `gift_type` tinyint NOT NULL DEFAULT '1' COMMENT '礼品码类型',
  `bundle_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'bundle_id',
  `effective_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生效时间',
  `expire_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '过期时间',
  `gift_status` tinyint NOT NULL DEFAULT '0' COMMENT '状态',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '数据版本锁',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_gift_code` (`gift_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='礼品码表';



# 转储表 wt_legion
# ------------------------------------------------------------

DROP TABLE IF EXISTS `wt_legion`;

CREATE TABLE `wt_legion` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '军团ID',
  `tag` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '8位简码',
  `legion_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '军团名',
  `legion_banner` json DEFAULT NULL COMMENT '军团旗帜',
  `slogan` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标语',
  `limit_trophy` int NOT NULL DEFAULT '0' COMMENT '奖杯限制',
  `limit_week` int NOT NULL DEFAULT '12' COMMENT '周数限制',
  `members` json NOT NULL COMMENT '成员',
  `open_lvl` tinyint NOT NULL DEFAULT '0' COMMENT '开放级别',
  `trophy_score` int NOT NULL DEFAULT '0' COMMENT '奖杯分',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `version` bigint NOT NULL COMMENT '版本锁',
  `war_score` int NOT NULL DEFAULT '0' COMMENT '军团战分数',
  `join_war` tinyint NOT NULL DEFAULT '0' COMMENT '是否加入军团战',
  `tank_num_limit` tinyint NOT NULL DEFAULT '0' COMMENT '加入军团的坦克限制',
  PRIMARY KEY (`id`),
  UNIQUE KEY `tag` (`tag`),
  KEY `legion_name` (`legion_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='军团表';



# 转储表 wt_order
# ------------------------------------------------------------

DROP TABLE IF EXISTS `wt_order`;

CREATE TABLE `wt_order` (
  `id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单ID',
  `player_id` bigint NOT NULL COMMENT '玩家ID',
  `bundle_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'BUNDLE ID',
  `order_status` tinyint NOT NULL COMMENT '订单类型',
  `product_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '产品ID',
  `pay_type` tinyint NOT NULL COMMENT '支付类型',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `trans_id` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '交易ID',
  `order_info` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '订单信息',
  `reward_info` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '奖励信息',
  `version` bigint DEFAULT '0' COMMENT '数据版本',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `player_id` (`player_id`) USING BTREE,
  KEY `trans_id` (`trans_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='订单表';



# 转储表 wt_player
# ------------------------------------------------------------

DROP TABLE IF EXISTS `wt_player`;

CREATE TABLE `wt_player` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '玩家ID',
  `visitor_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '访问ID',
  `tag` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '8位简码',
  `player_lvl` int NOT NULL DEFAULT '1' COMMENT '玩家等级',
  `player_exp` int NOT NULL DEFAULT '0' COMMENT '玩家经验值',
  `nickname` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '玩家名称',
  `icons` json NOT NULL COMMENT '头像',
  `gold_pool` int NOT NULL DEFAULT '600' COMMENT '金币池',
  `gold_pool_ts` int NOT NULL COMMENT '金币池上次更新的秒数',
  `diamond` int NOT NULL DEFAULT '0' COMMENT '钻石',
  `gold` bigint NOT NULL DEFAULT '0' COMMENT '金币',
  `legion_id` int NOT NULL COMMENT '军团ID',
  `legion_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '军团名',
  `legion_position` tinyint NOT NULL COMMENT '军团职位',
  `tier` int NOT NULL DEFAULT '1' COMMENT '战场级别',
  `trophy` int NOT NULL DEFAULT '0' COMMENT '奖杯数',
  `trophy_road` json NOT NULL COMMENT '荣耀之路',
  `vip` json NOT NULL,
  `chest_info` json NOT NULL COMMENT '宝箱信息',
  `garage` json NOT NULL COMMENT '车库',
  `inventory` json NOT NULL COMMENT '库存信息',
  `path_of_valor` json NOT NULL COMMENT '英勇之路',
  `ip_region` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP地区',
  `device_region` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '设备地区',
  `rename_times` int NOT NULL DEFAULT '0' COMMENT '改名次数',
  `settlement_trophy` int NOT NULL DEFAULT '0' COMMENT '结算奖杯数',
  `statistics_info` json NOT NULL COMMENT '统计信息',
  `guide_info` json NOT NULL COMMENT '新手信息',
  `status` int NOT NULL DEFAULT '1' COMMENT '账号状态',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `last_login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后登录时间',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '锁版本',
  `join_war` tinyint NOT NULL DEFAULT '0' COMMENT '是否加入军团战',
  `tank_team` json DEFAULT NULL COMMENT '军团战坦克编组',
  `competitive_rank` json DEFAULT NULL COMMENT '竞技模式排行数据',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `tag` (`tag`) USING BTREE,
  KEY `version` (`version`) USING BTREE,
  KEY `visitor_id` (`visitor_id`),
  KEY `legion_id` (`legion_id`),
  KEY `trophy` (`trophy`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='玩家表';



# 转储表 wt_social
# ------------------------------------------------------------

DROP TABLE IF EXISTS `wt_social`;

CREATE TABLE `wt_social` (
  `id` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '社交账号ID',
  `player_id` bigint NOT NULL COMMENT '玩家ID',
  `social_type` tinyint NOT NULL COMMENT '社交账号类型',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`,`player_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci ROW_FORMAT=DYNAMIC COMMENT='玩家与社交账号绑定信息表';




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
