package auth

import (
	"final-project-prakerja-golang-batch-11/configs"
	"final-project-prakerja-golang-batch-11/middleware"
	"final-project-prakerja-golang-batch-11/models/database"
	"final-project-prakerja-golang-batch-11/models/response"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(context echo.Context) error {
	var request database.User
	context.Bind(&request)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	request.Password = string(hashPassword)

	result := configs.DB.Create(&request)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), fmt.Sprintf("Duplicate entry '%s' for key 'idx_users_email", request.Email)) {
			return context.JSON(http.StatusConflict, response.Base{
				Status:  false,
				Message: "User with that email already exists",
				Data:    nil,
			})
		} else {
			return context.JSON(http.StatusInternalServerError, response.Base{
				Status:  false,
				Message: "Failed add data to database",
				Data:    nil,
			})
		}
	}

	authResponse := response.Auth{
		Id:          request.Id,
		Name:        request.Name,
		Email:       request.Name,
		PhoneNumber: request.PhoneNumber,
		Token:       middleware.GenerateTokenJWT(request.Id),
	}

	return context.JSON(http.StatusOK, response.Base{
		Status:  true,
		Message: "Success Register",
		Data:    authResponse,
	})
}
