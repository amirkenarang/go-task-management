package models

import (
	"time"

	"example.com/task-managment/internal/db"
)

type Task struct {
	ID          int64      `json:"id"`
	Title       string     `binding:"required" json:"title"`       // Title of the task
	Description string     `binding:"required" json:"description"` // Detailed description of the task
	Status      string     `binding:"required" json:"status"`      // Status: "To-Do", "In Progress", "Completed"
	Priority    string     `binding:"required" json:"priority"`    // Priority: "Low", "Medium", "High"
	DueDate     *time.Time `binding:"required" json:"due_date"`    // Optional due date for the task
	UserID      int64      `json:"user_id"`                        // Foreign key: the user to whom the task is assigned
	ProjectID   int64      `json:"project_id"`                     // Foreign key: the project this task belongs to
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

var tasks = []Task{}

func (t *Task) Save() error {
	query := `
	INSERT INTO tasks(title, description, status, priority, due_date, user_id, project_id) 
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(t.Title, t.Description, t.Status, t.Priority, t.DueDate, t.UserID, t.ProjectID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	t.ID = id
	return err
}

func GetAllTasks() ([]Task, error) {
	query := "SELECT id, title, description, status, priority, due_date, user_id, project_id, created_at, updated_at FROM tasks"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.Priority,
			&task.DueDate,
			&task.UserID,
			&task.ProjectID,
			&task.CreatedAt,
			&task.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTaskById(id int64) (*Task, error) {
	query := "SELECT id, title, description, status, priority, due_date, user_id, project_id, created_at, updated_at FROM tasks WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var task Task

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.Priority,
		&task.DueDate,
		&task.UserID,
		&task.ProjectID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (task Task) Update() error {
	query := `
	UPDATE tasks
	SET title = ?, description = ?, status = ?, priority = ?, due_date = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(task.Title, task.Description, task.Status, task.Priority, task.DueDate, task.ID)

	return err
}

func (task Task) Delete() error {
	query := `DELETE FROM tasks WHERE id = ?;`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(task.ID)

	return err
}
