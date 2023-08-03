-- +migrate Up

CREATE TABLE `books` (
    `book_id` int NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `author` varchar(255) NOT NULL,
    `quantity` int NOT NULL,
    PRIMARY KEY(`book_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +migrate Down

DROP TABLE books;
