package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

type AuthUser struct {
	Email  string `json:"email"`
	UserId int64  `json:"user_id"`
}

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(token string) (AuthUser, error) {
	var authUser AuthUser

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Unexpected signing method")
		}

		return []byte([]byte(os.Getenv("JWT_SECRET"))), nil
	})

	if err != nil {
		return authUser, errors.New("Could not parse token.")
	}
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return authUser, errors.New("Invalid token!")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return authUser, errors.New("Could not parse claims.")
	}

	email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	authUser = AuthUser{
		Email:  email,
		UserId: userId,
	}

	return authUser, nil

}
