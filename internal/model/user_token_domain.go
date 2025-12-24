package model

import (
	"fmt"
	"os"
	"time"

	"github.com/HavocJean/study-go/internal/config/rest_error"
	"github.com/golang-jwt/jwt"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (u *userDomain) GenerateToken() (string, *rest_error.RestError) {
	secret := os.Getenv("JWT_SECRET_KEY")

	claims := jwt.MapClaims{
		"user_id": u.id,
		"email":   u.email,
		"name":    u.name,
		"age":     u.age,
		"exp":     time.Now().Add(time.Hour * 24),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_error.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt, err=%s", err.Error()),
		)
	}

	return tokenString, nil
}
