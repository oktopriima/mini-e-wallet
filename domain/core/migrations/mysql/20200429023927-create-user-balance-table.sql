-- +migrate Up
CREATE TABLE IF NOT EXISTS `e-wallet`.`user_balances`
(
    `id`              INT UNSIGNED     NOT NULL AUTO_INCREMENT,
    `user_id`         INT(10) UNSIGNED NOT NULL,
    `balance`         FLOAT            NOT NULL,
    `balance_achieve` FLOAT            NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_user_balances_users1_idx` (`user_id` ASC),
    CONSTRAINT `fk_user_balances_users1`
        FOREIGN KEY (`user_id`)
            REFERENCES `e-wallet`.`users` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `e-wallet`.`user_balances`;