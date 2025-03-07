package models

import "time"

type User struct {
	ID        int64
	Email     string     `binding:"required" json:"email"`
	Name      string     `json:"name"`
	Role      string     `json:"role"`
	Password  string     `binding:"required" json:"password"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
