
CREATE TABLE IF NOT EXISTS tasks (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT NULL,
    status ENUM('To-Do', 'In Progress', 'Completed') NOT NULL DEFAULT 'To-Do',
    priority ENUM('Low', 'Medium', 'High') NOT NULL DEFAULT 'Medium',
    due_date DATETIME NULL,
    user_id BIGINT,
    project_id BIGINT,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
);


-- If you are using SQLite, use this query:
-- CREATE TABLE IF NOT EXISTS tasks (
-- 		id BIGINT AUTO_INCREMENT PRIMARY KEY,
-- 		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
-- 		updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
-- 		deleted_at DATETIME NULL,
-- 		title VARCHAR(255) NOT NULL,
-- 		description TEXT NULL,
-- 		status ENUM('To-Do', 'In Progress', 'Completed') NOT NULL DEFAULT 'To-Do',
-- 		priority ENUM('Low', 'Medium', 'High') NOT NULL DEFAULT 'Medium',
-- 		due_date DATETIME NULL,
-- 		user_id BIGINT,
-- 		project_id BIGINT,
-- 		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
-- 	);