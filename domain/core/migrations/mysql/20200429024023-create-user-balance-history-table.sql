-- +migrate Up
CREATE TABLE IF NOT EXISTS `e-wallet`.`user_balance_histories`
(
    `id`              INT UNSIGNED             NOT NULL AUTO_INCREMENT,
    `user_balance_id` INT UNSIGNED             NOT NULL,
    `balance_before`  FLOAT                    NOT NULL,
    `balance_after`   FLOAT                    NOT NULL,
    `activity`        VARCHAR(255)             NOT NULL,
    `type`            ENUM ('debit', 'credit') NOT NULL,
    `ip`              VARCHAR(45)              NOT NULL,
    `location`        VARCHAR(255)             NOT NULL,
    `user_agent`      VARCHAR(255)             NOT NULL,
    `author`          VARCHAR(255)             NOT NULL,
    PRIMARY KEY (`id`),
    INDEX `fk_user_balance_histories_user_balances1_idx` (`user_balance_id` ASC),
    CONSTRAINT `fk_user_balance_histories_user_balances1`
        FOREIGN KEY (`user_balance_id`)
            REFERENCES `e-wallet`.`user_balances` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `e-wallet`.`user_balance_histories`;