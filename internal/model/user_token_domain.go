package model

import (
	"fmt"
	"os"
	"strings"
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

func VerifyToken(token string) (UserDomainInterface, *rest_error.RestError) {
	secret := os.Getenv(JWT_SECRET_KEY)

	parsedToken, err := jwt.Parse(RemoveBearerPrefix(token), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_error.NewBadRequestError("invalid token")
	})

	if err != nil {
		return nil, rest_error.NewUnathorizedRequestError("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, rest_error.NewUnathorizedRequestError("invalid token")
	}

	return &userDomain{
		id:    claims["user_id"].(string),
		email: claims["email"].(string),
		name:  claims["name"].(string),
		age:   int8(claims["age"].(float64)),
	}, nil
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix("Bearer ", token)
	}

	return token
}
