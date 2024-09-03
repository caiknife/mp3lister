CREATE TABLE `charge_refund`
(
    `id`             bigint unsigned                                               NOT NULL AUTO_INCREMENT,
    `game_center_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `player_id`      bigint                                                        NOT NULL DEFAULT '0',
    `total_charge`   double                                                        NOT NULL DEFAULT '0',
    `diamonds`       bigint                                                        NOT NULL DEFAULT '0',
    `acquired`       tinyint(1)                                                    NOT NULL DEFAULT '0',
    `created_at`     datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     datetime                                                      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at`     datetime                                                               DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `game_center_id` (`game_center_id`),
    KEY `player_id` (`player_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;
