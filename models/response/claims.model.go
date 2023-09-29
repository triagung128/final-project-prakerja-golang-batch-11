package response

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	UserId uint `json:"userId"`
	jwt.RegisteredClaims
}
