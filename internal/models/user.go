package models

import (
	"errors"

	"example.com/task-managment/internal/db"
	"example.com/task-managment/internal/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required" json:"email"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Password string `binding:"required" json:"password"`
}

func (user User) Save() error {
	query := `INSERT INTO users(email, name, role, password) VALUES (?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, user.Name, user.Role, hashPassword)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = userId

	return err
}

func (user *User) ValidateCredentioals() error {
	query := `SELECT id, email, name, role, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Role, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid password")
	}

	return nil
}
