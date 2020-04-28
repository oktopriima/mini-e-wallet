-- +migrate Up
CREATE TABLE IF NOT EXISTS `e-wallet`.`users`
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `username`    VARCHAR(255) NOT NULL,
    `email`       VARCHAR(255) NOT NULL,
    `password`    VARCHAR(255) NOT NULL,
    `last_login`  TIMESTAMP    NOT NULL,
    `is_verified` TINYINT      NOT NULL,
    `is_active`   TINYINT      NOT NULL,
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `email_UNIQUE` (`email` ASC)
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `mkproject`.`users`;