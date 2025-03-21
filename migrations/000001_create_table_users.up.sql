CREATE TABLE IF NOT EXISTS users (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role ENUM('user', 'admin') DEFAULT 'user',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);


-- If you are using SQLite, use this query:
-- CREATE TABLE IF NOT EXISTS users (
-- 		id BIGINT AUTO_INCREMENT PRIMARY KEY,
-- 		email VARCHAR(255) NOT NULL UNIQUE,
-- 		name VARCHAR(255) NOT NULL,
-- 		password VARCHAR(255) NOT NULL,
-- 		role ENUM('user', 'admin') DEFAULT 'user',
-- 		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
-- 		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
-- 	); 