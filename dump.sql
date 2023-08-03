CREATE DATABASE IF NOT EXISTS lms;

USE lms;

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
    `user_id` int NOT NULL AUTO_INCREMENT,
    `full_name` varchar(255) NOT NULL,
    `username` varchar(255) NOT NULL,
    `type` char(60) NOT NULL,
    `hash` char(60) NOT NULL,
    PRIMARY KEY(`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `books`;

CREATE TABLE `books` (
    `book_id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `author` varchar(255) NOT NULL,
    `quantity` int NOT NULL,
    PRIMARY KEY(`book_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

DROP TABLE IF EXISTS `requests`;

CREATE TABLE `requests` (
    `request_id` int NOT NULL AUTO_INCREMENT,
    `book_id` int NOT NULL,
    `user_id` int NOT NULL,  
    `state` enum('requested', 'owned') NOT NULL,
    `req_type` enum('borrow', 'return', 'accepted'),  
    PRIMARY KEY(`request_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO users (username,full_name,hash,type) VALUES ("admin","iamadmin","$2a$10$dTT9W.1YqBhy760t54QUV.3ueDACZG/LbELi2PKaXE3PE3/83EYwW","admin");