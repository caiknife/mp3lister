DROP TABLE IF EXISTS `songs`;
CREATE TABLE `songs`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at`  datetime                     DEFAULT NULL,
    `updated_at`  datetime                     DEFAULT NULL,
    `deleted_at`  datetime                     DEFAULT NULL,
    `title`       varchar(200)        NOT NULL DEFAULT '' COMMENT '歌曲名称',
    `artist`      varchar(200)        NOT NULL DEFAULT '' COMMENT '艺术家',
    `album`       varchar(200)        NOT NULL DEFAULT '' COMMENT '专辑',
    `bpm`         int          NOT NULL DEFAULT '0' COMMENT 'BPM',
    `origin_file` varchar(500)        NOT NULL DEFAULT '' COMMENT '源文件路径',
    PRIMARY KEY (`id`),
    KEY `idx_songs_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
