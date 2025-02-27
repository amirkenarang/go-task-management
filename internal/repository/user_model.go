package repository

import (
	"database/sql"
	"errors"

	"example.com/task-managment/internal/models"
	"example.com/task-managment/internal/utils"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Save(user *models.User) error {
	query := `INSERT INTO users(email, name, role, password) VALUES (?, ?, ?, ?)`
	stmt, err := r.DB.Prepare(query)

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
	if err != nil {
		return err
	}

	user.ID = userId

	return nil
}

func (r *UserRepository) ValidateCredentioals(user *models.User) error {
	query := `SELECT id, email, name, role, password FROM users WHERE email = ?`
	row := r.DB.QueryRow(query, user.Email)

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
