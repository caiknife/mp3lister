DROP TABLE IF EXISTS `songs`;
CREATE TABLE `songs`
(
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `created_at`  DATETIME                 DEFAULT NULL,
    `updated_at`  DATETIME                 DEFAULT NULL,
    `deleted_at`  DATETIME                 DEFAULT NULL,
    `title`       VARCHAR(200)     NOT NULL DEFAULT '' COMMENT '歌曲名称',
    `artist`      VARCHAR(200)     NOT NULL DEFAULT '' COMMENT '艺术家',
    `album`       VARCHAR(200)     NOT NULL DEFAULT '' COMMENT '专辑',
    `bpm`         VARCHAR(5)      NOT NULL DEFAULT '' COMMENT 'BPM',
    `origin_file` VARCHAR(500)    NOT NULL DEFAULT '' COMMENT '源文件路径',
    PRIMARY KEY (`id`),
    KEY `idx_songs_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
