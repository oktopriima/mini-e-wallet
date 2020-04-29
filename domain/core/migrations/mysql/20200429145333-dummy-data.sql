-- +migrate Up

INSERT INTO `e-wallet`.users (id, username, email, password, last_login, is_verified, is_active, created_at, updated_at)
VALUES (1, 'oktoprima', 'okto@gmail.com', '$2a$10$BnOnekkuppb3nFcUMgNDJuuvT9f3VXUCdr5AgdUG6m0iGwyDGznFS',
        '2020-04-29 03:20:10', 0, 1, '2020-04-29 03:19:06', '2020-04-29 03:20:10');
INSERT INTO `e-wallet`.users (id, username, email, password, last_login, is_verified, is_active, created_at, updated_at)
VALUES (2, 'user_number_2', 'user2@gmail.com', '$2a$10$.pWdU9W99Bu8Vxkuzpfm8Oohpg3qOrMjzUrXc2irp2fk3zsTqdn3.',
        '2020-04-29 14:15:43', 0, 1, '2020-04-29 14:15:43', '2020-04-29 14:15:43');

INSERT INTO `e-wallet`.user_balances (id, user_id, balance, balance_achieve)
VALUES (1, 1, 75000, 75000);
INSERT INTO `e-wallet`.user_balances (id, user_id, balance, balance_achieve)
VALUES (2, 1, 50000, 25000);
INSERT INTO `e-wallet`.user_balances (id, user_id, balance, balance_achieve)
VALUES (3, 2, 25000, 25000);
INSERT INTO `e-wallet`.user_balances (id, user_id, balance, balance_achieve)
VALUES (4, 1, 25000, 25000);
INSERT INTO `e-wallet`.user_balances (id, user_id, balance, balance_achieve)
VALUES (5, 2, 50000, 25000);

INSERT INTO `e-wallet`.user_balance_histories (id, user_balance_id, balance_before, balance_after, activity, type, ip,
                                               location, user_agent, author)
VALUES (1, 1, 0, 75000, 'top up balance of Rp75,000.00 by oktoprima', 'credit', '172.20.10.3', '',
        'PostmanRuntime/7.24.1', 'oktoprima');
INSERT INTO `e-wallet`.user_balance_histories (id, user_balance_id, balance_before, balance_after, activity, type, ip,
                                               location, user_agent, author)
VALUES (2, 2, 75000, 50000, 'transfer balance of Rp25,000.00 to user_number_2 by oktoprima', 'debit', '172.20.10.3', '',
        'PostmanRuntime/7.24.1', 'oktoprima');
INSERT INTO `e-wallet`.user_balance_histories (id, user_balance_id, balance_before, balance_after, activity, type, ip,
                                               location, user_agent, author)
VALUES (3, 3, 0, 25000, 'transfer balance of Rp25,000.00 from oktoprima', 'credit', '172.20.10.3', '',
        'PostmanRuntime/7.24.1', 'user_number_2');
INSERT INTO `e-wallet`.user_balance_histories (id, user_balance_id, balance_before, balance_after, activity, type, ip,
                                               location, user_agent, author)
VALUES (4, 4, 50000, 25000, 'transfer balance of Rp25,000.00 to user_number_2 by oktoprima', 'debit', '172.20.10.3', '',
        'PostmanRuntime/7.24.1', 'oktoprima');
INSERT INTO `e-wallet`.user_balance_histories (id, user_balance_id, balance_before, balance_after, activity, type, ip,
                                               location, user_agent, author)
VALUES (5, 5, 25000, 50000, 'transfer balance of Rp25,000.00 from oktoprima', 'credit', '172.20.10.3', '',
        'PostmanRuntime/7.24.1', 'user_number_2');
-- +migrate Down
