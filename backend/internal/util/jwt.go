package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hendrialqori/war-ticket/backend/internal/entity"
)

func CreateToken(secretKey []byte, user *entity.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	result, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return result, nil
}

func VerififyToken(secretKey []byte, tokenString string) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("Invalid token")
	}

	return nil
}
