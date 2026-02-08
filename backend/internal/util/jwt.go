package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hendrialqori/war-ticket/backend/internal/entity"
)

type UserClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type JwtToken interface {
	Create(user *entity.User) (string, error)
	Verify(token string) (*UserClaims, error)
}

type JwtTokenImpl struct {
	SecretKey []byte
}

// Create implements [JwtToken].
func (j *JwtTokenImpl) Create(user *entity.User) (string, error) {

	claims := UserClaims{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.SecretKey)
}

func (j *JwtTokenImpl) Verify(token string) (*UserClaims, error) {
	claims := &UserClaims{}

	parseToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return j.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !parseToken.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func NewJwtToken(secretKey string) JwtToken {
	return &JwtTokenImpl{
		SecretKey: []byte(secretKey),
	}
}
