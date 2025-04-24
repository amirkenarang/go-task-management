package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"example.com/task-management/internal/cache"
	"example.com/task-management/internal/db"
	"example.com/task-management/internal/models"
	"example.com/task-management/internal/utils"
	"github.com/redis/go-redis/v9"
)

type UserRepository struct {
	DB    *sql.DB
	Redis *redis.Client
}

func NewUserRepository(db *sql.DB, redisClient *redis.Client) *UserRepository {
	return &UserRepository{
		DB:    db,
		Redis: redisClient,
	}
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

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	query := "SELECT id, email, name, role, password, created_at, updated_at FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.Role,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetUserById(id int64) (*models.User, error) {
	// Get user from Redis
	ctx := context.Background()
	cacheKey := fmt.Sprintf("user:%d", id)

	var user models.User
	err := cache.Get(ctx, cacheKey, &user)
	if err == nil {
		return &user, nil
	} else {
		log.Println("Key does not exists in cache!")
	}

	query := "SELECT id, email, name, role, password, created_at, updated_at FROM users WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.Role,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	// Save in reids
	cache.Set(ctx, cacheKey, user, cache.Expiration(cache.ONE_DAY))

	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	query := `
	UPDATE users
	SET email = ?, name = ?, role = ?, password = ?
	WHERE id = ?
	`

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Email, user.Name, user.Role, hashPassword, user.ID)

	return err
}

func (r *UserRepository) Delete(id int64) error {
	query := `DELETE FROM users WHERE id = ?;`

	stmt, err := r.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
