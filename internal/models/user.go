package models

type User struct {
	ID       int64
	Email    string `binding:"required" json:"email"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `binding:"required" json:"password"`
}
