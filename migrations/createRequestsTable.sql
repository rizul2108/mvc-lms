-- +migrate Up

CREATE TABLE `requests` (
    `request_id` int NOT NULL AUTO_INCREMENT,
    `book_id` int NOT NULL,
    `user_id` int NOT NULL,  
    `state` enum('Requested', 'Owned') NOT NULL,
    `req_type` enum('Borrow', 'return', 'Accepted'),  
    PRIMARY KEY(`request_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


-- +migrate Down

DROP TABLE requests;
