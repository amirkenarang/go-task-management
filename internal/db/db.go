package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		// Error log, is better that panic
		log.Fatal(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		name TEXT NOT NULL,
		password TEXT NOT NULL,
		role TEXT DEFAULT 'user',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		log.Fatal(err)
	}

	createTasksTable := `
    CREATE TABLE IF NOT EXISTS tasks (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
        deleted_at DATETIME,
        title TEXT NOT NULL,
        description TEXT,
        status TEXT NOT NULL DEFAULT 'To-Do',
        priority TEXT NOT NULL DEFAULT 'Medium',
        due_date DATETIME,
        user_id INTEGER,
        project_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)
    );`

	_, err = DB.Exec(createTasksTable)

	if err != nil {
		log.Fatal(err)
	}

}
