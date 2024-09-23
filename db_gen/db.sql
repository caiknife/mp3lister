DROP TABLE IF EXISTS `books`;
CREATE TABLE `books`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime                     DEFAULT NULL,
    `title`      varchar(200)        NOT NULL DEFAULT '' COMMENT '书籍名称',
    `author`     varchar(200)        NOT NULL DEFAULT '' COMMENT '作者',
    `genre`      varchar(200)        NOT NULL DEFAULT '' COMMENT '分类',
    PRIMARY KEY (`id`),
    KEY `title` (`title`),
    KEY `deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `cars`;
CREATE TABLE `cars`
(
    `id`           bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`   datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`   datetime                     DEFAULT NULL,
    `type`         varchar(200)        NOT NULL DEFAULT '' COMMENT '车辆类型',
    `fuel`         varchar(200)        NOT NULL DEFAULT '' COMMENT '燃料',
    `transmission` varchar(200)        NOT NULL DEFAULT '' COMMENT '传动结构',
    `brand`        varchar(200)        NOT NULL DEFAULT '' COMMENT '品牌',
    `model`        varchar(200)        NOT NULL DEFAULT '' COMMENT '型号',
    `year`         int(11)             NOT NULL DEFAULT '0' COMMENT '生产年份',
    PRIMARY KEY (`id`),
    KEY `deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `movies`;
CREATE TABLE `movies`
(
    `id`         bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime                     DEFAULT NULL,
    `name`       varchar(200)        NOT NULL DEFAULT '' COMMENT '电影名称',
    `genre`      varchar(200)        NOT NULL DEFAULT '' COMMENT '分类',
    `version`    int(11)                      DEFAULT NULL COMMENT '乐观锁版本号',
    PRIMARY KEY (`id`),
    KEY `name` (`name`),
    KEY `deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 101
  DEFAULT CHARSET = utf8mb4;

DROP TABLE IF EXISTS `songs`;
CREATE TABLE `songs`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at`  datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  datetime            NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`  datetime                     DEFAULT NULL,
    `title`       varchar(200)        NOT NULL DEFAULT '' COMMENT '歌曲名称',
    `artist`      varchar(200)        NOT NULL DEFAULT '' COMMENT '艺术家',
    `album`       varchar(200)        NOT NULL DEFAULT '' COMMENT '专辑',
    `bpm`         int(11)             NOT NULL DEFAULT '0' COMMENT 'BPM',
    `origin_file` varchar(500)        NOT NULL DEFAULT '' COMMENT '源文件路径',
    PRIMARY KEY (`id`),
    KEY `title` (`title`),
    KEY `origin_file` (`origin_file`),
    KEY `deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
