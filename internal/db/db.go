package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	_ "github.com/mattn/go-sqlite3"    // SQLite driver
)

var DB *sql.DB

func InitDB() {
	dbDriver := os.Getenv("DB_DRIVER")

	var dsn string
	var err error

	switch dbDriver {
	case "mysql":
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		DB, err = sql.Open("mysql", dsn)
	case "sqlite":
		dsn = os.Getenv("DB_SQLITE_FILE")
		DB, err = sql.Open("sqlite3", dsn)
	default:
		log.Fatal("Unsupported DB driver")
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	log.Println("Database connected successfully!")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// createTables()
}

func createTables() {

	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL,
		role ENUM('user', 'admin') DEFAULT 'user',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		log.Fatal(err)
	}

	createTasksTable := `CREATE TABLE IF NOT EXISTS tasks (
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
	);`

	_, err = DB.Exec(createTasksTable)

	if err != nil {
		log.Fatal(err)
	}

}
