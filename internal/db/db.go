package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"example.com/task-management/internal/utils"
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

	utils.LogSuccess("Database connected successfully!")

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
