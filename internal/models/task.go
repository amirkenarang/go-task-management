package models

import (
	"time"
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

// var tasks = []Task{}
