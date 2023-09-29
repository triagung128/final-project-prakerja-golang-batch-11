package middleware

import (
	"final-project-prakerja-golang-batch-11/models/response"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenJWT(userId uint) string {
	claims := response.Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resultJWT, _ := token.SignedString([]byte(os.Getenv("SECRET_JWT")))
	return resultJWT
}
