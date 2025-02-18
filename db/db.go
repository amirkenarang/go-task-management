package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
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
        project_id INTEGER
    );`

	_, err := DB.Exec(createTasksTable)

	if err != nil {
		panic("Could not create tasks table.")
	}

}
