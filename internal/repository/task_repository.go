package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"example.com/task-managment/internal/cache"
	"example.com/task-managment/internal/db"
	"example.com/task-managment/internal/models"
	"github.com/redis/go-redis/v9"
)

type TaskRepository struct {
	DB    *sql.DB
	Redis *redis.Client
}

// These functions are constructor functions in Go.
func NewTaskRepository(db *sql.DB, redisClient *redis.Client) *TaskRepository {
	return &TaskRepository{
		DB:    db,
		Redis: redisClient,
	}
}

func (r *TaskRepository) Save(task *models.Task) error {
	query := `
	INSERT INTO tasks(title, description, status, priority, due_date, user_id, project_id)
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(task.Title, task.Description, task.Status, task.Priority, task.DueDate, task.UserID, task.ProjectID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	task.ID = id

	// Store task in Redis cache
	ctx := context.Background()
	cacheKey := fmt.Sprintf("task:%d", task.ID)
	err = cache.Set(ctx, cacheKey, task, cache.Expiration(cache.ONE_DAY))
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) GetAllTasks() ([]models.Task, error) {
	query := "SELECT id, title, description, status, priority, due_date, user_id, project_id, created_at, updated_at FROM tasks"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
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

func (r *TaskRepository) GetTaskById(id int64) (*models.Task, error) {
	// Get task from Redis
	ctx := context.Background()
	cacheKey := fmt.Sprintf("task:%d", id)

	var task models.Task
	err := cache.Get(ctx, cacheKey, &task)
	if err == nil {
		return &task, nil
	} else {
		log.Println("Key does not exists in cache!")
	}

	query := "SELECT id, title, description, status, priority, due_date, user_id, project_id, created_at, updated_at FROM tasks WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	err = row.Scan(
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

	// Save in reids
	cache.Set(ctx, cacheKey, task, cache.Expiration(cache.ONE_DAY))

	return &task, nil
}

func (r *TaskRepository) Update(task *models.Task) error {
	query := `
	UPDATE tasks
	SET title = ?, description = ?, status = ?, priority = ?, due_date = ?
	WHERE id = ?
	`

	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(task.Title, task.Description, task.Status, task.Priority, task.DueDate, task.ID)

	return err
}

func (r *TaskRepository) Delete(id int64) error {
	query := `DELETE FROM tasks WHERE id = ?;`

	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
