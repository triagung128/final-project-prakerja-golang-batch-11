package utils

import (
	"final-project-prakerja-golang-batch-11/models/response"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetUserId(context echo.Context) uint {
	user := context.Get("user").(*jwt.Token)
	claims := user.Claims.(*response.Claims)
	return claims.UserId
}
