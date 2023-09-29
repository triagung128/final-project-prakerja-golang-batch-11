package middleware

import (
	"final-project-prakerja-golang-batch-11/models/response"
	"os"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetConfig() echojwt.Config {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(response.Claims)
		},
		SigningKey: []byte(os.Getenv("SECRET_JWT")),
	}

	return config
}
