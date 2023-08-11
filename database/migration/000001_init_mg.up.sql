
CREATE TABLE `users` (
    `userID` int NOT NULL AUTO_INCREMENT,
    `fullName` varchar(255) NOT NULL,
    `username` varchar(255) NOT NULL,
    `type` char(60) NOT NULL,
    `hash` char(60) NOT NULL,
    PRIMARY KEY(`userID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `books` (
    `bookID` int NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `author` varchar(255) NOT NULL,
    `quantity` int NOT NULL,
    PRIMARY KEY(`bookID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `requests` (
    `requestID` int NOT NULL AUTO_INCREMENT,
    `bookID` int NOT NULL,
    `userID` int NOT NULL,  
    `state` enum('Requested', 'Owned') NOT NULL,
    `requestType` enum('Borrow', 'return', 'Accepted'),  
    `requestDate` DATETIME ,
    PRIMARY KEY(`requestID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


