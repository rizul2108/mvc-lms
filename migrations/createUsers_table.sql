-- +migrate Up

CREATE TABLE `users` (
    `user_id` int NOT NULL AUTO_INCREMENT,
    `full_name` varchar(255) NOT NULL,
    `username` varchar(255) NOT NULL,
    `type` char(60) NOT NULL,
    `hash` char(60) NOT NULL,
    PRIMARY KEY(`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +migrate Down

DROP TABLE users;
